# Go commands - Reference

## Usage
- `go <command> [arguments]`

## Go Commands
- `go version `    # print Go version
- `go clean   `    # remove object files and cached files
- `go build   `    # compile packages and dependencies
- `go fmt     `    # gofmt (reformat) package sources
- `go doc     `    # show documentation for package or symbol
- `go env     `    # print Go environment information
- `go fix     `    # update packages to use new APIs
- `go generate`    # generate Go files by processing source
- `go get     `    # add dependencies to current module and install them
- `go install `    # compile and install packages and dependencies
- `go list    `    # list packages or modules
- `go run     `    # compile and run Go program
- `go test    `    # test packages
- `go tool    `    # run specified go tool
- `go vet     `    # report likely mistakes in packages


## Go Commands : test
- `go test`        # to execute the test files in the current directory 
- `go test -cover` # to execute tests with coverage details displayed

- `go test -bench=.` # to perform benchmarking

- `go test -cover -coverprofile=c.out`  # to generate test coverage profile as a file
- `go tool cover -html=c.out -o coverage.html` # to generate html file from coverage profile file 


 