package osptime

import (
	"fmt"
	"strings"
	"time"
)

type OpenStackTime time.Time

const layout = "2006-01-02T15:04:05-0700"

var zero = (time.Time{}).UnixNano()

func (t *OpenStackTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*t = OpenStackTime(time.Time{})
		return
	}
	t0, err := time.Parse(layout, s)
	*t = OpenStackTime(t0)
	return
}

func (t *OpenStackTime) MarshalJSON() ([]byte, error) {
	if (time.Time(*t)).UnixNano() == zero {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", (time.Time(*t)).Format(layout))), nil
}

func (t *OpenStackTime) String() string {
	return (time.Time(*t)).String()
}
