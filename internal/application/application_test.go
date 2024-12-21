package application_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nastts/rpn/internal/application"
)

func CalcHandlerTestOk(t *testing.T){
	tests := []struct{
		reg string
		expected string
	}{
		{`{"expression":"1+1"}`, `{"result":"2"}`},
		{`{"expression":"3+3*6"}`, `{"result":"21"}`},
		{`{"expression":"1+8/2*4"}`, `{"result":"17"}`},
		{`{"expression":"(1+1)*2"}`, `{"result":"4"}`},
		{`{"expression":"10-2+3"}`, `{"result":"11"}`},
		{`{"expression":"5*(2+3)"}`, `{"result":"25"}`},
		{`{"expression":"(8/4)+(3*2)"}`, `{"result":"8"}`},
		{`{"expression":"7-(3+2)"}`, `{"result":"2"}`},
		{`{"expression":"6/2*(1+2)"}`, `{"result":"9"}`},
		{`{"expression":"(3+5)*(2-1)"}`, `{"result":"8"}`},
		{`{"expression":"(10-3)*(5+ 2)"}`, `{"result":"49"}`},
		{`{"expression":"(6+2)*(3/2)"}`, `{"result":"12"}`},
		{`{"expression":"(5 + 5) / (10 / 2)"}`, `{"result":"2"}`},
		{`{"(8 + 4) * (2 - 1)"}`, `{"result":"12"}`},
	}
	for _, testCase := range tests{
		request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", nil)
		w := httptest.NewRecorder()
		application.CalcHandler(w, request)
		if w.Code != http.StatusOK {
			t.Errorf("Wrong status code, expected 200, got: %d", w.Code)
		}else if w.Body.String() != testCase.expected{
			t.Errorf("wrong data, expected '%s', but got '%s'", testCase.expected, w.Body.String())
		}
	}
}


func CalcHandlerTestErr(t *testing.T){
	tests := []string{
		`{"expression":"1+"}`,
		`{"expression":"3+3*************************************6"}`,
		`{"expression":"1+8/|2*4"}`,
		`{"expression":"(1+)*2"}`,
		`{"expression":""}`,
		`{"expression":"5(2+3)"}`,
		`{"expression":"(84)(32)"}`,
	}
	for _, testCase := range tests{
		request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", nil)
		w := httptest.NewRecorder()
		application.CalcHandler(w, request)
		if w.Code != http.StatusUnprocessableEntity {
			t.Errorf("Wrong status code, expected 422, got: %s", testCase)
		}
	}
}
