# vela-sam
windows sam dump

## install

```go
    package main

    import (
        "github.com/vela-security/vela-public/assert"
        sam "github.com/vela-security/vela-sam"
    )

    func use(xEnv assert.Environment) {
        sam.WithEnv(xEnv)
    }
	
	//build
	//go mod tidy
```
## example
```lua
    vela.sam_dump("xxxxx.xxx.exe" , function(raw) -- raw:string
        print(raw)
        vela.Debug("%s" , raw) 
    end)
```