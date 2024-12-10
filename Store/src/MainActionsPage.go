package src

import (
	"fmt"
)

//All about players actions is in this file

func mainPage(company *Company, flag *bool) {
	var action string
	for action != "end" {
		action = getMainAction()

		switch action {
		case "c":
			company.CompanyPage()
		case "w":

		case "s":
			StatShower(company.stores)
		case "n":
			return
		case "end":
			*flag = false
		}
	}

}
func showMainActions() {
	fmt.Printf("\tHere actions you can do(action - shortcut):\n\tCompany menu - c\n\tNext day - n\n\tEnd Game - end\n\tEnter shortcut: ")
}

func getMainAction() (action string) {
	showMainActions()
	_, err := fmt.Scan(&action)
	if err != nil {
		panic(err)
	}
	ClearScreen()
	return action
}
