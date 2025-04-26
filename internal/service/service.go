package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

func CheckMorse(s string) string {
	arr := strings.Split(s, " ")
	for _, l := range arr {
		for _, c := range l {
			if c != '.' && c != '-' {
				return morse.ToMorse(s)
			}
		}
	}
	return morse.ToText(s)
}
