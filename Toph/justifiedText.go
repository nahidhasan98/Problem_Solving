package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var s string

	//reading a line with spaces
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}

	//removing extra spaces from begining & ending
	s = strings.TrimSpace(s)

	//replacing all consecutive spaces with a single space
	var reg = regexp.MustCompile(`[ ]{2,}`)
	s = reg.ReplaceAllString(s, " ")

	//replacing a long word (longer than 8 characters) with "#bigword"
	list := strings.Split(s, " ")
	s = ""
	for i := 0; i < len(list); i++ {
		if len(list[i]) > 8 {
			list[i] = "#bigword"
		}
		s += list[i]

		if i != len(list)-1 {
			s += " "
		}
	}

	//dividing into lines
	var idx = 0
	for i := 0; i < len(s); i++ {
		var start = idx
		var end = min(idx+8, len(s)-1)

		for j := end; j >= start; j-- {
			if s[j] == ' ' {
				s = s[:j] + "\n" + s[j+1:]
				idx = j + 1
				break
			}
		}
	}

	fmt.Printf("%s\n", s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
