package main

import (
	"fmt"
	"net/http"
	"splendor/utils"
)

type GameStatus struct {
	ConnectStatus string
	SessionID     string
	TableID       string
	PlayerID      string
	UserName      string
	*Client
}

func (g *GameStatus) Info() string {
	ret := fmt.Sprintf("[%+v]", g.UserName)
	ret += fmt.Sprintf("状态: %+v; ", g.ConnectStatus)
	ret += fmt.Sprintf("会话: %+v; ", utils.CompressUuid(g.SessionID))
	ret += fmt.Sprintf("房间: %+v; ", utils.CompressUuid(g.TableID))
	ret += fmt.Sprintf("玩家: %+v;\n ", utils.CompressUuid(g.PlayerID))
	return ret
}

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

	return content, nil
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

func (c *Client) TableInfo() (string, error) {
	resp, err := c.SendRequest("table_info", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := c.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}
