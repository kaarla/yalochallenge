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
	     "expression": "(age) > 18",
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

    fmt.Println("\n\nPrueba 4.\n resultado esperado:\n {\"ages\":\"false\",\"transition\":\"0\"}\n Resultado: ")
    logicExpression4 := `{
	     "expression": "(15) >= 18 && 18 == 18",
	     "save": "ages",
	     "transitions": {
		       "isTrue": 1,
		       "isFalse": 0,
		       "isError": -1
	     },
	      "context": {}
      }`
    var logicE4 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression4), &logicE4)
    logicE4.Evaluate()

    fmt.Println("\n\nPrueba 5.\n resultado esperado:\n {\"trueAges\":\"true\",\"transition\":\"1\"}\n Resultado: ")
    logicExpression5 := `{
	     "expression": "(15) <= 18 || mayor != 18",
	     "save": "trueAges",
	     "transitions": {
		       "isTrue": 1,
		       "isFalse": 0,
		       "isError": -1
	     },
	      "context": {
          "mayor":18
        }
      }`
    var logicE5 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression5), &logicE5)
    logicE5.Evaluate()

    fmt.Println("\n\nPrueba 6.\n resultado esperado:\n {\"falseAges\":\"false\",\"transition\":\"0\"}\n Resultado: ")
    logicExpression6 := `{
	     "expression": "(15) < 18 || mayor != 18",
	     "save": "trueAges",
	     "transitions": {
		       "isTrue": 1,
		       "isFalse": 0,
		       "isError": -1
	     },
	      "context": {
          "mayor":18
        }
      }`
    var logicE6 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression6), &logicE6)
    logicE6.Evaluate()

    fmt.Println("\n\nPrueba 7.\n resultado esperado:\n {\"menorQ\":\"true\",\"transition\":\"1\"}\n Resultado: ")
    logicExpression7 := `{
	     "expression": "(age) < 18",
	     "save": "menorQ",
	     "transitions": {
		   "isTrue": 1,
		   "isFalse": 3,
		   "isError": 4
	    },
	     "context": {
		       "age": 15
	     }
    }`
    var logicE7 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression7), &logicE7)
    logicE7.Evaluate()

    fmt.Println("\n\nPrueba 8.\n resultado esperado:\n {\"menorOI\":\"true\",\"transition\":\"1\"}\n Resultado: ")
    logicExpression8 := `{
	     "expression": "(age) <= 18",
	     "save": "menorOI",
	     "transitions": {
		   "isTrue": 1,
		   "isFalse": 3,
		   "isError": 4
	    },
	     "context": {
		       "age": 15
	     }
    }`
    var logicE8 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression8), &logicE8)
    logicE8.Evaluate()

    fmt.Println("\n\nPrueba 9.\n resultado esperado:\n {\"igual\":\"false\",\"transition\":\"3\"}\n Resultado: ")
    logicExpression9 := `{
	     "expression": "(age) == 18",
	     "save": "menorOI",
	     "transitions": {
		   "isTrue": 1,
		   "isFalse": 3,
		   "isError": 4
	    },
	     "context": {
		       "age": 15
	     }
    }`
    var logicE9 evaluator.LogicEx
    json.Unmarshal([]byte(logicExpression9), &logicE9)
    logicE9.Evaluate()

}
