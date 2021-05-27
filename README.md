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
    - `touch ~/.profile` (Only if `.profile` doesn't exist [to understand more please use this command `man bash`])
    - `vi ~/.profile`
    -> paste :
                
            export GOPATH=$HOME/go
            export PATH=$PATH:$GOPATH/bin
    - `source $HOME/.profile` (for validation)
  
1- First Program :

  - go run and go build 
  
  There are two similar commands available via go: 
  `go run` and `go build`. Each command takes either a single Go file, a list of Go files, or the name of a package.

  ++ Fact :

  - Golang is a compiled programming languages but if we want to execute the code and see the result without building the binary,
    we could run directly `go run`.
  - Most of the time you want to build a binary for later use. That’s where you use the go build command.

2- Install Third Party Tools :

+ One thing to remember about go package management is that is different from any other tool that do the same thing :
as NPM or Maven, go install new packages from there repositories as example below :
  
  + $ `go install github.com/rakyll/hey@latest` ( @latest is for version that you want to install)
  + check the bin by typing `ls $HOME/go/bin`
  + also we can run -> `hey https://www.golang.org`
  
  -< to do code formatting and linting :
  - run $ `go fmt`
  
  NB: There’s an enhanced version of `go fmt` available called `goimports` that also cleans up your import statements. 
      It puts them in alphabetical order, removes unused imports, and attempts to guess any unspecified imports. 
      Its guesses are sometimes inaccurate, so you should insert imports yourself.

  + You can download goimports with the command 
    $ `go install golang.org/x/tools/cmd/goimports@latest`
    You run it across your project with the command:
    $ `goimports -l -w .`
    
    ++ The `-l` flag tells goimports to print the files with incorrect formatting to the console. 
       The `-w` flag tells goimports to modify the files in-place. 
       The `.` specifies the files to be scanned: everything in the current directory and all of its subdirectories.
    
    NB: Always run go fmt or goimports before compiling.
  
  - Code linting :
  
  $ `go install golang.org/x/lint/golint@latest`
  and run $ `golint ./...` (to target the entire project)
  
  - Code vetting :

  ++ There is another class of errors that developers run into. 
  The code is syntactically valid, but there are mistakes that are not what you meant to do. 
  This includes things like passing the wrong number of parameters to formatting methods or 
  assigning values to variables that are never used. The go tool includes a command called 
  go vet to detect these kinds of errors. Run go vet on your code with the command:
  
  $ `go vet ./...`

  + You can run Code Formatting and linting and vetting at the same time with one tool :
  `golangci-lint`
    
    - you can install it thought `https://golangci-lint.run/usage/install/`
    - once installed run $ `golangci-lint run`

  +++ Because golangci-lint runs so many tools (as of this writing, it runs 10 different linters by default, 
      and allows you to enable another 50), it’s inevitable that your team disagrees with some of its suggestions. 
      You can configure which linters are enabled and which files they analyze by including a file named .golangci.yml 
     at the root of your project. The documentation for the file format is found at https://github.com/golangci/golangci-lint#config-file.
