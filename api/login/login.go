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
	"crypto/tls"
	"fmt"
	"github.com/patrickmn/go-cache"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/cache"
	"go-MyVIT/api/login/captcha"
	"go-MyVIT/api/status"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Response struct {
	Regno  string              `json:"regno"`
	Campus string              `json:"campus"`
	Status status.StatusStruct `json:"status"`
}

var Sessionname string

/*
Creates a new StudLogin object and Starts logging in
@return Response struct
@param Registration_Number Password
*/
func NewLogin(bow *browser.Browser, reg, pass, baseuri string, cac *cache.Cache) *Response {
	stats := make(chan int)
	go DoLogin(bow, reg, pass, stats, baseuri, cac)
	success := <-stats
	var stt status.StatusStruct

	if success == 1 {
		stt = status.Success()
	} else if success == 2 {
		stt = status.ServerError()
	} else {
		stt = status.CredentialError()
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
func DoLogin(bow *browser.Browser, reg, pass string, stats chan int, baseuri string, cac *cache.Cache) {

	if bow.Open("https://vtop.vit.ac.in/student/captcha.asp") != nil {
		stats <- 2
	} else {
		out, err := os.Create("api/login/" + reg + ".bmp")
		bow.Download(out)
		out1 := captcha.Parse(reg)
		go os.Remove("api/login/" + reg + ".bmp")
		if err != nil {
			stats <- 0
		} else {
			v := url.Values{}
			v.Set("regno", reg)
			v.Add("passwd", pass)
			v.Add("vrfcd", out1)
			v.Add("message", "")
			bow.PostForm(baseuri+"/student/stud_login_submit.asp", v)
			stud_home := baseuri + "/student/stud_home.asp"
			home := baseuri + "/student/home.asp"
			u := bow.Url().String()
			if u == stud_home || u == home {
				cac.Set(reg, &cacheSession.MemCache{Regno: reg, MemCookie: bow.SiteCookies()}, cache.DefaultExpiration)
				stats <- 1
			} else {
				stats <- 0
			}
		}
	}
}

/*
Method to login into vtopbeta2
Uses a different captcha parser
*/

func LoginVtopBeta(client http.Client, regNo, psswd string) http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	jar, _ := cookiejar.New(nil)
	client = http.Client{
		Jar:       jar,
		Transport: tr,
	}
	postData := url.Values{}
	req, _ := http.NewRequest("GET", "https://vtopbeta.vit.ac.in/vtop/", strings.NewReader(postData.Encode()))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_10_5) AppleWebKit (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html := string(body)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((html)))
	var base64 string
	doc.Find("img[alt='vtopCaptcha']").Each(func(i int, inp *goquery.Selection) {
		b64, _ := inp.Attr("src")
		b64 = strings.Split(b64, ",")[1]
		base64 = b64
	})

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("python", "parseIt.py", base64)
	cmd.Dir = dir + "/api/login/CaptchaVtopBeta/"
	output, err2 := cmd.CombinedOutput()
	if err2 != nil {
		fmt.Println(fmt.Sprint(err2) + ":" + string(output))
	}

	var captcha string
	if err2 == nil {
		captchaFile, err3 := ioutil.ReadFile(cmd.Dir + "output.txt")
		if err3 == nil {
			captcha = string(captchaFile)
			fmt.Println("Captcha is", captcha)
		} else {
			fmt.Println(err3)
		}

	}
	postData.Add("uname", ""+regNo)
	postData.Add("passwd", ""+psswd)
	postData.Add("captchaCheck", captcha)
	//fmt.Println("Post data = ",postData)
	PostData2 := strings.NewReader("uname="+regNo+"&passwd="+psswd+"&captchaCheck=" + captcha)
	req2, _ := http.NewRequest("POST", "https://vtopbeta.vit.ac.in/vtop/processLogin", PostData2)
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err = client.Do(req2)
	body,_=ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	html = string(body)
	//fmt.Println("The login's response")
	//fmt.Println("HTML on login :- ",html)
	return client

}
