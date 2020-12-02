package main

import (
	"fmt"
	"strconv"
	"strings"

	sc "github.com/aronfan/shmcore"
)

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
	case "dump":
		return pc.dump()
	}
	return nil
}

func (pc *pipecmd) dump() error {
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
	bytes, err := sc.GetShmBytesByKey(key)
	if err != nil {
		return err
	}
	if bytes == 0 {
		return fmt.Errorf("shm not exist")
	}

	return nil
}

func pipe(cmd string) error {
	pc := newPipeCommand(cmd)
	return pc.dispatch()
}
