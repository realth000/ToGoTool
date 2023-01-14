package rpm

import (
	"regexp"
	"strings"
)

var (
	rpmNameRegexp = regexp.MustCompile(`^[^ \d]+.+-.+-.+$`)
)

// SplitRpmName splits a full name of package then returns the
// actual name, package version and packager release.
// e.g.: openssh-8.8p1-7.fc37 => openssh, 8.2p1, 7.fc37
func SplitRpmName(fullName string) (name string, version string, release string) {
	name, version, release = "", "", ""
	// Rpm package name not starts with 0-9 and contains more than two "-".
	if !rpmNameRegexp.MatchString(fullName) {
		list := strings.Split(fullName, "-")
		if len(list) == 1 {
			name = list[0]
		} else if len(list) == 2 {
			name = list[0]
			version = list[1]
		} else if len(list) == 3 {
			name = list[0]
			version = list[1]
			release = list[2]
		} else if len(list) > 4 {
			name = strings.Join(list[:len(list)-2], "-")
			version = list[len(list)-2]
			release = list[len(list)-1]
		}
		return name, version, release
	}
	t := strings.Split(fullName, `-`)
	l := len(t)
	if l >= 3 {
		release = t[l-1]
		version = t[l-2]
		name = strings.Join(t[0:l-2], `-`)
	}
	return name, version, release
}
