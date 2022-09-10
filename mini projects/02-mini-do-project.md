# Prometheus & Grafana Monitoring On Kubernets using Helm


[Live Demo](http://expanding-cards.azurewebsites.net)<br>
Here is the project which I have deployed using Azure App services - [Expanding Cards](https://github.com/measutosh/frontend-gallery/tree/main/expanding-cards)

### Possible ways to configure Grafana & Prometheus in Kubernets Cluster

- Manually deploying all deployments, services, ConfigMaps, secrets and configuration files.
- More efficient way is using Helm Chart to deploy Prometheus Operator.

Steps followed :-

- helm commands used to install prometheus

  ```bash
  helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
  helm repo update
  helm install prometheus prometheus-community/prometheus
  ```

- The output should be something like this

  ```powershell
  PS C:\Users\measutosh> helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
  "prometheus-community" has been added to your repositories

  PS C:\Users\measutosh> helm repo update
  Hang tight while we grab the latest from your chart repositories...
  ...Successfully got an update from the "jenkins" chart repository
  ...Successfully got an update from the "prometheus-community" chart repository
  Update Complete. ⎈Happy Helming!⎈

  PS C:\Users\measutosh> helm install prmt prometheus-community/prometheus      
  NAME: prmt
  LAST DEPLOYED: Sat Sep 10 11:13:58 2022
  NAMESPACE: default
  STATUS: deployed
  REVISION: 1
  TEST SUITE: None
  NOTES:
  The Prometheus server can be accessed via port 80 on the following DNS name from within your cluster:
  prmt-prometheus-server.default.svc.cluster.local

  Get the Prometheus server URL by running these commands in the same shell:
    export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
    kubectl --namespace default port-forward $POD_NAME 9090


  The Prometheus alertmanager can be accessed via port 80 on the following DNS name from within your cluster:
  prmt-prometheus-alertmanager.default.svc.cluster.local


  Get the Alertmanager URL by running these commands in the same shell:
    export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=alertmanager" -o jsonpath="{.items[0].metadata.name}")       
    kubectl --namespace default port-forward $POD_NAME 9093
  #################################################################################
  ######   WARNING: Pod Security Policy has been moved to a global property.  #####
  ######            use .Values.podSecurityPolicy.enabled with pod-based      #####
  ######            annotations                                               #####
  ######            (e.g. .Values.nodeExporter.podSecurityPolicy.annotations) #####
  #################################################################################


  The Prometheus PushGateway can be accessed via port 9091 on the following DNS name from within your cluster:
  prmt-prometheus-pushgateway.default.svc.cluster.local


  Get the PushGateway URL by running these commands in the same shell:
    export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=pushgateway" -o jsonpath="{.items[0].metadata.name}")        
    kubectl --namespace default port-forward $POD_NAME 9091

  For more information on running Prometheus, visit:
  https://prometheus.io/
  ```

- Verify by checking all the details

  ```powershell
  PS C:\Users\measutosh> kubectl get all
  NAME                                                READY   STATUS             RESTARTS         AGE
  pod/prmt-kube-state-metrics-d6bd4f5d9-kv5sm         1/1     Running            0                13m
  pod/prmt-prometheus-pushgateway-7cf88c767f-xwfqn    1/1     Running            0                13m
  pod/prmt-prometheus-server-868cb86b69-xvhrf         0/2     CrashLoopBackOff   14 (2m52s ago)   13m
  pod/prmt-prometheus-alertmanager-565d646757-9pdrm   0/2     CrashLoopBackOff   14 (2m37s ago)   13m
  pod/prmt-prometheus-node-exporter-xj6sv             1/1     Running            0                13m

  NAME                                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
  service/kubernetes                      ClusterIP   10.43.0.1       <none>        443/TCP    86m
  service/prmt-prometheus-server          ClusterIP   10.43.33.21     <none>        80/TCP     13m
  service/prmt-prometheus-node-exporter   ClusterIP   10.43.154.177   <none>        9100/TCP   13m
  service/prmt-prometheus-pushgateway     ClusterIP   10.43.83.146    <none>        9091/TCP   13m
  service/prmt-kube-state-metrics         ClusterIP   10.43.183.58    <none>        8080/TCP   13m
  service/prmt-prometheus-alertmanager    ClusterIP   10.43.196.36    <none>        80/TCP     13m

  NAME                                           DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
  daemonset.apps/prmt-prometheus-node-exporter   1         1         1       1            1           <none>          13m

  NAME                                           READY   UP-TO-DATE   AVAILABLE   AGE
  deployment.apps/prmt-kube-state-metrics        1/1     1            1           13m
  deployment.apps/prmt-prometheus-pushgateway    1/1     1            1           13m      
  deployment.apps/prmt-prometheus-alertmanager   0/1     1            0           13m      
  deployment.apps/prmt-prometheus-server         0/1     1            0           13m      

  NAME                                                      DESIRED   CURRENT   READY   AGE
  replicaset.apps/prmt-prometheus-server-868cb86b69         1         1         0       13m
  replicaset.apps/prmt-prometheus-alertmanager-565d646757   1         1         0       13m
  replicaset.apps/prmt-kube-state-metrics-d6bd4f5d9         1         1         1       13m
  replicaset.apps/prmt-prometheus-pushgateway-7cf88c767f    1         1         1       13m
  ```

- Check all the services to verify the new service

  ```powershell
  PS C:\Users\measutosh> kubectl get svc
  NAME                            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
  kubernetes                      ClusterIP   10.43.0.1       <none>        443/TCP    89m
  prmt-prometheus-server          ClusterIP   10.43.33.21     <none>        80/TCP     16m
  prmt-prometheus-node-exporter   ClusterIP   10.43.154.177   <none>        9100/TCP   16m
  prmt-prometheus-pushgateway     ClusterIP   10.43.83.146    <none>        9091/TCP   16m
  prmt-kube-state-metrics         ClusterIP   10.43.183.58    <none>        8080/TCP   16m
  prmt-prometheus-alertmanager    ClusterIP   10.43.196.36    <none>        80/TCP     16m
  ```
  
- By default the service is _ClusterIP_, to access Prometheus from outside, a _NodePort_ needs to be created

  ```powershell
  PS C:\Users\measutosh> kubectl expose service prmt-prometheus-server --type=NodePort --target-port=9090 --name=prmt-prometheus-server-ext
  service/prmt-prometheus-server-ext exposed
  ```
