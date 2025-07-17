package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"strings"
	"github.com/icholy/digest"
)

//SendHTTPRequest ... send request to the server with encoded token value
func SendHTTPRequest(url string, encodedToken string, wg *sync.WaitGroup, timeout int){

    if timeout == 0 {
        timeout = 10
    }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedToken))
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	resp.Body.Close()
	defer wg.Done()
	if resp.StatusCode == 200{
		fmt.Println()
		PrintSuccess("Found valid credentials", Base64Decode(encodedToken), false, false)
		fmt.Println()
		fmt.Println()
		os.Exit(0)
	}else{
		PrintFailed("failed", Base64Decode(encodedToken), false, true)
	}
}

// SendHTTPRequestWithAuthType ... send request with either Basic or Digest auth
func SendHTTPRequestWithAuthType(url string, encodedToken string, wg *sync.WaitGroup, timeout int, authType string) {
    switch authType {
    case "basic":
        SendHTTPRequest(url, encodedToken, wg, timeout)
    case "digest":
        SendDigestHTTPRequest(url, encodedToken, wg, timeout)
    default:
        log.Fatalf("Unsupported auth type: %s", authType)
    }
}

// SendDigestHTTPRequest ... stub for digest auth, to be implemented
func SendDigestHTTPRequest(url string, encodedToken string, wg *sync.WaitGroup, timeout int) {
	defer wg.Done()
	if timeout == 0 {
		timeout = 10
	}
	// Decode the token to get username and password
	decoded := Base64Decode(encodedToken)
	parts := strings.SplitN(decoded, ":", 2)
	if len(parts) != 2 {
		PrintFailed("failed", "Invalid token format", false, true)
		return
	}
	username := parts[0]
	password := parts[1]

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &digest.Transport{
			Username: username,
			Password: password,
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request. ", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println()
		PrintSuccess("Found valid credentials", decoded, false, false)
		fmt.Println()
		fmt.Println()
		os.Exit(0)
	} else {
		PrintFailed("failed", decoded, false, true)
	}
}