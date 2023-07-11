# Cross compile each of binaries we want to distribute. Store them in the docs
# folder because that's where we serve the website from. I'll eventually create
# a little JS tool that automatically selects the best OS and architecture for
# you (but also allows you to change it).

# Mac Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o docs/bin/darwin-arm64/ask
# Mac Intel
GOOS=darwin GOARCH=amd64 go build -o docs/bin/darwin-amd64/ask
# Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o docs/bin/linux-amd64/ask
# Linux 32-bit
GOOS=linux GOARCH=386 go build -o docs/bin/linux-386/ask
# Windows 64-bit x86
GOOS=windows GOARCH=amd64 go build -o docs/bin/windows-amd64/ask.exe
# Windows 64-bit Arm
GOOS=windows GOARCH=arm64 go build -o docs/bin/windows-arm64/ask.exe
# Windows 32-bit x86
GOOS=windows GOARCH=386 go build -o docs/bin/windows-386/ask.exe
