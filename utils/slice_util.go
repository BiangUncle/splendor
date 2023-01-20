package utils

import "math/rand"

func ShuffleSlice(s []interface{}) []interface{} {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}
