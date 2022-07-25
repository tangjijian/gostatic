package struction

type Student struct {
	name string
}

func (s *Student) Get() string {
	return s.name
}
func (s *Student) Set(name string) {
	s.name = name
}
