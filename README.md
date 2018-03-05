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
You can launch locally with:

```
# local
git clone https://github.com/valentin2105/Nemo.git
cd Nemo/ && wget https://...
chmod +x Nemo
./Nemo --kubeconfig /home/user/.kube/config
```

Or you can run it in your Kubernetes cluster using `ServiceAccount` and `RBAC` in combination with `Helm`.

```
# Helm
git clone https://github.com/valentin2105/Nemo.git
cd Nemo/helm
helm install -n nemo --namespace nemo-ui .
```

## Screenshots

<br>
<img src="https://i.imgur.com/Xc5y7Im.png" width="646" height="440">

## How to build it
`Nemo` uses `GoDeps` to fetch Go dependencies:

```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
dep ensure
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Nemo d
./Nemo --kubeconfig /home/user/.kube/config
```
