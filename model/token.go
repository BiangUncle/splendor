package model

import "math/rand"

const (
	EmeraldTokenNumber   = 7
	DiamondTokenNumber   = 7
	SapphireTokenNumber  = 7
	OnyxTokenNumber      = 7
	RubyTokenNumber      = 7
	GoldJokerTokenNumber = 5
)

type Token struct {
}

func GetNewTokens() []Token {
	return nil
}

func ShuffleToken(tokens []Token) []Token {
	rand.Shuffle(len(tokens), func(i, j int) {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	})
	return tokens
}
