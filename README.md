# AskAI

A command-line utility for asking AI questions.  It sends your question to the OpenAI API.

This utility does not have a stable usage yet. The command-line options are likely to change. This is an early proof of concept (PoC).

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

AskAI is written in Go. Once you've installed Go you can compile `ask` by cloning this repo and then running the following command in the root directory.

```
go build
```

There is also a `build.sh` shell script that cross-compiles `ask` for lots of different architectures. Run the following command to build all the architectures. 

```
./build.sh
```


## Arguments

```
Usage: ask [options] "question"
Options:
  --help  Display help for the list of command-line arguments
  -q      Repeat the question before answering it
```


## Future

Some things I might want to do in the future.

- Add a license
- Build binaries for popular operating systems
- Add an option to change the model
- Maybe support other AI API's


## The Name

I use the names _AskAI_ and _Ask_ interchangably. The official name of the project is _AskAI_ but I use _Ask_ for short and for the command-line filename. I do so because _Ask_ is not specific enough for search queries or conversations but it works well as the name of the command-line binary.
