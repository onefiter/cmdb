package example

import (
	"strconv"
	"testing"
)

func Test_strconv_ParseInt(t *testing.T) {

	code := "-1"

	c, err := strconv.ParseInt(code, 10, 64)

	if err != nil {
		t.Log(err)
	}

	t.Log("code:", c)

}
