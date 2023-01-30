package main

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"splendor/utils"
	"time"
)

type GameStatus struct {
	ConnectStatus string
	SessionID     string
	TableID       string
	PlayerID      string
	UserName      string
	*Client
	*GameCron
}

func (g *GameStatus) Info() string {
	ret := fmt.Sprintf("[%+v]", g.UserName)
	ret += fmt.Sprintf("状态: %+v; ", g.ConnectStatus)
	ret += fmt.Sprintf("会话: %+v; ", utils.CompressUuid(g.SessionID))
	ret += fmt.Sprintf("房间: %+v; ", utils.CompressUuid(g.TableID))
	ret += fmt.Sprintf("玩家: %+v;\n ", utils.CompressUuid(g.PlayerID))
	return ret
}

func (g *GameStatus) AskWhichTurn() (string, error) {
	resp, err := g.SendRequest("cur_player", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := g.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) JoinGame() (string, error) {
	resp, err := g.SendRequest("join", map[string]any{
		"username": "biang",
	})

	g.Cookies = resp.Cookies()

	content, err := g.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) Alive() (string, error) {
	resp, err := g.SendRequest("alive", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := g.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) Leave() (string, error) {
	resp, err := g.SendRequest("leave", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := g.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) TableInfo() (string, error) {
	resp, err := g.SendRequest("table_info", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := g.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) NextTurn() (string, error) {
	resp, err := g.SendRequest("next_turn", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := g.ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) KeepALive() (string, error) {
	resp, err := g.SendRequest("keep_a_live", map[string]any{})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("错误码: %+v", resp.StatusCode))
	}
	return "ok", nil
}

func (g *GameStatus) IfMyTurn() (bool, error) {
	content, err := g.AskWhichTurn()
	if err != nil {
		return false, err
	}

	curPlayerID := gjson.Get(content, "current_player_id").String()
	if curPlayerID != g.PlayerID {
		time.Sleep(time.Second)
		return false, nil
	}
	return true, nil
}
