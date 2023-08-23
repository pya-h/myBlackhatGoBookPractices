package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func ExtractResponse(res *http.Response) {
	// response Body is of type ReadCloser
	// u need to close it after using
	defer res.Body.Close()
	log.Println(res.Status)
	log.Println(res.StatusCode)
	if body, err := ioutil.ReadAll(res.Body); err == nil {
		log.Println(string(body))
	} else {
		log.Panicln(err)
	}
	fmt.Println("-----------------------------------------------------------------------------------------------------")
}

func main() {
	if res, err := http.Get("http://google.com/robots.txt"); err == nil {
		defer res.Body.Close()
		ExtractResponse(res)
	} else {
		log.Println("GET:", err)
	}

	if res, err := http.Head("http://google.com/robots.txt"); err == nil {
		// response Body is of type ReadCloser
		// u need to close it after using
		defer res.Body.Close()
		ExtractResponse(res)
	} else {
		log.Println("HEAD:", err)
	}

	form := url.Values{}

	form.Add("foo", "bar")

	if res, err := http.Post("http://google.com/robots.txt",
		"application/x-www-form-urlencoded", strings.NewReader(form.Encode())); err == nil {
		// response Body is of type ReadCloser
		// u need to close it after using
		defer res.Body.Close()
		ExtractResponse(res)
	} else {
		log.Println("POST:", err)
	}

	// or use FormPost
	if res, err := http.PostForm("http://google.com/robots.txt", form); err == nil {
		// response Body is of type ReadCloser
		// u need to close it after using
		defer res.Body.Close()
		ExtractResponse(res)
	} else {
		log.Println("PostForm: ", err)
	}

	var client http.Client

	if req, err := http.NewRequest("DELETE", "http://google.com/robots.txt", nil); err == nil {
		if res, err := client.Do(req); err == nil {
			ExtractResponse(res)
		} else {
			log.Println("DELETE: ", err)

		}
	}

	form.Add("Another", "Foo")

	if req, err := http.NewRequest("PUT", "http://google.com/robots.txt", strings.NewReader(form.Encode())); err == nil {
		if res, err := client.Do(req); err == nil {

			ExtractResponse(res)
		} else {
			log.Println("PUT: ", err)

		}
	}
}
