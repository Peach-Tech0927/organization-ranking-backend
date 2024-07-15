package githubQuery

import (

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"os"
)

const githubAPI = "https://api.github.com/graphql"

func GetContributions(username string) (int, error) {
	//仮
	jst := time.FixedZone("JST", 9*60*60)
	now := time.Now().In(jst)
	oneDayAgo := now.AddDate(0, 0, -1)
	oneWeekAgo := now.AddDate(0, 0, -7)

	token := os.Getenv("GITHUB_TOKEN")
	variables := Variables{
		Username: username,
		From:     oneWeekAgo.Format(time.RFC3339),
		To:       oneDayAgo.Format(time.RFC3339),
	}

	requestBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", githubAPI, bytes.NewBuffer(body))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	var response GraphQLResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return response.Data.User.ContributionsCollection.ContributionCalendar.TotalContributions, nil
}