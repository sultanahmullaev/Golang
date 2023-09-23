package main

import (
	"Assignment1/manager"
	"Assignment1/programmer"
	"Assignment1/tester"
	"fmt"
)

func main() {
	var manager1 = manager.Manager{}
	manager1.SetPosition("Manager")
	manager1.SetSalary(4200)
	manager1.SetAddress("St. George, bld. 36, apt. 86")
	fmt.Println(manager1.GetPosition())
	fmt.Println(manager1.GetSalary())
	fmt.Println(manager1.GetAddress())

	var tester1 = tester.Tester{}
	tester1.SetPosition("Tester")
	tester1.SetSalary(3500)
	tester1.SetAddress("St. George, bld. 36, apt. 86")
	fmt.Println(tester1.GetAddress())

	var programmer1 = programmer.Programmer{}
	programmer1.SetPosition("Main")
	programmer1.SetSalary(500000)
	programmer1.SetAddress("St. George, bld. 36, apt. 86")
	fmt.Println(programmer1.GetSalary())
}
