package cmds

import (
	"math/rand"
)

func fortunate() string {
	fortune := [7]string{"大凶", "凶", "末吉", "小吉", "中吉", "大吉", "仙草吉"}
	randomfortune := fortune[rand.Intn(6)]
	return randomfortune
}
