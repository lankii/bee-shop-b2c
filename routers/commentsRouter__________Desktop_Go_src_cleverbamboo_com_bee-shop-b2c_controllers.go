package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdPositionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AreaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleCategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:AttributeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:BrandController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CartItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ConsultationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponCodeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:CouponController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCenterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryCorpController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DeliveryTemplateController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:DepositController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FeedbackController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:FriendLinkController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GiftItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:LogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberAttributeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MemberRankController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:MessageController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:NavigationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:OrderLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ParameterGroupController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PaymentMethodController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PluginConfigController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PmsProjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductCategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ProductNotifyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReceiverController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RefundsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RegisterCodeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReturnsItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ReviewController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SeoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:ShippingMethodController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SnController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:SpecificationValueController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"] = append(beego.GlobalControllerRouter["cleverbamboo.com/bee-shop-b2c/controllers:TagController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
