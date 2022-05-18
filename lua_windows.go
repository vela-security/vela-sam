package sam

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
	"github.com/vela-security/vela-public/pipe"
)

var xEnv assert.Environment

func samDumpL(L *lua.LState) int {
	exe := L.CheckString(1)
	pip := pipe.NewByLua(L, pipe.Seek(1))

	if err := checksum(exe); err != nil {
		L.RaiseError("windows %s checksum fail %v", exe, err)
		return 0
	}

	dump(exe, pip, xEnv.Clone(L))
	return 0
}

func WithEnv(env assert.Environment) {
	xEnv = env

	xEnv.Set("sam_dump", lua.NewFunction(samDumpL))

}
