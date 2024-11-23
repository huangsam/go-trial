package basic

type Person struct {
	Name string
	age  int32
}

func (p Person) GetAge() int32 {
	return p.age
}

func (p *Person) SetAge(newAge int32) {
	p.age = newAge
}

func (p Person) IsOlderThan(target int32) bool {
	return p.age >= target
}
