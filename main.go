package main

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"fmt"
	"time"
	"flag"
)

var timeoutDelay int = 0
var statusCode string = "200"
var port = flag.String("port", "6686", "port on which application will run")

type MainController struct {
	beego.Controller
}

type DataController struct {
	beego.Controller
}

func init() {
	beego.Router("/", &MainController{})
	beego.Router("/api/setter", &DataController{})
}

func (c *MainController) Get() {
	time.Sleep(time.Duration(timeoutDelay) * time.Second)
	c.Abort(statusCode)
}

func (this *DataController) Get() {
	timeout := strings.Trim(this.GetString("timeout"), "")
	sCode := strings.Trim(this.GetString("statusCode"), "")
	fmt.Println(timeout, statusCode)
	if timeout != "" {
		tout, err := strconv.Atoi(timeout)
		if err == nil {
			timeoutDelay = tout
		}
	}
	if sCode != "" {
		statusCode = sCode
	}
	this.ServeJSON()
}

func main() {
	flag.Parse()
	beego.Run(":" + *port)
}
