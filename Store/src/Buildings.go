package src

import (
	"fmt"
	"time"
)

type Buildings struct {
	stores []*Store
}

func (b *Buildings) showStat() {
	storesAmount := len(b.stores)

	if storesAmount == 0 {
		fmt.Print("\tYou don't have any stores yet!")
		time.Sleep(2 * time.Second)
		ClearScreen()
	}

	for i := 0; i < storesAmount; i++ {
		b.stores[i].showStat()
		CountDown := make(chan int)
		go ShowCountDown(2, CountDown)
		<-CountDown
		ClearScreen()
	}

}
