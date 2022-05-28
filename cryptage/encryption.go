package cryptage

import (
	"github.com/restartfu/ars-shell-crypt/security"
	"github.com/restartfu/ars-shell-crypt/standard"
	"strings"
)

// Encrypt encrypte le string donné selon le niveau de sécurité donné.
func Encrypt(str string, security security.Level) string {
	newStr := ""
	for _, i := range str {
		newStr += string(encryptChar(i, security.Spacing))
	}
	return newStr + security.FirstStrBreak + newStr[:1] + security.SecondStrBreak
}

// Decrypt decrypte le string donné selon le niveau de sécurité donné.
func Decrypt(str string, security security.Level) string {
	newStr := ""
	hash := strings.Split(str, "`")[0]
	for _, char := range hash {
		newStr += string(decryptChar(char, security.Spacing))
	}
	return newStr
}

// encryptChar encrypte le caractère donné selon l'espacement donné.
func encryptChar(char rune, spacing int) rune {
	var newChar rune
	for i, ch := range standard.CharList {
		if ch == char {
			if i >= spacing {
				for true {
					if i+spacing <= len(standard.CharList)-1 {
						newChar = standard.CharList[i+spacing]
						break
					} else if (i+spacing)-len(standard.CharList) <= len(standard.CharList)-1 {
						newChar = standard.CharList[(i+spacing)-len(standard.CharList)]
						break
					}
					break
				}
			} else if i < spacing {
				for true {
					if i+spacing <= len(standard.CharList)-1 {
						newChar = standard.CharList[i+spacing]
						break
					}
					break
				}
			}
		}
	}
	return newChar
}

// decryptChar décrypte le caractère donné selon l'espacement donné.
func decryptChar(char rune, spacing int) rune {
	var newChar rune
	for i, ch := range standard.CharList {
		if ch == char {
			if i >= spacing {
				for true {
					if i-spacing <= len(standard.CharList)-1 {
						newChar = standard.CharList[i-spacing]
						break
					}
					break
				}
			} else if i < spacing {
				for true {
					if i-spacing <= len(standard.CharList)-1 {
						newChar = standard.CharList[i-spacing]
						break
					} else if (i-spacing)+len(standard.CharList) <= len(standard.CharList)-1 {
						newChar = standard.CharList[(i-spacing)+len(standard.CharList)]
						break
					}
					break
				}
			}
		}
	}
	return newChar
}
