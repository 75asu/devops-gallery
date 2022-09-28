# End To END CI/CD Pipeline, Jenkins Instance On Minikube With Dynamic Agents, Dockerization Of App
<br>
## This is how it looks once setup is done completely
<br>
<!-- ![Video Demo](./04-devops-project-ss.mp4) -->
<video src='./04-devops-project-ss.mp4' width=180/>
<br>

***
> This whole project has been done in GitLap private repo, to showcase everything I have stored the soruce files here.
***

## This CI/CD pipeline does the following things in order

- **Triggers the pipeline when any commit goes to any of the 3 branches of a GitLab private repository.**
- **The pipeline gets triggered inside Jenkins which is hosted on a minikube cluster.**
- **Just after the trigger is set off, the following things happen in order :-**
	- **Checkout of the repo into a Dynamic slave node of jenkins created in minikube.**
	- **A docker instance gets kicked off, build an image of the React app codebase using the dockerile present in the repo.**
	- **The docker image is being pushed to Docker Hub.**
	- **Whenever this job gets failed or succeed, a mail goes to selected persons.**
	- **Once the job is finished the dynamic agent also gets decommissioned.**


### Setup a Jenkins Instance in minikube cluster

- For this, these steps have been followed
	- Create K8s Cluster in your local system using minikube.
	- Create a new context and switch to it using the kubeconfig file, here I have used the default config.
    - Command to see the `kubeconfig` file - 

        ```bash
        measutosh@crispy MINGW64 ~
        $ kubectl config view
        apiVersion: v1
        clusters:
        - cluster:
            certificate-authority: C:\Users\measutosh\.minikube\ca.crt
            extensions:
            - extension:
                last-update: Wed, 14 Sep 2022 10:06:09 GMT
                provider: minikube.sigs.k8s.io
                version: v1.26.1
            name: cluster_info
            server: https://127.0.0.1:57769
        name: minikube
        contexts:
        - context:
            cluster: minikube
            extensions:
            - extension:
                last-update: Wed, 14 Sep 2022 10:06:09 GMT
                provider: minikube.sigs.k8s.io
                version: v1.26.1
            name: context_info
            namespace: default
            user: minikube
        name: minikube
        current-context: minikube
        kind: Config
        preferences: {}
        users:
        - name: minikube
        user:
            client-certificate: C:\Users\measutosh\.minikube\profiles\minikube\client.crt
            client-key: C:\Users\measutosh\.minikube\profiles\minikube\client.key
        ```
        
    - Used the following command to create a new context, to switch into that context

        ```bash
        measutosh@crispy MINGW64 ~
        $ kubectl config current-context
        minikube

        measutosh@crispy MINGW64 ~
        $ kubectl config use-context devops-test
        Switched to context "devops-test".
        ```

