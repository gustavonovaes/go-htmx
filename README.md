
## Tools to install

```bash
# for language server
go install golang.org/x/tools/gopls@latest 

# for hot reloading
go install github.com/air-verse/air@latest 

# for formatting
go install github.com/segmentio/golines@latest 
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
```

## Running with Docker

```bash
docker build -t go-htmx .
docker run -p 8080:8080 go-htmx
```


## Treemap
![googlechrome github io_lighthouse_treemap__gzip=1](https://github.com/user-attachments/assets/db25958b-4cfb-4be9-9955-51eff8e5a9f0)
