package team_leader

type TeamLeader struct {
	position string
	salary   int
	address  string
}

func (temp *TeamLeader) GetPosition() string {
	return temp.position
}

func (temp *TeamLeader) GetSalary() int {
	return temp.salary
}

func (temp *TeamLeader) GetAddress() string {
	return temp.address
}

func (temp *TeamLeader) SetPosition(position string) {
	temp.position = position
}

func (temp *TeamLeader) SetSalary(salary int) {
	temp.salary = salary
}

func (temp *TeamLeader) SetAddress(address string) {
	temp.address = address
}
