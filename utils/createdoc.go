package utils

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func (task Task) ToHtml() string {
	html := "<div id=\"task\">\n"
	html += "<p style=\"text-align:left;\">"
	html += "<b>" + task.Id + ".</b> - " + fmt.Sprint(task.Points) + " pont\n"

	if len(task.Author) > 0 {
		html += "<span style=\"float:right;\">"

		byperson := true
		for _, name := range strings.Split(task.Author, " ") {
			if unicode.IsLower(rune(name[0])) {
				byperson = false
			}
		}

		if byperson {
			html += "Javasolta: "
		}
		html += task.Author

		html += "</span>\n"
	}

	html += "</p>\n"

	html += task.Description + "\n"

	html += "</div>\n"

	return html
}

func (ps Problemset) ToHtml() string {
	data, err := os.ReadFile("sablon.html")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	html := string(data)

	html += "<h3>" + ps.Chategory + " pontverseny " + strings.ToLower(ps.Month) + "</h3>\n"

	for _, task := range ps.Tasks {
		html += task.ToHtml()
	}
	html += "</body></html>"

	return html
}

func (ps Problemset) SavePdf(html string) {
	// HOST HTML
	ts := httptest.NewServer(htmlHandler(html))
	defer ts.Close()

	// SAVE AS PDF
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuffer []byte
	if err := chromedp.Run(ctx, htmlToPdfTask(ts.URL, &pdfBuffer)); err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("Komal%s%s.pdf", ps.Month, strings.ReplaceAll(ps.Chategory, "/", ""))
	if err := ioutil.WriteFile(filename, pdfBuffer, 0644); err != nil {
		log.Fatal(err)
	}
}

func htmlHandler(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, strings.TrimSpace(content))
	})
}

func htmlToPdfTask(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Sleep(2 * time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithMarginTop(0.5).
				WithMarginLeft(0.5).
				WithMarginRight(0.5).
				WithPrintBackground(true).
				Do(ctx)

			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
