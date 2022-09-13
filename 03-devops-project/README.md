# Go WebApp with Dockerized Image and Deployd in Kubernetes

<!-- ## This is how it looks once setup is done completely

![Video Demo](./03-devops-project-ss.gif) -->


> Steps to prepare the project :-
# Small microsrevices built in Go(work in progress)

### Golang Setup

- After installation of Golang, verify the variables _GOROOT_, _GOPATH_, _GOBIN_ are set or not.

  ```bash
  measutosh@crispy MINGW64 ~
  $ echo $GOROOT


  measutosh@crispy MINGW64 ~
  $ echo $GOPATH
  C:\Users\measutosh\go

  measutosh@crispy MINGW64 ~
  $ echo $GOBIN
  ```

- See what are the various available for _Go_.

  ```bash
  measutosh@crispy MINGW64 ~
  $ go env
  set GO111MODULE=
  set GOARCH=amd64
  set GOBIN=
  set GOCACHE=C:\Users\measutosh\AppData\Local\go-build
  set GOENV=C:\Users\measutosh\AppData\Roaming\go\env  
  set GOEXE=.exe
  set GOEXPERIMENT=
  set GOFLAGS=
  set GOHOSTARCH=amd64
  set GOHOSTOS=windows
  set GOINSECURE=
  set GOMODCACHE=C:\Users\measutosh\go\pkg\mod
  set GONOPROXY=
  set GONOSUMDB=
  set GOOS=windows
  set GOPATH=C:\Users\measutosh\go
  set GOPRIVATE=
  set GOPROXY=https://proxy.golang.org,direct
  set GOROOT=C:\Program Files\Go
  set GOSUMDB=sum.golang.org
  set GOTMPDIR=
  set GOTOOLDIR=C:\Program Files\Go\pkg\tool\windows_amd64
  set GOVCS=
  set GOVERSION=go1.19
  set GCCGO=gccgo
  set GOAMD64=v1
  set AR=ar
  set CC=gcc
  set CXX=g++
  set CGO_ENABLED=1
  set GOMOD=NUL
  set GOWORK=
  set CGO_CFLAGS=-g -O2
  set CGO_CPPFLAGS=
  set CGO_CXXFLAGS=-g -O2
  set CGO_FFLAGS=-g -O2
  set CGO_LDFLAGS=-g -O2
  set PKG_CONFIG=pkg-config
  set GOGCCFLAGS=-m64 -mthreads -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=C:\Users\MEASUT~1\AppData\Local\Temp\go-build1806508971=/tmp/go-build -gno-record-gcc-switches
  ```

- Create a basic folder structure for a Go webapp and add the location of this path in _GOPATH_ in _System Environment Variables_. Open a new terminal session to verify whether the paths for the variables are set properly or not. 

  ```bash
  measutosh@crispy MINGW64 ~
  $ echo $GOPATH
  C:\Users\measutosh\Documents\VSCode\DevOps Projects\devops-gallery\03-devops-project\go-devops

  measutosh@crispy MINGW64 ~
  $ echo $GOBIN
  C:\Users\measutosh\Documents\VSCode\DevOps Projects\devops-gallery\03-devops-project\go-devops\bin
  ```

- Once the code is done then build the project. Create a _dockerignore_ file and then create the docker file to dockerize it.
- Add content to the _dockerfile_ and meanwhile start the _minikube cluster_ too.
