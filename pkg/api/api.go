package api

import (
	"net/http"
)

const FormatDate string = "20060102"

func Init(router *http.ServeMux) {
	fs := http.FileServer(http.Dir("./web"))
	router.HandleFunc("/api/nextdate", NextDayHandler)
	router.Handle("/css/", fs)
	router.Handle("/js/", fs)
	router.Handle("/", fs)
}
