# go-uselibrary
A go app to use the firstlibrary go package

This app uses the library at : https://github.com/thisisdavidbell/go-examples/gofirstlibrary

## Building a Go library, and using it

* Create a new git repo, and clone it into 'src/github.com/thisisdavidbell'. Use repo: `https://github.com/thisisdavidbell/go-firstlibrary`
* Create new go file. In this case you now already have [reverse.go](reverse.go)
* build the package: `go install` in go-firstlibrary.
  * Note: `go build` does nothing. `go install` produces a package in the pkg dir.
* Create a new go file `go-uselibrary-hw` to use this library
  * You can just clone the following example into 'src/github.com/thisisdavidbell': `https://github.com/thisisdavidbell/go-uselibrary`
  * this gives you `go-uselibrary-hw.go`
* Ensure that the import statement in `go-uselibrary-hw.go` matches the name of the library you created above.
* build the go application: `go install` in go-uselibrary-hw dir, or `go install github.com/thisisdavidbell/go-uselibrary`.
  * Note that the command is `go-uselibrary` - the name of the directory, not the file itself.
* test the app: `go-use-lbrary`
  * Remember this is picked up from the bin dir, as it was added to PATH.

You should now have a dir structure like:
```
bin/
    go-uselibrary                 # command executable
pkg/
    darwin_amd64/          # this will reflect your OS and architecture
        github.com/thisisdavidbell/
            go-firstlibrary.a  # package object
src/
    github.com/thisisdavidbell/
        go-uselibrary/
            go-uselibrary-hw.go      # command source
        go-firstlibrary/
            reverse.go    # package source
```
Note: Go command executables are statically linked - that is the dependencies are pulled in at compile time, and included in the executable. Therefore, the package objects need not be present to run Go programs.

Package names notes:
* The first statement in a Go source file must be
```
package <package name>
```
where name is the package's default name for imports. (All files in a package must use the same name.)
* Go's convention is that the package name is the last element of the import path: the package imported as "crypto/rot13" should be named rot13.
* Executable commands must always use package main.
* There is no requirement that package names be unique across all packages linked into a single binary, only that the import paths (their full file names) be unique.
