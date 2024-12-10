package src

//All about company is in this file

import (
	"fmt"
	"time"
)

type Company struct {
	name       string
	state      []*Worker
	stores     *Buildings
	money, age int
	finance    *Finance
}

func CreateCompany() *Company {
	newCompany := new(Company)
	newCompany.getName()
	newCompany.money = 100000
	newCompany.stores = new(Buildings)
	newCompany.finance = new(Finance)
	return newCompany
}

func (c *Company) showStat() {
	fmt.Printf("\tCompany name: %s, age: %d days\n\tMoney: %d\n\tState: %d workers\n", c.name, c.age, c.money, len(c.state))

	countDown := make(chan int)
	go ShowCountDown(3, countDown)
	<-countDown

	ClearScreen()
}

func (c *Company) hireWorkers() { //need to upgrade(too long)
	offer := checkOffers()
	var decision string

	for {
		fmt.Println("\tYou have a new offer for your company:")
		offer.showStat()
		decision = getDecision()
		if decision != "+" && decision != "-" {
			ClearScreen()
			fmt.Println("\tIncorrect input!")
			time.Sleep(1 * time.Second)
			ClearScreen()
		} else {
			break
		}
	}

	ClearScreen()

	switch decision {
	case "+":
		c.state = append(c.state, offer)

	}
}

func (c *Company) fireWorkers() {
	if len(c.state) == 0 {
		fmt.Println("\tYou don't have any workers now!")
		time.Sleep(3 * time.Second)
		ClearScreen()
		return
	}
	worker := selectWorker(c)

	var newState []*Worker

	for i := 0; i < len(c.state); i++ {
		if worker != c.state[i] {
			newState = append(newState, c.state[i])
		}
	}
	c.state = newState
}

func (c *Company) ShowWorkers() {
	stateLength := len(c.state)

	for i := 0; i < stateLength; i++ {
		fmt.Printf("\tWorker %d:\n", i+1)
		StatShower(c.state[i])
	}
	if len(c.state) == 0 {
		fmt.Println("\tNobody works in your company!")
	}
}

func (c *Company) working() {
	c.age++
}

func selectWorker(c *Company) *Worker {
	for i := 0; i < len(c.state); i++ {
		StatShower(c.state[i])
		decision := getDecision()
		ClearScreen()
		if decision == "+" {
			return c.state[i]
		}
	}
	fmt.Println("\tThere is no more workers!")
	time.Sleep(2 * time.Second)
	ClearScreen()
	return nil
}

func getDecision() string {
	var decision string

	fmt.Print("\tDo you want to choice this worker(enter '+' or '-'): ")
	_, err := fmt.Scan(&decision)
	if err != nil {
		panic(err)
	}

	return decision
}

func (c *Company) getName() {
	fmt.Print("\tEnter name of your company: ")
	_, err := fmt.Scan(&c.name)

	if err != nil {
		panic(err)
	}
	ClearScreen()
}

func checkOffers() *Worker {
	return generateWorker()
}
