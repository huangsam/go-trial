package basic

type Person struct {
	Name string
	age  int32
}

func GetAge(p Person) int32 {
	return p.age
}

func SetAge(p *Person, newAge int32) {
	p.age = newAge
}
