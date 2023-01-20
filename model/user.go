package model

type User struct {
	Tokens           TokenStack        // 宝石列表
	Bonuses          TokenStack        // 奖励列表
	DevelopmentCards []DevelopmentCard // 发展卡列表
	NobleTitles      []NobleTile       // 贵族
	Prestige         int               // 声望
}

func (u *User) AddTokens(tokens TokenStack) {
	u.Tokens.Add(tokens)
	return
}
