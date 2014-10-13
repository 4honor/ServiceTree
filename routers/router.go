// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ServiceTree/controllers"

	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/favor",
			beego.NSInclude(
				&controllers.FavorController{},
			),
		),

		beego.NSNamespace("/group",
			beego.NSInclude(
				&controllers.GroupController{},
			),
		),

		beego.NSNamespace("/idalloc",
			beego.NSInclude(
				&controllers.IdallocController{},
			),
		),

		beego.NSNamespace("/machine",
			beego.NSInclude(
				&controllers.MachineController{},
			),
		),

		beego.NSNamespace("/monitor",
			beego.NSInclude(
				&controllers.MonitorController{},
			),
		),

		beego.NSNamespace("/subsys",
			beego.NSInclude(
				&controllers.SubsysController{},
			),
		),

		beego.NSNamespace("/tag_meta",
			beego.NSInclude(
				&controllers.TagMetaController{},
			),
		),

		beego.NSNamespace("/tag_value",
			beego.NSInclude(
				&controllers.TagValueController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/view",
			beego.NSInclude(
				&controllers.ViewController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
