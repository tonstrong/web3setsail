package test

type Demo struct {
	Name string
	Age  int
}

var demo Demo = Demo{}

func Student() Demo {
	return demo
}

func SetStudent(p Demo) {
	demo.Name = p.Name
	demo.Age = p.Age
}
