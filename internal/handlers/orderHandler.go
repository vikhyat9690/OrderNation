package handlers

import (
	"log"
	"net/http"
	"ordernationn/internal/domain"
	"ordernationn/internal/service"
	"strings"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// package handlers

// The method returns a function that matches the http.HandlerFunc signature.
func (o *OrderHandler) PlaceOrder() http.HandlerFunc {
	// This is the function that http.HandleFunc will receive and call on a request.
	return func(w http.ResponseWriter, r *http.Request) {
		// --- 1. Validate HTTP Method ---
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// --- 2. Decode/Parse Request Body ---
		var order domain.Order
		// NOTE: You need logic to decode the JSON body into the 'order' struct.
		// I'm omitting the full decoding logic for brevity, but it's essential.
		// Example:
		// if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		//     http.Error(w, "Invalid request body", http.StatusBadRequest)
		//     return
		// }

		// Get the Context from the request
		ctx := r.Context()

		// --- 3. Call Service Logic ---
		result, err := o.service.PlaceOrder(ctx, order)

		// --- 4. Handle Errors ---
		if err != nil {
			// Log the detailed error
			log.Printf("Error placing order: %v", err)

			// Return a public-facing error message and status code
			if strings.Contains(err.Error(), "not available") {
				http.Error(w, err.Error(), http.StatusConflict) // e.g., 409 Conflict
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		// --- 5. Write Successful Response ---
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201 Created
		// In a real application, you would encode a response struct, e.g.,
		// json.NewEncoder(w).Encode(map[string]string{"message": result})
		w.Write([]byte(result))
	}
}
