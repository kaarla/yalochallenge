package evaluator

import(
  "encoding/json"
  "strconv"
  "fmt"
)

type LogicEx struct{
  Expression string           `json:"expression"`
  Save string                 `json:"save"`
  Transitions TransitionsL    `json:"transitions"`
  Context map[string]int      `json:"context"`
}

type TransitionsL struct{
  IsTrue int    `json:"isTrue"`
  IsFalse int   `json:"isFalse"`
  IsError int   `json:"isError"`
}

type OutputL struct{
  Result bool
  Transition int
}

const LogicBinaryOp string = ">G<m==d||&&"

func (logicE LogicEx)Evaluate(){
  post, err := InfixToPostfix(logicE.Expression, logicE.Context, 0)
  if err == nil{
    tree := NewInterpTree(post)
    fmt.Println(logicE.createResponse(tree.Root.elem))
  }else{
    fmt.Println(logicE.createResponse(err.Error()))
  }
}

func (logicE *LogicEx)createResponse(result string) string{
  response := make(map[string]string)
  response[logicE.Save] = result

  switch result {
  case "false":
    response["transition"] = strconv.Itoa(logicE.Transitions.IsFalse)
    break
  case "true":
    response["transition"] = strconv.Itoa(logicE.Transitions.IsTrue)
    break
  default:
    response["transition"] = strconv.Itoa(logicE.Transitions.IsError)
  }
  json, _ := json.Marshal(response)
  return string(json)
}
