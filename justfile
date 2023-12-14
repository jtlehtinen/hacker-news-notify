set windows-shell := ['cmd', '/C']
set shell := ['bash', '-c']

default: run

# Compile and run the application
run:
  @go run ./cmd

# Build release executable for Windows
build:
  @go build -ldflags=all='-H=windowsgui' -o hnn.exe ./cmd
