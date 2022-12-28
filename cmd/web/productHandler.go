package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const MAX_FORM_ACCEPT = 10 << 20 //10MB
const PRODUCT_IMAGE_FOLDER  = "/home/workspace/code/ReactPractice/Image/"

var wg sync.WaitGroup

func (app *AppConfig) productAdd (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	err := r.ParseMultipartForm(MAX_FORM_ACCEPT) 
	if err != nil {
		app.clientError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(r.PostForm.Get("info"))
	wg.Add(2)
	go func (size string) {
		for i := 0; i < 5; i++ {
			num := strconv.Itoa(i)
			file, _, err := r.FormFile("file"+ size +num)
			fmt.Println("file"+num)
			if err != nil {
				app.clientError(w, r, http.StatusBadRequest, err.Error())
				return
			}
			body, err := ioutil.ReadAll(file)
			if err != nil {
				app.clientError(w, r, http.StatusBadRequest, err.Error())
				return
			}
			err = ioutil.WriteFile("./file"+ size +num+".jpeg", body, 0644)
			if err != nil {
				app.clientError(w, r, http.StatusBadRequest, err.Error())
				return
			}
		}
		wg.Done()
	}("Lg")
	go func (size string) {
		for i := 0; i < 5; i++ {
			num := strconv.Itoa(i)
			file, _, err := r.FormFile("file"+ size +num)
			fmt.Println("file"+num)
			if err != nil {
				app.clientError(w, r, http.StatusBadRequest, err.Error())
				return
			}
			body, err := ioutil.ReadAll(file)
			if err != nil {
				app.clientError(w, r, http.StatusBadRequest, err.Error())
				return
			}
			err = ioutil.WriteFile("./file"+ size +num+".jpeg", body, 0644)
			if err != nil {
				app.clientError(w, r, http.StatusBadRequest, err.Error())
				return
			}
		}
		wg.Done()
	}("Sm")
	wg.Wait()
} 