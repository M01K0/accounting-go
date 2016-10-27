package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexyslozada/accounting-go/common"
	"github.com/alexyslozada/accounting-go/dao/executedao"
	"github.com/alexyslozada/accounting-go/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CostCenterInsert(w http.ResponseWriter, r *http.Request) {
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

func CostCenterUpdate(w http.ResponseWriter, r *http.Request) {
	var ccr CostCenterResource
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.DisplayError(w, err, fmt.Sprintf("El id del centro de costo nó es válido, %s", vars["id"]), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&ccr)
	if err != nil {
		common.DisplayError(w, err, "Estructura no válida", http.StatusBadRequest)
		return
	}

	cc := &ccr.Data
	cc.ID = int16(id)
	err = executedao.CostCenterDAO.Update(cc)
	if err != nil {
		common.DisplayError(w, err, "No se pudo actualizar el registro", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(CostCenterResource{*cc})
	if err != nil {
		common.DisplayError(w, err, "Error inesperado al convertir a json la respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func CostCenterGetByID(w http.ResponseWriter, r *http.Request) {
	var ccr CostCenterResource
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.DisplayError(w, err, fmt.Sprintf("El id del centro de costo nó es válido, %s", vars["id"]), http.StatusBadRequest)
		return
	}

	cc, err := executedao.CostCenterDAO.GetByID(int16(id))
	ccr.Data = *cc
	if err != nil {
		common.DisplayError(w, err, "No se encontró el registro", http.StatusNotFound)
		return
	}

	j, err := json.Marshal(ccr)
	if err != nil {
		common.DisplayError(w, err, "Error inesperado al convertir a json la respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func CostCenterGetAll(w http.ResponseWriter, r *http.Request) {
	ccs := []models.CostCenter{}
	err := errors.New("")
	hasPagination := common.QueryValues(r, "pagination")
	if hasPagination != nil {
		page := common.QueryValues(r, "page")[0]
		limit := common.QueryValues(r, "limit")[0]
		orderBy := common.QueryValues(r, "orderby")[0]
		orderType := common.QueryValues(r, "ordertype")[0]
		p, _ := strconv.Atoi(page)
		l, _ := strconv.Atoi(limit)
		o, _ := strconv.Atoi(orderBy)
		ccs, err = executedao.CostCenterDAO.GetAllPagination(p, l, o, orderType)
		if err != nil {
			common.DisplayError(w, err, "No se encontraron registros", http.StatusNotFound)
			return
		}
	} else {
		ccs, err = executedao.CostCenterDAO.GetAll()
		if err != nil {
			common.DisplayError(w, err, "No se encontraron registros", http.StatusNotFound)
			return
		}
	}

	j, err := json.Marshal(CostCentersResource{Data: ccs})
	if err != nil {
		common.DisplayError(w, err, "Error inesperado al convertir a json la respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func CostCenterDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.DisplayError(w, err, fmt.Sprintf("El id del centro de costo nó es válido, %s", vars["id"]), http.StatusBadRequest)
		return
	}

	err = executedao.CostCenterDAO.Delete(int16(id))
	if err != nil {
		common.DisplayError(w, err, "No se pudo borrar el registro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
