# Exercise 1.01

Steps to run exercise 1.01

```bash
kubectl create deployment log-output --image=jk0der/log-output:exercise-1.01
```

Verify log-output is running with `kubectl get pods`, which should output something similar to below:

```bash
NAME                                 READY   STATUS    RESTARTS   AGE
log-output-5dd864485b-mbh95          1/1     Running   0          11m
```

Check the logs of the pod with `kubectl logs log-output-5dd864485b-mbh95`
```bash
2024-12-07T16:52:24Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:52:29Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:52:34Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:52:39Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:52:44Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:52:49Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:52:54Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:52:59Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:53:04Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
2024-12-07T16:53:09Z: 92e5602d-ec23-4cb5-90ad-0fba8ef39fc3
```
