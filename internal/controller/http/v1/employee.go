package v1

import (
	"net/http"
	"strconv"

	"github.com/v1adhope/web-service-employees/internal/entity"
	"github.com/v1adhope/web-service-employees/internal/usecase"
)

type Routes struct {
	mux *http.ServeMux
	u   usecase.Employee
}

func NewRoutes(r *Routes) {
	r.mux.HandleFunc("/v1/add/", r.add)
	r.mux.HandleFunc("/v1/delete/", r.delete)
	r.mux.HandleFunc("/v1/all-by-company/", r.getAllByCompanyID)
	r.mux.HandleFunc("/v1/all-by-deportament/", r.getAllByDeportamentName)
	r.mux.HandleFunc("/v1/update-part/", r.updatePart)
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

func (rs *Routes) getAllByCompanyID(w http.ResponseWriter, r *http.Request) {
	rawID := r.URL.Query().Get("id")

	if rawID == "" {
		JSON(w, http.StatusBadRequest, map[string]any{
			"msg": "query is empty",
		})

		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		JSON(w, http.StatusBadRequest, map[string]any{
			"msg": "invalid id",
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
	name := r.URL.Query().Get("name")

	if name == "" {
		JSON(w, http.StatusBadRequest, map[string]any{
			"msg": "query is empty",
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
