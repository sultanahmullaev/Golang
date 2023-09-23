package programmer

type Programmer struct {
	position string
	salary   int
	address  string
}

func (temp *Programmer) GetPosition() string {
	return temp.position
}

func (temp *Programmer) GetSalary() int {
	return temp.salary
}

func (temp *Programmer) GetAddress() string {
	return temp.address
}

func (temp *Programmer) SetPosition(position string) {
	temp.position = position
}

func (temp *Programmer) SetSalary(salary int) {
	temp.salary = salary
}

func (temp *Programmer) SetAddress(address string) {
	temp.address = address
}
