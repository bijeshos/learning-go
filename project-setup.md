# Project setup
- Run `go env -w GOPATH=/path/to/gopath`
    - to set the GOPATH location

- Run `go env -w GOBIN=/path/to/my/bin`
    - to set the binary path

- Run `go mod init github.com/bijeshos/learning-go`
    - to initialize a module

- Run `go build`
- Run `go build github.com/bijeshos/learning-go`
    - to compile packages and dependencies

- Run `go install`
- Run `go install github.com/bijeshos/learning-go`
    - to install project's binaries in workspace's bin directory

- Run `go list -m all`
    - to view final versions that will be used in a build for all direct and indirect dependencies

- Run `go get golang.org/x/tools/cmd/godoc`
    - to get go doc