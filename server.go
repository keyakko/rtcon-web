package main

import (
  "net/http"
  "fmt"
  "os"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

type PostData struct {
  Type string `json: "type" form: "type" query: "type"`
  Msg  string `json: "msg"  form: "msg"  query: "msg"`
}


func OutputLog(msg string) {
  fmt.Print("\033[106;30m[LOG]\033[0m  ")
  fmt.Println(msg)
}
func OutputWarning (msg string) {
  fmt.Print("\033[043;30m[WARN]\033[0m ")
  fmt.Println(msg)
}
func OutputInfo (msg string) {
  fmt.Print("\033[104;30m[INFO]\033[0m ")
  fmt.Println(msg)
}
func OutputError (msg string) {
  fmt.Print("\033[041;30m[ERR]\033[0m  ")
  fmt.Println(msg)
}

func ConsoleGetter (c echo.Context) error {
  d := new(PostData)
  if err := c.Bind(d); err != nil {
    return c.String(http.StatusOK, "err")
  }

  if d.Type == "log" {
    OutputLog(d.Msg)
  } else if d.Type == "warn" {
    OutputWarning(d.Msg)
  } else if d.Type == "info" {
    OutputInfo(d.Msg)
  } else if d.Type == "err"  {
    OutputError(d.Msg)
  }

  return c.String(http.StatusOK, "done")
}

func main () {
  e := echo.New()
  e.Use(middleware.CORS())
  e.GET("/", ConsoleGetter)
  e.POST("/", ConsoleGetter)

  port := os.Getenv("PORT")
  if port == "" {
    port = "5500"
    fmt.Printf("Using default port: %s", port)
  }
  e.Start(fmt.Sprintf(":%s", port))
}
