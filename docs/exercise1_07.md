# Exercise 1.07

## Commits
- [95fc876](https://github.com/joelkur/devops-with-kubernetes-mooc/commit/95fc876f94a088939b4ed00542319ddc66cfc2cc) - exercise 1.07: add service and ingress to log-output
- [acb01cb](https://github.com/joelkur/devops-with-kubernetes-mooc/commit/acb01cb74d93bf3d57378a229c3d0c4f1c41831d) - exercise 1.07: update log-output deployment tag
- [2d24447](https://github.com/joelkur/devops-with-kubernetes-mooc/commit/2d244476a22fc4d352bb37de93f453cb065ae841) - exercise 1.07: add server to log-output app that responds with current status
- [b13b41f](https://github.com/joelkur/devops-with-kubernetes-mooc/commit/b13b41f07ce6119c7dd45b57d0dd4581ccf64b43) - exercise 1.06: add service for todo-app

## Description

Below steps assume that k3d cluster has been created with
```bash
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```

Steps to run exercise 1.07

```bash
kubectl apply -f log-output/manifests/deployment.yaml
kubectl apply -f log-output/manifests/service.yaml
kubectl apply -f log-output/manifests/ingress.yaml
```

Making a GET http request to http://localhost:8081 should now respond with the current status
```bash
$ curl http://localhost:8081
2024-12-08T08:21:51Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
```

And the logs of the pod should print the output with timestamp every 5 seconds, as before:
```bash
kubectl logs log-output-dep-7985dcf44b-tmhtc | head
2024-12-08T08:11:31Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:11:36Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:11:41Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:11:46Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:11:51Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:11:56Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:12:01Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:12:06Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:12:11Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
2024-12-08T08:12:16Z: 038dba11-e959-4646-9e3a-2c1eb3bd6b34
```

