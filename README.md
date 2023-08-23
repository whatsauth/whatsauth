# whatsauth

golang websocket package for whatsapp authentication

## Usage

```sh
go get github.com/whatsauth/whatsauth
```

```go
package controller

import (
 "encoding/json"
 "fmt"
 "strconv"

 "github.com/gin-gonic/gin"
 "github.com/whatsauth/whatsauth"
)

func WsWhatsAuthQR(c *gin.Context) {
 roomid:=whatsauth.ServeWs(c.Writer, c.Request)
}

func PostWhatsAuthRequest(c *gin.Context) {
  var req whatsauth.WhatsauthRequest
  c.BindJSON(&req)
  status := whatsauth.SendStructTo(req.Uuid, infologin)
  c.JSON(200, status)
}
```

## Tagging

develop and publish new version of whatsauth

```sh
git tag v0.2.6
git push origin --tags
go list -m github.com/whatsauth/whatsauth@v0.2.6
```

## Environment

Setting up environment

```sh
GOPROXY=proxy.golang.org
```
