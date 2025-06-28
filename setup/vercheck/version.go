package vercheck

import (
	"fmt"
	"strconv"
	"strings"
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
