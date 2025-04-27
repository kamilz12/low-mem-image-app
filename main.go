//go:build !windows
package main

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed index.html
var indexTmpl string

//go:embed weather.html
var weatherTmpl string

var (
	author   = "Kamil Ziolkowski"
	owmKey   = os.Getenv("OWM_KEY")
	port     = os.Getenv("PORT") // domyślnie 8000
	client   = &http.Client{Timeout: 5 * time.Second}
	cityList = [][2]string{
		{"PL", "Warsaw"}, {"DE", "Berlin"},
		{"US", "New York"}, {"JP", "Tokyo"},
	}
)

func main() {
	if port == "" {
		port = "8000"
	}
	log.Printf("Start %s | Autor: %s | Port: %s",
		time.Now().UTC().Format(time.RFC3339), author, port)

	tmpl := template.Must(template.New("i").Parse(indexTmpl))
	template.Must(tmpl.New("w").Parse(weatherTmpl))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			http.Redirect(w, r, "/weather/"+r.FormValue("pair"), http.StatusSeeOther)
			return
		}
		tmpl.ExecuteTemplate(w, "i", cityList)
	})

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		pair := r.URL.Path[len("/weather/"):]
		cc, city, _ := cut(pair, '|')
		url := "https://api.openweathermap.org/data/2.5/weather?q=" +
			city + "," + cc + "&units=metric&lang=pl&appid=" + owmKey
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "WeatherApp/1.0")
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), 500); return
		}
		defer res.Body.Close()
		var j map[string]any
		json.NewDecoder(res.Body).Decode(&j)
		if res.StatusCode != 200 {
			http.Error(w, "API error", 500); return
		}
		tmpl.ExecuteTemplate(w, "w", j)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func cut(s string, sep byte) (left, right string, ok bool) {
	if i := len(s); i > 0 {
		for j := 0; j < i; j++ {
			if s[j] == sep {
				return s[:j], s[j+1:], true
			}
		}
	}
	return
}