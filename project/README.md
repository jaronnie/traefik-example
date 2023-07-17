# project

## use docker as provider

```shell
make deploy
```

open `localhost:8080` in browser, you can see web UI.

```shell
# use nginx as console
curl localhost/gateway/app1/api/v1/echo
curl localhost/gateway/app2/api/v1/echo
curl localhost/gateway/app3/api/v1/echo
curl localhost/gateway/app4/api/v1/echo # error, not found

# use caddy as console
curl localhost:81/gateway/app1/api/v1/echo
curl localhost:81/gateway/app2/api/v1/echo
curl localhost:81/gateway/app3/api/v1/echo
curl localhost:81/gateway/app4/api/v1/echo # error, not found
```
