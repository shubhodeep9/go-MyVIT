/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/


package login

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"net/url"
	"os"
	"os/exec"
	)



type Response struct {
	Regno  string `json:"regno"`
	Status int 	`json:"status"`
}

var Sessionname string

/*
Creates a new StudLogin object and Starts logging in
@return Response struct
@param Registration_Number Password
*/
func NewLogin(bow *browser.Browser, reg, pass, baseuri string) *Response {
	status := make(chan int)
	go DoLogin(bow,reg,pass,status, baseuri)
	success := <-status
	return &Response{
		Regno:  reg,
		Status: success,
	}
}

/*
Parses the captcha using parse.py and creates a session,
Using that session user is logged in.
@param bow(surf Browser) registration_no password status(channel for goroutine)
@return void
*/
func DoLogin(bow *browser.Browser, reg, pass string,status chan int, baseuri string) {
	bow.Open(baseuri+"/student/captcha.asp")
	out,_ :=os.Create("api/login/captcha_student.bmp")
	bow.Download(out)
	out1, err := exec.Command("python","api/login/parse.py").Output()
	if err != nil {
		status <- 0
	} else {
		capt := string(out1)[:len(out1)-1]
		v:= url.Values{}
		v.Set("regno",reg)
		v.Add("passwd",pass)
		v.Add("vrfcd",capt)
		v.Add("message","")
		bow.PostForm(baseuri+"/student/stud_login_submit.asp",v)
		stud_home := "/student/stud_home.asp"
		home := "/student/home.asp"
		u := bow.Url().EscapedPath()
		if u == stud_home || u == home {
			status <-1
		} else {
			status <- 0
		}
	}
}

