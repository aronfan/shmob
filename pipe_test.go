package main

import "testing"

func TestPipe(t *testing.T) {
	pc := newPipeCommand("key=255&op=stat")
	t.Logf("%+v", pc.params)
	err := pc.dispatch()
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Logf("test shm=255 ok")

	pc = newPipeCommand("key=256&op=stat")
	t.Logf("%+v", pc.params)
	err = pc.dispatch()
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Logf("test shm=256 ok")
}
