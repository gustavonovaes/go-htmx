
```
go install golang.org/x/tools/gopls@latest
go install github.com/air-verse/air@latest
go install github.com/segmentio/golines@latest
```


```
    "emeraldwalk.runonsave": {
        "commands": [
            {
                "cmd": "$HOME/go/bin/golines ${file} -w",
                "match": "\\.go$"
            }
        ]
    },
```