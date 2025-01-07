package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Configuration
const (
	endpointsFile   = "endpoints.txt" // File containing the list of endpoints
	telegramBotToken = "xxxxx"        // Replace with your bot token
	telegramChatID   = "yyyyyy"       // Replace with your chat ID
)

// Function to send a message to Telegram
func sendTelegramMessage(message string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
	payload := fmt.Sprintf("chat_id=%s&text=%s", telegramChatID, message)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		fmt.Println("Error creating Telegram request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending Telegram message:", err)
		return
	}
	defer resp.Body.Close()
}

// Function to check endpoints
func checkEndpoints() {
	file, err := os.Open(endpointsFile)
	if err != nil {
		fmt.Println("Error opening endpoints file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		endpoint := strings.TrimSpace(scanner.Text())
		if endpoint == "" {
			continue
		}

		resp, err := http.Get(endpoint)
		if err != nil {
			sendTelegramMessage(fmt.Sprintf("ALERT: Endpoint %s is unreachable. Error: %v", endpoint, err))
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			sendTelegramMessage(fmt.Sprintf("ALERT: Endpoint %s returned status code %d", endpoint, resp.StatusCode))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading endpoints file:", err)
	}
}

func main() {
	checkEndpoints()
}

