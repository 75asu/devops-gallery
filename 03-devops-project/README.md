# Go WebApp with Dockerized Image and Deployed in Kubernetes

<!-- ## This is how it looks once setup is done completely

![Video Demo](./03-devops-project-ss.gif) -->


> Steps to prepare the project :-

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

- See all the docker containers for confirmation

  ```bash
  measutosh@crispy MINGW64 ~
  $ eval $(minikube docker-env)

  measutosh@crispy MINGW64 ~
  $ docker ps
  CONTAINER ID   IMAGE                          COMMAND                  CREATED          STATUS          PORTS     NAMES
  c0133df060e8   kubernetesui/metrics-scraper   "/metrics-sidecar"       12 minutes ago   Up 12 minutes             k8s_dashboard-metrics-scraper_dashboard-metrics-scraper-78dbd9dbf5-p8j92_kubernetes-dashboard_a8bd4b5f-e384-4fc7-a57f-9699a8485673_0
  9e27426bebd9   kubernetesui/dashboard         "/dashboard --insecu…"   12 minutes ago   Up 12 minutes             k8s_kubernetes-dashboard_kubernetes-dashboard-5fd5574d9f-r885b_kubernetes-dashboard_6eef73af-d8fe-41d0-a0f4-c64f96c6eed7_0
  a3cd5e4ce29e   6e38f40d628d                   "/storage-provisioner"   12 minutes ago   Up 12 minutes             k8s_storage-provisioner_storage-provisioner_kube-system_b631381c-bc31-4269-91bb-6dfd90d617f1_1
  fd5f208b57fd   a4ca41631cc7                   "/coredns -conf /etc…"   13 minutes ago   Up 13 minutes             k8s_coredns_coredns-6d4b75cb6d-ww4mk_kube-system_b3cdff18-81b7-47c0-9ae9-93f624c20819_0
  3001dfc696d3   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_dashboard-metrics-scraper-78dbd9dbf5-p8j92_kubernetes-dashboard_a8bd4b5f-e384-4fc7-a57f-9699a8485673_0
  d07ecbdf4fbc   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_kubernetes-dashboard-5fd5574d9f-r885b_kubernetes-dashboard_6eef73af-d8fe-41d0-a0f4-c64f96c6eed7_0
    6837ce85a12a   2ae1ba6417cb                   "/usr/local/bin/kube…"   13 minutes ago   Up 13 minutes             k8s_kube-proxy_kube-proxy-mjrzn_kube-system_61781a67-98fd-4c4d-acde-3feefe7e4c1d_0
  0ee30dc0424e   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_coredns-6d4b75cb6d-ww4mk_kube-system_b3cdff18-81b7-47c0-9ae9-93f624c20819_0
  bc5e65e561dd   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_kube-proxy-mjrzn_kube-system_61781a67-98fd-4c4d-acde-3feefe7e4c1d_0
  1e96420f4a13   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_storage-provisioner_kube-system_b631381c-bc31-4269-91bb-6dfd90d617f1_0
  7a69fa2d5df4   aebe758cef4c                   "etcd --advertise-cl…"   13 minutes ago   Up 13 minutes             k8s_etcd_etcd-minikube_kube-system_906edd533192a4db2396a938662a5271_0
  cd35426a7fc0   3a5aa3a515f5                   "kube-scheduler --au…"   13 minutes ago   Up 13 minutes             k8s_kube-scheduler_kube-scheduler-minikube_kube-system_2e95d5efbc70e877d20097c03ba4ff89_0
  05eec57345d6   d521dd763e2e                   "kube-apiserver --ad…"   13 minutes ago   Up 13 minutes             k8s_kube-apiserver_kube-apiserver-minikube_kube-system_af8a252bb89a737e9c95199d01283487_0
  dcf0864fc145   586c112956df                   "kube-controller-man…"   13 minutes ago   Up 13 minutes             k8s_kube-controller-manager_kube-controller-manager-minikube_kube-system_76444121a189d8a30add20fb32ab6d4e_0
  ec7bf1d5f39f   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_kube-scheduler-minikube_kube-system_2e95d5efbc70e877d20097c03ba4ff89_0
  54c76d32cd20   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_kube-apiserver-minikube_kube-system_af8a252bb89a737e9c95199d01283487_0
  7719e1c00808   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_etcd-minikube_kube-system_906edd533192a4db2396a938662a5271_0
  8f232ca42e27   k8s.gcr.io/pause:3.6           "/pause"                 13 minutes ago   Up 13 minutes             k8s_POD_kube-controller-manager-minikube_kube-system_76444121a189d8a30add20fb32ab6d4e_0
  ```

