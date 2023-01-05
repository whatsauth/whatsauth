# whatsauth

golang package for whatsapp authentication

## Usage

### Initialize your module

```sh
go mod init example.com/my-whatsauth-demo
```

### Get the gogis module

Note that you need to include the **v** in the version tag.

```sh
go get github.com/whatsauth/whatsauth
```

```go
package main

import (
    "fmt"

    "github.com/whatsauth/whatsauth"
)

func main() {
    
}
```

## To Contribute this Repo

```sh
go test
```

## Tagging

develop and publish new version of gogis

```sh
git tag v0.1.2
git push origin --tags
go list -m github.com/whatsauth/whatsauth@v0.0.1
```

## Environment

Setting up environment

```sh
GOPROXY=proxy.golang.org
```
