# Creating a Go application

This tutoril walks through the creation of a hello world application

## Simple Hello

* Install go: https://golang.org/doc/install?download=go1.8.darwin-amd64.pkg
* Create hello.go: [hello.go](hello.go)
* Compile it: `go build hello.go`
  * Note: Later, we will use `go install ...` to build the executable and store it in a more formal directory structure.
* Run the resulting file: `./hello`

## Building and Installing Go application into bin directory
* Setup GOPATH - this is the env var used to point at your workspace
* Create a directory, your workspace: `mkdir go-hw`
* Setup GOPATH - this is the env var used to point at your workspace: `export $GOPATH=/Users/davidbell/projects/go-hw`
* Create the basic directory structure:
```
  cd go-hw
  mkdir src
  mkdir bin
  mkdir pkg
```
* The one workspace contains many 'repositories'. It is common to reflect this in the path. If a repository comes from github, use github.com/user in the path of the source:
```
  mkdir -p src/github.com/thisisdavidbell
  cd src/github.com/thisisdavidbell
  git clone <this repo>
```
* You would now create the go program, but `src/github.com/thisisdavidbell/go-helloworld/hello.go` should now exist on your file system
* Compile the program and place it in /bin. From anywhere on the system: `go install github.com/thisisdavidbell/go-helloworld`.
  * Note the path is relative to '$GOPATH' and is the dir containing the go file(s), but not including the file
