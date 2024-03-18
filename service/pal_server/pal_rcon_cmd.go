package pal_server

import (
	"io"
	"net/http"
	"strings"
)

type PalRconCmd struct{}

// Send 发送RCON命令
func (s *PalRconCmd) Send(request string) (response string, err error) {
	url := "http://127.0.0.1:3000/command"
	method := "POST"
	payload := strings.NewReader(`{"command":"` + request + `"}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer panneltoken")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
