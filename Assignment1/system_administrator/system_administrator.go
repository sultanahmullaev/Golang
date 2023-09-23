package system_administrator

type SystemAdministrator struct {
	position string
	salary   int
	address  string
}

func (temp *SystemAdministrator) GetPosition() string {
	return temp.position
}

func (temp *SystemAdministrator) GetSalary() int {
	return temp.salary
}

func (temp *SystemAdministrator) GetAddress() string {
	return temp.address
}

func (temp *SystemAdministrator) SetPosition(position string) {
	temp.position = position
}

func (temp *SystemAdministrator) SetSalary(salary int) {
	temp.salary = salary
}

func (temp *SystemAdministrator) SetAddress(address string) {
	temp.address = address
}
