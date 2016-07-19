package routers

import (
	"github.com/xiaojianwu/dd4x/web80/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
