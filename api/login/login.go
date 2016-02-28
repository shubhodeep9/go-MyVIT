package login

import (
	"net/http"
	"os/exec"
	"io/ioutil"
	"strings"
	"github.com/headzoo/surf/browser"
	"fmt"
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
	regno string
	password string
	Session string
	Cookies []*http.Cookie
}

/*
Creates a new StudLogin object and Starts logging in
@return Login struct
@param Registration_Number Password
*/
func NewLogin(bow *browser.Browser,reg,pass string) {
	newLogin := &Login{
		regno: reg,
		password: pass,
		Session: "",
	}
	status := make(chan int)
	go newLogin.DoLogin(status)
	success := <-status
	if success==1 {
		fmt.Println("Success")
		bow.SetSiteCookies(newLogin.GetCookies())
	} else {
		fmt.Println("Try again")
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
	exec.Command("python","login/login.py",l.regno,l.password).Output()
	dat, _ := ioutil.ReadFile("login/cookie.txt")
	if strings.Contains(string(dat),"14BCS0002") {
		index := strings.Index(string(dat),"ASPSESSIONIDQAWCRCSR")+21
		l.Session = string(dat)[index:index+24]
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
	session := &http.Cookie {
		Name: "ASPSESSIONIDQAWCRCSR",
		Value: l.Session,
		Path: "/",
		Domain: "academics.vit.ac.in",
	}
	logregno := &http.Cookie {
		Name: "logstudregno",
		Value: l.regno,
		Path: "/student",
		Domain: "academics.vit.ac.in",
	}
	l.Cookies= append(l.Cookies,session)
	l.Cookies = append(l.Cookies,logregno)
}

func (l *Login) GetCookies() []*http.Cookie {
	return l.Cookies
}
