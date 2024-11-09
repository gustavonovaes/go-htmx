
## Tools to install

```bash
# for language server
go install golang.org/x/tools/gopls@latest 

# for hot reloading
go install github.com/air-verse/air@latest 

# for formatting
go install github.com/segmentio/golines@latest 

# for static analysis
go install github.com/jondot/goweight@latest
```

## Configure VSCode to run `golines` on save

Install the extension [emeraldwalk/vscode-runonsave](https://github.com/emeraldwalk/vscode-runonsave) and add the following configuration to the `settings.json` file:

```json
"emeraldwalk.runonsave": {
    "commands": [
        {
            "cmd": "$HOME/go/bin/golines ${file} -w",
            "match": "\\.go$"
        }
    ]
},
```

## Try

[https://go-htmx-18ejig.fly.dev/](https://go-htmx-18ejig.fly.dev/)

## Running

```bash
git clone https://github.com/gustavonovaes/go-htmx
go mod download
air

# example output
2024/11/09 13:53:29 INFO: POST 24.52µs /api/count - RAM 0.35 MB
2024/11/09 13:53:29 INFO: POST 36.289µs /api/count - RAM 0.36 MB
2024/11/09 13:53:29 INFO: POST 16.689µs /api/count - RAM 0.37 MB
2024/11/09 13:53:30 INFO: POST 98.349µs /api/count - RAM 0.37 MB
2024/11/09 13:53:30 INFO: POST 26.1µs /api/count - RAM 0.38 MB
2024/11/09 13:53:31 INFO: GET 600.571522ms /api/count - RAM 0.38 MB
2024/11/09 13:53:32 INFO: GET 600.720792ms /api/count - RAM 0.39 MB
2024/11/09 13:53:32 INFO: POST 31.92µs /api/count - RAM 0.39 MB
2024/11/09 13:53:32 INFO: POST 94.319µs /api/count - RAM 0.40 MB
2024/11/09 13:53:32 INFO: POST 32.65µs /api/count - RAM 0.40 MB
```

## Running with Docker

```bash
docker build -t go-htmx .
docker run -p 8080:8080 go-htmx

# or with docker-compose
docker-compose up

# image size 
> docker image ls localhost/go-htmx_app:latest
REPOSITORY             TAG         IMAGE ID      CREATED        SIZE
localhost/go-htmx_app  latest      1bf6cefd9391  6 minutes ago  7.19 MB
```

## Static analysis

```bash
goweight gustavonovaes.dev/go-htmx/internal/core | head -5
```
Example output:
```bash
   13 MB runtime
  8.7 MB net/http
  4.4 MB net
  4.2 MB crypto/tls
  3.0 MB reflect
```

## Treemap
![googlechrome github io_lighthouse_treemap__gzip=1](https://github.com/user-attachments/assets/db25958b-4cfb-4be9-9955-51eff8e5a9f0)
