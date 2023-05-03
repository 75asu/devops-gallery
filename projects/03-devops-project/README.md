# Go WebApp with Dockerized Image and Deployed in Kubernetes

## This is how it looks once setup is done completely

![Video Demo](./03-devops-project-ss.gif)


## Steps to prepare the project :-

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

### Container Setup

- Once the code is done then build the project. Create a _dockerignore_ file and then create the docker file to dockerize it.
- Add content to the _dockerfile_ and meanwhile start the _minikube cluster_ too. Run the `eval $(minikube docker-env)` command to returns a set of Bash environment variable exports to configure your local environment to re-use the Docker daemon inside the Minikube instance.

  ```bash
  measutosh@crispy MINGW64 ~
  $ minikube status
  minikube
  type: Control Plane   
  host: Running
  kubelet: Running      
  apiserver: Running    
  kubeconfig: Configured
  ```

- Build the container using docker file, run it and verify from the logs.

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker build -t go-app-ms:latest .
  [+] Building 17.3s (11/11) FINISHED
  => [internal] load build definition from Dockerfile                         1.5s 
  => => transferring dockerfile: 849B                                         0.7s 
  => [internal] load .dockerignore                                            0.7s 
  => => transferring context: 34B                                             0.2s 
  => [internal] load metadata for docker.io/library/golang:latest             9.7s 
  => [1/6] FROM docker.io/library/golang:latest@sha256:d3ca1795fc42c82a42442  0.1s 
  => => resolve docker.io/library/golang:latest@sha256:d3ca1795fc42c82a42442  0.1s 
  => [internal] load build context                                            0.3s 
  => => transferring context: 1.21kB                                          0.1s 
  => CACHED [2/6] WORKDIR /app                                                0.0s 
  => CACHED [3/6] COPY go.mod go.sum ./                                       0.0s 
  => CACHED [4/6] RUN go mod download                                         0.0s 
  => [5/6] COPY . .                                                           0.6s 
  => [6/6] RUN go build -o main .                                             4.1s 
  => exporting to image                                                       0.5s 
  => => exporting layers                                                      0.4s 
  => => writing image sha256:b30cb0fff7db468b34e2cc88231996c5386cbfffa648018  0.0s 
  => => naming to docker.io/library/go-app-ms:latest                          0.0s

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker run -d -p 80:80 --name web go-app-ms:latest
  49e7ef9f33206d6dc23482db9a2eb001900d912808fb795d4c3a35e579a8501f

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker logs -f web
  2022/09/13 18:10:28 Web server has started at <http://localhost>

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker logs -f web
  2022/09/13 18:10:28 Web server has started at http://localhost
  2022/09/13 18:13:43 Homepage available at http://localhost/
  2022/09/13 18:15:08 Fetching the details http://localhost/details
  49e7ef9f3320 172.17.0.2

  ```

- If it gives the expecetd results in the proper API endpoints then setup  a _docker-compose.yaml_ file and add the content required to create a docker container with dockerfile. Build the app, run it using the compose file.

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker-compose build
  [+] Building 23.2s (11/11) FINISHED
  => [internal] load build definition from dockerfile                         2.4s 
  => => transferring dockerfile: 849B                                         1.4s 
  => [internal] load .dockerignore                                            2.3s 
  => => transferring context: 65B                                             0.8s 
  => [internal] load metadata for docker.io/library/golang:latest            14.3s 
  => [1/6] FROM docker.io/library/golang:latest@sha256:d3ca1795fc42c82a42442  0.5s 
  => => resolve docker.io/library/golang:latest@sha256:d3ca1795fc42c82a42442  0.5s 
  => [internal] load build context                                            4.4s 
  => => transferring context: 1.99MB                                          3.8s 
  => CACHED [2/6] WORKDIR /app                                                0.0s 
  => CACHED [3/6] COPY go.mod go.sum ./                                       0.0s 
  => CACHED [4/6] RUN go mod download                                         0.0s 
  => CACHED [5/6] COPY . .                                                    0.0s 
  => CACHED [6/6] RUN go build -o main .                                      0.0s 
  => exporting to image                                                       0.2s 
  => => exporting layers                                                      0.0s 
  => => writing image sha256:b30cb0fff7db468b34e2cc88231996c5386cbfffa648018  0.0s 
  => => naming to docker.io/library/go-app-ms:latest                          0.0s

  Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker-compose up -d
  [+] Running 1/1
  - Container go-microservices-web-1  S...                                    7.3s 

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker-compose stop
  [+] Running 1/1
  - Container go-microservices-web-1  S...                                    2.3s


  ```

### Minikube Deployment

- Verify again whether the API endpoints are giving results or not. I yes then stop docker compose and prepare for the minikube deployment.Make a folder to hold all the kubernetes deployment objects, create abd add content to _deployment.yml and service.yml_, apply those files, verify the deployment

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ kubectl apply -f deployments/
  deployment.apps/web-deployment created
  service/web-service created

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ kubectl get pods
  NAME                              READY   STATUS   RESTARTS   AGE
  web-deployment-594c4bf755-4jqr2   0/1     Error    0          12s
  web-deployment-594c4bf755-4zvp6   0/1     Error    0          12s
  web-deployment-594c4bf755-twknl   0/1     Error    0          12s

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ kubectl get pods
  NAME                              READY   STATUS             RESTARTS      AGE
  web-deployment-594c4bf755-4jqr2   0/1     CrashLoopBackOff   1 (8s ago)    21s    
  ago)   28s                                                          ago)   21s
  ago)   28s                                                          go)    21s    
  ago)   28s
  ago)   28s
  ago)   28s
  ago)   28s
  ago)   28s
  NAME          TYPE        CLUSTER-IP      EXTERNAL-IP  Off   1 (14s ago)   28s
  NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
  NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
  NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
  NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
  NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
  kubernetes    ClusterIP   10.96.0.1       <none>        443/TCP        11h
  web-service   NodePort    10.107.47.180   <none>        80:31768/TCP   42s

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ minikube service web-service
  |-----------|-------------|-------------|---------------------------|
  | NAMESPACE |    NAME     | TARGET PORT |            URL            |
  |-----------|-------------|-------------|---------------------------|
  | default   | web-service |          80 | http://192.168.49.2:31768 |
  |-----------|-------------|-------------|---------------------------|
  🏃  Starting tunnel for service web-service.
  |-----------|-------------|-------------|------------------------|
  | NAMESPACE |    NAME     | TARGET PORT |          URL           |
  |-----------|-------------|-------------|------------------------|
  | default   | web-service |             | http://127.0.0.1:51698 |
  |-----------|-------------|-------------|------------------------|
  🎉  Opening service default/web-service in default browser...
  ❗  Because you are using a Docker driver on windows, the terminal needs to be open to run it.
  ```



