package main

import(
  "github.com/kaarla/yalochallenge/evaluator"
  "encoding/json"
  "fmt"
)

func main(){
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
    var arithE evaluator.Arithmetic
    err1 := json.Unmarshal([]byte(arithmeticExpression1), &arithE)
    if err1 != nil{
      err1 = fmt.Errorf("NaN")
    }
    arithE.Evaluate(err1)

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
arithE2.Evaluate(err)
}
