package httpinterface

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Option func(*http.Request)

func DoHttpReq(method, uri string, opts ...Option) ([]byte, error) {
	req, err := http.NewRequest(strings.ToUpper(method), uri, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	for _, v := range opts {
		v(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func AddGetParams(param map[string]string) Option {
	return func(req *http.Request) {
		q := req.URL.Query()
		for k, v := range param {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func AddJsonPostParams(param interface{}) Option {
	return func(req *http.Request) {
		data, err := json.Marshal(param)
		if err == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Body = io.NopCloser(bytes.NewReader(data))
		} else {
			fmt.Printf("err:%v", err)
		}
	}
}

func AddFormPostParams(param map[string]string) Option {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		reqBody := strings.Builder{}
		for k, v := range param {
			str := fmt.Sprintf("%s=%v&", k, v)
			reqBody.WriteString(str) //(k + "=" + v + "&")
		}
		req.Body = io.NopCloser(bytes.NewReader([]byte(reqBody.String())))
	}
}

func AddCookies(cookies []*http.Cookie) Option {
	return func(r *http.Request) {
		if len(cookies) > 0 {
			for _, cookie := range cookies {
				r.AddCookie(cookie)
			}
		}
	}
}

func AddHeaders(headers map[string]string) Option {
	return func(r *http.Request) {
		for k, v := range headers {
			r.Header.Add(k, v)
		}
	}
}

// https://echo.free.beeceptor.com?x=y
func Do() {
	type Post struct {
		Name  string `json:"name"`
		Ctime string `json:"starttm"`
		Age   int32
		Say   int8 `json:"-"`
		Flag  bool
	}

	byteArr, err := DoHttpReq(
		"post", "https://postman-echo.com/post",
		AddGetParams(map[string]string{"x": "y"}),
		// AddFormPostParams(map[string]string{"name": "z3", "starttm": "2020-02-2", "ag+e": "18"}),
		AddJsonPostParams(Post{Name: "z3", Ctime: "2020-02-2", Age: 18}),
		AddCookies([]*http.Cookie{{Name: "a", Value: "1"}, {Name: "b", Value: "2"}}),
		AddHeaders(map[string]string{"c": "3", "d": "4"}),
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("res = %v", string(byteArr))
	}
}

type Xhttp struct {
	Method     string
	Url        string
	QueryParam map[string]string
	Param      map[string]interface{}
	Cookies    []*http.Cookie
	Headers    map[string]string
}

// xhttp:=new(Xhttp)
// xhttp.AddCookie()
// xhttp.AddHeaders()
// xhttp.post()

// type Hook interface{
// 	func before(req) bool {

// 	}

// 	func after(resp) bool {

// 	}
// }