- Deploy the Jenkins server on your namespace and expose Jenkins.The following steps have been followed to this.
	- Create a Namespace
	- Create a service account with Kubernetes admin permissions.
	- Create local persistent volume for persistent Jenkins data on Pod restarts.
	- Create a deployment YAML and deploy it.
	- Create a service YAML and deploy it.
	- Access the Jenkins application on a Node Port.
	```bash

	measutosh@crispy MINGW64 ~
	$ kubectl create namespace devops-tools
	namespace/devops-tools created

	measutosh@crispy MINGW64 ~
	$ kubectl apply -f serviceAccount.yml
	clusterrole.rbac.authorization.k8s.io/jenkins-admin created
	serviceaccount/jenkins-admin created
	clusterrolebinding.rbac.authorization.k8s.io/jenkins-admin created

	measutosh@crispy MINGW64 ~
	$ kubectl create -f jenkins-volume.yaml
	storageclass.storage.k8s.io/local-storage created
	persistentvolume/jenkins-pv-volume created
	persistentvolumeclaim/jenkins-pv-claim created

	measutosh@crispy MINGW64 ~
	$ kubectl apply -f jenkins-deployment.yaml
	deployment.apps/jenkins created

	measutosh@crispy MINGW64 ~
	$ kubectl get deployments -n devops-tools
	NAME      READY   UP-TO-DATE   AVAILABLE   AGE
	jenkins   0/1     1            0           27s


	measutosh@crispy MINGW64 ~
	$ kubectl  describe deployments --namespace=devops-tools
	Name:                   jenkins
	Namespace:              devops-tools
	CreationTimestamp:      Thu, 15 Sep 2022 07:00:34 +0000
	Labels:                 <none>
	Annotations:            deployment.kubernetes.io/revision: 1
	Selector:               app=jenkins-server
	Replicas:               1 desired | 1 updated | 1 total | 0 available | 1 unavailable
	StrategyType:           RollingUpdate
	MinReadySeconds:        0
	RollingUpdateStrategy:  25% max unavailable, 25% max surge
	Pod Template:
	Labels:           app=jenkins-server
	Service Account:  jenkins-admin
	Containers:
	jenkins:
		Image:       jenkins/jenkins:lts
		Ports:       8080/TCP, 50000/TCP
		Host Ports:  0/TCP, 0/TCP
		Limits:
		cpu:     1
		memory:  2Gi
		Requests:
		cpu:        500m
		memory:     500Mi
		Liveness:     http-get http://:8080/login delay=90s timeout=5s period=10s #success=1 #failure=5
		Readiness:    http-get http://:8080/login delay=60s timeout=5s period=10s #success=1 #failure=3
		Environment:  <none>
		Mounts:
		/var/jenkins_home from jenkins-data (rw)
	Volumes:
	jenkins-data:
		Type:       PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
		ClaimName:  jenkins-pv-claim
		ReadOnly:   false
	Conditions:
	Type           Status  Reason
	----           ------  ------
	Available      False   MinimumReplicasUnavailable
	Progressing    True    ReplicaSetUpdated
	OldReplicaSets:  <none>
	NewReplicaSet:   jenkins-fd5fdf49f (1/1 replicas created)
	Events:
	Type    Reason             Age   From                   Message
	----    ------             ----  ----                   -------
	Normal  ScalingReplicaSet  60s   deployment-controller  Scaled up replica set jenkins-fd5fdf49f to 1
	
	
	measutosh@crispy MINGW64 ~
	$ kubectl apply -f jenkins-service.yaml
	service/jenkins-service created


	measutosh@crispy MINGW64 ~
	$ minikube service jenkins-service -n devops-tools
	|--------------|-----------------|-------------|-------------------------|
	|  NAMESPACE   |      NAME       | TARGET PORT |           URL           |
	|--------------|-----------------|-------------|-------------------------|
	| devops-tools | jenkins-service |        8080 | http://172.17.0.2:32000 |
	|--------------|-----------------|-------------|-------------------------|
	🏃  Starting tunnel for service jenkins-service.
	|--------------|-----------------|-------------|------------------------|
	|  NAMESPACE   |      NAME       | TARGET PORT |          URL           |
	|--------------|-----------------|-------------|------------------------|
	| devops-tools | jenkins-service |             | http://127.0.0.1:63351 |
	|--------------|-----------------|-------------|------------------------|
	🎉  Opening service devops-tools/jenkins-service in default browser...
	❗  Because you are using a Docker driver on windows, the terminal needs to be open to run it.

	measutosh@crispy MINGW64 ~
	$ kubectl get pods --namespace=devops-tools
	NAME                      READY   STATUS    RESTARTS   AGE
	jenkins-fd5fdf49f-9wwlb   1/1     Running   0          14m
	```

