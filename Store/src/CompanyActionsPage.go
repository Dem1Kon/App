package src

import (
	"fmt"
)

func (c *Company) CompanyPage() {
	var action string

	for {
		action = getCompanyAction()

		switch action {
		case "c":
			StatShower(c)
		case "w":
			c.WorkerPage()
		case "fin":
			StatShower(c.finance)
		case "b":
			return
		}
	}

}

func getCompanyAction() (action string) {
	showCompanyAction()
	_, err := fmt.Scan(&action)
	if err != nil {
		panic(err)
	}
	ClearScreen()
	return action
}

func showCompanyAction() {
	fmt.Print("\tHere actions you can do(action - shortcut):\n\tShow company card - c\n\tWorkers - w\n\tShow finances - fin\n\tBack - b\n\tEnter shortcut: ")
}
