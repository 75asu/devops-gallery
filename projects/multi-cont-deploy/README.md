# Multi-Container Full Stack App Deployment On Minikube

![flow-diagram](./images/diagram-multi-cont-deploy.PNG)

## Tools used

- Docker
- Minikube

<br>

## To spinup the app locally in minikube follow the steps

<br>

![demo-diagram](./images/demo-multi-cont-deploy.PNG)
- clone this repo to local
- enter into cloned folder
- run `minikube apply -f manifets`
- similarly to delete the app run `minikube delete -f manifest`

### Steps followed to build this

- Client and its service configuration
- Postgres, its persistentVolumeClaim, service configuration
- Server, its service configuration
- Ingress Nginx controller configuration