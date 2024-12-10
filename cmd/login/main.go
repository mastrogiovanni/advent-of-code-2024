package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/mastrogiovanni/advent-of-code-2024/src/utility"
)

func main() {
	problem := utility.GetAOCProblem(2024, 1)
	problem = strings.ReplaceAll(problem, "\n", "")
	log.Println(problem)
	var validID = regexp.MustCompile(`<pre><code>(.*?)</code>`)
	for _, item := range validID.FindAll([]byte(problem), -1) {
		log.Println(item)
	}
}
