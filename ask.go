package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/codazoda/lil/env"
	"github.com/codazoda/lil/file"
	"github.com/sashabaranov/go-openai"
)

func main() {

	// Define variables
	var system string

	// Setup the command line arguments we accept
	repeatQuestion := flag.Bool("r", false, "Repeat the question before answering it")
	systemName := flag.Bool("n", false, "Output the name of the system role file (without extension)")
	systemRole := flag.String("s", "", "Read the specified file into the system role")

	// Parse the command line
	flag.Parse()

	// Grab the first non-flag command-line argument
	question := flag.Arg(0)

	// If the question is still empty use any data piped from stdin
	if len(question) < 1 {
		// Check if data was piped into stdin
		fileInfo, _ := os.Stdin.Stat()
		if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
			// Read the piped data from stdin
			var err error
			question, err = readStdin()
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				return
			}
		}
	}

	// If no question was asked, provide help output
	if len(question) < 1 {
		showSyntax()
	}

	// Set the system role to a default or load the contents of a text file
	if *systemRole == "" {
		system = "You are a helpful assistant."
	} else {
		var err error
		system, err = readFile(*systemRole)
		if err != nil {
			fmt.Printf("File error: %v\n", err)
			return
		}
	}

	// Call the OpenAI API and return the response
	token := os.Getenv("OPENAI_API_KEY")
	model := env.GetEnv("OPENAI_MODEL", "gpt-3.5-turbo")
	url := env.GetEnv("OPENAI_BASE_URL", "https://api.openai.com/v1")
	config := openai.DefaultConfig(token)
	config.BaseURL = url
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	// If the flag was present, repeat the question in the output
	if *repeatQuestion {
		fmt.Printf("\"%s\"\n\n", question)
	}

	// If the flag was present, add the system role name to the beginning of
	// the output
	if *systemName {
		fmt.Printf("%s: ", *systemRole)
	}

	// Output the answer
	fmt.Println(resp.Choices[0].Message.Content)

}

func readStdin() (string, error) {
	// Read all data from standard input
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("error reading from stdin: %v", err)
	}

	// Convert byte data to string
	return string(data), nil
}

func readFile(name string) (string, error) {
	// Find the XDG config directory
	configPath, err := file.GetConfigDir("ask")
	if err != nil {
		fmt.Println("Cannot find config directory", err)
		return "", err
	}
	// Open the specified system (preamble) file
	file, err := os.Open(filepath.Join(configPath, "system", name+".txt"))
	if err != nil {
		return "", err
	}
	// Read the bytes from the file
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		return "", err
	}
	// Close the file
	file.Close()
	// Convert the byte array into a string
	fileString := buf.String()
	// Return the contents of the file as a string
	return fileString, nil
}

func showSyntax() {
	fmt.Println("Usage: ask [options] \"question\"")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}
