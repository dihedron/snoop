package osptime

import (
	"encoding/json"
	"log/slog"
	"testing"
)

func TestOpenStackTime(t *testing.T) {
	type Test struct {
		Time OpenStackTime
	}

	var data = `
		{"Time": "2021-09-09T07:52:34.990592+0000"}
	`

	a := Test{}
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		t.Fatal(err)
	}
	slog.Debug("after unmarshalling", "elapsed", a.Time.String())
	//fmt.Println(json.Unmarshal([]byte(data), &a))
	//fmt.Println(a.Time.String())

}
