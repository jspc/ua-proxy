[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjspc%2Fua-proxy.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjspc%2Fua-proxy?ref=badge_shield)

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



## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjspc%2Fua-proxy.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjspc%2Fua-proxy?ref=badge_large)