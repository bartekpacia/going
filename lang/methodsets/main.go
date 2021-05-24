package main

type Worker interface {
	DoWork(data []byte) (int, error)
	Identify() string
}

type Human struct {
	FirstName string
	LastName  string
}

func (m Human) DoWork(data []byte) (int, error) {
	return 7, nil
}

func (m Human) Identify() string {
	return m.FirstName + m.LastName
}

type Machine struct {
	ID string
}

func (m Machine) DoWork(data []byte) (int, error) {
	return 7, nil
}

func (m Machine) Identify() string {
	return m.ID
}

func main() {
	var humanWorker Worker
	humanWorker = Human{FirstName: "John", LastName: "Doe"}
	humanWorker.DoWork([]byte{2, 4, 8})
	humanWorker.Identify()

	var machineWorker Worker
	machineWorker = Machine{ID: "25"}
	machineWorker.DoWork([]byte{2, 4, 8})
	machineWorker.Identify()
}
