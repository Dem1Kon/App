package src

import "fmt"

func (c *Company) WorkerPage() {
	var action string

	for {
		action = getWorkerAction()

		switch action {
		case "s":
			showAllWorkers(c)
		case "h":
		case "f":
		case "b":
			return
		}
	}
}
func getWorkerAction() (action string) {
	showWorkerAction()

	_, err := fmt.Scan(&action)
	if err != nil {
		panic(err)
	}
	ClearScreen()
	return action
}

func showAllWorkers(c *Company) {
	stateLength := len(c.state)

	for i := 0; i < stateLength; i++ {

	}
}

func showWorkerAction() {
	fmt.Print("\tHere actions you can do(action - shortcut):\n\t Show all workers - s\n\tHire workers - h\n\tFire workers - f\n\tBack - b\n\tEnter shortcut: ")
}
