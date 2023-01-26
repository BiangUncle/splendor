package main

import (
	"fmt"
	"net/http"
)

func (c *Client) SendRequest(uri string, args map[string]any) (*http.Response, error) {
	url := c.ConstructURL(uri, args)
	req, err := c.ConstructRequest(url, c.Cookies)
	if err != nil {
		return nil, err
	}
	return c.Send(req)
}

func (c *Client) JoinGame() (string, error) {
	resp, err := c.SendRequest("join", map[string]any{
		"username": "biang",
	})

	c.Cookies = resp.Cookies()

	content, err := c.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	fmt.Println(content)

	return "ok", nil
}

func (c *Client) Alive() (string, error) {
	resp, err := c.SendRequest("alive", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := c.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (c *Client) Leave() (string, error) {
	resp, err := c.SendRequest("leave", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := c.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}
