package githubQuery
//   fromToは要相談
import "encoding/json"

const query = `
query($username: String!, $from: DateTime!, $to: DateTime!) {
  user(login: $username) {
    contributionsCollection(from: $from, to: $to) {
		contributionCalendar {
			totalContributions
      }
    }
  }
}`

type Variables struct {
	Username string `json:"username"`
	From     string `json:"from"`
	To       string `json:"to"`
}

type GraphQLRequest struct {
	Query     string   `json:"query"`
	Variables Variables `json:"variables"`
}

type ContributionCalendar struct {
	TotalContributions int `json:"totalContributions"`
}

type ContributionsCollection struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}

type User struct {
	ContributionsCollection ContributionsCollection `json:"contributionsCollection"`
}

type Data struct {
	User User `json:"user"`
}

type GraphQLResponse struct {
	Data Data `json:"data"`
}