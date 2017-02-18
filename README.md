# go-MyVIT
The student login API for the app MyVIT in Go programming language<br />
[![Build Status](https://travis-ci.org/shubhodeep9/go-MyVIT.svg?branch=master)](https://travis-ci.org/shubhodeep9/go-MyVIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/shubhodeep9/go-MyVIT)](https://goreportcard.com/report/github.com/shubhodeep9/go-MyVIT)

##API URLs
```sh
#login
curl https://myffcs.in:10443/campus/[campus]/login --data "regNo=[reg]&psswd=[psswd]"
#refresh
curl https://myffcs.in:10443/campus/[campus]/refresh --data "regNo=[reg]&psswd=[psswd]"
#course page
curl https://myffcs.in:10443/campus/vellore/coursepage/data --data "regNo=[reg]&psswd=[psswd]&crs=[crscode]&slt=[slots]&fac=[facid]"
```
##Custom installation
Dependencies are already satisfied in Godep folder in api.<br />
Extra dependency to be installed: <br />
```sh
$ go get github.com/tools/godep
$ godep get
```
Set Environment variables: <br />
```sh
$ export SEM=[FS or WS]
$ export VITMONGOURL=[mlab uri]
$ export VITKEY=[KEY for gcm registration]
```
Execution<br />
```sh
$ bee run
```
> Use cron jobs for continuous serving, a module already has been built [MyVIT-Cron](https://github.com/shubhodeep9/MyVIT-Cron)

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
