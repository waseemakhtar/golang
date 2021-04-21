package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://www.git.example.com/"

type Client struct {
	Username string
	Password string
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (s *Client) GetUser(id int) (*User, error) {
	url := fmt.Sprintf(baseURL+"/api/v4//users/%d", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("req: %#v\n", req)
	bytes, err := s.doRequest(req)
	if err != nil {
		fmt.Print("Error in doRequest\n")
		return nil, err
	}
	fmt.Printf("User1: %#v\n", bytes)
	var user User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	fmt.Printf("User2: %#v\n", user)
	return &user, nil
}

func main() {
	client := NewBasicAuthClient("username", "password")

	// Fetch a user
	user, _ := client.GetUser(1250)

	fmt.Printf("User: %#v\n", user)
}
