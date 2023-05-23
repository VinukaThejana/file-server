package middleware

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

// IP is used to extract the ip addreses of the machines that requested for the
// file service
func IP(c *fiber.Ctx) error {
	color.Blue(fmt.Sprintf("Connection request from IP address : %s\n", c.IP()))
	return c.Next()
}
