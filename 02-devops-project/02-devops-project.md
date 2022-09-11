# Prometheus & Grafana Monitoring On Kubernets using Helm

## This is how it looks once setup is done completely

![Video Demo](./02-devops-project-ss.webm)

## Possible ways to configure Grafana & Prometheus in Kubernets Cluster

- Manually deploying all deployments, services, ConfigMaps, secrets and configuration files.
- More efficient way is using Helm Chart to deploy Prometheus Operator.


> Steps followed :-


### Prometheus Setup

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
  LAST DEPLOYED: Sun Sep 11 13:30:46 2022
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
  NAME                                                READY   STATUS              RESTARTS   AGE
  pod/prmt-kube-state-metrics-d6bd4f5d9-l2g47         0/1     ContainerCreating   0          26s
  pod/prmt-prometheus-alertmanager-565d646757-qrfzt   0/2     ContainerCreating   0          26s
  pod/prmt-prometheus-node-exporter-h6gj7             0/1     ContainerCreating   0          26s
  pod/prmt-prometheus-pushgateway-7cf88c767f-6fcxp    0/1     ContainerCreating   0          26s

  service/kubernetes                      ClusterIP   10.96.0.1       <none>        443/TCP    23m
  service/prmt-kube-state-metrics         ClusterIP   10.102.77.21    <none>        8080/TCP   27s
  service/prmt-prometheus-alertmanager    ClusterIP   10.106.143.96   <none>        80/TCP     26s
  service/prmt-prometheus-node-exporter   ClusterIP   10.97.149.109   <none>        9100/TCP   27s
  service/prmt-prometheus-pushgateway     ClusterIP   10.103.35.123   <none>        9091/TCP   27s
  service/prmt-prometheus-server          ClusterIP   10.110.91.110   <none>        80/TCP     27s

  NAME                                           DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE

  deployment.apps/prmt-kube-state-metrics        0/1     1            0           26s
  deployment.apps/prmt-prometheus-alertmanager   0/1     1            0           26s
  deployment.apps/prmt-prometheus-pushgateway    0/1     1            0           26s
  deployment.apps/prmt-prometheus-server         0/1     1            0           26s

  NAME                                                      DESIRED   CURRENT   READY   AGE
  replicaset.apps/prmt-kube-state-metrics-d6bd4f5d9         1         1         0       26s
  replicaset.apps/prmt-prometheus-alertmanager-565d646757   1         1         0       26s
  replicaset.apps/prmt-prometheus-pushgateway-7cf88c767f    1         1         0       26s
  replicaset.apps/prmt-prometheus-server-868cb86b69         1         1         0       26s
  ```

- Check all the services to verify the new service

  ```powershell
  PS C:\Users\measutosh> kubectl get svc
  NAME                            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
  kubernetes                      ClusterIP   10.96.0.1       <none>        443/TCP    25m
  prmt-prometheus-alertmanager    ClusterIP   10.106.143.96   <none>        80/TCP     2m18s
  prmt-prometheus-node-exporter   ClusterIP   10.97.149.109   <none>        9100/TCP   2m19s
  prmt-prometheus-pushgateway     ClusterIP   10.103.35.123   <none>        9091/TCP   2m19s
  prmt-prometheus-server          ClusterIP   10.110.91.110   <none>        80/TCP     2m19s  
  ```
  
- By default the service is _ClusterIP_, to access Prometheus from outside, a _NodePort_ needs to be created

  ```powershell
  PS C:\Users\measutosh> kubectl expose service prmt-prometheus-server --type=NodePort --target-port=9090 --name=prmt-prometheus-server-ext
  service/prmt-prometheus-server-ext exposed
  ```

- Verify the creation of _NOdePort_

  ````powershell
  PS C:\Users\measutosh> kubectl get svc
  NAME                            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
  kubernetes                      ClusterIP   10.96.0.1       <none>        443/TCP        26m
  prmt-kube-state-metrics         ClusterIP   10.102.77.21    <none>        8080/TCP       2m53s
  prmt-prometheus-alertmanager    ClusterIP   10.106.143.96   <none>        80/TCP         2m52s
  prmt-prometheus-node-exporter   ClusterIP   10.97.149.109   <none>        9100/TCP       2m53s
  prmt-prometheus-pushgateway     ClusterIP   10.103.35.123   <none>        9091/TCP       2m53s
  prmt-prometheus-server          ClusterIP   10.110.91.110   <none>        80/TCP         2m53s
  prmt-prometheus-server-ext      NodePort    10.96.78.135    <none>        80:31735/TCP   19s
  ```

