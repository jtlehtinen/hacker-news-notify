set windows-shell := ['cmd', '/C']
set shell := ['bash', '-c']

default: run

# Compile and run
run:
  go run ./cmd

# Build release version
build:
  go build -ldflags=all='-H=windowsgui' -o hnn.exe ./cmd
