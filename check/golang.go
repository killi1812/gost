package check

import (
	"os/exec"

	"github.com/killi1812/gost/util/cerror"
)

var golangBin = ""

func GetGolangBin() string {
	return golangBin
}

func LoadGolangBin() error {
	// Check if go can be found in standard shell
	cmd := exec.Command("go", "version")
	data, err := cmd.Output()
	if err != nil {
		return isMissing(err)
	}
	// TODO: parse data for version
	rez := string(data)
	println(rez)

	// TODO: set go bin variable
	return nil
}

func Init() {
	err := LoadGolangBin()
	if err != nil {
		panic(err)
	}
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
