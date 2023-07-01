package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {

	// Setup the command line arguments we accept
	repeatQuestion := flag.Bool("r", false, "Repeat the question before answering it")
	//question := flag.Bool("q", false, "Send this question")

	// Parse the command line and grab the first non-flag command-line argument
	flag.Parse()
	question := flag.Arg(0)

	// If no question was asked, provide help output
	if len(question) < 1 {
		showSyntax()
	}

	// Call the OpenAI API and return the response
	token := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
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

	// If the -q flag was present, repeat the question in the output
	if *repeatQuestion {
		fmt.Printf("\"%s\"\n\n", question)
	}

	// Output the answer
	fmt.Println(resp.Choices[0].Message.Content)

}

func showSyntax() {
	fmt.Println("Usage: ask [options] \"question\"")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}
