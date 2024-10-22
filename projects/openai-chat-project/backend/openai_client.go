package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OpenAIRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Struct for OpenAI API response
type OpenAIResponse struct {
	Choices []struct {
		Message OpenAIMessage `json:"message"`
	} `json:"choices"`
}

// Load environment variables from .env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file located, attempting to read from OS")
	}
}

// SendMessageToOpenAI sends a user message to OpenAI API and returns the response
func SendMessageToOpenAI(message string) (string, error) {
	// Load environment variables
	loadEnv()

	// OpenAI API URL
	apiUrl := "https://api.openai.com/v1/chat/completions"

	// Get the OpenAI API key from environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	// Create the request body
	requestBody := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []OpenAIMessage{
			{Role: "user", Content: message},
		},
	}

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response body to OpenAIResponse struct
	var openaiResponse OpenAIResponse
	err = json.Unmarshal(body, &openaiResponse)
	if err != nil {
		return "", err
	}

	// Return the content of the first choice message
	if len(openaiResponse.Choices) > 0 {
		return openaiResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI API")
}
