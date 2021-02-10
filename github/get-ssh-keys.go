package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// APISSHKeys Github SSH Keys type
type APISSHKeys = []struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

// GetSSHKeys Get SSH keys for an user
func GetSSHKeys(user string) (APISSHKeys, error) {

	var keys APISSHKeys

	userURLEncoded := url.QueryEscape(user)
	res, err := http.Get("https://api.github.com/users/" + userURLEncoded + "/keys")

	if err != nil {
		return keys, err
	}

	if res.StatusCode != 200 {
		return keys, fmt.Errorf("Can't get your github user. GitHub API status code %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&keys)

	if err != nil {
		return keys, err
	}

	return keys, nil
}
