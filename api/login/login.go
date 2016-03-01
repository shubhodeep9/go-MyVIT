package login

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"net/url"
	"os"
	"os/exec"
	)

/*
Type structure of Login
@param Registration_no Password Session(ASPSESSIONIDQAWCRCSR) Cookies(@keys ASPSESSIONIDQAWCRCSR logstudregno )
*/

type Response struct {
	Regno  string `json:"regno"`
	Status int 	`json:"status"`
}

var Sessionname string

/*
Creates a new StudLogin object and Starts logging in
@return Login struct
@param Registration_Number Password
*/
func NewLogin(bow *browser.Browser, reg, pass string) *Response {
	status := make(chan int)
	go DoLogin(bow,reg,pass,status)
	success := <-status
	return &Response{
		Regno:  reg,
		Status: success,
	}
}

/*
Executes the login.py script to create a session,
login.py:
	@Libraries: Mechanize BeautifulSoup
	@param regno password
	@returns ASPSESSIONIDQAWCRCSR to cookie.txt
*/
func DoLogin(bow *browser.Browser, reg, pass string,status chan int) {
	bow.Open("https://academics.vit.ac.in/student/captcha.asp")
	out,_ :=os.Create("api/login/captcha_student.bmp")
	bow.Download(out)
	out1, _ := exec.Command("python","api/login/parse.py").Output()
	capt := string(out1)[:len(out1)-1]
	v:= url.Values{}
	v.Set("regno",reg)
	v.Add("passwd",pass)
	v.Add("vrfcd",capt)
	v.Add("message","")
	bow.PostForm("https://academics.vit.ac.in/student/stud_login_submit.asp",v)
	stud_home := "/student/stud_home.asp"
	home := "/student/home.asp"
	u := bow.Url().EscapedPath()
	if u == stud_home || u == home {
		status <-1
	} else {
		status <- 0
	}
}

