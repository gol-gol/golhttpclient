package golhttpclient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultHttpMethod = "GET"
)

type Request struct {
	Protocol      string
	Method        string
	Url           string
	Path          string
	Params        map[string]string
	Headers       map[string]string
	Body          *bytes.Buffer
	SkipSSLVerify bool
}

func (req *Request) Fetch() (*http.Response, error) {
	if req.Method == "" {
		req.Method = DefaultHttpMethod
	}
	return req.fetch()
}

func (req *Request) GetBytes() (body []byte, err error) {
	req.Method = "GET"
	body, err = req.fetchBytes()
	return
}

func (req *Request) Get() (body string, err error) {
	req.Method = "GET"
	body, err = req.fetchBody()
	return
}

func (req *Request) Put() (body string, err error) {
	req.Method = "PUT"
	body, err = req.fetchBody()
	return
}

func (req *Request) Post() (body string, err error) {
	req.Method = "POST"
	body, err = req.fetchBody()
	return
}

func (req *Request) Patch() (body string, err error) {
	req.Method = "PATCH"
	body, err = req.fetchBody()
	return
}

func (req *Request) Delete() (body string, err error) {
	req.Method = "DELETE"
	body, err = req.fetchBody()
	return
}

func (req *Request) Head() (*http.Response, error) {
	req.Method = "HEAD"
	return req.fetch()
}

func (req *Request) Options() (*http.Response, error) {
	req.Method = "OPTIONS"
	return req.fetch()
}

func (req *Request) fetch() (resp *http.Response, err error) {
	httpClient := &http.Client{
		Transport: customRoundTripper(req.SkipSSLVerify),
	}

	var r *http.Request
	if req.Body == nil {
		r, err = http.NewRequest(req.Method, "", nil)
	} else {
		r, err = http.NewRequest(req.Method, "", req.Body)
	}
	if err != nil {
		return
	}
	r.URL = req.getURL()
	req.setHttpHeaders(r)

	resp, err = httpClient.Do(r)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (req *Request) fetchBytes() (body []byte, err error) {
	resp, err := req.fetch()
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func (req *Request) fetchBody() (body string, err error) {
	bodyText, err := req.fetchBytes()
	if err != nil {
		return
	}
	body = string(bodyText)
	return
}

func customRoundTripper(skipVerify bool) (customTransport http.RoundTripper) {
	customTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: skipVerify},
		ExpectContinueTimeout: 1 * time.Second,
	}
	return
}

func (req *Request) getURL() (uri *url.URL) {
	var getParamsURI string
	var _val string
	for key, val := range req.Params {
		_val = url.QueryEscape(val)
		if getParamsURI == "" {
			getParamsURI = fmt.Sprintf("%s=%s", key, _val)
		} else {
			getParamsURI = fmt.Sprintf("%s&%s=%s", getParamsURI, key, _val)
		}
	}
	requestUrl := fmt.Sprintf("%s?%s", req.Url, getParamsURI)
	uri, err := url.Parse(requestUrl)

	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (req *Request) setHttpHeaders(r *http.Request) (err error) {
	basicAuth := strings.Split(req.Headers["basicAuth"], ":")
	if len(basicAuth) > 1 {
		apiUsername, apiPassword := basicAuth[0], strings.Join(basicAuth[1:], ":")
		r.SetBasicAuth(apiUsername, apiPassword)
	}
	for header, value := range req.Headers {
		r.Header.Add(header, value)
	}
	return
}
