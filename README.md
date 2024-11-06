
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

## Running

```bash
air
```

## Running with Docker

```bash
docker build -t go-htmx .
docker run -p 8080:8080 go-htmx
```
