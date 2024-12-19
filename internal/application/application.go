package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	

	"github.com/nastts/rpn/pkg/calculation"
)

type Answer struct{
	Expression string `json:"expression"`
}
type Response struct{
	Result string `json:"result"`
}
type Error struct{
	Error string `json:"error"`
}


func CalcHandler(w http.ResponseWriter, r *http.Request){
	var rs Response
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		mainError := "method not allowed"
		output := Error{Error: mainError}
		jsonOutput, _ := json.Marshal(output)
		w.Write(jsonOutput)
		return
	}
	answer := new(Answer)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
		mainError := "expression is not valid"
		output := Error{Error: mainError}
		jsonOutput, _ := json.Marshal(output)
		w.Write(jsonOutput)
		return
	}
	res, err := calculation.Calc(answer.Expression)
	if err != nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
		mainError := "expression is not valid"
		output := Error{Error: mainError}
		jsonOutput, _ := json.Marshal(output)
		w.Write(jsonOutput)
		return
	} 
	rs.Result = fmt.Sprintf("%.2f", res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rs)
	

}
	