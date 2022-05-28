# ARS_SHELL_CRYPT

This is a Go implementation of https://github.com/Neptune-IT/ARS_SHELL_CRYPT

```go
package main

import (
	"fmt"
	"github.com/restartfu/ars-shell-crypt/cryptage"
	"github.com/restartfu/ars-shell-crypt/security"
)

func main() {
	// this will encrypt the given string "test".
	e := cryptage.Encrypt("test", security.LevelDefault)
	// print out the result
	fmt.Printf("Encrypted: %s\n", e)
	// this will decrypt the given encrypted string.
	d := cryptage.Decrypt(e, security.LevelDefault)
	// print out the result
	fmt.Printf("Decrypted: %s\n", d)
}
```