package echo

import "testing"

func TestEcho(t *testing.T) {
	hello := "helloworld"
	v := Echo(hello)
	if v != hello {
		t.Error("Expected ", hello, ", got ", v)
	}
}
