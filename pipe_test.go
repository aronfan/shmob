package main

import "testing"

func TestStat(t *testing.T) {
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

func BenchmarkStat(b *testing.B) {
	pc := newPipeCommand("key=255&op=stat")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pc.dispatch()
	}
}
