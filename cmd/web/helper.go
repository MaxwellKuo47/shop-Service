package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

//clientError help to log information and return specific status code
func (app *AppConfig) clientError(w http.ResponseWriter, r *http.Request, status int, logString string) {
	outputText := http.StatusText(status)
	app.clientErrLogger.Printf("ClientError => %q -> %q response '%d %s' : %q", r.Method, r.URL.Path, status, outputText, logString)
	http.Error(w, outputText, status)
}

//serverError help to log the error and return status code InternalServerError
func (app *AppConfig) serverError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("ServerError => %q -> %q : %q\n%s", r.Method, r.URL.Path, err.Error(), debug.Stack())
	app.serverErrLogger.Output(2, trace)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

//handleUploadImage parsing form files and copy to the specific folder.
func handleUploadImage(uuids []string, fileNames []string, time time.Time, desFolder string, r *http.Request) ([]string, error) {
	timeString := time.Format("20060102150405")
	var storeNames []string
	for index, fileName := range fileNames {
		file, _, err := r.FormFile(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		storeName := timeString + uuids[index] + ".jpeg"
		fp, err := os.Create(desFolder + storeName)
		if err != nil {
			return nil, err
		}
		defer fp.Close()
		_, err = io.Copy(fp, file)
		if err != nil {
			return nil, err
		}
		storeNames = append(storeNames, storeName)
	}
	return storeNames, nil
}
