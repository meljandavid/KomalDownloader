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
	pages := utils.RetrieveHtml(m, em, pw, set)

	// PROCESS PAGE
	for _, cht := range set {
		mm, _ := strconv.Atoi(m)
		ps := utils.Problemset{Month: utils.Months[mm], Tasks: []utils.Task{}, Chategory: cht}
		ps.MakeProblemset(pages[utils.ChategoryToSubject[cht]])

		// HTML FROM THE TEMPLATE
		html := ps.ToHtml()

		// Convert to PDF
		ps.SavePdf(html)

		fmt.Println(cht + " feladatok lementve!")
	}
}
