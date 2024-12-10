package src

import "fmt"

type Finance struct {
	allExpenses, allEarns, salaryExpenses, storesExpenses int
	totalIncome                                           string
}

func (f *Finance) showStat() {
	fmt.Printf("\tTotal earnings: %d\n\tExpenses:\tSalary: %d \tRent: %d\n\tTotal Expenses: %d\n\tTotal income(loss): %s", f.allEarns, f.salaryExpenses, f.storesExpenses, f.allExpenses, getTotalIncome(f))

	CountDown := make(chan int)
	go ShowCountDown(3, CountDown)
	<-CountDown
	ClearScreen()
}

func getTotalIncome(f *Finance) (totalIncome string) {
	total := f.allEarns - f.allExpenses
	if total < 0 {
		totalIncome = fmt.Sprintf("(%d)", total)
	} else {
		totalIncome = fmt.Sprintf("%d", total)
	}
	return totalIncome
}
