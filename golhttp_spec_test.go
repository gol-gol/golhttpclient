package golhttpclient

import (
	"io/ioutil"
	"testing"
)

var (
	specPost = `POST /what/is/path?user=me HTTP/1.0
Host: 10.0.0.1:8080
Postman-Token: xxx-post

`

	specGet = `GET /?statuses=read,liked&amp; page=1 HTTP/1.1
Host: 1.1.2.0
cache-control: no-cache

`
	specBody = `{
    "start_date": "2029-04-19T00:00:00Z",
    "end_date": "2039-04-19T00:00:00Z",
    "name": "Alice",
    "friends": [
        "Bob",
        "Eve"
    ]
}`
)

func TestUnmarshalPost(t *testing.T) {
	var reqPost Request
	Unmarshal([]byte(specPost), &reqPost)

	if reqPost.Method != "POST" {
		t.Errorf("FAILED to parse HTTP Method: %s", reqPost.Method)
	}
	if reqPost.Protocol != "HTTP/1.0" {
		t.Errorf("FAILED to parse HTTP Protocol: %s", reqPost.Protocol)
	}
	if reqPost.Path != "/what/is/path" {
		t.Errorf("FAILED to parse HTTP Request Path: %s", reqPost.Path)
	}
	if reqPost.Params["user"] != "me" {
		t.Errorf("FAILED to parse HTTP URL Params: %v", reqPost.Params)
	}
	if reqPost.Headers["Host"] != "10.0.0.1:8080" || reqPost.Headers["Postman-Token"] != "xxx-post" {
		t.Errorf("FAILED to parse HTTP Headers: %v", reqPost.Headers)
	}

	body, _ := ioutil.ReadAll(reqPost.Body)
	if string(body) != "" {
		t.Errorf("FAILED to parse HTTP Body: %v", body)
	}
}

func TestUnmarshalPostWithBody(t *testing.T) {
	var reqPost Request
	Unmarshal([]byte(specPost+specBody), &reqPost)

	if reqPost.Method != "POST" {
		t.Errorf("FAILED to parse HTTP Method: %s", reqPost.Method)
	}
	if reqPost.Protocol != "HTTP/1.0" {
		t.Errorf("FAILED to parse HTTP Protocol: %s", reqPost.Protocol)
	}
	if reqPost.Path != "/what/is/path" {
		t.Errorf("FAILED to parse HTTP Request Path: %s", reqPost.Path)
	}
	if reqPost.Params["user"] != "me" {
		t.Errorf("FAILED to parse HTTP URL Params: %v", reqPost.Params)
	}
	if reqPost.Headers["Host"] != "10.0.0.1:8080" ||
		reqPost.Headers["Postman-Token"] != "xxx-post" {
		t.Errorf("FAILED to parse HTTP Headers: %v", reqPost.Headers)
	}

	bodyBytes, _ := ioutil.ReadAll(reqPost.Body)
	body := string(bodyBytes)
	if body != specBody {
		t.Errorf("FAILED to parse HTTP Body: %s", body)
	}
}

func TestUnmarshalGet(t *testing.T) {
	var reqGet Request
	Unmarshal([]byte(specGet), &reqGet)

	if reqGet.Method != "GET" {
		t.Errorf("FAILED to parse HTTP Method: %s", reqGet.Method)
	}
	if reqGet.Protocol != "HTTP/1.1" {
		t.Errorf("FAILED to parse HTTP Protocol: %s", reqGet.Protocol)
	}
	if reqGet.Path != "/" {
		t.Errorf("FAILED to parse HTTP Request Path: %s", reqGet.Path)
	}
	if reqGet.Params["statuses"] != "read,liked" || reqGet.Params["amp;page"] != "1" {
		t.Errorf("FAILED to parse HTTP URL Params: %v", reqGet.Params)
	}
	if reqGet.Headers["Host"] != "1.1.2.0" || reqGet.Headers["cache-control"] != "no-cache" {
		t.Errorf("FAILED to parse HTTP Headers: %v", reqGet.Headers)
	}

	body, _ := ioutil.ReadAll(reqGet.Body)
	if string(body) != "" {
		t.Errorf("FAILED to parse HTTP Body: %v", body)
	}
}

func TestUnmarshalGetWithBody(t *testing.T) {
	var reqGet Request
	Unmarshal([]byte(specGet+specBody), &reqGet)

	if reqGet.Method != "GET" {
		t.Errorf("FAILED to parse HTTP Method: %s", reqGet.Method)
	}
	if reqGet.Protocol != "HTTP/1.1" {
		t.Errorf("FAILED to parse HTTP Protocol: %s", reqGet.Protocol)
	}
	if reqGet.Path != "/" {
		t.Errorf("FAILED to parse HTTP Request Path: %s", reqGet.Path)
	}
	if reqGet.Params["statuses"] != "read,liked" || reqGet.Params["amp;page"] != "1" {
		t.Errorf("FAILED to parse HTTP URL Params: %v", reqGet.Params)
	}
	if reqGet.Headers["Host"] != "1.1.2.0" || reqGet.Headers["cache-control"] != "no-cache" {
		t.Errorf("FAILED to parse HTTP Headers: %v", reqGet.Headers)
	}

	bodyBytes, _ := ioutil.ReadAll(reqGet.Body)
	body := string(bodyBytes)
	if body != specBody {
		t.Errorf("FAILED to parse HTTP Body: %s", body)
	}
}
