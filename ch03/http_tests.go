package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	if res, err := http.Get("http://google.com/robots.txt"); err == nil {
		defer res.Body.Close()
		log.Println(res)
	} else {
		log.Println(err)
	}

	if res, err  := http.Head("http://google.com/robots.txt"); err == nil {
		defer res.Body.Close()
		log.Println(res)
	} else {
		log.Println(err)
	}

	form := url.Values{}

	form.Add("foo", "bar")

	if res, err := http.Post("http://google.com/robots.txt",
		"application/x-www-form-urlencoded", strings.NewReader(form.Encode())); err == nil {
			defer res.Body.Close()
			log.Println(res)
	} else {
		log.Println(err)
	}

	// or use FormPost
	if res, err := http.PostForm("http://google.com/robots.txt", form); err == nil {
		defer res.Body.Close()
		log.Println(res)
	} else {
		log.Println(err)
	}
}
