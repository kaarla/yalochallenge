package evaluator

import(
  "encoding/json"
  "strconv"
  "fmt"
)

type Arithmetic struct{
  Expression string           `json:"expression"`
  Save string                 `json:"save"`
  Transitions TransitionsA    `json:"transitions"`
  Context map[string]int      `json:"context"`
}

type TransitionsA struct{
  Next int    `json:"next"`
  Error int   `json:"error"`
}

type OutputA struct{
  Result float32
  Transition int
}

//binary operators
const ArithmeticBinaryOp string = "+-*/^"
//unary operators
const ArithmeticUnaryOp string = "!"
//only arithmetic symbols
var ArithmeticSymbols string = ArithmeticBinaryOp + ArithmeticUnaryOp
//valid symbols
var ArithmeticValidSymbols string = ArithmeticSymbols + "()."

//function to evaluate an expression
func (arithE *Arithmetic)Evaluate(error error){
  if error != nil{
    fmt.Println(arithE.createResponse(error.Error(), true))
    return
  }
  post, err := InfixToPostfix(arithE.Expression, arithE.Context, 1)
  if err == nil{
    tree := NewInterpTree(post)
    fmt.Println(arithE.createResponse(tree.Root.elem, false))
  }else{
    fmt.Println(arithE.createResponse(err.Error(), true))
  }
}

//function to build a json response and return it as an string
func (arithE *Arithmetic)createResponse(result string, err bool) string{
  response := make(map[string]string)
  response[arithE.Save] = result
  if err{
    response["transition"] = strconv.Itoa(arithE.Transitions.Error)
  }else{
    response["transition"] = strconv.Itoa(arithE.Transitions.Next)
  }
  json, _ := json.Marshal(response)
  return string(json)
}
