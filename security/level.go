package security

import (
	"math/rand"
	"time"

	"github.com/restartfu/ars-shell-crypt/standard"
)

var (
	LevelDefault = New(1)
	LevelWeak    = New(2)
	LevelMedium  = New(3)
	LevelStrong  = New(4)
	// LevelHash est indéchriffrable
	LevelHash = New(5)
)

// Level est une structure de niveau de sécurité
type Level struct {
	Spacing        int
	FirstStrBreak  string
	SecondStrBreak string
}

// New retourne un nouveau Level
func New(level int) Level {
	var spacing int
	var firstStrBreak, secondStrBreak string
	if level <= 4 {
		spacing = level * 10
		firstStrBreak = genSecString(level)
		secondStrBreak = genSecString(level)
	} else {
		spacing = level * 9
	}
	return Level{
		Spacing:        spacing,
		FirstStrBreak:  firstStrBreak,
		SecondStrBreak: secondStrBreak,
	}
}

// genSecString retourne une chaine de caractère aléatoire
func genSecString(level int) string {
	rand.Seed(time.Now().UnixNano())
	var secStr = "`"
	for i := 0; i != level; i++ {
		secStr += string(standard.CharList[rand.Intn(len(standard.CharList)-1)])
	}
	return secStr
}
