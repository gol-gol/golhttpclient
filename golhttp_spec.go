package golhttpclient

import (
	"bufio"
	"bytes"
	"strings"
)

const (
	initSpec    = 0
	methodSpec  = 1
	headersSpec = 2
	bodySpec    = 3
)

func Unmarshal(spec []byte, req *Request) {
	state := 0
	withNewLine := false
	var body bytes.Buffer
	scanner := bufio.NewScanner(bytes.NewReader(spec))
	req.Params = map[string]string{}
	req.Headers = map[string]string{}

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			if state == initSpec {
				state = methodSpec
			} else if state == headersSpec {
				state = bodySpec
			}
			continue
		} else if state <= methodSpec {
			unmarshalMethodUrlAndParams(req, txt)
			state = headersSpec
		} else if state == headersSpec {
			unmarshalHeader(req, txt)
		} else if state == bodySpec {
			if withNewLine {
				body.WriteString("\n")
			} else {
				withNewLine = true
			}
			body.WriteString(txt)
		}
	}
	req.Body = &body
}

func unmarshalMethodUrlAndParams(req *Request, txt string) {
	txtParts := strings.Split(txt, " ")
	lastIndex := len(txtParts) - 1
	req.Method = txtParts[0]
	req.Protocol = txtParts[lastIndex]

	_url := strings.Join(txtParts[1:lastIndex], "")
	_url = strings.Replace(_url, "&amp; ", "&", -1)
	_urlParts := strings.Split(_url, "?")
	req.Path = _urlParts[0]
	if len(_urlParts) == 1 {
		return
	}
	for _, param := range strings.Split(_urlParts[1], "&") {
		getKeyVal := strings.Split(param, "=")
		if len(getKeyVal) == 1 {
			req.Params[getKeyVal[0]] = ""
			continue
		}
		req.Params[getKeyVal[0]] = strings.Join(getKeyVal[1:len(getKeyVal)], "=")
	}
}

func unmarshalHeader(req *Request, txt string) {
	txtParts := strings.Split(txt, ":")
	headerName := strings.TrimSpace(txtParts[0])
	headerValue := strings.TrimSpace(
		strings.Join(txtParts[1:len(txtParts)], ":"),
	)
	req.Headers[headerName] = headerValue
}
