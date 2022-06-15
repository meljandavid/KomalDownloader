package main

import (
	"fmt"
	"meljandavid/komaldownloader/frontend"
	"meljandavid/komaldownloader/utils"
	"strconv"
)

func main() {
	// INPUT - TODO?: gui
	em, pw, m, set := frontend.GetCreds()

	// LOGIN & RETRIEVE PAGE
	pagehtml := utils.RetrieveHtml(m, em, pw)

	// PROCESS PAGE
	for _, cht := range set {
		mm, _ := strconv.Atoi(m)
		ps := utils.Problemset{Month: utils.Months[mm%100], Tasks: []utils.Task{}, Chategory: cht}
		ps.MakeProblemset(pagehtml)

		// HTML FROM THE TEMPLATE
		html := ps.ToHtml()

		// Convert to PDF
		ps.SavePdf(html)

		fmt.Println(cht + " feladatok lementve!")
	}
}
