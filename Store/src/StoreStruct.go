package src

import "fmt"

type StoreWorker interface {
	earnMoney() int
	getInfo()
}

type Store struct {
	name                                            string
	lastDayMoney, allTimeMoney, rent, workersAmount int
}

func (s *Store) showStat() {
	fmt.Printf("\tStore name: %s, Earned: \n\tLast day: %d\n\tAll time: %d\n", s.name, s.lastDayMoney, s.allTimeMoney)

	ClearScreen()
}
