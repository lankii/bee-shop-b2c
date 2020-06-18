package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api/admin:AccountController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api/admin:AccountController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api/admin:AccountController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api/admin:AccountController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api/admin:IndexController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api/admin:IndexController"],
        beego.ControllerComments{
            Method: "GetAllAdmins",
            Router: "/getAll",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
