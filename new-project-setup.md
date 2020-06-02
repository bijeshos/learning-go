# New project setup

- Ensure that Go is installed and setup in your machine. You would need to setup a code editor as well. If not already done, please refer [here](https://github.com/bijeshos/learning-go/blob/master/prerequisites.md) for instructions to do so.

- Following are a few of the main environment variables needed. If not set, Go will assume default values. If you prefer to set them manually, follow these steps:
    - Set `GOPATH` location 
        - Run `go env -w GOPATH=/path/to/gopath` in a terminal/command prompt
    - Set `GOBIN` location        
        - Run `go env -w GOBIN=/path/to/my/bin`
    - (*Ensure to replace directory paths as needed*)

- Create a new directory to keep your project code (let's call that directory `my-first-go-project`). 

- Open a terminal/command prompt in the newly created directory
    
- Initialize a module 
    - Run `go mod init github.com/your-github-handle/my-first-go-project`
    - (*Ensure to replace __your-github-handle__ with correct value*)

- Now, you can add programs to the project

- To compile packages and dependencies
    - Run `go build github.com/your-github-handle/my-first-go-project`

- To run the generated binary
    - Run `./my-first-go-project` (on Linux)
    - Run `my-first-go-project.exe` (on Windows)

- To install project's binaries in workspace's bin directory
    - Run `go install github.com/your-github-handle/my-first-go-project`
   
