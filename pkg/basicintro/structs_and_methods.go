package basicintro

// Person represents a person with a name and age.
type Person struct {
	Name string
	Age  int
	Address
}

type Address struct {
	Street string
	City   string
	Zip    string
}

// IsOlderThan returns true if the person's age is greater than or equal to the target age.
func (p Person) IsOlderThan(target int) bool {
	return p.Age > target
}
