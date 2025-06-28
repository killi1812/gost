// Package vercheck is used to check golang version
package vercheck

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/killi1812/gost/util/cerror"
)

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
	// Check if go can be found in standard shell path
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
