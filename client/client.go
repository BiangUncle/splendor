package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	Address string
	Session sessions.Session
	Cookies []*http.Cookie
}

func ConstructClient() *Client {
	c := &Client{
		client:  &http.Client{},
		Address: "127.0.0.1:8765",
	}
	return c
}

func (c *Client) ConstructURL(uri string, args map[string]any) string {
	url := "http://" + c.Address + "/"

	if uri != "" {
		url = url + uri
	}

	if len(args) == 0 {
		return url
	}

	url = url + "?"

	for k, v := range args {
		url = url + fmt.Sprintf("%+v=%+v?", k, v)
	}

	url = url[:len(url)-1]

	return url
}

func (c *Client) ConstructRequest(url string, cookies []*http.Cookie) (*http.Request, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}

	return req, nil
}

func (c *Client) Send(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ExtractBodyContent(resp *http.Response) (string, error) {
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (c *Client) SendRequest(uri string, args map[string]any) (*http.Response, error) {
	url := c.ConstructURL(uri, args)
	req, err := c.ConstructRequest(url, c.Cookies)
	if err != nil {
		return nil, err
	}
	return c.Send(req)
}

func CheckRespStatusCode(resp *http.Response) (int, string, error) {
	content, err := ExtractBodyContent(resp)
	if err != nil {
		return 0, "", err
	}
	msg := gjson.Get(content, "msg").String()
	return resp.StatusCode, msg, nil
}
