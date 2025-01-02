package main

import (
	"github.com/Asefeh-J/OAuth_Go/api"
	"github.com/Asefeh-J/OAuth_Go/oauth"
)

func main() {
	oauth.InitOAuth()
	api.StartServer()
}
