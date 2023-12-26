# EDdsa Signature Implementation in Go

## Install

[Install Golang](https://go.dev/doc/install)

Install [Xcode Command Line Tools](https://developer.apple.com/downloads/)

```
go install golang.org/x/mobile/cmd/gomobile@latest
```

```
gomobile init
```

## Build

### iOS 

```
gomobile bind -target=ios .
```

This command will generate a binary according to the target platform.

Drag the binary into project and build.