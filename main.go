package main

import (
	"fmt"
	"net/http"

	gd "github.com/bolknote/go-gd"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", createImage)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func createImage(c echo.Context) error {
	pict := gd.CreateFromJpeg("images/bg.jpg")

	head := gd.CreateFromJpeg("images/head.jpg")

	destroy := func() {
		fmt.Print("释放")
		pict.Destroy()
		head.Destroy()
	}

	defer destroy()

	// white := pict.ColorAllocate(255, 255, 255)

	head.CopyMerge(pict, 10, 10, 0, 0, 132, 132, 100)

	// zhFont := "font/AdobeHeitiStd-Regular.otf"
	// enBoldFont := "font/Helvetica Bold.ttf"
	// enFont := "font/Helvetica.ttf"

	// pict.StringFT(white, zhFont, 30, 0, 200, 200, "我试试中文的字体!")
	// pict.StringFT(white, enFont, 20, 0, 100, 50, "12:02:02")
	// pict.StringFT(white, enBoldFont, 30, 0, 100, 100, "2020-12-12")

	uid := uuid.Must(uuid.NewV4()).String()

	path := "outs/" + uid + ".jpg"
	pict.Jpeg(path, 100)

	return c.String(http.StatusOK, path)
}
