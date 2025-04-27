package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	introspectionEndpoint = "https://{{ DOMAIN }}/oauth/v2/introspect"
	revokationEndpoint    = "https://{{ DOMAIN }}/oauth/v2/revoke"

	clientID     = "super_secret_client_id"
	clientSecret = "super_secret_client_secret"
)

func VerifyJWT(accessToken string) (*AuthInfo, error) {
	response, err := verifyAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	if !response.Active {
		return nil, fmt.Errorf("token is not active")
	}

	var roles []string
	for key := range response.Roles {
		roles = append(roles, key)
	}

	info := AuthInfo{
		Name:     response.Name,
		Username: response.Username,
		Email:    response.Email,
		Verified: true,
		Roles:    roles,
	}

	return &info, nil
}

func verifyAccessToken(token string) (*IntrospectionResponse, error) {
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", introspectionEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("introspection failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	var introspectionResponse IntrospectionResponse
	err = json.Unmarshal(body, &introspectionResponse)
	if err != nil {
		return nil, err
	}

	return &introspectionResponse, nil
}

func RevokeAccessToken(token string) (*int, error) {
	data := url.Values{}
	data.Set("token", token)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", revokationEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Token revokation failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	return &resp.StatusCode, nil
}

type AuthInfo struct {
	Name     string
	Username string
	Email    string
	Verified bool
	Roles    []string
}
type IntrospectionResponse struct {
	Active        bool                         `json:"active"`
	Scope         string                       `json:"scope"`
	ClientID      string                       `json:"client_id"`
	Username      string                       `json:"username"`
	Email         string                       `json:"email"`
	Name          string                       `json:"name"`
	ExpiresAt     int64                        `json:"exp"`
	Roles         map[string]map[string]string `json:"urn:oauth2:iam:org:project:304089514807181689:roles"`
	FamilyName    string                       `json:"family_name"`
	GivenName     string                       `json:"given_name"`
	Locale        string                       `json:"locale"`
	EmailVerified bool                         `json:"email_verified"`
}
