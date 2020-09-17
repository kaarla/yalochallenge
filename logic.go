package main

import(
  "github.com/kaarla/yalochallenge/evaluator"
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
    var logicE evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression1), &logicE)
    logicE.Evaluate()

}
