package main

import(
  "github.com/kaarla/yalochallenge/evaluator"
  "encoding/json"
  "fmt"
)

func main(){
    fmt.Println("Prueba 1.\n resultado esperado:\n {\"transition\":\"102\",\"result\":\"NaN\"}\n Resultado: ")
    arithmeticExpression1 := ` {
	     "expression": "(str/2)",
	     "save": "result",
	     "transitions": {
		   "next": 101,
		   "error": 102
	    },
	     "context": {
		       "str": "string-value"
	     }
     }`

    var arithE evaluator.Arithmetic
    err1 := json.Unmarshal([]byte(arithmeticExpression1), &arithE)
    if err1 != nil{
      err1 = fmt.Errorf("NaN")
    }
    arithE.Evaluate(err1)

    fmt.Println("\n\nPrueba 2.\n resultado esperado:\n {\"transition\":\"1\",\"result\":\"0.01836547291\"}\n Resultado obtenido: ")
    arithmeticExpression2 := `{
	     "expression": "value/(99**2)",
	      "save": "result",
	      "transitions": {
		    "next": 1,
		    "error": 2
	     },
	      "context": {
		        "value": 180
	       }
    }`

    var arithE2 evaluator.Arithmetic
    err := json.Unmarshal([]byte(arithmeticExpression2), &arithE2)
    if err != nil{
      err = fmt.Errorf("NaN")
    }
    arithE2.Evaluate(err)

    fmt.Println("\n\nPrueba 3.\n resultado esperado:\n {\"transition\":\"25\",\"result\":\"5\"}\n Resultado obtenido: ")
    arithmeticExpression3 := `{
	     "expression": "(10/2)",
	     "save": "result",
	     "transitions": {
		       "next": 25,
		       "error": 50
	     },
	     "context": {}
       }`

    var arithE3 evaluator.Arithmetic
    err3 := json.Unmarshal([]byte(arithmeticExpression3), &arithE3)
    if err3 != nil{
      err3 = fmt.Errorf("NaN")
    }
    arithE3.Evaluate(err3)

    fmt.Println("\n\nPrueba 4.\n resultado esperado:\n {\"transition\":\"8\",\"fact\":\"120\"}\n Resultado obtenido: ")
    arithmeticExpression4 := `{
	     "expression": "(10/2)!",
	     "save": "fact",
	     "transitions": {
		       "next": 8,
		       "error": 50
	     },
	     "context": {}
       }`

    var arithE4 evaluator.Arithmetic
    err4 := json.Unmarshal([]byte(arithmeticExpression4), &arithE4)
    if err4 != nil{
      err4 = fmt.Errorf("NaN")
    }
    arithE4.Evaluate(err4)

    fmt.Println("\n\nPrueba 5.\n resultado esperado:\n {\"transition\":\"6\",\"threeOps\":\"3.5568743e+14\"}\n Resultado obtenido: ")
    arithmeticExpression5 := `{
	     "expression": "((8+9)!)-dos",
	     "save": "threeOps",
	     "transitions": {
		       "next": 6,
		       "error": 9
	     },
	     "context": {
         "dos": 2
       }
      }`

    var arithE5 evaluator.Arithmetic
    err5 := json.Unmarshal([]byte(arithmeticExpression5), &arithE5)
    if err5 != nil{
      err5 = fmt.Errorf("NaN")
    }
    arithE5.Evaluate(err5)

}
