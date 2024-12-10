package src

import (
	"fmt"
	"os"
	"os/exec"
	time2 "time"
)

//Global functions are in this file

type Viewer interface {
	showStat()
}

func StatShower(obj Viewer) {
	obj.showStat()
}

func ClearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls") // For Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ShowCountDown(time int, CountDown chan int) {
	fmt.Println()
	for time >= 0 {
		fmt.Printf("\r\t%d secons left!", time)
		time2.Sleep(time2.Second)
		time--
	}
	close(CountDown)
}
