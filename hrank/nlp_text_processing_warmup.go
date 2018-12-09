// In the work solution to hackerrank.com AI\NLP\ Text Processing Warmup
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	wordDay := "(([12][0-9]|[3][01]|[1-9])(st|nd|rd|th)?)"
	numDay := "([12][0-9]|[3][01]|[1-9])"
	wordMonth := "(January|Febrauray|March|April|May|June|July|August|September|October|November|December)|(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sept|Oct|Nov|Dec)"
	numMonth := "([1][012]|[1-9])"
	shortYear := "[0-9]{2}"
	longYear := "[0-9]{4}"
	numDate := "[^0-9A-Za-z_](([12][0-9]|[3][01]|[1-9])([-][\\])([1][012]|[1-9]))।(([1][012]|[1-9]))([-][\\]([-][\\])([12][0-9]|[3][01]|[1-9]))([0-9]{2}।[0-9]{4})[^0-9A-Za-z_]"
	longDate := "[^0-9A-Za-z_](([12][0-9]|[3][01]|[1-9])(st|nd|rd|th)?)\\s((January|Febrauray|March|April|May|June|July|August|September|October|November|December)|(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sept|Oct|Nov|Dec))\\s([0-9]{2})|([0-9]{4})[^0-9A-Za-z_]"

	numDateRegex := regexp.MustCompile(numDate)
	longDateRegex := regexp.MustCompile(longDate)
	aRegex := regexp.MustCompile("[^0-9A-Za-z_]a[^0-9A-Za-z_]")
	anRegex := regexp.MustCompile("[^0-9A-Za-z_]and[^0-9A-Za-z_]")
	theRegex := regexp.MustCompile("[^0-9A-Za-z_]the[^0-9A-Za-z_]")
	var numLines int
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d", &numLines)
	for i := 0; i < numLines; i++ {
		line, _ := reader.ReadString('\n')
		// fmt.Println(line)
		fmt.Println(len(aRegex.FindAllString(line, -1)))
		fmt.Println(len(anRegex.FindAllString(line, -1)))
		fmt.Println(len(theRegex.FindAllString(line, -1)))
		fmt.Println(len(numDateRegex.FindAllString(line, -1)) + len(longDateRegex.FindAllString(line, -1)))
	}
}
