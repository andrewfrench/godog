package godog

type ScenarioState map[interface{}]interface{}

// Set saves the value 'v' at key 'k'
func (s *ScenarioState) Set(k, v interface{}) {
	(*s)[k] = v
}

// Get returns the value at key 'k'. The returned value
// may be nil.
func (s *ScenarioState) Get(k interface{}) interface{} {
	return (*s)[k]
}
