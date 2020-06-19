package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"],
        beego.ControllerComments{
            Method: "GetAllAdmins",
            Router: "/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/api:AdminController"],
        beego.ControllerComments{
            Method: "UpdateAdmin",
            Router: "/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
