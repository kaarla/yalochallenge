package main

import(
  // "fmt"
  "encoding/json"
)

func main(){
    logicExpression1 := `{
	                        "expression": "(min > amount)",
	                        "save": "validAmount",
                          "transitions": {
		                           "isTrue": 5,
		                           "isFalse": 10,
		                           "isError": 25
	                        },
                          "context": {
		                          "amount": 150,
		                          "min": 45
	                        }
                        }`
    var logicE LogicEx

    json.Unmarshal([]byte(logicExpression1), &logicE)
    logicE.Evaluate()

    arithmeticExpression1 := `{
	"expression": "(str/2)",
	"save": "result",
	"transitions": {
		"next": 101,
		"error": 102
	},
	"context": {
		"str": "string-value"
	}
}
`
    var arithE Arithmetic
    json.Unmarshal([]byte(arithmeticExpression1), &arithE)
    arithE.Evaluate()
}
