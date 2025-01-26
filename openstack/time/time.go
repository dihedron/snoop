package time

import (
	"fmt"
	"strings"
	gotime "time"
)

type OpenStackTime gotime.Time

const layout = "2006-01-02T15:04:05-0700"

var zero = (gotime.Time{}).UnixNano()

func (t *OpenStackTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*t = OpenStackTime(gotime.Time{})
		return
	}
	t0, err := gotime.Parse(layout, s)
	*t = OpenStackTime(t0)
	return
}

func (t *OpenStackTime) MarshalJSON() ([]byte, error) {
	if (gotime.Time(*t)).UnixNano() == zero {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", (gotime.Time(*t)).Format(layout))), nil
}

func (t *OpenStackTime) String() string {
	return (gotime.Time(*t)).String()
}
