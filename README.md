# AskAI

A command-line utility for asking AI questions.  It sends your question to the OpenAI API.

![Demo of Ask Command Line](docs/image/ask.gif)

The command-line options are becoming more stable now but this is still beta.

Usage:
```
ask "Write a summary for me"
```

You can also set a system message, which reads from a file:
```
ask -s future "What does the future hold?"
```

You can now pipe questions in via standard input:
```
ask < question.txt
```


## Getting Started

- Add the OPENAI_API_KEY to your environment variables
- Compile ask

Ask reads the OpenAPI key from the OPENAI_API_KEY environment variable and uses that to authenticate with the API. So, add your key to the environment with something like the following.

```
export OPENAI_API_KEY=your_key_goes_here
```

Alternatively, add the above export statement to the `~/.bashrc`, `~/.zshrc`, or equivelent file.


## Configuration

You can store preamble text in the `~/.config/ask/system/` directory as text files. For example, create a file called `future.txt` there and put the following text into it.

```
It's the year 2053.
```

Now you can run the following command to ask questions using the _future_ preamble.

```
ask -s future "What does the future hold?"
```

Ask uses the [XDG Directory Specification](https://xdgbasedirectoryspecification.com/) to store files in a way that is cross platform friendly.

This directory is called `system` because OpenAI refers to the preamble as the _system_ message. By default it's something like, "You are a helpful assistant." Setting this is extremely useful for fine tuning the answers you get from the AI.


## Compiling

There is also a `build.sh` shell script that cross-compiles `ask` for lots of different architectures. This puts all the builds in the `./docs` folder. Run the following command to build all the architectures. 

```
./build.sh
```

You can also compile `ask` by by running the following command in the root directory. This builds for the current architecture and outputs the file in the current directory.

```
go build
```


## Arguments

```
Usage: ask [options] "question"
Options:
  --help  Display help for the list of command-line arguments
  -q      Repeat the question before answering it
  -s      Read the specified file into the system role as a preamble
```


## Future

Some things I might want to do in the future.

- Add a license
- Build binaries for popular operating systems
- Add an option to change the model
- Maybe support other AI API's


## The Name

I use the names _AskAI_ and _Ask_ interchangably. The official name of the project is _AskAI_ but I use _Ask_ for short and for the command-line filename. I do so because _Ask_ is not specific enough for search queries or conversations but it works well as the name of the command-line binary.
