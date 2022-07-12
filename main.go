package main

import (
	"bufio"
	"context"
	"fmt"
	"meljandavid/komaldownloader/frontend"
	"meljandavid/komaldownloader/utils"
	"os"
	"strconv"

	"github.com/chromedp/chromedp"
)

func main() {
	// INPUT - TODO?: gui
	em, pw, m, set := frontend.GetCreds()

	// LOGIN & RETRIEVE PAGES
	pages := utils.RetrieveHtml(m, em, pw, set)
	mm, _ := strconv.Atoi(m)

	// PROCESS PAGES
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	for _, cht := range set {
		ps := utils.Problemset{Month: utils.Months[mm], Tasks: []utils.Task{}, Chategory: cht}
		ps.MakeProblemset(pages[utils.ChategoryToSubject[cht]])

		// HTML FROM THE TEMPLATE
		html := ps.ToHtml()

		// Convert to PDF
		ps.SavePdf(html, ctx)

		fmt.Println(cht + " feladatok lementve")
	}

	// PAUSE
	fmt.Println("\nA kért pontversenyek feladatsorai sikeresen le lettek mentve!")
	fmt.Println("(nyomjon ENTER-t a kilépéshez vagy zárja be a konzolt)")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
