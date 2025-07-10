package person

type Person struct {
	Name string
	Age int
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) IsUnderAge() bool {
	if p.Age < 18 {
		return true
	}
	return false
}