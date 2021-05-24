package main

import "fmt"

type Worker interface {
	DoWork(data []byte) (int, error)
	String() string
}

type Human struct {
	FirstName string
	LastName  string
}

func (m Human) DoWork(data []byte) (int, error) {
	return 7, nil
}

func (h Human) String() string {
	return h.FirstName + " " + h.LastName
}

type Machine struct {
	ID string
}

func (m Machine) DoWork(data []byte) (int, error) {
	return 7, nil
}

func (m Machine) String() string {
	return m.ID
}

func main() {
	var humanWorker Worker
	humanWorker = Human{FirstName: "John", LastName: "Doe"}
	humanWorker.DoWork([]byte{2, 4, 8})

	var machineWorker Worker
	machineWorker = Machine{ID: "25"}
	machineWorker.DoWork([]byte{2, 4, 8})

	var stringer1 fmt.Stringer = humanWorker
	var stringer2 fmt.Stringer = machineWorker
	fmt.Println(stringer1, stringer2)
}
