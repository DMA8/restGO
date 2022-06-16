package entrypoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testTask/internal/domain"
)

func (h *Handler) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func (h *Handler) getByIDs(w http.ResponseWriter, r *http.Request) {
	var InpMsg domain.Props
	var InpIDs []int
	defer r.Body.Close()
	initHeaders(w)
	json.NewDecoder(r.Body).Decode(&InpMsg)
	for _, prop := range InpMsg.Props {
		InpIDs = append(InpIDs, prop.ID)
	}
	ans, err := h.useCase.GetByIDs(InpIDs)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error while getting props %s", err.Error())))
	}
	ansBytes, err := json.Marshal(ans)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error while getting props %s", err.Error())))

	}
	writeAnswer(w, http.StatusOK, string(ansBytes))
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var InpMsg domain.Props
	initHeaders(w)
	json.NewDecoder(r.Body).Decode(&InpMsg)
	InpMsg.ConverTime()
	err := h.useCase.CreateProps(InpMsg)
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, fmt.Sprintf("couldn't insert %s", err.Error()))
		return
	}
	writeAnswer(w, http.StatusOK, "insert succeed")

}

func (h *Handler) create1(w http.ResponseWriter, r *http.Request) {
	var InpMsg domain.Props
	initHeaders(w)
	json.NewDecoder(r.Body).Decode(&InpMsg)
	InpMsg.ConverTime()
	err := h.useCase.CreateProp(InpMsg)
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, fmt.Sprintf("couldn't insert %s", err.Error()))
		return
	}
	writeAnswer(w, http.StatusOK, "insert succeed")

}

func (h * Handler) updateProps(w http.ResponseWriter, r *http.Request) {
	var InpMsg domain.Props
	initHeaders(w)
	json.NewDecoder(r.Body).Decode(&InpMsg)
	InpMsg.ConverTime()
	err := h.useCase.UpdateProps(InpMsg)
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, fmt.Sprintf("couldn't insert %s", err.Error()))
		return
	}
	writeAnswer(w, http.StatusOK, "insert succeed")
}