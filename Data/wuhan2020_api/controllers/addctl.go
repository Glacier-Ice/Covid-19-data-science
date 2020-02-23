package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/arrebole/wuhan2020_api/model"
	"github.com/arrebole/wuhan2020_api/service"
)

// AddCtl ...
func AddCtl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.Write(model.FailResp("Only post method allowed"))
		return
	}

	var postdata model.PostData
	err := json.NewDecoder(r.Body).Decode(&postdata)
	if err != nil {
		w.Write(model.FailResp("Type error"))
		return
	}

	if !postdata.IsLegal() {
		w.Write(model.FailResp("Missing data field"))
		return
	}
	if service.SaveArchive(postdata.GetArchive()) == 1 {
		w.Write(model.SuccessResp("update success"))
		return
	}
	service.SaveLog(postdata.GetLog(r.RemoteAddr))
	w.Write(model.SuccessResp("add success"))
}
