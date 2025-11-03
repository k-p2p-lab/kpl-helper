package curl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Curl is a simple HTTP client function to send requests.
func Curl(url, method string, body ...string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var reqBody io.Reader = nil
	if len(body) > 0 {
		reqBody = bytes.NewBuffer([]byte(body[0]))
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "Go-Curl/1.0")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(respBody), nil
}

// GetTF converts a boolean value to its string representation: "t" for true, "f" for false.
func GetTF(value bool) string {
	if value {
		return "t"
	} else {
		return "f"
	}
}
