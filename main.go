package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "https://translate.googleapis.com/translate_a/single?client=gtx&dt=t&sl=%s&tl=%s&q="

var sourceLang string
var targetLang string

func init() {
	var trFlag bool
	flag.BoolVar(&trFlag, "tr", false, "Translate from Turkish to English")

	flag.Parse()

	if trFlag {
		sourceLang = "tr"
		targetLang = "en"
	} else {
		sourceLang = "en"
		targetLang = "tr"
	}
}

func translate(text string) ([]string, error) {
	query := url.QueryEscape(text)
	url := fmt.Sprintf(baseURL, sourceLang, targetLang) + query
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body error: %v", err)
	}

	var result []interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %v", err)
	}

	var translations []string
	if len(result) > 0 {
		sentences, ok := result[0].([]interface{})
		if ok {
			for _, sentence := range sentences {
				firstSentence, ok := sentence.([]interface{})
				if ok && len(firstSentence) > 0 {
					translatedText, ok := firstSentence[0].(string)
					if ok {
						translations = append(translations, translatedText)
					}
				}
			}
		}
	}

	if len(translations) == 0 {
		return nil, fmt.Errorf("translation not found")
	}

	return translations, nil
}

func main() {
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide text to translate.")
		return
	}

	text := strings.Join(flag.Args(), " ")
	translatedTexts, err := translate(text)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, translatedText := range translatedTexts {
		fmt.Printf("- %s\n", translatedText)
	}
}
