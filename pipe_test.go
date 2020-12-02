package main

import "testing"

func TestPipe(t *testing.T) {
	pc := newPipeCommand("op=dump&key=1")
	t.Logf("%+v", pc.params)
	err := pc.dispatch()
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	t.Logf("test1 ok")
}
