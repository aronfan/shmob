package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	sc "github.com/aronfan/shmcore"
)

type bstat struct {
	bytes uint32
	count uint32
	frees uint32
}

type pipecmd struct {
	params map[string]string
}

func newPipeCommand(cmd string) *pipecmd {
	ss := strings.Split(cmd, "&")
	params := make(map[string]string)
	for i := 0; i < len(ss); i++ {
		s := ss[i]
		kv := strings.Split(s, "=")
		if len(kv) >= 2 {
			params[kv[0]] = kv[1]
		} else {
			params[kv[0]] = ""
		}
	}
	return &pipecmd{params: params}
}

func (pc *pipecmd) dispatch() error {
	op, ok := pc.params["op"]
	if !ok {
		return fmt.Errorf("op not exist")
	}

	switch op {
	case "stat":
		return pc.stat()
	}
	return nil
}

func (pc *pipecmd) stat() error {
	ok := sc.ResumeEnabled()
	if !ok {
		return fmt.Errorf("resume not enabled")
	}
	s, ok := pc.params["key"]
	if !ok {
		return fmt.Errorf("key not exist")
	}
	if s == "" {
		return fmt.Errorf("key is empty")
	}
	k, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	key := uint32(k)
	err = sc.IsShmExist(key)
	if err != nil {
		return err
	}

	seg, err := sc.NewSegment(key, 0)
	if err != nil {
		return err
	}

	err = seg.Attach()
	if err != nil {
		return err
	}

	m := make(map[uint16]*bstat)
	seg.Observe(
		func(shead *sc.SegmentHead) {
		},
		func(index uint16, bhead *sc.BucketHead) {
			m[index] = &bstat{bytes: bhead.GetBytes(), count: bhead.GetCount(), frees: 0}
		},
		func(hindex uint16, uindex uint32, unit *sc.BucketUnit) {
			l := unit.GetLen()
			if l == 0 {
				stat, ok := m[hindex]
				if ok {
					stat.frees++
				}
			}
		},
	)

	seg.Detach()

	fmt.Fprintf(os.Stderr, "shm key=%d\n", key)
	fmt.Fprintln(os.Stderr, "bytes\tfree\ttotal")
	for _, v := range m {
		fmt.Printf("%d\t%d\t%d\n", v.bytes, v.frees, v.count)
	}

	return nil
}

func pipe(cmd string) error {
	pc := newPipeCommand(cmd)
	return pc.dispatch()
}
