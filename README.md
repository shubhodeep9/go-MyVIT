# go-MyVIT
The student login API for the app MyVIT in Go programming language<br />
[![Build Status](https://travis-ci.org/shubhodeep9/go-MyVIT.svg?branch=master)](https://travis-ci.org/shubhodeep9/go-MyVIT)

##API URLs
<ul>
<li>
LogIn: http://myffcs.in:8080/campus/[campus]/login --data "regNo=[reg]&psswd=[psswd]"
</li>
<li>
Refresh: http://myffcs.in:8080/campus/[campus]/refresh --data "regNo=[reg]&psswd=[psswd]"
</li>
<li>
CoursePage: http://myffcs.in:8080/campus/vellore/coursepage/data --data "regNo=[reg]&psswd=[psswd]&crs=[crscode]&slt=[slots]&fac=[facid]"
</li>
</ul>
##Custom installation
Dependencies are already satisfied in Godep folder in api.<br />
Extra dependency to be installed: <br />
```sh
$ go get github.com/astaxie/beego
$ go get github.com/patrickmn/go-cache
$ go get github.com/smartystreets/goconvey/convey
$ go get github.com/hotei/bmp
$ go get github.com/alexjlockwood/gcm
$ go get gopkg.in/mgo.v2
```
Execution<br />
```sh
$ bee run
```

##Features
<ul>
<li>CaptchaParser(Karthik Balakrishnan Algorithm implemented on Golang)</li>
<li>100% Go Code</li>
<li>Beautiful goroutines written</li>
</ul>

## Contributors
<a href="https://github.com/shubhodeep9">Shubhodeep Mukherjee</a><br />
<a href="https://github.com/JiraiyaFool">Ujjwal Ayyangar</a>

## Credits
Twenty One Pilots, without their songs this project would have been a disaster.
