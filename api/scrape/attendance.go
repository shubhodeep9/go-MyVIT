package scrape

func ShowAttendance(bow *browser.Browser,regno, password, baseuri string) {
	response := login.NewLogin(bow,regno,password,baseuri)
	status := "Success" 
	if response.Status != 1 {
		status = "Failure"
	} else {
		bow.Open(baseuri+"/student/timetable_ws.asp")
		//Twice loading due to Redirect policy defined by academics.vit.ac.in
		bow.Open(baseuri+"/student/timetable_ws.asp")
	}
}