- To get the password, opened the log of the pod

	```bash
	measutosh@crispy MINGW64 ~
	$ kubectl logs jenkins-fd5fdf49f-9wwlb --namespace=devops-tools
	Running from: /usr/share/jenkins/jenkins.war
	webroot: EnvVars.masterEnvVars.get("JENKINS_HOME")
	2022-09-15 07:00:41.880+0000 [id=1]     INFO    winstone.Logger#logInternal: Beginning extraction from war file
	2022-09-15 07:00:45.560+0000 [id=1]     WARNING o.e.j.s.handler.ContextHandler#setContextPath: Empty contextPath
	2022-09-15 07:00:46.028+0000 [id=1]     INFO    org.eclipse.jetty.server.Server#doStart: jetty-10.0.11; built: 2022-06-21T21:12:44.640Z; git: d988aa016e0bb2de6fba84c1659049c72eae3e32; jvm 11.0.16.1+1
	2022-09-15 07:00:47.772+0000 [id=1]     INFO    o.e.j.w.StandardDescriptorProcessor#visitServlet: NO JSP Support for /, did not find org.eclipse.jetty.jsp.JettyJspServlet
	2022-09-15 07:00:48.321+0000 [id=1]     INFO    o.e.j.s.s.DefaultSessionIdManager#doStart: Session workerName=node0
	2022-09-15 07:00:51.350+0000 [id=1]     INFO    hudson.WebAppMain#contextInitialized: Jenkins home directory: /var/jenkins_home found at: EnvVars.masterEnvVars.get("JENKINS_HOME")
	2022-09-15 07:00:52.638+0000 [id=1]     INFO    o.e.j.s.handler.ContextHandler#doStart: Started w.@4d8286c4{Jenkins v2.361.1,/,file:///var/jenkins_home/war/,AVAILABLE}{/var/jenkins_home/war}
	2022-09-15 07:00:52.732+0000 [id=1]     INFO    o.e.j.server.AbstractConnector#doStart: Started ServerConnector@e84a8e1{HTTP/1.1, (http/1.1)}{0.0.0.0:8080}
	2022-09-15 07:00:52.840+0000 [id=1]     INFO    org.eclipse.jetty.server.Server#doStart: Started Server@32c8e539{STARTING}[10.0.11,sto=0] @14739ms2022-09-15 07:00:52.852+0000 [id=23]    INFO    winstone.Logger#logInternal: Winstone Servlet Engine running: controlPort=disabled
	2022-09-15 07:00:54.223+0000 [id=28]    INFO    jenkins.InitReactorRunner$1#onAttained: Started initialization
	2022-09-15 07:00:54.335+0000 [id=29]    INFO    jenkins.InitReactorRunner$1#onAttained: Listed all plugins
	2022-09-15 07:00:58.948+0000 [id=28]    INFO    jenkins.InitReactorRunner$1#onAttained: Prepared all plugins
	2022-09-15 07:00:58.976+0000 [id=29]    INFO    jenkins.InitReactorRunner$1#onAttained: Started all plugins
	2022-09-15 07:00:59.146+0000 [id=29]    INFO    jenkins.InitReactorRunner$1#onAttained: Augmented all extensions
	2022-09-15 07:00:59.430+0000 [id=29]    INFO    jenkins.InitReactorRunner$1#onAttained: System config loaded
	2022-09-15 07:00:59.431+0000 [id=29]    INFO    jenkins.InitReactorRunner$1#onAttained: System config adapted
	2022-09-15 07:00:59.432+0000 [id=29]    INFO    jenkins.InitReactorRunner$1#onAttained: Loaded all jobs
	2022-09-15 07:01:00.524+0000 [id=29]    INFO    jenkins.InitReactorRunner$1#onAttained: Configuration for all jobs updated
	WARNING: An illegal reflective access operation has occurred
	WARNING: Illegal reflective access by org.codehaus.groovy.vmplugin.v7.Java7$1 (file:/var/jenkins_home/war/WEB-INF/lib/groovy-all-2.4.21.jar) to constructor java.lang.invoke.MethodHandles$Lookup(java.lang.Class,int)
	WARNING: Please consider reporting this to the maintainers of org.codehaus.groovy.vmplugin.v7.Java7$1
	WARNING: Use --illegal-access=warn to enable warnings of further illegal reflective access operations
	WARNING: All illegal access operations will be denied in a future release
	2022-09-15 07:01:03.629+0000 [id=42]    INFO    hudson.model.AsyncPeriodicWork#lambda$doRun$1: Started Download metadata
	2022-09-15 07:01:03.649+0000 [id=42]    INFO    hudson.util.Retrier#start: Attempt #1 to do the action check updates server
	2022-09-15 07:01:04.931+0000 [id=29]    INFO    jenkins.install.SetupWizard#init:

	*************************************************************
	*************************************************************
	*************************************************************

	Jenkins initial setup is required. An admin user has been created and a password generated.
	Please use the following password to proceed to installation:

	831ab0cbb89e460f81bbe77237ac3edb

	This may also be found at: /var/jenkins_home/secrets/initialAdminPassword

	*************************************************************
	*************************************************************
	*************************************************************

	2022-09-15 07:02:49.515+0000 [id=28]    INFO    jenkins.InitReactorRunner$1#onAttained: Completed initialization
	2022-09-15 07:02:49.604+0000 [id=22]    INFO    hudson.lifecycle.Lifecycle#onReady: Jenkins is fully up and running
	2022-09-15 07:03:01.583+0000 [id=42]    INFO    h.m.DownloadService$Downloadable#load: Obtained the updated data file for hudson.tasks.Maven.MavenInstaller
	2022-09-15 07:03:01.584+0000 [id=42]    INFO    hudson.util.Retrier#start: Performed the action check updates server successfully at the attempt #1
	2022-09-15 07:03:01.599+0000 [id=42]    INFO    hudson.model.AsyncPeriodicWork#lambda$doRun$1: Finished Download metadata. 117,969 ms
	```



