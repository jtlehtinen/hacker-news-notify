# hacker-news-notify

System tray application for Windows that displays a toast notification when best/new/top story listings change in Hacker News.

## Requirements

- [just](https://github.com/casey/just) is used as a command runner.
- [rsrc](https://github.com/akavel/rsrc) exe should be in PATH. It is used to create a `syso` file including the application icon. This allows to embed the icon in the exe.

If you don't want to use either of these then to build run the following in the root folder.

```sh
$ go build -ldflags=all='-H=windowsgui' -o hnn.exe ./cmd
```
