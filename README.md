<h1 style='text-align:center !importan; display:block !important;'>Private & Public Key - PPK</h1>

Generating private and public key file for ssh conection.

## How To Use?
***Download package***

```go get github.com/taluttasgiran/ppk```

This codes, generate two file into your project directory.
```
    package main
    
    import "github.com/taluttasgiran/ppk"
    
    func main() {
    	privateKey := ppk.GetPrivateKey()
    	ppk.SavePrivateKey("private", privateKey)
    	ppk.SavePublicKey("public", privateKey)
    }
```

