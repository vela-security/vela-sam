package sam

import (
	"fmt"
	"github.com/vela-security/vela-public/lua"
)

type user struct {
	name string
	id   string
	nt   string
	lm   string
}

func (u user) String() string                         { return fmt.Sprintf("%p", &u) }
func (u user) Type() lua.LValueType                   { return lua.LTObject }
func (u user) AssertFloat64() (float64, bool)         { return 0, false }
func (u user) AssertString() (string, bool)           { return "", false }
func (u user) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (u user) Peek() lua.LValue                       { return u }

func (u user) isEmpty() bool {
	return u.name == ""
}

func (u user) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "name":
		return lua.S2L(u.name)
	case "id":
		return lua.S2L(u.id)
	case "nt_hash":
		return lua.S2L(u.nt)
	case "lm_hash":
		return lua.S2L(u.lm)
	default:
		return lua.LNil
	}

}