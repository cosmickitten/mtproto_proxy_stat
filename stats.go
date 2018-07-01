package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// User struct contains num field
type User struct {
	Num string
}

// HTML struct contains html template of index page
type HTML struct {
	IndexPage string
}

// Users contain count of current users
var Users User

// IndexTemplate contain index page html template
var IndexTemplate = HTML{
	IndexPage: `<!DOCTYPE html>
	<html>
	
	<head>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css" integrity="sha384-9gVQ4dYFwwWSjIDZnLEWnxCjeSWFphJiwGPXr1jddIhOegiu1FwO5qRGvFXOdJZ4"
			crossorigin="anonymous">
		<title>Stats</title>
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	</head>
	
	<body>
		<div class="container-fluid">
			<div class="row justify-content-center text-center" style="margin-top: 20%">
				<h1>Count of current users of MTProto Proxy: {{.Num}}</h1>
			</div>
		</div>
	</body>
	
	</html>`,
}

// CurrenUsers fetching stats from mtproto proxy
func CurrenUsers() (err error) {
	response, err := http.Get(`http://localhost:2398/stats`)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	defer response.Body.Close()
	stat := strings.Split(string(body), "\n")
	for _, item := range stat {
		if strings.HasPrefix(item, `total_special_connections`) {
			Users.Num = strings.Split(item, "\t")[1]
		}
	}
	return nil
}

func sendStat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.Must(template.New("indexpage").Parse(IndexTemplate.IndexPage))
		t.Execute(w, Users)
	}
}

func init() {
	go func() {
		for {
			defer runtime.GC()
			CurrenUsers()
			time.Sleep(10 * time.Second)
		}
	}()
}

func main() {
	http.HandleFunc("/", sendStat)
	http.ListenAndServe(":80", nil)
}