- Get the URL of exposed prometheus server instance

  ```powershell
  PS C:\Users\measutosh> minikube service prmt-prometheus-server-ext
  |-----------|----------------------------|-------------|---------------------------|
  | NAMESPACE |            NAME            | TARGET PORT |            URL            |
  |-----------|----------------------------|-------------|---------------------------|
  | default   | prmt-prometheus-server-ext |          80 | http://192.168.49.2:31735 |
  |-----------|----------------------------|-------------|---------------------------|
  🏃  Starting tunnel for service prmt-prometheus-server-ext.
  |-----------|----------------------------|-------------|------------------------|
  | NAMESPACE |            NAME            | TARGET PORT |          URL           |
  |-----------|----------------------------|-------------|------------------------|
  | default   | prmt-prometheus-server-ext |             | http://127.0.0.1:61679 |
  |-----------|----------------------------|-------------|------------------------|
  🎉  Opening service default/prmt-prometheus-server-ext in default browser...
  ❗  Because you are using a Docker driver on windows, the terminal needs to be open to run it.
  ```

- In the UI, under the _Status_ many things can be observed and under the _Graphs_ from main page if you will search _node_, the many data would have started getting captured by Prometheus, now we need to provide these data to Grafana.

***

### Grafana Setup

- Using helm, installing Grafaa

  ```powershell
  helm repo add grafana https://grafana.github.io/helm-charts
  helm repo update
  helm install my-release grafana/grafana
  ```

- The output should look something like this

  ```powershell
  PS C:\Users\measutosh> helm repo add grafana https://grafana.github.io/helm-charts
  "grafana" has been added to your repositories
  PS C:\Users\measutosh> helm repo update
  Hang tight while we grab the latest from your chart repositories...
  ...Successfully got an update from the "jenkins" chart repository
  ...Successfully got an update from the "grafana" chart repository
  ...Successfully got an update from the "prometheus-community" chart repository
  Update Complete. ⎈Happy Helming!⎈
  PS C:\Users\measutosh> helm install grafana grafana/grafana
  W0911 14:05:11.793392    4372 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
  W0911 14:05:11.821015    4372 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
  W0911 14:05:12.321678    4372 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
  W0911 14:05:12.333211    4372 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
  NAME: grafana
  LAST DEPLOYED: Sun Sep 11 14:05:10 2022
  NAMESPACE: default
  STATUS: deployed
  REVISION: 1
  NOTES:
  1. Get your 'admin' user password by running:

    kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

  2. The Grafana server can be accessed via port 80 on the following DNS name from within your cluster:

    grafana.default.svc.cluster.local

    Get the Grafana URL to visit by running these commands in the same shell:

      export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}")
      kubectl --namespace default port-forward $POD_NAME 3000

  3. Login with the password from step 1 and the username: admin
  #################################################################################
  ######   WARNING: Persistence is disabled!!! You will lose your data when   #####
  ######            the Grafana pod is terminated.                            #####
  #################################################################################
  ```

- Check all the releases by helm

  ```powershell
  PS C:\Users\measutosh> helm list
  NAME    NAMESPACE       REVISION        UPDATED                                 STATUS          CHART                   APP VERSION
  grafana default         1               2022-09-11 14:05:10.5934876 +0000 UTC   deployed        grafana-6.37.3          9.1.4
  prmt    default         1               2022-09-11 13:30:46.3147468 +0000 UTC   deployed        prometheus-15.12.0      2.36.2
  ```

- Check all the pods for verification

  ```powershell
  PS C:\Users\measutosh> kubectl get all
  NAME                                                READY   STATUS    RESTARTS   AGE
  pod/grafana-68ccc465d8-zjzb8                        1/1     Running   0          4m23s
  pod/prmt-kube-state-metrics-d6bd4f5d9-l2g47         1/1     Running   0          38m
  pod/prmt-prometheus-alertmanager-565d646757-qrfzt   2/2     Running   0          38m
  pod/prmt-prometheus-node-exporter-h6gj7             1/1     Running   0          38m
  pod/prmt-prometheus-pushgateway-7cf88c767f-6fcxp    1/1     Running   0          38m
  pod/prmt-prometheus-server-868cb86b69-twtsn         2/2     Running   0          38m

  NAME                                    TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
  service/grafana                         ClusterIP   10.111.124.200   <none>        80/TCP         4m24s
  service/kubernetes                      ClusterIP   10.96.0.1        <none>        443/TCP        62m
  service/prmt-kube-state-metrics         ClusterIP   10.102.77.21     <none>        8080/TCP       38m
  service/prmt-prometheus-alertmanager    ClusterIP   10.106.143.96    <none>        80/TCP         38m
  service/prmt-prometheus-node-exporter   ClusterIP   10.97.149.109    <none>        9100/TCP       38m
  service/prmt-prometheus-pushgateway     ClusterIP   10.103.35.123    <none>        9091/TCP       38m
  service/prmt-prometheus-server          ClusterIP   10.110.91.110    <none>        80/TCP         38m
  service/prmt-prometheus-server-ext      NodePort    10.96.78.135     <none>        80:31735/TCP   36m

  NAME                                           DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
  daemonset.apps/prmt-prometheus-node-exporter   1         1         1       1            1           <none>          38m

  NAME                                           READY   UP-TO-DATE   AVAILABLE   AGE
  deployment.apps/grafana                        1/1     1            1           4m23s
  deployment.apps/prmt-kube-state-metrics        1/1     1            1           38m
  deployment.apps/prmt-prometheus-alertmanager   1/1     1            1           38m
  deployment.apps/prmt-prometheus-pushgateway    1/1     1            1           38m
  deployment.apps/prmt-prometheus-server         1/1     1            1           38m

  NAME                                                      DESIRED   CURRENT   READY   AGE
  replicaset.apps/grafana-68ccc465d8                        1         1         1       4m23s
  replicaset.apps/prmt-kube-state-metrics-d6bd4f5d9         1         1         1       38m
  replicaset.apps/prmt-prometheus-alertmanager-565d646757   1         1         1       38m
  replicaset.apps/prmt-prometheus-pushgateway-7cf88c767f    1         1         1       38m
  replicaset.apps/prmt-prometheus-server-868cb86b69         1         1         1       38m
  ```

