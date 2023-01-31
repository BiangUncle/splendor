package main

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"splendor/utils"
	"sync"
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

func ConstructGameStatus(c *Client) *GameStatus {
	g := &GameStatus{
		Client: c,
		GameCron: &GameCron{
			stop: make(chan struct{}),
			wg:   &sync.WaitGroup{},
		},
	}

	return g
}

func (g *GameStatus) Info() []string {
	return []string{
		fmt.Sprintf("玩家[%+v] ", g.UserName),
		fmt.Sprintf("状态[%+v] ", g.ConnectStatus),
		fmt.Sprintf("会话[%+v] ", utils.CompressUuid(g.SessionID)),
		fmt.Sprintf("房间[%+v] ", utils.CompressUuid(g.TableID)),
		fmt.Sprintf("玩家[%+v] ", utils.CompressUuid(g.PlayerID)),
	}
}

func (g *GameStatus) ShowPlayerInfo() {
	fmt.Println("============================================================")
	infoRow := ""
	for _, info := range g.Info() {
		infoRow += info
	}
	fmt.Println(infoRow)
	fmt.Println("============================================================")
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

// IsOurTurn 检查是否为自己的回合，异常打印
func (g *GameStatus) IsOurTurn() bool {
	yes, err := g.IfMyTurn()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !yes {
		fmt.Println("不是你的回合")
		return false
	}
	return true
}

func (g *GameStatus) IfMyTurn() (bool, error) {
	content, err := g.AskWhichTurn()
	if err != nil {
		return false, err
	}
	fmt.Println(content)

	curPlayerID := gjson.Get(content, "current_player_id").String()
	if curPlayerID != g.PlayerID {
		return false, nil
	}
	return true, nil
}

func (g *GameStatus) TakeThreeTokens(tokensString string) (string, error) {
	_, err := g.SendRequest("take_three_tokens", map[string]any{
		"tokens": tokensString,
	})
	if err != nil {
		return "", err
	}
	return "ok", nil
}

func (g *GameStatus) TakeDoubleTokens(tokenId int) (string, error) {
	_, err := g.SendRequest("take_double_tokens", map[string]any{
		"token_id": tokenId,
	})
	if err != nil {
		return "", err
	}
	return "ok", nil
}
