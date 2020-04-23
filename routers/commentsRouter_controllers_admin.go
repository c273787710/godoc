package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["godoc/controllers/admin:CateController"] = append(beego.GlobalControllerRouter["godoc/controllers/admin:CateController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godoc/controllers/admin:CateController"] = append(beego.GlobalControllerRouter["godoc/controllers/admin:CateController"],
        beego.ControllerComments{
            Method: "DelCate",
            Router: `/del`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godoc/controllers/admin:CateController"] = append(beego.GlobalControllerRouter["godoc/controllers/admin:CateController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godoc/controllers/admin:CateController"] = append(beego.GlobalControllerRouter["godoc/controllers/admin:CateController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godoc/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["godoc/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["godoc/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["godoc/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
