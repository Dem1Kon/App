package src

//There is the game process file

func StartGame() {
	playerCompany := CreateCompany()
	flag := true
	for flag {
		mainPage(playerCompany, &flag)
	}

}
