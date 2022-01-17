package main

import (
	"math/rand"
	"strings"
	"time"
)

const allCharacter = "ABCDEFFGIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func generate(isCorrect bool) string {
	var output strings.Builder
	rand.Seed(time.Now().Unix())
	strLength := rand.Intn(20) //  max length of random string
	output.Reset()
	for i := 0; i < strLength; i++ {
		if isCorrect {
			randomChar := rand.Intn(len(allCharacter))
			output.WriteString(string(allCharacter[randomChar]))
		} else {
			output.WriteString(string(allCharacter[i]))
		}

	}

	return output.String()
}
