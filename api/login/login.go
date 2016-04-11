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
	"fmt"
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

type MemCache struct {
	Regno     string
	MemCookie []*http.Cookie
}

type Response struct {
	Regno  string `json:"regno"`
	Status int    `json:"status"`
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
func DoLogin(bow *browser.Browser, reg, pass string, status chan int, baseuri string, cac *cache.Cache) {

	fmt.Println(bow.Open("https://academics.vit.ac.in/student/captcha.asp"))
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
		fmt.Println(bow.Url())
		cac.Set(reg, &MemCache{Regno: reg, MemCookie: bow.SiteCookies()}, cache.DefaultExpiration)
		stud_home := "/student/stud_home.asp"
		home := "/student/home.asp"
		u := bow.Url().EscapedPath()
		if u == stud_home || u == home {
			status <- 1
		} else {
			status <- 0
		}
	}
}
