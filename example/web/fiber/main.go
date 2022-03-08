// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-fiber/boot"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample rk-demo server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Register handler
	entry := rkfiber.GetFiberEntry("greeter")
	entry.App.Get("/v1/greeter", Greeter)
	// This is required!!!
	entry.RefreshFiberRoutes()

	boot.WaitForShutdownSig(context.TODO())
}

// Greeter handler
// @Summary Greeter
// @Id 1
// @Tags Hello
// @version 1.0
// @Param name query string true "name"
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(ctx *fiber.Ctx) error {
	ctx.Response().SetStatusCode(http.StatusOK)
	return ctx.JSON(&GreeterResponse{
		Message: fmt.Sprintf("Hello %s!", ctx.Query("name")),
	})
}

type GreeterResponse struct {
	Message string
}
