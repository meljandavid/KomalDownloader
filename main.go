package main

import (
	"bufio"
	"fmt"
	"meljandavid/komaldownloader/frontend"
	"meljandavid/komaldownloader/utils"
	"os"
	"strconv"
)

func main() {
	// INPUT - TODO?: gui
	em, pw, m, set := frontend.GetCreds()

	// LOGIN & RETRIEVE PAGES
	pages := utils.RetrieveHtml(m, em, pw, set)

	// PROCESS PAGES
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

	// PAUSE
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
