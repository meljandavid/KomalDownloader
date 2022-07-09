package utils

var Months = [13]string{"", "Januar", "Februar", "Marcius", "Aprilis", "Majus",
	"Junius", "Juilius", "Augusztus", "Szeptember", "Oktober", "November", "December"}

var ChategoryToSubject = map[string]string{"K": "mat", "C": "mat", "B": "mat", "A": "mat",
	"M": "fiz", "G": "fiz", "P": "fiz", "I": "inf", "I/S": "inf", "S": "inf"}

type Task struct {
	Id          string
	Description string
	Points      int
	Author      string
}

type Problemset struct {
	Month     string
	Chategory string
	Tasks     []Task
}
