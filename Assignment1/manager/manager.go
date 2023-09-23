package manager

type Manager struct {
	Position string
	salary   int
	address  string
}

func (temp *Manager) GetPosition() string {
	return temp.Position
}

func (temp *Manager) GetSalary() int {
	return temp.salary
}

func (temp *Manager) GetAddress() string {
	return temp.address
}

func (temp *Manager) SetPosition(position string) {
	temp.Position = position
}

func (temp *Manager) SetSalary(salary int) {
	temp.salary = salary
}

func (temp *Manager) SetAddress(address string) {
	temp.address = address
}
