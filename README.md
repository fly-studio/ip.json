# ip.json(Golang)
A ip.json web server via [纯真IP](http://www.cz88.net/ip/) 库

## Print like 
```
{
    "ip":"116.211.100.249",
    "url":"/",
    "country":"本地",
    "area":"本地"
}
```

## CLI

```
./ip.json --qqwry /path/to/qqwry.dat --port 8080
```

- **qqwry**: the path of qqwry.dat. default: current path
- **port**: your server's listening port. default: 2060


## Nginx conf

```
location = /ip.json
{
        proxy_pass http://127.0.0.1:2060;

        proxy_redirect off;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $remote_addr;
}

```

Visit: `http://your.domain/ip.json`