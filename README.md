# gohttphw
A go http server returning hello world, later to be used with docker

# Creating a web service, using built in net/http
* Checkout out [go-http-hw.go](go-http-hw.go)
* Build and run go-http-how.go. Either
  * Use `go build` or `go build github.com/thisisdavidbell/gohttphw` to create executable in current dir, then run the executable created in the local dir: `./gohttphw`
  * Or use `go run <path if not curr dir/>go-http-hw.go` to compile and run and throw away executable.
  * OR use `go install` to build the executable and place it in the $GOPATH/bin directory.

# Useful links
* basic http example: https://golang.org/doc/articles/wiki/
* Docker: https://blog.golang.org/docker
* Recommended router when built in one not enough: http://www.gorillatoolkit.org/pkg/mux