- Follow the simillar process for Grafana too to get the exposed IP

  ```powershell
  PS C:\Users\measutosh> kubectl expose service grafana --type=NodePort --target-port=3000 --name=grafana-ext
  service/grafana-ext exposed
  PS C:\Users\measutosh> minikube service grafana-ext
  |-----------|-------------|-------------|---------------------------|
  | NAMESPACE |    NAME     | TARGET PORT |            URL            |
  |-----------|-------------|-------------|---------------------------|
  | default   | grafana-ext |          80 | http://192.168.49.2:30113 |
  |-----------|-------------|-------------|---------------------------|
  🏃  Starting tunnel for service grafana-ext.
  |-----------|-------------|-------------|------------------------|
  | NAMESPACE |    NAME     | TARGET PORT |          URL           |
  |-----------|-------------|-------------|------------------------|
  | default   | grafana-ext |             | http://127.0.0.1:62004 |
  |-----------|-------------|-------------|------------------------|
  🎉  Opening service default/grafana-ext in default browser...
  ❗  Because you are using a Docker driver on windows, the terminal needs to be open to run it.
  ```

- To get the username and password for the _Grafana Dashboard_ check the secrets from the _yaml file_ of Grafana.

  ```powershell
  PS C:\Users\measutosh> kubectl get secret --namespace default grafana -o yaml
  apiVersion: v1
  data:
    admin-password: RGgwNldhNmluVVNBZVRVOURGYzc5bnUyQzJadjVOVVFXRkpadVk5Sw==
    admin-user: YWRtaW4=
    ldap-toml: ""
  kind: Secret
  metadata:
    annotations:
      meta.helm.sh/release-name: grafana
      meta.helm.sh/release-namespace: default
    creationTimestamp: "2022-09-11T14:05:12Z"
    labels:
      app.kubernetes.io/instance: grafana
      app.kubernetes.io/managed-by: Helm
      app.kubernetes.io/name: grafana
      app.kubernetes.io/version: 9.1.4
      helm.sh/chart: grafana-6.37.3
    name: grafana
    namespace: default
    resourceVersion: "2968"
    uid: 25c5fe41-1dd0-49ad-a22d-a30878ad67d7
  type: Opaque
  ```

- Here the username and the password stays encrypted, those needs to be decrypted first.

  ```powershell
  PS C:\Users\measutosh> echo "YWRtaW4=" | openssl base64 -d ; echo
  admin
  cmdlet Write-Output at command pipeline position 1
  Supply values for the following parameters:
  InputObject[0]:
  PS C:\Users\measutosh> echo "RGgwNldhNmluVVNBZVRVOURGYzc5bnUyQzJadjVOVVFXRkpadVk5Sw==" | openssl base64 -d ; echo
  Dh06Wa6inUSAeTU9DFc79nu2C2Zv5NUQWFJZuY9K
  cmdlet Write-Output at command pipeline position 1
  Supply values for the following parameters:
  InputObject[0]: 
  PS C:\Users\measutosh>
  ```

***

### Prometheus & Grafana Connection

- From _Data Sources_ in Grafana, Prometheus can be selected. There add the IP address from the terminal which was revealed in the terminal earlier at the below part, save it and test it.

  ```powershell
  PS C:\Users\measutosh> minikube service prmt-prometheus-server-ext
  |-----------|----------------------------|-------------|---------------------------|
  | NAMESPACE |            NAME            | TARGET PORT |            URL            |
  |-----------|----------------------------|-------------|---------------------------|
  | default   | prmt-prometheus-server-ext |          80 | http://192.168.49.2:31735 |
  |-----------|----------------------------|-------------|---------------------------|
  ```

- Dashboards in Grafana can be created from scratch or can be imported from _[Grafana's official website](https://grafana.com/grafana/dashboards/6417-kubernetes-cluster-prometheus/)_. For this choose _import_ from the left side of the dashboard, add the _id_, select the resource type as _prometheus_, then add the dashboards and it will load the dashboard with beatiful graphs.
