package basicintro

// Person represents a person with a name and age.
type Person struct {
	Name string
	age  int32
}

// GetAge returns the age of the person.
func (p Person) GetAge() int32 {
	return p.age
}

// SetAge sets the age of the person.
func (p *Person) SetAge(newAge int32) {
	p.age = newAge
}

// IsOlderThan returns true if the person's age is greater than or equal to the target age.
func (p Person) IsOlderThan(target int32) bool {
	return p.age >= target
}
