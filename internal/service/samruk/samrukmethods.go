package samruk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"parser/internal/models"
)

func (s SamrukService) GetLotsByFilter(filter models.Filter) ([]models.Result, error) {
	return nil, nil
}

func (s SamrukService) GetLotsById(id string) (models.Result, error) {
	parsedURL, err := url.Parse(s.URLs.LotURL)
	if err != nil {
		return models.Result{}, err
	}

	values := parsedURL.Query()

	values.Add("size", "1")
	values.Add("page", "0")

	parsedURL.RawQuery = values.Encode()

	reqBody := fmt.Sprintf(`{"tenderSubjectTypes": [],
	"advertStatus": "ALL",
	"lotStatus": "ALL",
	"query": %s
	}`, id)

	req, err := http.Post(parsedURL.String(), "application/json", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return models.Result{}, err
	}

	if req.StatusCode != http.StatusOK {
		return models.Result{}, errors.New("Bad status code")
	}

	var res []models.Result
	if err := json.NewDecoder(req.Body).Decode(&res); err != nil {
		return models.Result{}, err
	}

	return res[0], nil
}
