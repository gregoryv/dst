package dst

import (
	"bytes"
	"testing"

	"github.com/gregoryv/golden"
)

func TestErrors(t *testing.T) {
	// check readability of the actual errors
	all := []error{
		SetContent(nil, "badfile.x"),
		SetFloat(nil, "v"),
		SetInt(nil, "v"),
		SetURL(nil, "htt p://"),
		SetDuration(nil, "v"),
		SetUint32(nil, "v"),
		SetBool(nil, "NO"),
	}
	var buf bytes.Buffer
	for _, err := range all {
		buf.WriteString(err.Error())
		buf.WriteString("\n")
	}
	golden.Assert(t, buf.String())
}
