package services

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

type DetectiveService struct {
	db      *cache.Cache
	BaseURL string
}

func NewDetectiveService() *DetectiveService {
	return &DetectiveService{
		db:      cache.New(time.Minute*15, time.Minute*15),
		BaseURL: os.Getenv("DETECTIVE_URL"),
	}
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

	cached, found := detectiveService.db.Get(key)

	if found {
		return cached.([]Breach), nil
	} else {
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

		detectiveService.db.Set(key, response.Result, cache.DefaultExpiration)

		return response.Result, nil
	}
}
