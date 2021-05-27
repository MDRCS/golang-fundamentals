# Golang Fundamentals
Basic To Advanced Level on Golang.

+ Setup golang environment :

    - $ `brew install go`
    - Check golang version by typing $ `go version`

    If everything is set up correctly, you should see something like this printed:
    `go version go1.15.2 darwin/amd64` 
  
    - To investigate `go` command if it comes from the correct path :
    - $ `which go`
      `-> /usr/local/bin/go`
    
    - When you install a new package go, it should be under this Folder :
    - $ `echo $HOME/go` and source code under `echo $HOME/go/src`.
    - Compiled binaries under `echo $HOME/go/bin`.
    
    + To setup golang path for future dependencies :
    - touch ~/.profile (Only if `.profile` doesn't exist [to understand more please use this command `man bash`])
    - vi ~/.profile
    -> paste :
                
            export GOPATH=$HOME/go
            export PATH=$PATH:$GOPATH/bin

1- First Program :

  - go run and go build 
  
  There are two similar commands available via go: 
  `go run` and `go build`. Each command takes either a single Go file, a list of Go files, or the name of a package.

  + Fact :

  - Golang is a compiled programming languages but if we want to execute the code and see the result without building the binary,
    we could run directly `go run`.
  - Most of the time you want to build a binary for later use. That’s where you use the go build command.


