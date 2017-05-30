package scrape

import (
	//"go-MyVIT/api/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"fmt"
	"go-MyVIT/api/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"go-MyVIT/api/status"
	//"strconv"
	"strings"
	//"sync"
)

type MenuStruct struct {
	Status status.StatusStruct `json:"status"`
	Menu   []SubMenu           `json:"menu"`
}

type Entity struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
type SubMenu struct {
	Name    string   `json:"name"`
	Link    string   `json:"link"`
	Content []Entity `json:"content"`
}

func Parser(text, prefix string) []SubMenu {
	splitted := strings.Split(text, "\n")
	toggle := 0
	subMenus := []SubMenu{}
	tempy := SubMenu{}
	for i, val := range splitted {
		inSplit := strings.Split(val, ",")

		insideSelf := false

		if len(inSplit) > 3 && i > 8 {
			if !strings.Contains(inSplit[0], "p0") {
				insideSelf = true
			}

			if inSplit[3] == `"main"` {
				if insideSelf == false {
					temp := SubMenu{
						Name: strings.Trim(inSplit[1][1:], "\""),
						Link: prefix + strings.Trim(inSplit[2][1:], `\"`),
					}
					subMenus = append(subMenus, temp)

				} else {
					temp := Entity{
						Name: strings.Trim(inSplit[1][1:], "\""),
						Link: prefix + strings.Trim(inSplit[2][1:], `\"`),
					}
					//fmt.Println(temp)
					tempy.Content = append(tempy.Content, temp)
					fmt.Println(tempy)

				}
			} else {
				toggle = (toggle + 1) % 2
				if toggle == 0 {
					subMenus = append(subMenus, tempy)
					tempy = SubMenu{}
				}
				tempy.Name = strings.Trim(inSplit[1][1:], `\"`)
				tempy.Link = prefix + strings.Trim(inSplit[2][1:], `\"`)

			}

		}
	}
	return subMenus
}

func ShowMenu(bow *browser.Browser, reg, baseuri string, found bool) *MenuStruct {
	temp := []SubMenu{}
	stat := status.Success()
	if false {
		stat = status.SessionError()
	} else {
		//https://vtop.vit.ac.in/student/stud_menu.asp
		prefix := baseuri + "/student/"
		bow.Open(baseuri + "/student/stud_menu.asp")
		if bow.Open(baseuri+"/student/stud_menu.asp") == nil {
			temp = Parser(bow.Body(), prefix)
		}
	}

	return &MenuStruct{
		Status: stat,
		Menu:   temp,
	}
}
