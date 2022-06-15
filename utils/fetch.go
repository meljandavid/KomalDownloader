package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func RetrieveHtml(chategory, em, pw string) string {
	psurl := fmt.Sprintf("https://www.komal.hu/feladat?a=honap&h=%d%s&t=mat&l=hu",
		time.Now().Year(), chategory)

	logindata := url.Values{
		"a":     {"login"},
		"email": {em},
		"pwd":   {pw},
	}
	resp, err := http.PostForm("https://www.komal.hu/u", logindata)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	bodybytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	if strings.Contains(string(bodybytes), "beállításaim") {
		fmt.Println("Sikeres bejelentkezes")
	} else {
		fmt.Println("Sikertelen bejelentkezes")
	}

	req, _ := http.NewRequest("GET", psurl, nil)
	for _, cookie := range resp.Cookies() {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	pagehtml := string(data)

	return pagehtml
}
