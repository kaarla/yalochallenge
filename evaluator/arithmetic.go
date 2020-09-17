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

const ArithmeticBinaryOp string = "+-*/^"
const ArithmeticUnaryOp string = "!"
var ArithmeticSymbols string = ArithmeticBinaryOp + ArithmeticUnaryOp
var ArithmeticValidSymbols string = ArithmeticSymbols + "()."

func (arithE *Arithmetic)Evaluate(){
  post, err := InfixToPostfix(arithE.Expression, arithE.Context, 1)
  if err == nil{
    tree := NewInterpTree(post)
    fmt.Println(arithE.createResponse(tree.Root.elem, false))
  }else{
    fmt.Println(arithE.createResponse(err.Error(), true))
  }
}

func (arithE *Arithmetic)createResponse(result string, err bool) string{
  response := make(map[string]string)
  _, error := strconv.Atoi(result)
  if error != nil {
    response[arithE.Save] = "NaN"
  }else{
    response[arithE.Save] = result
  }
  if err{
    response["transition"] = strconv.Itoa(arithE.Transitions.Next)
  }else{
    response["transition"] = strconv.Itoa(arithE.Transitions.Error)
  }
  json, _ := json.Marshal(response)
  return string(json)
}
