package src

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generateWorker() *Worker {
	w := new(Worker)

	w.generateName()
	w.generateExp()
	w.generateAge()
	w.power = 100
	w.generateSalary()

	return w
}

type Worker struct {
	name, status                   string
	age, salary, power, experience int
}

func (w *Worker) Work() {
	if w.power >= 20 {
		w.power -= 20
		w.experience += rand.Intn(50)
	} else {
		fmt.Print("Worker got sick!")
		w.power -= 100
	}
}

func (w *Worker) Relax() {
	w.power += 40
}

func (w *Worker) showStat() {
	fmt.Printf("\tMain information: %s, %d years\t\tSalary per month: %d\t\tPower: %d\tExperience: %d\n", w.name, w.age, w.salary, w.power, w.experience)
}

func (w *Worker) generateName() {
	w.name = randomFromFile("Names.txt")
}

func (w *Worker) generateAge() {
	w.age = rand.Intn(30) + 17
}

func (w *Worker) generateExp() {
	w.experience = rand.Intn(100)
}

func (w *Worker) generateSalary() {
	w.salary = (rand.Intn(40) + 10) * 1000
}

func randomFromFile(fileName string) string {
	file, err := os.OpenFile("Names.txt", os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	randomInt := (rand.Intn(31))
	rand.Seed(time.Now().UnixNano())
	counter := 0
	for scan.Scan() {
		if counter == randomInt {
			return scan.Text()
		}
		counter++
	}
	return "Dominik"
}
