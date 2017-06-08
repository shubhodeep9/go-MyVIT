package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["go-MyVIT/controllers:CoursePageController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:CoursePageController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:FacultyController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:FacultyController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:GcmMessageController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:GcmMessageController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})
	beego.GlobalControllerRouter["go-MyVIT/controllers:MenuController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:MenuController"],
			beego.ControllerComments{
				"Get",
				`/`,
				[]string{"get"},
				nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})
	
	beego.GlobalControllerRouter["go-MyVIT/controllers:CalCoursesController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:CalCoursesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})
	
	
	beego.GlobalControllerRouter["go-MyVIT/controllers:FacultyPhotosController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:FacultyPhotosController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})
	beego.GlobalControllerRouter["go-MyVIT/controllers:ExamScheduleController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:ExamScheduleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:LeaveRequestController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:LeaveRequestController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:MarksController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:MarksController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:PicController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:PicController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:RefreshController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:RefreshController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:SpotlightController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:SpotlightController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})
	beego.GlobalControllerRouter["go-MyVIT/controllers:MessageController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:MessageController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:PersonalDetailsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:PersonalDetailsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:EducationalDetailsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:EducationalDetailsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})
	beego.GlobalControllerRouter["go-MyVIT/controllers:FamilyDetailsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:FamilyDetailsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})
	beego.GlobalControllerRouter["go-MyVIT/controllers:Timetable2Controller"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:Timetable2Controller"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:RoomAllotController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:RoomAllotController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:HostelDetailsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:HostelDetailsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:StatsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:StatsController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

}
