package model

import (
	"errors"
	"fmt"
	"splendor/utils"
	"time"
)

const HandCardUpperBound = 3
const TokensNumberUpperLimit = 10

var GlobalPlayer = make(map[string]*Player)
var PlayerLastLoginTime = make(map[string]time.Time)

type Player struct {
	Name             string               `json:"name"`              // 玩家名字
	Tokens           TokenStack           `json:"tokens"`            // 宝石列表
	Bonuses          TokenStack           `json:"bonuses"`           // 奖励列表
	DevelopmentCards DevelopmentCardStack `json:"development_cards"` // 发展卡列表
	HandCards        DevelopmentCardStack `json:"hand_cards"`        // 手中的发展卡
	NobleTitles      NobleTilesStack      `json:"noble_titles"`      // 贵族
	Prestige         int                  `json:"prestige"`          // 声望
	PlayerID         string               `json:"player_id"`         // 玩家ID
}

// CreatePlayer 创建一个玩家
func CreatePlayer() *Player {
	return &Player{
		Tokens:           CreateEmptyTokenStack(),
		Bonuses:          CreateEmptyTokenStack(),
		DevelopmentCards: make(DevelopmentCardStack, 0),
		HandCards:        make(DevelopmentCardStack, 0),
		NobleTitles:      make(NobleTilesStack, 0),
		Prestige:         0,
	}
}

// CreatePlayerID 生成一个玩家ID
func CreatePlayerID() string {
	return utils.GetUuidV4()
}

// JoinNewPlayer 在全局对象里面增加一个玩家
func JoinNewPlayer() (*Player, string, error) {

	player := CreatePlayer()
	playerID := CreatePlayerID()
	player.PlayerID = playerID

	if _, ok := GlobalPlayer[playerID]; ok {
		return nil, "", errors.New(fmt.Sprintf("这个玩家已经已经有了"))
	}
	GlobalPlayer[playerID] = player
	PlayerLastLoginTime[playerID] = time.Now()

	utils.SystemPrintf("玩家: %s 登陆游戏\n", player.Name)

	return player, playerID, nil
}

// KeepALive 保持活度
func KeepALive(playerID string) {
	PlayerLastLoginTime[playerID] = time.Now()
}

// CheckIfOnline 确认是否在线
func CheckIfOnline(playerID string) bool {
	if lastTime, ok := PlayerLastLoginTime[playerID]; ok {
		now := time.Now()
		// 大于 10s 就是没激活
		if now.Sub(lastTime) > 10*time.Second {
			return false
		}
		return true
	}
	return false
}

// CheckAllPlayerNetStatus 检查所有玩家的在线状态
func CheckAllPlayerNetStatus() []*Player {
	var offlinePlayerIDs []*Player

	for playerID, player := range GlobalPlayer {
		if !CheckIfOnline(playerID) {
			fmt.Printf("玩家 %+v 已离线\n", player.Name)
			offlinePlayerIDs = append(offlinePlayerIDs, player)
		}
	}

	return offlinePlayerIDs
}

// ClearOfflinePlayer 清理离线用户
func ClearOfflinePlayer() {
	// 检查
	offlinePlayers := CheckAllPlayerNetStatus()
	// 清理
	for _, player := range offlinePlayers {
		fmt.Printf("清理离线用户: %+v\n", player.Name)
		DeleteGlobalPlayer(player.PlayerID)
	}
}

// DeleteGlobalPlayer 删除玩家全局信息
func DeleteGlobalPlayer(playerID string) {
	delete(GlobalPlayer, playerID)
	delete(PlayerLastLoginTime, playerID)
}

// GetGlobalPlayer 根据ID获取玩家对象
func GetGlobalPlayer(playerID string) (*Player, error) {
	if player, ok := GlobalPlayer[playerID]; ok {
		return player, nil
	}
	return nil, errors.New(fmt.Sprintf("没有这个ID的玩家, playerID = %+v", playerID))
}

// AddTokens 玩家获取宝石
func (p *Player) AddTokens(tokens TokenStack) {
	p.Tokens.Add(tokens)
	return
}

// AddDevelopmentCard 玩家获取发展卡
func (p *Player) AddDevelopmentCard(card *DevelopmentCard) {
	// 发展卡增加这个
	p.DevelopmentCards = append(p.DevelopmentCards, card)
	p.Prestige += card.Prestige
	p.Bonuses[card.BonusType]++

	return
}

