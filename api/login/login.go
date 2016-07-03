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
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/cache"
	"net/url"
	"os"
	"os/exec"
)

type Response struct {
	Regno  string       `json:"regno"`
	Campus string       `json:"campus"`
	Status StatusStruct `json:"status"`
}

type StatusStruct struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var Sessionname string

/*
Creates a new StudLogin object and Starts logging in
@return Response struct
@param Registration_Number Password
*/
func NewLogin(bow *browser.Browser, reg, pass, baseuri string, cac *cache.Cache) *Response {
	status := make(chan int)
	go DoLogin(bow, reg, pass, status, baseuri, cac)
	success := <-status
	var stt StatusStruct
	if success == 1 {
		stt = StatusStruct{
			Message: "Successful Execution",
			Code:    0,
		}
	} else {
		stt = StatusStruct{
			Message: "UnSuccessful Execution",
			Code:    12,
		}
	}
	return &Response{
		Regno:  reg,
		Campus: "vellore",
		Status: stt,
	}
}

/*
Parses the captcha using parse.py and creates a session,
Using that session user is logged in.
@param bow(surf Browser) registration_no password status(channel for goroutine)
@return void
*/
func DoLogin(bow *browser.Browser, reg, pass string, status chan int, baseuri string, cac *cache.Cache) {

	bow.Open("https://vtop.vit.ac.in/student/captcha.asp")
	out, _ := os.Create("api/login/" + reg + ".bmp")
	bow.Download(out)
	out1, err := exec.Command("python", "api/login/parse.py", reg).Output()
	go os.Remove("api/login/" + reg + ".bmp")
	if err != nil {
		status <- 0
	} else {

		capt := string(out1)[:len(out1)-1]
		v := url.Values{}
		v.Set("regno", reg)
		v.Add("passwd", pass)
		v.Add("vrfcd", capt)
		v.Add("message", "")
		bow.PostForm(baseuri+"/student/stud_login_submit.asp", v)
		stud_home := "/student/stud_home.asp"
		home := "/student/home.asp"
		u := bow.Url().EscapedPath()
		if u == stud_home || u == home {
			cac.Set(reg, &cacheSession.MemCache{Regno: reg, MemCookie: bow.SiteCookies()}, cache.DefaultExpiration)
			status <- 1
		} else {
			status <- 0
		}
	}
}
