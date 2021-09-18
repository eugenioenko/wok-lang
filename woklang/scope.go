package woklang

// Scope holds the structure of the environment of the runtime data state
type Scope struct {
	values map[string]WokData
	parent *Scope
}

func NewScope(parent *Scope) *Scope {
	values := make(map[string]WokData)
	scope := &Scope{parent: parent, values: values}
	return scope
}

func (scope *Scope) Get(key string) (WokData, bool) {
	value, ok := scope.values[key]
	if ok {
		return value, true
	}
	if scope.parent != nil {
		return scope.parent.Get(key)
	}
	return nil, false
}

func (scope *Scope) Has(key string) bool {
	_, ok := scope.values[key]
	return ok
}

func (scope *Scope) Set(key string, value WokData) {
	scope.values[key] = value
}
