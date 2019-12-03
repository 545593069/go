package main


type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{id: id, name: name}
}

func (s *student) Id() int64 {
	return s.id
}

func (s *student) SetId(id int64) {
	s.id = id
}

func (s *student) Name() string {
	return s.name
}

func (s *student) SetName(name string) {
	s.name = name
}