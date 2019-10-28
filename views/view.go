package views

import "github.com/labstack/echo"

func Index(ctx echo.Context) error {

	return ctx.File("./static/test.html")
}
