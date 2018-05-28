package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type GPSdata struct {
	Latitude  string
	Longitude string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	//fmt.Fprintf(w, "I love GO %s\n", html.EscapeString(r.URL.Path[1:]))
	if r.Method == "GET" {
		//fmt.Println("method:", r.Method) //获取请求的方法
		//fmt.Println("username", r.Form["username"])
		//fmt.Println("password", r.Form["password"])
		uploadTemplate := template.Must(template.ParseFiles("welcome.gtpl"))
		uploadTemplate.Execute(w, nil)

		for k, v := range r.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)

		//未知类型的推荐处理方法

		var f interface{}
		json.Unmarshal(result, &f)
		m := f.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

		//结构已知，解析到结构体

		var s GPSdata
		json.Unmarshal([]byte(result), &s)
		fmt.Fprintf(w, "Latitude: %s\n", s.Latitude)
		fmt.Fprintf(w, "Longitude: %s\n", s.Longitude)
		//fmt.Println(s.ServersID)
	}
}
