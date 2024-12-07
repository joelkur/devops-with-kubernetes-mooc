# Exercise 1.02

Steps to run exercise 1.02

```bash
kubectl create deployment todo-app --image=jk0der/todo-app:exercise-1.02
```

Verify todo-app is running with `kubectl get pods`, which should output something similar to below:

```bash
NAME                                 READY   STATUS    RESTARTS   AGE
todo-app-b67749cd7-2v7vw             1/1     Running   0          13s
```

Check the logs of the pod with `kubectl logs todo-app-b67749cd7-2v7vw`
```bash
2024/12/07 17:47:01 Server started in port 3000
```