- Build the container using docker file, run it and verify from the logs.

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker build -t go-app-normal:latest .
  Sending build context to Docker daemon  1.996MB
  Step 1/8 : FROM golang:latest
  ---> 6d7adf071e0a
  Step 2/8 : WORKDIR /app
  ---> Using cache
  ---> 08602fb114e6
  Step 3/8 : COPY go.mod go.sum ./
  ---> Using cache
  ---> 50826d07248f
  Step 4/8 : RUN go mod download
  ---> Using cache
  ---> d83c0e30b4ed
  Step 5/8 : COPY . .
  ---> 4d9ebc13bfd7
  Step 6/8 : RUN go build -o main .
  ---> Running in 8e6ea83f19fd
  Removing intermediate container 8e6ea83f19fd
  ---> 381e9d41243d
  Step 7/8 : EXPOSE 80
  ---> Running in 2716e3bdaf0b
  Removing intermediate container 2716e3bdaf0b
  ---> 605a0bc97fa1
  Step 8/8 : ENTRYPOINT [ "./main" ]
  ---> Running in c984499a05f2
  Removing intermediate container c984499a05f2
  ---> ae58f0067e61
  Successfully built ae58f0067e61
  Successfully tagged go-app-normal:latest
  SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is recommended to double check and reset permissions for sensitive files and directories.

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker run -d -p 80:80 --name goweb go-app-normal:latest
  10dca664e2c04cc9d4dc84ef5555b6f27baca5c41ef8c1336140f01e2972636a

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker logs -f goweb
  2022/09/13 07:25:12 Web server has started at http://localhost
  ```

- `minikube ip` will give the ip and this ip will be able to show the app in browser. But the size of the image will be too big.

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker image ls
  REPOSITORY                                TAG       IMAGE ID       CREATED         SIZE
  go-app-normal                             latest    ae58f0067e61   2 hours ago     1GB
  ```

- So to avoid this multi-stage build needs to be done, _goditroless image_ will be used and the current image needs to be removed.

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devo Projects/devops-gallery/03-devops-project/go-devo                                                  ps/src/go-microservices (main)ps/src/go-microservices (main)
  $ docker rm -f goweb
  goweb
  ```

- After refering to the [repo](https://github.com/GoogleContainerTools/distroless), modify the docker file accordingly, follow the simillar process mentioned above to expose the port This time the size will be less.

  ```bash
  $ docker image ls
  REPOSITORY                                TAG       IMAGE ID       CREATED         SIZE  
  go-app-ms                                 latest    dedc4e2031c9   2 hours ago     9.33MB
  ```

- Create a _docker-compose.yaml_ file and add the content required to create a docker container with dockerfile. Build the app, run it using the compose file.

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker-compose build
  [+] Building 2.7s (0/2)
  => [internal] load build definition from dockerfile                                                                                      2.6s 
  => [internal] load .dockerignore                                                                                     [+] Building 48.6s (14/14) FINISHED
  => [internal] load build definition from dockerfile  4.0s
  => => transferring dockerfile: 579B                  0.7s
  => [internal] load .dockerignore                     3.4s
  => => transferring context: 65B                      0.9s
  => [internal] load metadata for gcr.io/distroless/s  0.0s
  => [internal] load metadata for docker.io/library/g  0.0s
  => [builder 1/6] FROM docker.io/library/golang:late  9.2s
  => [internal] load build context                     9.6s
  => => transferring context: 1.99MB                   5.5s
  => [stage-1 1/2] FROM gcr.io/distroless/static-debi  1.1s
  => [builder 2/6] WORKDIR /app                        1.3s
  => [builder 3/6] COPY go.mod go.sum ./               1.1s
  => [builder 4/6] RUN go mod download                13.0s
  => [builder 5/6] COPY . .                            1.8s
  => [builder 6/6] RUN go build -o main .             12.2s
  => [stage-1 2/2] COPY --from=builder /app/main .     1.9s
  => exporting to image                                1.9s
  => => exporting layers                               1.3s
  => => writing image sha256:1716a13d7c8307cd10000f1d  0.4s
  => => naming to docker.io/library/go-app-ms:latest   0.2s

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker-compose up -d
  [+] Running 2/2
  - Network go-microservices_web      Created          1.4s
  - Container go-microservices-web-1  Started          5.7s

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ docker-compose ps
  NAME                     COMMAND             SERVICE             STATUS              PORTS
  go-microservices-web-1   "/main"             web                 restarting
  ```

- Make a folder to hold all the kubernetes deployment objects, create abd add content to _deployment.yml and service.yml_, apply those files, verify the deployment

  ```bash
  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ kubectl apply -f deployments/
  deployment.apps/web-deployment created
  service/web-service created

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ kubectl get pods
  NAME                              READY   STATUS             RESTARTS      AGE
  web-deployment-594c4bf755-5zxz4   0/1     CrashLoopBackOff   3 (28s ago)   76s
  web-deployment-594c4bf755-ltdgn   0/1     CrashLoopBackOff   3 (28s ago)   76s
  web-deployment-594c4bf755-nrvhr   0/1     CrashLoopBackOff   3 (26s ago)   76s

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ kubectl get svc
  NAME          TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
  kubernetes    ClusterIP   10.96.0.1      <none>        443/TCP        5h16m
  web-service   NodePort    10.102.10.57   <none>        80:32700/TCP   84s

  measutosh@crispy MINGW64 ~/Documents/VSCode/DevOps Projects/devops-gallery/03-devops-project/go-devops/src/go-microservices (main)
  $ minikube service web-service
  |-----------|-------------|-------------|---------------------------|
  | NAMESPACE |    NAME     | TARGET PORT |            URL            |
  |-----------|-------------|-------------|---------------------------|
  | default   | web-service |          80 | http://192.168.49.2:32700 |
  |-----------|-------------|-------------|---------------------------|
  🏃  Starting tunnel for service web-service.
  |-----------|-------------|-------------|------------------------|
  | NAMESPACE |    NAME     | TARGET PORT |          URL           |
  |-----------|-------------|-------------|------------------------|
  | default   | web-service |             | http://127.0.0.1:64660 |
  |-----------|-------------|-------------|------------------------|
  🎉  Opening service default/web-service in default browser...
  ❗  Because you are using a Docker driver on windows, the terminal needs to be open to run it.
  ```



