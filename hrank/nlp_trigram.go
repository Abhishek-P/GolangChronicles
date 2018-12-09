//Working solution to https://www.hackerrank.com/challenges/the-trigram/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	wordRegexp := regexp.MustCompile("[^0-9A-Za-z_]?[a-zA-z]+?[^0-9A-Za-z_]")
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	s := readLine(reader)
	var tokens []string
	// fmt.Println(s)
	seenAlready := false
	for !(seenAlready == true && s == "") {
		if s != "" {
			newTokens := wordRegexp.FindAllString(s, -1)
			tokens = append(tokens, newTokens...)
			seenAlready = false
		} else {
			seenAlready = true
		}
		s = readLine(reader)
	}

	sentences := make(map[int][]string)
	prev := -1 // because when the first sentence is to be split previous is -1
	tokens2 := tokens
	count := 0
	for i, a := range tokens {
		if a[len(a)-1] == []byte(".")[0] {
			token := a[:len(a)-1]
			tokens2[i] = token
			sentences[count] = tokens2[prev+1 : i+1]
			count += 1
			prev = i
		}
	}
	// for i,a := range tokens2 {
	//     fmt.Printn(i,a)
	// }
	// fmt.Printf("Len of Sentences %d",len(sentences))

	trigramCounts := make(map[string]int)

	highest := ""
	// fmt.Printf("count: %d, len(sentences): %d", count, len(sentences))
	for i := 0; i < len(sentences); i++ {
		sentence := sentences[i]
		// fmt.Println(sentence)
		for j := 0; j < len(sentence)-2; j++ {
			trigram := trimAndLower(sentence[j]) + " " + trimAndLower(sentence[j+1]) + " " + trimAndLower(sentence[j+2])

			_, exists := trigramCounts[trigram]

			if exists == false {
				trigramCounts[trigram] = 0
			}

			trigramCounts[trigram] += 1
			// fmt.Printf("highest yet %s: %d\n",highest, trigramCounts[highest])
			// fmt.Printf("trigram: %s: %d\n",trigram, trigramCounts[trigram])

			if trigramCounts[trigram] > trigramCounts[highest] {
				highest = trigram
			}
		}
	}

	// fmt.Println(tokens)
	// fmt.Println(trigramCounts)
	fmt.Println(highest)
}

func trimAndLower(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
