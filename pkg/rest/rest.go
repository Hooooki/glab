/* MIT License

Copyright (c) 2023 Fragan Gourvil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE */

package rest

import (
	"fmt"
	"gitlab-environment/pkg/config"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// GitlabUri : Gitlab base URI
const (
	GitlabUri string = "https://gitlab.com/api/v4/projects"
)

type Methods string

// HttpMethods : Available HTTP Methods
const (
	GET    Methods = "GET"
	POST   Methods = "POST"
	DELETE Methods = "DELETE"
	PUT    Methods = "PUT"
)

// Request : Lite declaration of a HTTP Request
type Request struct {
	Headers  map[string]string
	Endpoint string
	Config   config.Config
	Method   Methods
	Body     io.Reader
}

type Response struct {
	Body       []byte
	StatusCode int
}

// init : Generate a new HttpRequest using the net/http package
func (r Request) init() (*http.Request, error) {

	var endpoint = r.Endpoint

	if strings.HasPrefix(endpoint, "/") {
		endpoint = strings.TrimPrefix(endpoint, "/")
	}

	req, err := http.NewRequest(string(r.Method), fmt.Sprintf("%v/%v/%v",
		GitlabUri,
		r.Config.Context.CurrentProject.Id,
		endpoint,
	), r.Body)

	if nil != err {
		return req, err
	}

	for header, value := range r.Headers {
		req.Header.Add(header, value)
	}

	req.Header.Add("PRIVATE-TOKEN", r.Config.Context.CurrentProject.Token)
	return req, err

}

// Send : Send a new HTTP request based on the Request data
func (r Request) Send() *Response {
	client := &http.Client{}

	if r.Method == POST || r.Method == PUT {

		if nil == r.Headers {
			header := make(map[string]string)
			header["Content-Type"] = "application/x-www-form-urlencoded"
			r.Headers = header
		}

	}

	req, err := r.init()

	if err != nil {
		log.Fatalln(fmt.Sprintf("Error on creating Request %v", err.Error()))
		os.Exit(1)
	}

	resp, err := client.Do(req)
	// Defer allow us to execute a function at the end of the method
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("Error when closing Body")
		}
	}(resp.Body)

	b, err := io.ReadAll(resp.Body)

	if nil != err {
		panic(err)
	}

	return &Response{
		Body:       b,
		StatusCode: resp.StatusCode,
	}
}
