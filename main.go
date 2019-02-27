package main

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"fmt"
)

var timeoutDelay int = 0

type MainController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

type DataController struct {
	beego.Controller
}

type PostDataController struct {
	beego.Controller
}

func init() {
	beego.Router("/", &MainController{})
	beego.Router("/test", &TestController{})
	beego.Router("/api/data", &DataController{})
	beego.Router("/api/setter", &PostDataController{})
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *TestController) Get() {
	c.TplName = "test.html"
}

func (c *DataController) Get() {
	c.Data["json"] = map[string]string{"timeout": strconv.Itoa(timeoutDelay)}
	c.ServeJSON();
}

func (this *PostDataController) Post() {
	timeout := strings.Trim(this.GetString("timeout"), "")
	fmt.Println(timeout)
	if timeout == "" {
		this.Data["json"] = map[string]string{"msg": "Invalid timeout value."}
	} else {
		tout, err := strconv.Atoi(timeout)
		if err == nil {
			timeoutDelay = tout
		}
		this.Data["json"] = map[string]string{"msg": "Timeout updated to " + timeout}
	}
	this.ServeJSON()
}

func main() {
	beego.Run()
}
