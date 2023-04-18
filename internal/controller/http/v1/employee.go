// INFO: No error handling, no validation
package v1

import (
	"net/http"

	"github.com/v1adhope/web-service-employees/internal/entity"
	"github.com/v1adhope/web-service-employees/internal/usecase"
)

type Routes struct {
	mux *http.ServeMux
	u   usecase.Employee
}

func NewRoutes(r *Routes) {
	r.mux.HandleFunc("/v1/employee/", r.employee)
	r.mux.HandleFunc("/v1/employees-by-company/", r.getAllByCompanyID)
	r.mux.HandleFunc("/v1/employees-by-deportament/", r.getAllByDeportamentName)
}

func (rs *Routes) employee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		rs.add(w, r)
	case http.MethodDelete:
		rs.delete(w, r)
	case http.MethodPatch:
		rs.updatePart(w, r)
	default:
		rs.writeNotFound(w)
	}
}

func (rs *Routes) add(w http.ResponseWriter, r *http.Request) {
	var emp entity.Employee

	err := BindJSON(w, r, &emp)
	if err != nil {
		return
	}

	id, err := rs.u.Create(r.Context(), emp)
	if err != nil {
		JSON(w, http.StatusTeapot, map[string]any{
			"msg": err.Error(),
		})

		return
	}

	JSON(w, http.StatusCreated, map[string]any{
		"id": id,
	})
}

func (rs *Routes) delete(w http.ResponseWriter, r *http.Request) {
	dto := deleteDTO{}

	err := BindJSON(w, r, &dto)
	if err != nil {
		return
	}

	err = rs.u.DeleteByID(r.Context(), dto.ID)
	if err != nil {
		JSON(w, http.StatusTeapot, map[string]any{
			"msg": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rs *Routes) updatePart(w http.ResponseWriter, r *http.Request) {
	var emp entity.Employee

	err := BindJSON(w, r, &emp)
	if err != nil {
		return
	}

	err = rs.u.UpdateByID(r.Context(), emp)
	if err != nil {
		JSON(w, http.StatusTeapot, map[string]any{
			"msg": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rs *Routes) getAllByCompanyID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rs.writeNotFound(w)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		JSON(w, http.StatusBadRequest, map[string]any{
			"msg": _errQueryIsEmpty.Error(),
		})

		return
	}

	emp, err := rs.u.GetByCompanyID(r.Context(), id)
	if err != nil {
		JSON(w, http.StatusTeapot, map[string]any{
			"msg": err.Error(),
		})

		return
	}

	JSON(w, http.StatusOK, emp)
}

func (rs *Routes) getAllByDeportamentName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rs.writeNotFound(w)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		JSON(w, http.StatusBadRequest, map[string]any{
			"msg": _errQueryIsEmpty.Error(),
		})

		return
	}

	emp, err := rs.u.GetByDeportamentName(r.Context(), name)
	if err != nil {
		JSON(w, http.StatusTeapot, map[string]any{
			"msg": err.Error(),
		})

		return
	}

	JSON(w, http.StatusOK, emp)
}
