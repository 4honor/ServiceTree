// @APIVersion 0.0.1
// @Title 服务树接口列表
// @Description  服务树对外接口
// @Contact yujinqiu@diditaxi.com.cn
// @TermsOfServiceUrl http://odin.xiaojukeji.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ServiceTree/controllers"

	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/:subsys", &controllers.PageController{})

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

		beego.NSNamespace("/tag_key",
			beego.NSInclude(
				&controllers.TagKeyController{},
			),
		),
        beego.NSNamespace("/tree",
            beego.NSInclude(
                &controllers.TreeController{}, 
            ),
        ),
	)
	beego.AddNamespace(ns)
}
