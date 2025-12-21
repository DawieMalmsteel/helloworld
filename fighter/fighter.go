// Package fighter defines fighter
package fighter

var Goku = Fighter{
	name: "Goku",
	age:  24,
}

type Fighter struct {
	name string
	age  int
}

func (f *Fighter) GetName() string {
	return f.name
}

func (f *Fighter) GetAge() int {
	return f.age
}

func (f *Fighter) ChangeName(val string) {
	f.name = val
}

func (f *Fighter) ChangeAge(val int) {
	f.age = val
}
