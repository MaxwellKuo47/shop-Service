package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const MAX_FORM_ACCEPT = 10 << 20 //10MB
const PRODUCT_IMAGE_FOLDER = "/home/workspace/code/ReactPractice/Image/"

//productAdd handle products adding proocess
func (app *AppConfig) productAdd(w http.ResponseWriter, r *http.Request) {
	type ColorSizeElement struct {
		Color        string `json:"color"`
		AmountSizeOS int    `json:"OS"`
		AmountSizeS  int    `json:"S"`
		AmountSizeM  int    `json:"M"`
		AmountSizeL  int    `json:"L"`
		AmountSizeXL int    `json:"XL"`
	}
	type RequstJsonInfo struct {
		ProdName     string             `json:"productName"`
		ProdDes      string             `json:"productDes"`
		ProdMktPrice int                `json:"productMktPrice"`
		ProdSalPrice int                `json:"productSalPrice"`
		ProdAmount   []ColorSizeElement `json:"productAmount"`
		ProdTags     []string           `json:"productTags"`
	}

	err := r.ParseMultipartForm(MAX_FORM_ACCEPT)
	if err != nil {
		app.clientError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var reqJSON RequstJsonInfo
	err = json.Unmarshal([]byte(r.PostForm.Get("info")), &reqJSON)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	fmt.Println(reqJSON)

	//info should contian files name (Request side should change)
	//Here using the front-end side uuid(fileIDs) as the file name for temp unique file name solution
	fileIDs := strings.Split(r.PostForm.Get("fileIDs"), ",")
	var lgFiles, smFiles []string //files identifier for formValue
	for i := 0; i < len(fileIDs); i++ {
		numStr := strconv.Itoa(i)
		lgFiles = append(lgFiles, "fileLg"+numStr) //fileLg1, fileLg2...(Which is form key of files set at front-end)
		smFiles = append(smFiles, "fileSm"+numStr) //fileSm1, fileSm2...
	}

	//Time is used to add to the name.
	//storeNames contain the name we want to store in the DB
	curTime := time.Now()
	var storeNames []string
	storeNames, err = handleUploadImage(fileIDs, lgFiles, curTime, "imageLg/", r)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	//Both Large or Small images name are same
	_, err = handleUploadImage(fileIDs, smFiles, curTime, "imageSm/", r)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	fmt.Println(storeNames)

}
