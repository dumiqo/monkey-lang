package object

func NewEnvironment() *Enviroment {
	s := make(map[string]Object)
	return &Enviroment{store: s, outer: nil}
}

func NewEncosedEnvironment(out *Enviroment) *Enviroment {
	env := NewEnvironment()
	env.outer = out
	return env
}

type Enviroment struct {
	store map[string]Object
	outer *Enviroment
}

func (e *Enviroment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Enviroment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
