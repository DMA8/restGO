package entrypoint

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"testTask/internal/domain"
)

type MessageOut struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

type MessageProps struct {
	StatusCode int            `json:"status_code"`
	Props      []*domain.Prop `json:"props"`
	IsError    bool           `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func writeAnswer(writer http.ResponseWriter, status int, message string) {
	var errorFlag bool
	if status >= 400 {
		errorFlag = true
	}
	msg := MessageOut{
		StatusCode: status,
		Message:    message,
		IsError:    errorFlag,
	}
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(msg)
	if err != nil {
		log.Println("BAD json") //TODO
	}
}

func writeAnswerWithProps(writer http.ResponseWriter, status int, props []*domain.Prop) {
	var errorFlag bool
	if status >= 400 {
		errorFlag = true
	}
	msg := MessageProps{
		StatusCode: status,
		Props:      props,
		IsError:    errorFlag,
	}
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(msg)
	if err != nil {
		log.Println("BAD json") //TODO
	}
}

func validateInput(inp *domain.Props) error {
	err := errors.New("bad input")
	if inp.Props == nil {
		return err
	}
	for _, v := range inp.Props {
		if v.ID <= 0 {
			return err
		}
	}
	return inp.ConverTime()
}

func extractPropsFromCtxt(ctx context.Context) (*domain.Props, error) {
	propsCtx := ctx.Value(prop("prop"))
	switch val := propsCtx.(type) {
	case *domain.Props:
		return val, nil
	default:
		return nil, errors.New("couldn't extract domain from ctx")
	}
}