## Setup a multibranch CI/CD Pipeline on Jenkins that will have 3 stages (This pipeline will be triggered automatically when someone pushes to SCM repository.

- For this, the following steps have been followed 

  - In the cloned repo, created 2 other branches - **dev**, **test**
  - In Jenkins created a pipeline with multibranch option. In configure page - added github as branch source, added credentials.
  - Added my gitlab credentials on the folder level of the Jenkins multibranch pipeline.
  - Addded build branch as **master branch** in the regular expression.
  - Once the scan suceeded, laster added the `^main|dev|test.*` regex to **discover branches** option.


  ```bash
    Started
	[Fri Sep 16 10:06:28 UTC 2022] Starting branch indexing...
	> git --version # timeout=10
	> git --version # 'git version 2.30.2'
	using GIT_ASKPASS to set credentials 
	> git ls-remote --symref -- https://gitlab.com/measutosh/assignment_app.git # timeout=10
	Creating git repository in /var/jenkins_home/caches/git-62dda451cd517f388c57785609d6d9de
	> git init /var/jenkins_home/caches/git-62dda451cd517f388c57785609d6d9de # timeout=10
	Setting origin to https://gitlab.com/measutosh/assignment_app.git
	> git config remote.origin.url https://gitlab.com/measutosh/assignment_app.git # timeout=10
	Fetching & pruning origin...
	Listing remote references...
	> git config --get remote.origin.url # timeout=10
	> git --version # timeout=10
	> git --version # 'git version 2.30.2'
	using GIT_ASKPASS to set credentials 
	> git ls-remote -h -- https://gitlab.com/measutosh/assignment_app.git # timeout=10
	Fetching upstream changes from origin
	> git config --get remote.origin.url # timeout=10
	using GIT_ASKPASS to set credentials 
	> git fetch --tags --force --progress --prune -- origin +refs/heads/*:refs/remotes/origin/* # timeout=10
	Checking branches...
	Checking branch main
		‘Jenkinsfile’ not found
		Does not meet criteria
	Checking branch test
		‘Jenkinsfile’ not found
		Does not meet criteria
	Checking branch dev
		‘Jenkinsfile’ not found
		Does not meet criteria
	Processed 3 branches
	[Fri Sep 16 10:06:50 UTC 2022] Finished branch indexing. Indexing took 22 sec
	Finished: SUCCESS
	```

  - Will be using push notifications in SCM to trigger the pipeline in Jenkins. Chose the **Scan Multibranch Pipeline Triggers**, added the interval time as **1 minute**. 
  - The code will be pulled from the SCM repository to Jenkins .
  - The Docker image will be created in the second stage and pushed to a private repository. (used Docker hub)
  - Went through the code, it's a reactjs/nodejs app, so the **Dockerfile** should be written accordingly.

	```bash
	FROM node:lts-alpine
	ENV NODE_ENV development
	# Add a work directory
	WORKDIR /app
	# Cache and Install dependencies
	COPY app/package.json .
	COPY app/yarn.lock .
	RUN yarn install
	# Copy app files
	COPY . .
	# Expose port
	EXPOSE 3000
	# Start the app
	CMD [ "yarn", "start" ]
	```

  - Make Jenkins setup in such a way that whenever pipeline runs it will dynamically provision  containers as Jenkins worker nodes and the Jenkins worker must delete automatically when the job is finished.(covers concepts like:  `Docker in Docker with Jenkins Pod on Kubernetes, Jenkins Dynamic Agent`)
  - Installed the plugins - **Kubernetes Plugin, Docker Plugin, Git Plugin, Node and Label Parameter Plugin**, followed these steps to connect the minikube cluster with jenkins - **Manage Jenkins -> Manage Node and Cloud -> Configure Clouds -> Add a new cloud -> Kubernetes -> Kubernetes Cloud Details.**, but faced an error.

