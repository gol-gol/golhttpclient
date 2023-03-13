package golhttpclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	SERVER          = "127.0.0.1:18765"
	KEY_SERVER_ADDR = "serverAddr"
	RESPONSE_BODY   = "Hello, HTTP!\n"
)

func startServer(listenAt string) *http.Server {
	mux := http.NewServeMux()
	getHello := func(w http.ResponseWriter, r *http.Request) {
		reply := fmt.Sprintf("%s %s", r.Method, RESPONSE_BODY)
		io.WriteString(w, reply)
	}
	mux.HandleFunc("/hello", getHello)

	server := &http.Server{
		Addr:    listenAt,
		Handler: mux,
	}
	go server.ListenAndServe()
	time.Sleep(5 * time.Millisecond)
	return server
}

func TestFetch(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	resp, errResp := req.Fetch()
	if errResp != nil {
		t.Errorf("FAILED for making Fetch call\nError: %s", errResp.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	expected := fmt.Sprintf("GET %s", RESPONSE_BODY)
	if err != nil || string(bodyBytes) != expected {
		msg := fmt.Sprintf("FAILED for making Fetch GET call\nError: %v\nBody: %s", err, string(bodyBytes))
		t.Error(msg)
	}
	server.Close()
}

func TestGetBytes(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	expected := fmt.Sprintf("GET %s", RESPONSE_BODY)
	if bodyBytes, err := req.GetBytes(); err != nil || string(bodyBytes) != expected {
		msg := fmt.Sprintf("FAILED for making GET Bytes call\nError: %v\nBody: %s", err, string(bodyBytes))
		t.Error(msg)
	}
	server.Close()
}

func TestGet(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	expected := fmt.Sprintf("GET %s", RESPONSE_BODY)
	if body, err := req.Get(); err != nil || body != expected {
		msg := fmt.Sprintf("FAILED for making GET call\nError: %v\nBody: %s", err, body)
		t.Error(msg)
	}
	server.Close()
}

func TestPut(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	expected := fmt.Sprintf("PUT %s", RESPONSE_BODY)
	if body, err := req.Put(); err != nil || body != expected {
		msg := fmt.Sprintf("FAILED for making PUT call\nError: %v\nBody: %s", err, body)
		t.Error(msg)
	}
	server.Close()
}

func TestPost(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	expected := fmt.Sprintf("POST %s", RESPONSE_BODY)
	if body, err := req.Post(); err != nil || body != expected {
		msg := fmt.Sprintf("FAILED for making POST call\nError: %v\nBody: %s", err, body)
		t.Error(msg)
	}
	server.Close()
}

func TestPatch(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	expected := fmt.Sprintf("PATCH %s", RESPONSE_BODY)
	if body, err := req.Patch(); err != nil || body != expected {
		msg := fmt.Sprintf("FAILED for making PATCH call\nError: %v\nBody: %s", err, body)
		t.Error(msg)
	}
	server.Close()
}

func TestDelete(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	expected := fmt.Sprintf("DELETE %s", RESPONSE_BODY)
	if body, err := req.Delete(); err != nil || body != expected {
		msg := fmt.Sprintf("FAILED for making DELETE call\nError: %v\nBody: %s", err, body)
		t.Error(msg)
	}
	server.Close()
}

func TestHead(t *testing.T) {
	server := startServer(SERVER)
	req := Request{Url: fmt.Sprintf("http://%s/hello", SERVER)}
	resp, errResp := req.Head()
	if errResp != nil {
		t.Errorf("FAILED for making Fetch call\nError: %s", errResp.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(bodyBytes) != "" {
		msg := fmt.Sprintf("FAILED for making HEAD call\nError: %v\nBody: %s", err, string(bodyBytes))
		t.Error(msg)
	}
	server.Close()
}
