# Ask

This is a command-line utility for asking AI questions. It makes a call to the OpenAI API.

Usage:
```
ask "Write a summary for me"
```


## Getting Started

- Add the OPENAI_API_KEY to your environment variables
- Compile ask

Ask reads the OpenAPI key from the OPENAI_API_KEY environment variable and uses that to authenticate with the API. So, add your key to the environment with something like the following.

```
export OPENAI_API_KEY=your_key_goes_here
```

Alternatively, add the above export statement to the `~/.bashrc`, `~/.zshrc`, or equivelent file.


## Compiling

Ask is written in Go. Once you've installed Go you can compile `ask` by cloning this repo and then running the following command in the root directory.

```
go build
```


## Arguments

--help  help with the list of command-line arguments available
-q      repeat the question before answering it


## Future

Some things I might want to do in the future.

- Add a license
- Build binaries for popular operating systems
- Add an option to change the 
- Maybe support other AI API's
