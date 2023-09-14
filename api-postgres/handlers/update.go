package handlers

import (
	"api-postgres/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {

	//ex1
	in := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(in, 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//ex2
	// id, err := strconv.Atoi(chi.URLParam(r, "id"))
	// if err != nil {
	// 	log.Println("Erro ao fazer parse do id: %v", err)
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(id, todo)

	if err != nil {
		log.Println("Erro ao atualizar o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Println("Erro: foram atualizado  %d registro", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados atualizados com Sucesso!",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
