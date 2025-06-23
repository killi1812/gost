package check

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/killi1812/gost/util/cerror"
)

type version struct {
	Major int
	Minor int
	Patch int
}

func (v version) String() string {
	return fmt.Sprintf("go%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *version) Parse(data string) error {
	data = data[2:]
	parts := strings.Split(data, ".")

	// parse Major version
	{
		rez, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		v.Major = rez
	}

	// parse Minor version
	{
		rez, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}
		v.Minor = rez
	}

	// parse Patch version
	{
		rez, err := strconv.Atoi(parts[2])
		if err != nil {
			return err
		}
		v.Patch = rez
	}

	return nil
}

// Compare compares 2 versions
//
// -1 if version is lower then given
//
// 0 if version is same as given
//
// +1 if version is grater then given
//
// use as following ver.Compare(_MIN_GOLANG_VERSION) < 0
func (v version) Compare(other version) int {
	if v.Major < other.Major {
		return -1
	} else if v.Major > other.Major {
		return 1
	}

	if v.Minor < other.Minor {
		return -1
	} else if v.Minor > other.Minor {
		return 1
	}

	if v.Patch < other.Patch {
		return -1
	} else if v.Patch > other.Patch {
		return 1
	}

	return 0
}

var _MIN_GOLANG_VERSION = version{
	Major: 1,
	Minor: 24,
	Patch: 0,
}

var golangBin = ""

func GetGolangBin() string {
	return golangBin
}

func LoadGolangBin() error {
	// Check if go can be found in standard shell
	const _GO_PATH = "go"
	err := runGo(_GO_PATH)

	// Continue if error is ErrGoMissing
	if err != nil && !errors.Is(err, cerror.ErrGoMissing) {
		return err
	}

	// TODO: check other locations

	return nil
}

func runGo(path string) error {
	cmd := exec.Command(path, "version")
	data, err := cmd.Output()
	if err != nil {
		return isMissing(err)
	}

	// Parse version
	rez := string(data)
	println(rez)

	parts := strings.Split(rez, " ")

	ver := version{}
	if err := ver.Parse(parts[2]); err != nil {
		panic(err)
	}
	if ver.Compare(_MIN_GOLANG_VERSION) < 0 {
		return cerror.ErrGoVersionNotSupported
	}

	// Set go executable path
	golangBin = path
	fmt.Println(ver.String())
	return nil
}

func Init() {
	err := LoadGolangBin()
	if err != nil {
		panic(err)
	}
	println(GetGolangBin())
}

// isMissing is used to determin if error cerror.ErrGoMissing or unexpected error
func isMissing(err error) error {
	nerr, ok := err.(*exec.Error)
	if !ok {
		return err
	}

	tmp := nerr.Unwrap()

	if tmp.Error() == "executable file not found in $PATH" {
		return cerror.ErrGoMissing
	}

	return tmp
}
