# Exercise 1.05

## Commits

- [8c93e8c7eb7b86e3b37a87d9f182b30f42081a89](https://github.com/joelkur/devops-with-kubernetes-mooc/commit/8c93e8c7eb7b86e3b37a87d9f182b30f42081a89)
- [7029ae9ae9189570e7548db2a9b5a5f1ae93beb5](https://github.com/joelkur/devops-with-kubernetes-mooc/commit/7029ae9ae9189570e7548db2a9b5a5f1ae93beb5)
- [cfc0ebe6d4dee3785d336e7e97a71f125c163b85](https://github.com/joelkur/devops-with-kubernetes-mooc/commit/cfc0ebe6d4dee3785d336e7e97a71f125c163b85)

## Description

Steps to run exercise 1.05

```bash
kubectl apply -f todo-app/manifests/deployment.yaml
```

Open a port-forward from the pod port 8080 to a local port e.g. 8080
```bash
kubectl port-forward todo-app-59b485cc8f-jtwt2 8080:8080
```

Making a GET http request to the pod should now respond with html
```bash
$ curl http://localhost:8080
<html>
  <head>
    <title>todo-app</title>
  </head>
  <body>
    <h1>todo-app</h1>
  </body>
</html>
```

