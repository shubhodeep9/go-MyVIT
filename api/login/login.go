package login

import (
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

/*
Interface definition for StudentLogin,
*/
type StudLogin interface {
	DoLogin(status chan int)
	setSession()
	GetCookies() []*http.Cookie
}

/*
Type structure of Login
@param Registration_no Password Session(ASPSESSIONIDQAWCRCSR) Cookies(@keys ASPSESSIONIDQAWCRCSR logstudregno )
*/
type Login struct {
	regno    string
	password string
	Session  string
	Cookies  []*http.Cookie
}

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
	newLogin := &Login{
		regno:    reg,
		password: pass,
		Session:  "",
	}
	status := make(chan int)
	go newLogin.DoLogin(status)
	success := <-status
	if success == 1 {
		bow.SetSiteCookies(newLogin.GetCookies())
	}
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
func (l *Login) DoLogin(status chan int) {
	exec.Command("python", "api/login/login.py", l.regno, l.password).Output()
	//fmt.Println(s)
	dat, _ := ioutil.ReadFile("api/login/"+l.regno+".txt")
	exec.Command("rm","api/login/"+l.regno+".txt").Output()
	if strings.Contains(string(dat), "14BCS0002") {
		index := strings.Index(string(dat), "ASPSESSION") + 21
		Sessionname = string(dat)[index-21:index-1]
		l.Session = string(dat)[index : index+24]
		l.setSession()
		status <- 1
	} else {
		status <- 0
	}
}

/*
Sets ASPSESSIONIDQAWCRCSR logstudregno as cookies in http.CookieJar
Implements *http.Cookie struct
*/
func (l *Login) setSession() {
	session := &http.Cookie{
		Name:   Sessionname,
		Value:  l.Session,
		Path:   "/",
		Domain: "academics.vit.ac.in",
	}
	logregno := &http.Cookie{
		Name:   "logstudregno",
		Value:  l.regno,
		Path:   "/student",
		Domain: "academics.vit.ac.in",
	}
	l.Cookies = append(l.Cookies, session)
	l.Cookies = append(l.Cookies, logregno)
}

func (l *Login) GetCookies() []*http.Cookie {
	return l.Cookies
}
