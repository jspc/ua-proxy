ua-proxy
==

buidling
--

```bash
$ go build
```

runninnggn
--

```bash
$ sudo ./ua-proxy
```

(This listens  on port 666 for totally non-idiotic, edgelord reasons)

testing
--

```
$ go test
```

using
--

```bash
$ curl localhost:666 -A '{"method": "GET", "url": https://requestb.in/xroxujxr"}'
```

