package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Vractos/ecoffe-go/pkg/metrics"
	"github.com/Vractos/ecoffe-go/usecases/order"
	"github.com/go-chi/chi/v5"
)

func createOrder(service order.UseCase, logger metrics.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &order.CreateOrderDtoInput{}
		err := json.NewDecoder(r.Body).Decode(input)
		if err != nil {
			logger.Error("Error to decode body", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = service.CreateOrder(*input)
		if err != nil {
			logger.Error(
				"Fail to create order",
				err,
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func MakeOrderHandlers(r chi.Router, service order.UseCase, logger metrics.Logger) {
	r.Route("/orders", func(r chi.Router) {
		r.Post("/", createOrder(service, logger))
	})
}
