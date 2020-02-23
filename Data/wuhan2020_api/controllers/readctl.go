package controllers

import (
	"net/http"

	"github.com/arrebole/wuhan2020_api/model"
	"github.com/arrebole/wuhan2020_api/service"
)

// ReadCtl ...
func ReadCtl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	city, province := r.FormValue("city"), r.FormValue("province")
	if city == "" && province == "" {
		w.Write(model.FailResp("miss query city or province"))
		return
	}
	if province != "" {
		w.Write(model.ArchivesResp(service.GetArchivesByProvince(province)))
		return
	}
	w.Write(model.ArchivesResp(service.GetArchivesByCity(city)))
}
