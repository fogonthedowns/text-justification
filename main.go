package main

import "fmt"

func main() {
	in := []string{"This", "is", "an", "example", "of", "text", "justification."}
	out := fullJustify(in, 16)
	fmt.Println(out)
}

func fullJustify(words []string, maxWidth int) []string {
	var numChars int
	var numSpaces int
	line := make([]string, 0)
	ans := make([]string, 0)
	for _, word := range words {
		// makes a new line if the characters + spaces plus the new word is
		// greater than maxWidth
		if numChars+numSpaces+len(word) > maxWidth {
			numSpaces = maxWidth - numChars
			ans = append(ans, justify(line, maxWidth, numSpaces, false))
			numChars = 0
			numSpaces = 0
			line = make([]string, 0)
		}
		numChars += len(word)
		numSpaces++
		line = append(line, word)
	}
	numSpaces = maxWidth - numChars
	ans = append(ans, justify(line, maxWidth, numSpaces, true))
	return ans
}

func justify(words []string, maxWidth int, numSpaces int, isLastLine bool) string {
	s := ""
	// one space between words on the last line
	// long trailing space after last word
	if isLastLine || len(words) == 1 {
		for _, word := range words {
			s += word
			if len(s) < maxWidth {
				s += " "
			}
		}
		for len(s) < maxWidth {
			s += " "
		}
		return s
	}
	numSlots := len(words) - 1
	spaces := make([]string, len(words))
	spaces[0] = ""
	for i := 0; i < numSpaces%numSlots; i++ {
		spaces[i+1] += " "
	}
	for i := 0; i < numSpaces/numSlots; i++ {
		for j := 1; j < len(spaces); j++ {
			spaces[j] += " "
		}
	}
	for i, word := range words {
		s += spaces[i]
		s += word
	}
	return s
}
