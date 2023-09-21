package version

import "strconv"

type Version uint64

func (v Version) String() string {
	return strconv.Itoa(int(v))
}

func New(version uint64) (*Version, error) {

	v := Version(version)

	return &v, nil
}
