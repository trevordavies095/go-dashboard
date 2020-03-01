package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	Name string `json:"name"`
}

// Global constants

// Global variables
var tpl *template.Template
var config Config

func init() {
	// Local constants

	// Local variables

	/****** start init() ******/

	tpl = template.Must(template.ParseGlob("templates/*"))
	config = openConfig()

}

func main() {
	// Local constants

	// Local variables

	/****** start main() ******/

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	// Local constants

	// Local variables

	/****** start index() ******/
	tpl.ExecuteTemplate(w, "index.gohtml", config)
}

func openConfig() Config {
	// Local constants

	// Local variables
	configFile, err := os.Open("config.json")
	var config Config

	/****** start open_config() ******/

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(byteValue, &config)

	configFile.Close()
	return config
}
