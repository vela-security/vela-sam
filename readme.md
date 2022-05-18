# vela-sam
windows sam dump

## example
```
    vela.sam_dump("xxxxx.xxx.exe" , function(raw) -- raw:string
        print(raw)
        vela.Debug("%s" , raw) 
    end
```