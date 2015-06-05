package minimembuf_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/extemporalgenome/minimembuf"
)

func testBuffer(t *testing.T, rw io.ReadWriter) {
	const (
		input  = `hello small buffers`
		output = `"` + input + `"` + "\n"
	)
	fmt.Fprintf(rw, "%q\n", input)
	data, err := ioutil.ReadAll(rw)
	if err != nil {
		// really shouldn't happen
		t.Fatal(err)
	}
	if string(data) != output {
		t.Fatalf("got %q, want %q", input, output)
	}
}

func TestBuffer(t *testing.T) {
	testBuffer(t, new(minimembuf.Buffer))
}
