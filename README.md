# How to Write Go Code
* [How to Write Go Code](https://golang.org/doc/code.html)
```
    3  go mod init github.com/godsman-yang/hello
    4  cd hello
    6  go install github.com/godsman-yang/hello
   10  cd morestrings/
   11  go build
   13  cd ..
   14  go install github.com/godsman-yang/hello
   23  go clean -modcache
   25  cd morestrings/
   26  go test
```
