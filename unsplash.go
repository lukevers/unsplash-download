package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)


type Image struct {
	Format    string
	Width     int
	Height    int
	Filename  string
	Id        int
	Author    string
	Author_Url string
	Post_Url   string
}

func Download(image Image) {

	res, err := http.Get(image.Post_Url+"/download")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	err = ioutil.WriteFile("images/"+image.Filename, body, 0664)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Downloading "+images.Filename)
}

func main() {
	// GET all of our JSON
	res, err := http.Get("https://unsplash.it/list")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var images[]Image
	err = json.Unmarshal(body, &images)
	if err != nil {
		log.Fatal(err)
	}

	for _, img := range images {
		Download(img)
	}
}
