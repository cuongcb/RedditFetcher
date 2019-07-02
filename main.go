package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	fmt.Println("Hello World.")
	reg, err := regexp.Compile("https?://imgur.com/a......")
	handleError(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.reddit.com/r/programming/comments/c7nu7u.json?sort=best&limit=10&depth=10", nil)
	req.Header.Set("User-Agent", "linux:go-postgrabber:v0.1")
	handleError(err)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()

		buf := new(bytes.Buffer)

		buf.ReadFrom(resp.Body)

		s := buf.String() // Does a complete copy of the bytes in the buffer.

		links := reg.FindString(stripchars(s, "\""))

		fmt.Printf("%v\n", links)

		fmt.Printf("%v\n", s)
	}

	fmt.Println(resp.Status)
	resp.Body.Close()
}