// AddNobleTile 玩家招待贵族
func (p *Player) AddNobleTile(noble *NobleTile) {
	// 发展卡增加这个
	p.NobleTitles = append(p.NobleTitles, noble)
	p.Prestige += noble.Prestige

	return
}

// HasEnoughToken 判断是否足够宝石
func (p *Player) HasEnoughToken(tokens TokenStack) bool {
	// 使用的黄金数量
	remainGoldJoker := p.Tokens[TokenIdxGoldJoker]

	for idx, tokensNum := range p.Tokens {
		// 不判断黄金
		if idx == TokenIdxGoldJoker {
			continue
		}
		// 判断现金加奖励够不够购买
		if tokensNum+p.Bonuses[idx] >= tokens[idx] {
			continue
		}
		// 判断加上黄金够不够，足够需要扣除黄金
		if tokensNum+p.Bonuses[idx]+remainGoldJoker >= tokens[idx] {
			remainGoldJoker -= tokens[idx] - (tokensNum + p.Bonuses[idx])
			continue
		}
		return false
	}

	return true
}

// PayToken 支付宝石
func (p *Player) PayToken(tokens TokenStack) (TokenStack, error) {
	// 使用的黄金数量
	remainGoldJoker := p.Tokens[TokenIdxGoldJoker]

	returnToken := CreateEmptyTokenStack()

	for idx, tokensNum := range p.Tokens {
		// 不判断黄金
		if idx == TokenIdxGoldJoker {
			continue
		}
		// 判断现金加奖励够不够购买
		if tokensNum+p.Bonuses[idx] >= tokens[idx] {
			needPay := tokens[idx] - p.Bonuses[idx] // 计算需要买多少
			returnToken[idx] += needPay
			continue
		}
		// 判断加上黄金够不够，足够需要扣除黄金
		if tokensNum+p.Bonuses[idx]+remainGoldJoker >= tokens[idx] {
			remainGoldJoker -= tokens[idx] - (tokensNum + p.Bonuses[idx])
			returnToken[idx] += tokensNum
			continue
		}
		return nil, errors.New(fmt.Sprintf("支付错误，根本不够啊，需要 %d, 只有 %d, 剩余黄金 %d。", tokens[idx], tokensNum+p.Bonuses[idx], remainGoldJoker))
	}

	// 计算需要返还的黄金
	returnToken[TokenIdxGoldJoker] = p.Tokens[TokenIdxGoldJoker] - remainGoldJoker

	// 角色扣除宝石
	err := p.Tokens.Minus(returnToken)
	if err != nil {
		return nil, err
	}

	return returnToken, nil
}

// AddHandCard 玩家获取手牌
func (p *Player) AddHandCard(card *DevelopmentCard) error {
	if len(p.HandCards) >= HandCardUpperBound {
		return errors.New(fmt.Sprintf("不能再增加手牌了，目前已经有 %d 张。", len(p.HandCards)))
	}
	// 发展卡增加这个
	p.HandCards = append(p.HandCards, card)
	return nil
}

// RemoveHandCard 移除手牌
func (p *Player) RemoveHandCard(cardIdx int) (*DevelopmentCard, error) {

	selectedIdx := -1
	var newHandCards []*DevelopmentCard

	for idx, card := range p.HandCards {
		if card.Idx == cardIdx {
			selectedIdx = idx
		} else {
			newHandCards = append(newHandCards, card)
		}
	}

	if selectedIdx == -1 {
		return nil, errors.New(fmt.Sprintf("移除手牌失败，cardIdx = %d, 现有手牌 %+v", cardIdx, p.HandCards))
	}

	ret := p.HandCards[selectedIdx]
	p.HandCards = newHandCards

	return ret, nil
}

// ReceiveNoble 招待贵族，如果不可以招待，返回 false
func (p *Player) ReceiveNoble(noble *NobleTile) (bool, error) {

	// 计算当前发展卡是否足够招待
	//existTokens := p.DevelopmentCards.ToTokenStack()
	existTokens := p.Bonuses // bonuses 和发展卡一个数量
	more := existTokens.MoreThan(noble.Acquires)
	if !more {
		return false, nil
	}

	// 将当前贵族加入自己名下
	p.AddNobleTile(noble)

	return true, nil
}

// TakeOutTokens 从手牌拿出宝石
func (p *Player) TakeOutTokens(t TokenStack) (TokenStack, error) {
	err := p.Tokens.Minus(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
