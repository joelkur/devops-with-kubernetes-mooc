# Exercise 1.06

Below steps assume that k3d cluster has been created with
```bash
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```

Steps to run exercise 1.06

```bash
kubectl apply -f todo-app/manifests/service.yaml
```

Making a GET http request to http://localhost:8082 should now respond with html
```bash
$ curl http://localhost:8082
<html>
  <head>
    <title>todo-app</title>
  </head>
  <body>
    <h1>todo-app</h1>
  </body>
</html>
```
