package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"service-transaction/transaction"

	"github.com/julienschmidt/httprouter"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type responseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// TransactionHandler represent the httphandler for transaction
type TransactionHandler struct {
	TUsecase transaction.Usecase
}

// NewTransactionHandler will initialize the articles/ resources endpoint
func NewTransactionHandler(router *httprouter.Router, tu transaction.Usecase) {
	handler := &TransactionHandler{
		TUsecase: tu,
	}
	// var w http.ResponseWriter
	// var r *http.Request
	router.GET("/", handler.FetchTransaction)
}

// FetchTransaction will fetch the transaction based on given params
func (t *TransactionHandler) FetchTransaction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	num, _ := strconv.Atoi(r.URL.Query().Get("num"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "ASC"
	}
	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	w.Header().Set("Content-Type", "application/json")
	var res responseData
	if sort != "asc" && sort != "ASC" && sort != "desc" && sort != "DESC" {
		res.Status = http.StatusBadRequest
		res.Message = "allowed sort is asc or desc"
	} else {
		listTrx, err := t.TUsecase.Fetch(ctx, num, offset, sort)

		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = err.Error()
		} else {
			res.Status = http.StatusOK
			res.Message = "success"
			res.Data = listTrx
		}
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
