# Nemo
[![Go Report Card](https://goreportcard.com/badge/github.com/valentin2105/Nemo)](https://goreportcard.com/report/github.com/valentin2105/Nemo)
[![Build Status](https://travis-ci.org/valentin2105/Nemo.svg?branch=master)](https://travis-ci.org/valentin2105/Nemo)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/dwyl/esta/issues)

<img src="https://i.imgur.com/AuRlZuJ.png">

`Nemo` (not the fish, the Captain) is a **Kubernetes UI** to list, describe and modify resources in your cluster.

## Features
- List most of Kubernetes resources
- Describe these resources
- Scale Up/Down your Deployments
- Delete resources
- Create All-in-one Kubernetes definitions (Deploy, Service, VolumeClaim...)

`Nemo` is writted in Golang and use a `kubeconfig` file to talk to Kubernetes.

## How to use it

You can launch easily with a simple `docker run` command :

```
 docker run -it -p 80:8080 -e KUBERNETES_SERVICE_HOST=k8s-api.domain.ltd \
    -e KUBERNETES_SERVICE_PORT=6443 \
    -e KUBERNETES_TOKEN=yourprivatetoken \
    valentinnc/nemo
```

Or you can launch it locally, without Docker :
```
# local
git clone https://github.com/valentin2105/Nemo.git
cd Nemo/ && wget https://...
chmod +x Nemo
./Nemo --kubeconfig /home/user/.kube/config
```

## Screenshots

<br>
<img src="https://i.imgur.com/Xc5y7Im.png" width="646" height="440">

## How to build it
`Nemo` uses `dep` to fetch Go dependencies:

```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
dep ensure
make build
./Nemo --kubeconfig /home/user/.kube/config
```
