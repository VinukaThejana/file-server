// File server
// Share files through a local network
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/VinukaThejana/file-server/config"
	"github.com/VinukaThejana/file-server/middleware"
	"github.com/VinukaThejana/file-server/utils"
	"github.com/VinukaThejana/go-utils/logger"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	flag "github.com/spf13/pflag"
)

var (
	log logger.Logger
	cfg config.Config
)

func init() {
	cfg.Load()

	var port, path string
	flag.StringVarP(&port, "port", "o", cfg.Port, "Specify the port that the file server is listening")
	flag.StringVarP(&path, "path", "p", cfg.Path, "Specify the location that needs to be served")
	flag.Parse()

	if port != cfg.Port {
		port = fmt.Sprintf(":%s", port)
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			path = scanner.Text()
		}
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			color.Red("The given path does not exist\n")
			os.Exit(0)
		}

		log.Errorf(err, nil)
	}

	cfg.Port = port
	cfg.Path = path

	// Display the IP addredd of this device through the interface that can connect
	// to the internet
	color.Red(fmt.Sprintf("\nhttp://%s%s", utils.GetOutboundIP().String(), cfg.Port))
	color.Red(fmt.Sprintf("Serving path [%s]\n", cfg.Path))
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:                 "file-server",
		DisableStartupMessage:   true,
		EnableTrustedProxyCheck: true,
		EnableIPValidation:      true,
	})
	app.Use(middleware.IP)

	app.Static("/", cfg.Path, fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	log.Errorf(app.Listen(cfg.Port), nil)
}
