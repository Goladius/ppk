#Private & Public Key - PPK

Generating private and public key file for ssh conection. If you learn how to use ssh with Golang you can check my this post:

[https://taluttasgiran.com/golang-ssh-connection-to-remote-server-via-private-key-file/](https://taluttasgiran.com/golang-ssh-connection-to-remote-server-via-private-key-file/)

## How To Use?
***Download package***

```go get github.com/goladius/ppk```

This codes, generate two file into your project directory.
```
    package main
    
    import "github.com/goladius/ppk"
    
    func main() {
    	privateKey := ppk.GetPrivateKey()
    	ppk.SavePrivateKey("private", privateKey)
    	ppk.SavePublicKey("private", privateKey)
    }
```

