# Stil Work In Progress

## Steps followed

- Create a new kubernetes cluster in AKS with 2 nodes and download the _kubeconfig_ file.
- Redirect the kubeconfig file content to _.kube_ location.

    ```bash
    cat end2end-jenkins-pipeline-cluster-config-file.yml > ~/.kube/config
    ```

- Download and installed kubectl CLI, helm CLI using chocolaty to setup helm.

    ```bash
    choco install kubernetes-cli
    choco install kubernetes-helm
    ```

### Jenkins Setup

***

- To install jenkins go to _artifacthub_ and use the commands to install jenkins using helm CLI

    ```bash
    helm repo add jenkins https://charts.jenkins.io
    helm repo update
    helm pull jenkins/jenkins --untar
    ```

- Open the jenkins folder from the download folder and replace the _servieType: ClusterIP_ to _servieType: LoadBalancer_
- The reason behind this is that - here we don't have any ingress controller. If we has ingress controller we would have make this as an internal service and use ingress resources to redirect the traffic. but in this case direct expose through loadbalancer has been done.Making the _servicePort: 80_ just to directly use LoadBalancer IP to access it because if we had used 8080 then LoadBalancer IP would have to be 8080 which is an extra port that is not needed.
- Now running the command to install jenkins on the cluster. This will check if new release has been there for chart or not. If yeas then it will upgrade then install or else only install.

    ```bash
    helm upgrade --install jenkins jenkins/
    ```

- Run few command to see some extra info about the instance, external IP of jenkins. Once external IP shows up, click it to see the jenkins instance on browser.

    ```bash
    kubectl get sts
    kubectl get svc
    ```

- To get the password of the jenkins instance, pick up the pod name, log into it and _cat_ the password file from it's suggested location form helm result that you may have got by running the previous commands.

    ```bash
    kubectl get po
    kubectl exec -it jenkins-0 --bash
    cat /run/secrets/additional/chart-admin-password && echo
    ```

- Go to _manage jenkins -> manage nodes and clouds_ in jenkins instance to see the node attached to it by the cluster to jenkins. In 
_manage jenkins -> manage nodes and clouds -> configure clouds_, it will be visibble that kubernetes is already setup there.