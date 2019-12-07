# nft over REST API

WIP

## Build

```bash
make build
```

## Run

```bash
./nft-http-api --unix /var/run/nft-http-api.sock

curl --unix-socket /var/run/nft-http-api.sock http://localhost/table
```

Or:

```bash
./nft-http-api --listen 127.0.0.1:4242 --tls-cert fixtures/cert.pem --tls-key fixtures/key.pem

curl https://localhostt:4242/table
```

## Contribute

```
# upload from mac to debian buster
rsync -avzhP . root@${REMOTE_SERVER_IP}:nft-http-api --del --exclude .git
```
