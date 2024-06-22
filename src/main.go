package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

func main() {

	s := echo.New()
	s.GET("/status", Handler)
	err := s.Start(":8080")
	if err!= nil {
        log.Fatal(err)
    }
	

}

func Handler(ctx echo.Context) error {

	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	duration := time.Until(d)

	s := fmt.Sprintf("Days remain: %d", int64(duration.Hours()/24))

	err := ctx.String(http.StatusOK, s)
	if err!= nil {
		return err
    }

	return nil
}
