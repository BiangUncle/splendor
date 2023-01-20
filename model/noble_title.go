package model

import "math/rand"

const NobleTitleNumber = 10

type NobleTitle struct {
}

func GetNewNobleTitles() []NobleTitle {
	return nil
}

func ShuffleNobleTitle(nobleTitles []NobleTitle) []NobleTitle {
	rand.Shuffle(len(nobleTitles), func(i, j int) {
		nobleTitles[i], nobleTitles[j] = nobleTitles[j], nobleTitles[i]
	})
	return nobleTitles
}
