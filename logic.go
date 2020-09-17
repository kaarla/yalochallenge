package main

import(
  "github.com/kaarla/yalochallenge/evaluator"
  "encoding/json"
  "fmt"
)

func main(){

    fmt.Println("Prueba 1.\n resultado esperado:\n {\"validAmount\":\"true\",\"transition\":\"5\"}\n Resultado: ")
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

    fmt.Println("\n\nPrueba 2.\n resultado esperado:\n {\"adult\":\"false\",\"transition\":\"23\"}\n Resultado: ")
    logicExpression2 := `{
	     "expression": "(age) >= 18",
	     "save": "adult",
	     "transitions": {
		   "isTrue": 15,
		   "isFalse": 23,
		   "isError": 45
	    },
	     "context": {
		       "age": 15
	     }
    }`
    var logicE2 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression2), &logicE2)
    logicE2.Evaluate()

    fmt.Println("\n\nPrueba 3.\n resultado esperado:\n {\"adult\":\"Uncaught ReferenceError: age is not defined\",\"transition\":\"45\"}\n Resultado: ")
    logicExpression3 := `{
	     "expression": "(age) >= 18",
	     "save": "adult",
	     "transitions": {
		       "isTrue": 15,
		       "isFalse": 23,
		       "isError": 45
	     },
	      "context": {}
      }`
    var logicE3 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression3), &logicE3)
    logicE3.Evaluate()

}
