package calculation_test

import (
	"testing"
	"github.com/nastts/rpn/pkg/calculation"

)

func TestCalc(t *testing.T){
	tests := []struct{
		expression string
		result float64
	}{
		{"1+1", 2},
		{"3+3*6", 21},
		{"1+8/2*4", 17},
		{"(1+1)*2", 4},
		{"10-2+3", 11},
		{"5*(2+3)", 25},
		{"(8/4)+(3*2)", 8},
		{"7-(3+2)", 2},
		{"6/2*(1+2)", 9},
		{"(3+5)*(2-1)", 8},
		{"(10-3)*(5+ 2)", 49},
		{"(6+2)*(3/2)", 12},
		{"(5 + 5) / (10 / 2)", 2},
		{"(8 + 4) * (2 - 1)", 12},
		
	}
	for _, testCase := range tests {
		
		result, _ := calculation.Calc(testCase.expression)
		if result != testCase.result {
			t.Errorf("%s=%v, want %v", testCase.expression, result, testCase.result)
		}
	}
}

func TestCalcErr(t *testing.T){
	tests := []struct{
		expression string
		result float64
	}{
		{"1+1", 1},
		{"3+3*6", 1},
		{"1+8/2*4", 7},
		{"(1+1)*2", 1},
		{"10-2+3", 1},
		{"5*(2+3)", 5},
		{"(8/4)+(3*2)", 0},
		{"7-(3+2)", 0},
		{"6/2*(1+2)", 90},
		{"(3+5)*(2-1)", 80},
		{"(10-3)*(5+ 2)", 4},
		{"(6+2)*(3/2)", 1},
		{"(5 + 5) / (10 / 2)", 26},
		{"(8 + 4) * (2 - 1)", 26},
		
	}
	for _, testCase := range tests {
		
		result, _ := calculation.Calc(testCase.expression)
		if result == testCase.result {
			t.Errorf("%s=%v, want %v", testCase.expression, result, testCase.result)
		}
	}
}

