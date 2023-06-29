# hello world

use docker as provider.

```shell
make deploy
```

open `localhost:8080` in browser, you can see web UI.

```shell
curl localhost/gateway/a
curl localhost/gateway/b
curl localhost/gateway/c
curl localhost/gateway/d # error, not found
```