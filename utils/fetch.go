package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func RetrieveHtml(month, em, pw string, set []string) map[string]string {
	logindata := url.Values{
		"a":     {"login"},
		"email": {em},
		"pwd":   {pw},
	}
	response, err := http.PostForm("https://www.komal.hu/u", logindata)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	bodybytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	if strings.Contains(string(bodybytes), "beállításaim") {
		fmt.Println("Sikeres bejelentkezes")
	} else {
		fmt.Println("Sikertelen bejelentkezes")
	}

	subjects := map[string]bool{}
	for _, cht := range set {
		subjects[ChategoryToSubject[cht]] = true
	}

	client := &http.Client{}

	pagehtmls := map[string]string{}
	for subject := range subjects {
		psurl := fmt.Sprintf("https://www.komal.hu/feladat?a=honap&h=%d%s&t=%s&l=hu",
			time.Now().Year(), month, subject)

		req, _ := http.NewRequest("GET", psurl, nil)
		for _, cookie := range response.Cookies() {
			req.AddCookie(cookie)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		pagehtml := string(data)

		pagehtmls[subject] = pagehtml
	}

	return pagehtmls
}
