// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cleverbamboo.com/bee-shop-b2c/api"
	"cleverbamboo.com/bee-shop-b2c/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/ad",
			beego.NSInclude(
				&controllers.AdController{},
			),
		),

		beego.NSNamespace("/ad_position",
			beego.NSInclude(
				&controllers.AdPositionController{},
			),
		),

		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
				&api.AdminController{},
			),
		),

		beego.NSNamespace("/area",
			beego.NSInclude(
				&controllers.AreaController{},
			),
		),

		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),

		beego.NSNamespace("/article_category",
			beego.NSInclude(
				&controllers.ArticleCategoryController{},
			),
		),

		beego.NSNamespace("/attribute",
			beego.NSInclude(
				&controllers.AttributeController{},
			),
		),

		beego.NSNamespace("/brand",
			beego.NSInclude(
				&controllers.BrandController{},
			),
		),

		beego.NSNamespace("/cart",
			beego.NSInclude(
				&controllers.CartController{},
			),
		),

		beego.NSNamespace("/cart_item",
			beego.NSInclude(
				&controllers.CartItemController{},
			),
		),

		beego.NSNamespace("/consultation",
			beego.NSInclude(
				&controllers.ConsultationController{},
			),
		),

		beego.NSNamespace("/coupon",
			beego.NSInclude(
				&controllers.CouponController{},
			),
		),

		beego.NSNamespace("/coupon_code",
			beego.NSInclude(
				&controllers.CouponCodeController{},
			),
		),

		beego.NSNamespace("/delivery_center",
			beego.NSInclude(
				&controllers.DeliveryCenterController{},
			),
		),

		beego.NSNamespace("/delivery_corp",
			beego.NSInclude(
				&controllers.DeliveryCorpController{},
			),
		),

		beego.NSNamespace("/delivery_template",
			beego.NSInclude(
				&controllers.DeliveryTemplateController{},
			),
		),

		beego.NSNamespace("/deposit",
			beego.NSInclude(
				&controllers.DepositController{},
			),
		),

		beego.NSNamespace("/feedback",
			beego.NSInclude(
				&controllers.FeedbackController{},
			),
		),

		beego.NSNamespace("/friend_link",
			beego.NSInclude(
				&controllers.FriendLinkController{},
			),
		),

		beego.NSNamespace("/gift_item",
			beego.NSInclude(
				&controllers.GiftItemController{},
			),
		),

		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsController{},
			),
		),

		beego.NSNamespace("/log",
			beego.NSInclude(
				&controllers.LogController{},
			),
		),

		beego.NSNamespace("/member",
			beego.NSInclude(
				&controllers.MemberController{},
			),
		),

		beego.NSNamespace("/member_attribute",
			beego.NSInclude(
				&controllers.MemberAttributeController{},
			),
		),

		beego.NSNamespace("/member_rank",
			beego.NSInclude(
				&controllers.MemberRankController{},
			),
		),

		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),

		beego.NSNamespace("/navigation",
			beego.NSInclude(
				&controllers.NavigationController{},
			),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),

		beego.NSNamespace("/order_item",
			beego.NSInclude(
				&controllers.OrderItemController{},
			),
		),

		beego.NSNamespace("/order_log",
			beego.NSInclude(
				&controllers.OrderLogController{},
			),
		),

		beego.NSNamespace("/parameter",
			beego.NSInclude(
				&controllers.ParameterController{},
			),
		),

		beego.NSNamespace("/parameter_group",
			beego.NSInclude(
				&controllers.ParameterGroupController{},
			),
		),

		beego.NSNamespace("/payment",
			beego.NSInclude(
				&controllers.PaymentController{},
			),
		),

		beego.NSNamespace("/payment_method",
			beego.NSInclude(
				&controllers.PaymentMethodController{},
			),
		),

		beego.NSNamespace("/permission",
			beego.NSInclude(
				&controllers.PermissionController{},
			),
		),

		beego.NSNamespace("/plugin_config",
			beego.NSInclude(
				&controllers.PluginConfigController{},
			),
		),

		beego.NSNamespace("/pms_project",
			beego.NSInclude(
				&controllers.PmsProjectController{},
			),
		),

		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{},
				&api.ProductController{},
			),
		),

		beego.NSNamespace("/product_category",
			beego.NSInclude(
				&controllers.ProductCategoryController{},
			),
		),

		beego.NSNamespace("/product_notify",
			beego.NSInclude(
				&controllers.ProductNotifyController{},
			),
		),

		beego.NSNamespace("/promotion",
			beego.NSInclude(
				&controllers.PromotionController{},
			),
		),

		beego.NSNamespace("/receiver",
			beego.NSInclude(
				&controllers.ReceiverController{},
			),
		),

		beego.NSNamespace("/refunds",
			beego.NSInclude(
				&controllers.RefundsController{},
			),
		),

		beego.NSNamespace("/register_code",
			beego.NSInclude(
				&controllers.RegisterCodeController{},
			),
		),

		beego.NSNamespace("/returns",
			beego.NSInclude(
				&controllers.ReturnsController{},
			),
		),

		beego.NSNamespace("/returns_item",
			beego.NSInclude(
				&controllers.ReturnsItemController{},
			),
		),

		beego.NSNamespace("/review",
			beego.NSInclude(
				&controllers.ReviewController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/seo",
			beego.NSInclude(
				&controllers.SeoController{},
			),
		),

		beego.NSNamespace("/shipping",
			beego.NSInclude(
				&controllers.ShippingController{},
			),
		),

		beego.NSNamespace("/shipping_item",
			beego.NSInclude(
				&controllers.ShippingItemController{},
			),
		),

		beego.NSNamespace("/shipping_method",
			beego.NSInclude(
				&controllers.ShippingMethodController{},
			),
		),

		beego.NSNamespace("/sn",
			beego.NSInclude(
				&controllers.SnController{},
			),
		),

		beego.NSNamespace("/specification",
			beego.NSInclude(
				&controllers.SpecificationController{},
			),
		),

		beego.NSNamespace("/specification_value",
			beego.NSInclude(
				&controllers.SpecificationValueController{},
			),
		),

		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
