package model

import "math/rand"

const DevelopmentCardNumber = 90

type DevelopmentCard struct {
}

func GetNewDevelopmentCards() []DevelopmentCard {
	return nil
}

func ShuffleDevelopmentCard(developmentCards []DevelopmentCard) []DevelopmentCard {
	rand.Shuffle(len(developmentCards), func(i, j int) {
		developmentCards[i], developmentCards[j] = developmentCards[j], developmentCards[i]
	})
	return developmentCards
}
