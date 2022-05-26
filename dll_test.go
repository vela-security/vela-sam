package sam

import (
	"syscall"
	"testing"
)

func Test_lsadump(t *testing.T) {
	lib := syscall.NewLazyDLL("d:\\github.com\\vela-dev\\share\\software\\lsadump.dll")

	main := lib.NewProc("GetSamAccounts")

	_, _, err := main.Call()
	t.Log(err)
}
