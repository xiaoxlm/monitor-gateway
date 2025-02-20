package main

import (
	"github.com/xiaoxlm/monitor-gateway/api/router"
)

// @title monitor-gateway
// @version 1.0
// @description This is a monitor gateway
// @termsOfService http://www.swagger.io/support

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	router.Start()
}
