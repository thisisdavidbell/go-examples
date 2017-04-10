# gotestrest - Create a REST service with go, and test at the REST level

# Creating a web service, using built in net/http
* fetch this repo to be able to use git, develop and run go: `go get `go get github.com/thisisdavidbell/go-examples/gotestrest`
* cd to github.com/thisisdavidbell/go-examples/gotestrest
  * Note `git status` runs fine.
* Build and run go-http-how.go. Either
  * Use `go build` or `go build github.com/thisisdavidbell/gohttphw` to create executable in current dir, then run the executable created in the local dir: `./gohttphw`
  * Or use `go run <path if not curr dir/>go-http-hw.go` to compile and run and throw away executable.
  * OR use `go install` to build the executable and place it in the $GOPATH/bin directory. Run the executable from the bin dir (add it to your ATH env var): `gohttphw`

# Dockerise go app
* create a Dockerfile which builds and runs the go executable. See [Dockerfile](Dockerfile)
* build the docker image: `docker build -t gohttphw-docker-image .`
* Run: `docker images`
* run the docker file: `docker run -p 8081:8080 -d gohttphw-docker-image`
  * Note: -d runs it in the background
* Run: `docker ps`
* call the service: `curl localhost:8081/hello`
* log into the running container: `docker exec -it <container id from docker ps> /bin/bash`
* Explore using: `ps -ef`, `ls -l /go`, then `exit`
* Stop the container: `docker stop <container id from docker ps>`
* Run: `docker ps`

# Useful links
* basic http example: https://golang.org/doc/articles/wiki/
* Docker: https://blog.golang.org/docker
* Recommended router when built in one not enough: http://www.gorillatoolkit.org/pkg/mux
