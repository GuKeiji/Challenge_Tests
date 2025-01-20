package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductsDefault_Get(t *testing.T) {
	repoMocked := repository.NewProductsMapMock()
	handlerTest := handler.NewProductsDefault(repoMocked)

	t.Run("Test erro query", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=abc", nil)

		if err != nil {
			t.Fatalf(err.Error())
		}

		recorder := httptest.NewRecorder()

		funcTest := handlerTest.Get()

		funcTest(recorder, req)

		expectedCode := http.StatusBadRequest
		require.Equal(t, expectedCode, recorder.Code)
	})

	t.Run("Get One Test success handler", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=2", nil)

		if err != nil {
			t.Fatalf(err.Error())
		}

		recorder := httptest.NewRecorder()

		funcTest := handlerTest.Get()

		funcTest(recorder, req)

		expectedCode := 200
		expectedBody := `{"data":{"2":{"id":2,"description":"Test id 2","price":30,"seller_id":2}},"message":"success"}`

		require.Equal(t, expectedCode, recorder.Code)
		require.Equal(t, expectedBody, recorder.Body.String())
	})

	t.Run("Get One Test failed handler", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products?id=99", nil)

		if err != nil {
			t.Fatalf(err.Error())
		}

		recorder := httptest.NewRecorder()

		funcTest := handlerTest.Get()

		funcTest(recorder, req)

		expectedCode := 200
		expectedBody := `{"data":{},"message":"success"}`

		require.Equal(t, expectedCode, recorder.Code)
		require.Equal(t, expectedBody, recorder.Body.String())
	})

	t.Run("Get All Test handler", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products", nil)

		if err != nil {
			t.Fatalf(err.Error())
		}

		recorder := httptest.NewRecorder()

		funcTest := handlerTest.Get()

		funcTest(recorder, req)

		expectedCode := 200
		expectedBody := `{
							"data": {
								"1": {
									"id": 1,
									"description": "Test id 1",
									"price": 20,
									"seller_id": 1
								},
								"2": {
									"id": 2,
									"description": "Test id 2",
									"price": 30,
									"seller_id": 2
								},
								"3": {
									"id": 3,
									"description": "Test id 3",
									"price": 40,
									"seller_id": 3
								}
							},
							"message": "success"
						}`

		require.Equal(t, expectedCode, recorder.Code)
		require.JSONEq(t, expectedBody, recorder.Body.String())
	})
}

func TestNewProductsDefault(t *testing.T) {
	repoMocked := repository.NewProductsMapMock()
	hd := handler.NewProductsDefault(repoMocked)

	require.NotNil(t, hd)
}
