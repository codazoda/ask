# Ask

This is a quick Go program to make a call to the OpenAI API to ask a question.

Format:
```
ask "Write a summary for me"
```

## Getting Started

Ask reads the OpenAPI key from the OPENAI_API_KEY environment variable and uses that to authenticate with the API. So, add your key to the environment with something like the following.

```
export OPENAI_API_KEY=your_key_goes_here
```

Alternatively, add the above export statement to the `~/.bashrc`, `~/.zshrc`, or equivelent file.

## Arguments

--help  help with the list of command-line arguments available
-q      repeat the question before answering it
