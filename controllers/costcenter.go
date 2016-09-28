package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/alexyslozada/accounting-go/common"
	"github.com/alexyslozada/accounting-go/dao/executedao"
)

func CostCenterCreate(w http.ResponseWriter, r *http.Request) {
	var ccr CostCenterResource
	err := json.NewDecoder(r.Body).Decode(&ccr)
	if err != nil {
		common.DisplayError(w, err, "Estructura no válida", http.StatusBadRequest)
		return
	}

	cc := &ccr.Data
	err = executedao.CostCenterDAO.Insert(cc)
	if err != nil {
		common.DisplayError(w, err, "No se pudo insertar el registro", http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(CostCenterResource{*cc})
	if err != nil {
		common.DisplayError(w, err, "Error inesperado al convertir a json la respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
