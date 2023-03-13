
## golhttpclient

> few helper/wrapper functions for HTTP Client
>
> extracted from an older pandora box of such packages at [abhishekkr/gol](https://github.com/abhishekkr/gol)

### Public Functions

* `(req *Request) Fetch() (*http.Response, error)`
* `(req *Request) GetBytes() (body []byte, err error)`
* `(req *Request) Get() (body string, err error)`
* `(req *Request) Put() (body string, err error)`
* `(req *Request) Post() (body string, err error)`
* `(req *Request) Patch() (body string, err error)`
* `(req *Request) Delete() (body string, err error)`
* `(req *Request) Head() (*http.Response, error)`
* `(req *Request) Options() (*http.Response, error)`

* `Unmarshal(spec []byte, req *Request)` to Unmarshal HTTP/1.1 Message Spec text as Bytes into golhttpclient.Request

---
