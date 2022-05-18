package sam

import (
	"bufio"
	"fmt"
	audit "github.com/vela-security/vela-audit"
	"github.com/vela-security/vela-public/auxlib"
	"github.com/vela-security/vela-public/lua"
	"github.com/vela-security/vela-public/pipe"
	"os/exec"
	"syscall"
	"time"
)

var (
	hashTab = map[string]bool{
		"b3ebe57160a9606563e587ed73ec6ca9": true,
		"310d3d16d623f38bfb42aabf3cd30e53": true,
		"bc4aef584fc07e0d1cd15579a7fa821c": true,
	}
)

func checksum(exe string) error {
	hash, err := auxlib.FileMd5(exe)
	if err != nil {
		return err
	}

	_, ok := hashTab[hash]
	if !ok {
		return fmt.Errorf("checksuam fail")
	}
	return nil
}

func newSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		HideWindow: true,
	}
}

func dump(exe string, px *pipe.Px, L *lua.LState) {
	cmd := exec.Command(exe)
	defer func() {
		if cmd.Process == nil {
			return
		}
		cmd.Process.Kill()
		audit.Debug("cmd window sam dump over.")
	}()

	cmd.SysProcAttr = newSysProcAttr()

	out, err := cmd.StdoutPipe()
	if err != nil {
		L.RaiseError("windows sam dump stdout pipe got fail %v", err)
		return
	}

	if e := cmd.Start(); e != nil {
		L.RaiseError("cmd window sam dump %v", err)
		return
	}

	verbose := bufio.NewScanner(out)
	co := xEnv.Clone(L)
	defer xEnv.Free(co)

	for verbose.Scan() {
		line := verbose.Text()

		px.Do(lua.S2L(line), co, func(err error) {
			audit.Errorf("cmd sam dump pipe call fail %v", err)
		})
		switch verbose.Err() {
		case nil:
			time.After(500 * time.Millisecond)
			continue
		default:
			xEnv.Errorf("cmd sam dump verbose error %v", err)
			return
		}
	}

	cmd.Wait()
}
