package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetStatusHedg(w http.ResponseWriter, r *http.Request) {
	accountValue, err := h.repository.GetAccount()
	if err != nil {
		http.Error(w, "Ошибка получения данных пользователя", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accountValue.SpotHedgingStatus); err != nil {
		http.Error(w, "Ошибка кодирования данных пользователя", http.StatusInternalServerError)
		return
	}
}
