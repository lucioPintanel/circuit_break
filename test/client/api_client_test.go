package client

import (
	"circuit-break/internal/client"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMakeRequest(t *testing.T) {
	// Cria um servidor mock
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	client := client.NewApiClient()
	response, err := client.MakeRequest(mockServer.URL) // Use a URL do servidor mock
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{"message": "success"}`
	if response != expected {
		t.Fatalf("Expected response %v, got %v", expected, response)
	}
}

func TestCircuitBreaker(t *testing.T) {
	// Configura um servidor mock que sempre retorna um erro
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
	}))
	defer mockServer.Close()

	client := client.NewApiClient()

	// Fazer múltiplas requisições para acionar o Circuit Breaker
	for i := 0; i < 5; i++ {
		_, err := client.MakeRequest(mockServer.URL)
		if err == nil {
			t.Fatalf("Expected error on iteration %d, got none", i)
		}
		if i >= 3 && err.Error() != "circuit breaker is open" {
			t.Fatalf("Expected circuit breaker to be open on iteration %d, got %v", i, err)
		}
	}

	// Aguarda o tempo de espera para o Circuit Breaker fechar o circuito
	time.Sleep(3 * time.Second)

	// Faz uma nova requisição para verificar se o circuito está fechado novamente
	_, err := client.MakeRequest(mockServer.URL)
	if err == nil {
		t.Fatalf("Expected error after reset, got none")
	}
}
