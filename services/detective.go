package services

import (
	"encoding/json"
	"net/http"
	"os"
)

type DetectiveService struct {
	BaseURL string
}

func NewDetectiveService() *DetectiveService {
	return &DetectiveService{BaseURL: os.Getenv("DETECTIVE_URL")}
}

type Breach struct {
	HasPassword bool     `json:"has_password"`
	Password    string   `json:"password"`
	SHA1        string   `json:"sha1"`
	Sources     []string `json:"sources"`
}

type Response struct {
	Success bool     `json:"success"`
	Found   int64    `json:"found"`
	Result  []Breach `json:"result"`
}

func (detectiveService *DetectiveService) Investigate(key string) ([]Breach, error) {
	var response Response
	req, err := http.NewRequest("GET", detectiveService.BaseURL, nil)

	if err != nil {
		return response.Result, err
	}

	req.Header.Set("X-RapidAPI-Key", os.Getenv("DETECTIVE_API_KEY"))
	req.Header.Set("X-RapidAPI-Host", "breachdirectory.p.rapidapi.com")

	query := req.URL.Query()

	query.Add("func", "auto")
	query.Add("term", key)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return response.Result, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return response.Result, err
	}

	return response.Result, nil
}
