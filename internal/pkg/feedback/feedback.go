package feedback

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Result ...
type Result struct {
	FeedbackTotal    int
	FeedbackPositive int
	FeedbackNegative int
	FeedbackNeutral  int
}

// Report ...
type Report map[string]Result

func (report Report) storeResult(model string, result Result) {
	report[model] = result
}

// ShowRating ...
func (report Report) ShowRating(model string) {
	ratingFound := false

	for m, r := range report {
		if m == model {
			fmt.Printf("Total Ratings:%v\tPositive:%v\tNegative:%v\tNeutral:%v", r.FeedbackTotal, r.FeedbackPositive, r.FeedbackNegative, r.FeedbackNeutral)
			ratingFound = true
		}
	}

	if !ratingFound {
		fmt.Printf("No rating for this vehicle")
	}
}

// Values ...
type Values struct {
	Models []Model `json:"values"`
}

// Model ...
type Model struct {
	Name     string   `json:"model"`
	Feedback []string `json:"feedback"`
}

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

func readJSONFile(feedFilePath string) Values {

	jsonFile, err := os.Open(feedFilePath)

	if err != nil {
		log.Fatal(err)
		//log.Fatal("File not found")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var content Values
	json.Unmarshal(byteValue, &content)

	return content
}

// ProcessRatingFeed ...
func ProcessRatingFeed(feedFilePath string) Report {
	f := readJSONFile(feedFilePath)
	report := make(Report)
	for _, v := range f.Models {
		var vehResult Result
		var vehRating rating

		for _, msg := range v.Feedback {
			if text := strings.Split(msg, ""); len(text) >= 5 {
				vehRating = 5.0
				vehResult.FeedbackTotal++

				for _, word := range text {
					switch s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r"); s {
					case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
						vehRating += extraPositive
					case "help", "helpful", "thanks", "thank you", "happy":
						vehRating += positive
					case "not helpful", "sad", "angry", "improve", "annoy":
						vehRating += negative
					case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
						vehRating += extraNegative
					}
				}

				switch {
				case vehRating > 8.0:
					vehResult.FeedbackPositive++
				case vehRating >= 4.0 && vehRating <= 8.0:
					vehResult.FeedbackNeutral++
				case vehRating < 4.0:
					vehResult.FeedbackNegative++
				}
			}
		}
		report.storeResult(v.Name, vehResult)
		//report[v.Name] = vehResult
	}

	return report
}
