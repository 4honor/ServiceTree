package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagValueController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:ViewController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:GroupController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:SubsysController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MonitorController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:FavorController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:TagMetaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:UserController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:UserController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:UserController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:UserController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:UserController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:IdallocController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:IdallocController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:IdallocController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:IdallocController"],
		beego.ControllerComments{
			"Alloc",
			`/:api`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"] = append(beego.GlobalControllerRouter["ServiceTree/controllers:MachineController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
