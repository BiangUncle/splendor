package service

import (
	"splendor/model"
)

func SetUp() error {
	developmentCards := model.GetNewDevelopmentCards()
	nobleTitles := model.GetNewNobleTitles()
	tokens := model.GetNewTokens() // todo: 是一个数组应该

	// todo 打乱

	// todo 每个级别展示4张

	// todo 打乱贵族牌，展示 +1 数量
}