- The URL which I provided was wrong, so to get the IPs correct this time used the commands

	```bash	
	measutosh@crispy MINGW64 ~
	$ minikube service jenkins-service
	|-----------|-----------------|-------------|---------------------------|
	| NAMESPACE |      NAME       | TARGET PORT |            URL            |
	|-----------|-----------------|-------------|---------------------------|
	| default   | jenkins-service |        8080 | http://192.168.49.2:32000 |
	|-----------|-----------------|-------------|---------------------------|
	🏃  Starting tunnel for service jenkins-service.
	|-----------|-----------------|-------------|------------------------|
	| NAMESPACE |      NAME       | TARGET PORT |          URL           |
	|-----------|-----------------|-------------|------------------------|
	| default   | jenkins-service |             | http://127.0.0.1:50267 |
	|-----------|-----------------|-------------|------------------------|
	🎉  Opening service default/jenkins-service in default browser...
	❗  Because you are using a Docker driver on windows, the terminal needs to be open to run it.

	measutosh@crispy MINGW64 ~
	$ kubectl cluster-info
	Kubernetes control plane is running at https://127.0.0.1:50162
	CoreDNS is running at https://127.0.0.1:50162/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
	```

	- URLs used while configuring Jenkins as dynamic agent were
		- kuberntes URL - `https.kubernetes.default`
		- jenkins URL - `http://192.168.49.2:32000`
		- jenkins tunnel url - `http://127.0.0.1:50267`
	
  - Send an email notification or slack message if the build succeeds or fails.

  - Docker image build and docker image push to docker hub have worked fine using below jenkinsfile code
	```groovy
	pipeline {
		agent {
			kubernetes {
				yaml '''
			apiVersion: v1
			kind: Pod
			spec:
			containers:
			- name: docker
				image: docker:latest
				command:
				- cat
				tty: true
				volumeMounts:
				- mountPath: /var/run/docker.sock
				name: docker-sock
			volumes:
			- name: docker-sock
				hostPath:
				path: /var/run/docker.sock    
			'''
			}
		}
		stages {
			stage('Checkout from SCM') {
				steps {
					checkout([$class: 'GitSCM', branches: [[name: '**']], extensions: [], userRemoteConfigs: [[credentialsId: 'gitlab-cred', url: 'https://gitlab.com/measutosh/assignment_app.git']]])
					sh 'ls -lart'
				}
			}
			stage('Build image, Login to Docker Hub, Push, Logout ') {
				steps {
					container('docker') {
						sh 'docker build -t measutosh/devops-test:v1.0 .'
						sh 'docker login -u measutosh -p Crimsy!23 docker.io'
						sh 'docker push measutosh/devops-test:v1.0'
						sh 'docker logout'
					}
				}
			}
		}
	}
	```


