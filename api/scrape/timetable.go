package main

import ( 
	//"fmt"
	"encoding/json"
	"os"
	)

type Timetable struct {
	Status string
	Time_table map[string]Contents
}

type Contents struct {
	Class_number string
	Course_code string
	Course_mode string
	Course_option string
}

func main() {
	con := Contents{
		Class_number: "234",
		Course_code: "423",
		Course_mode: "wer",
		Course_option: "dfg",
	}
	hmm := make(map[string]Contents)
	hmm["CSC"] = con
	time := &Timetable{
		Status: "success",
		Time_table: hmm,
	}
	b,_:=json.Marshal(time)
	//fmt.Println(json.Marshal(time))
	os.Stdout.Write(b)
	
}