package utility

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const SessionID = "53616c7465645f5f2afba6e806cf441d924af2497993f1e02ff76317f9cc12f567bcf75b04d3c25b182c55bf35faa526a5f6d721c0817db9f8ac985bf1be2221"

func GetAOCProblem(year, day int) string {
	resp, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day), nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp.Header.Add("Cookie", fmt.Sprintf("session=%s", SessionID))
	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func GetAOCInput(year, day int) string {
	resp, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp.Header.Add("Cookie", fmt.Sprintf("session=%s", SessionID))
	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}
