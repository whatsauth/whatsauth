# whatsauth

golang websocket package for whatsapp authentication

## Usage

```sh
go get github.com/whatsauth/whatsauth
```

In the main pakcage

```go
package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/informatics-research-center/auth-service/config"

	"github.com/whatsauth/whatsauth"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/informatics-research-center/auth-service/url"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New()
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
```

In the URL package

```go
package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gitlab.com/informatics-research-center/auth-service/controller"
	"os"
)

func Web(page *fiber.App) {
	//API from user whatsapp message from iteung gowa
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)                
	page.Post("/api/whatsauth/request/role", controller.PostWhatsAuthRole)

	//websocket whatsauth to serve wauthjs frontend
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR))

}
```

In the controller package

```go
package controller

import (
	"fmt"

	"github.com/aiteung/athelper"
	"gitlab.com/informatics-research-center/auth-service/model"

	"github.com/aiteung/atmodel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/watoken"
	"github.com/whatsauth/whatsauth"
	"gitlab.com/informatics-research-center/auth-service/config"
)

func WsWhatsAuthQR(c *websocket.Conn) { //simpati unpas lama
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error { //receiver whtasapp message token
	if string(c.Request().Host()) == config.Internalhost || string(c.Request().Host()) == "127.0.0.1:7777" {
		var req whatsauth.WhatsauthRequest
		var ntfbtn atmodel.NotifButton
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		app := watoken.GetAppSubDomain(req.Uuid)
		getapptried := 0
		for (getapptried < 17) && (app == "") {
			app = watoken.GetAppSubDomain(req.Uuid)
			getapptried += getapptried
		}

		if app == "siapbaak" {
			ntfbtn = whatsauth.RunModuleLegacy(req, config.PrivateKey, config.SiapUserTables[:], config.Ulbimssqlconn)
			fmt.Println(ntfbtn)
		} else if config.CheckIsAkademik(app) {
			ntfbtn = whatsauth.RunWithUsernames(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		} else {
			ntfbtn = whatsauth.RunWithUsernames(req, config.PrivateKey, config.AptimasTables[:], config.AptimasConn)
		}
		if app == "" {
			ntfbtn.Message.Message.FooterText = ntfbtn.Message.Message.FooterText + req.Uuid
		}
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}
}

func PostWhatsAuthRole(c *fiber.Ctx) error { //receiver whtasapp message token
	if string(c.Request().Host()) != config.Internalhost || string(c.Request().Host()) == "127.0.0.1:7777" {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}
	req := new(whatsauth.WhatsAuthRoles)
	err := c.BodyParser(req)
	if err != nil {
		return err
	}
	var ntfbtn atmodel.NotifButton
	app := watoken.GetAppSubDomain(req.Uuid)
	if app == "siapbaak" {
		ntfbtn := whatsauth.SelectedRoles(*req, config.PrivateKey, config.SiapUserTables[:], config.Ulbimssqlconn)
		fmt.Println(ntfbtn)
	} else if config.CheckIsAkademik(app) {
		ntfbtn = whatsauth.SelectedRoles(*req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
	} else {
		ntfbtn = whatsauth.SelectedRoles(*req, config.PrivateKey, config.AptimasTables[:], config.AptimasConn)
	}
	fmt.Printf("\nreturn button from auth : %+q \n", ntfbtn)
	return c.JSON(ntfbtn)
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
