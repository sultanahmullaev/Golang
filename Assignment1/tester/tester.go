package tester

type Tester struct {
	position string
	salary   int
	address  string
}

func (temp *Tester) GetPosition() string {
	return temp.position
}

func (temp *Tester) GetSalary() int {
	return temp.salary
}

func (temp *Tester) GetAddress() string {
	return temp.address
}

func (temp *Tester) SetPosition(position string) {
	temp.position = position
}

func (temp *Tester) SetSalary(salary int) {
	temp.salary = salary
}

func (temp *Tester) SetAddress(address string) {
	temp.address = address
}
