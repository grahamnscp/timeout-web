# webserver to simulate slow server response time, for testing ingress timeouts

Build the image:
```
# build image and push:
#  M1 macos
docker buildx build --push -t grahamh/timeout-web:1.0-arch64
docker run -d --rm -p 8080:8080 grahamh/timeout-web:1.0-arch64

#  linux
docker build -t grahamh/timeout-web:1.0 .
docker push grahamh/timeout-web:1.0
docker run -d --rm -p 8080:8080 grahamh/timeout-web:1.0
```

Start a local container instance to test:
```
# test run output (client)
curl http://localhost:8080/5
[10-30-2025 10:35:20.62 timeout-web] Host: 21830b6e054e, Recieved Request: gdesktop.glocal.lab:8080/5
[10-30-2025 10:35:25.62 timeout-web] Slept for: 5s

# (server logs)
docker logs 21830b6e054e -f
Serving HTTP on port :8080
[10-30-2025 10:35:20.62 timeout-web] Host: 21830b6e054e, Request: gdesktop.glocal.lab:8080/05
[10-30-2025 10:35:20.62 timeout-web] Sleeping for 5s..
[10-30-2025 10:35:25.62 timeout-web] done. Slept for: 5s
```

Default ingress-nginx timeout on rke2 kubernetes is 60s
```
curl http://mydnsname.example.com/61
<html>
<head><title>504 Gateway Time-out</title></head>
<body>
<center><h1>504 Gateway Time-out</h1></center>
<hr><center>nginx</center>
</body>
</html>
```

Change the app ingress settings in the deployment:
```
    annotations:
      kubernetes.io/ingress.class: nginx
      nginx.ingress.kubernetes.io/proxy-connect-timeout: "300"
      nginx.ingress.kubernetes.io/proxy-read-timeout: "300"
      nginx.ingress.kubernetes.io/proxy-send-timeout: "300"
```

