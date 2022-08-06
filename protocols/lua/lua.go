package lua

import lua "github.com/yuin/gopher-lua"

type Lua struct {
	state *lua.LState
}

func newLua() *Lua {
	l := &Lua{}
	l.state = lua.NewState()
	return l
}

func (l *Lua) Load(script string) error {
	return l.state.DoString(script)
}

func (l *Lua) Close() error {
	l.state.Close()
	return nil
}
