package game_designer

type GameDesigner struct {
	position string
	salary   int
	address  string
}

func (temp *GameDesigner) GetPosition() string {
	return temp.position
}

func (temp *GameDesigner) GetSalary() int {
	return temp.salary
}

func (temp *GameDesigner) GetAddress() string {
	return temp.address
}

func (temp *GameDesigner) SetPosition(position string) {
	temp.position = position
}

func (temp *GameDesigner) SetSalary(salary int) {
	temp.salary = salary
}

func (temp *GameDesigner) SetAddress(address string) {
	temp.address = address
}
