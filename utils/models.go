package utils

var Months = [13]string{"", "Januar", "Februar", "Marcius", "Aprilis", "Majus",
	"Junius", "Juilius", "Augusztus", "Szeptember", "Oktober", "November", "December"}

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
