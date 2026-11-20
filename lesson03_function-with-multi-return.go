package main

import (
	"strings"
)

func getInitials(s string) (string, string) {
	newStr := strings.ToUpper(s)
	names := strings.Split(newStr, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
