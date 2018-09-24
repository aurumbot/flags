package flags

import (
	"testing"
)

func TestNoCommand(t *testing.T) {
	argstr := `--name gabe miller -foo bar`
	myflags := Parse(argstr)
	for _, f := range myflags {
		t.Log(f)
	}
}
