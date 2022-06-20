package entrypoint

import (
	"fmt"
	"net/http"
)

func (h *Handler) getByIDs(w http.ResponseWriter, r *http.Request) {
	var InpIDs []uint
	defer r.Body.Close()
	InpMsg, err := extractPropsFromCtxt(r.Context())
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, string("couldn't extract from ctxt"))
		return
	}
	for _, prop := range InpMsg.Props {
		InpIDs = append(InpIDs, prop.ID)
	}
	ans, err := h.useCase.GetByIDs(InpIDs)
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, fmt.Sprintf("coudn't get props in DB %s",err.Error()))
		return
	}
	writeAnswerWithProps(w, http.StatusOK, ans)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	InpMsg, err := extractPropsFromCtxt(r.Context())
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, string("couldn't extract from ctxt"))
		return
	}
	err = h.useCase.CreateProps(InpMsg)
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, fmt.Sprintf("couldn't insert %s", err.Error()))
		return
	}
	writeAnswer(w, http.StatusOK, "insert succeed")
}

func (h *Handler) updateProps(w http.ResponseWriter, r *http.Request) {
	InpMsg, err := extractPropsFromCtxt(r.Context())
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, string("couldn't extract from ctxt"))
		return
	}
	err = h.useCase.UpdateProps(InpMsg)
	if err != nil {
		writeAnswer(w, http.StatusInternalServerError, fmt.Sprintf("couldn't update %s", err.Error()))
		return
	}
	writeAnswer(w, http.StatusOK, "update succeed")
}


// func (h *Handler) create1(w http.ResponseWriter, r *http.Request) {
// 	initHeaders(w)
// 	InpMsg, err := extractPropsFromCtxt(r.Context())
// 	if err != nil {
// 		writeAnswer(w, http.StatusInternalServerError, string("couldn't extract from ctxt"))
// 		return
// 	}
// 	err = h.useCase.CreateProp(InpMsg)
// 	if err != nil {
// 		writeAnswer(w, http.StatusInternalServerError, fmt.Sprintf("couldn't insert %s", err.Error()))
// 		return
// 	}
// 	writeAnswer(w, http.StatusOK, "insert succeed")
// }