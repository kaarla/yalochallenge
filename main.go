package main

import(
  "fmt"
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
    post1 := InfixToPostfix(logicE.Expression, logicE.Context, 0)
    fmt.Println("Expression postfix,", post1)
    tree1 := NewInterpTree(post1)
    fmt.Println("Res: ", tree1.Root.elem)

    arithmeticExpression1 := `{
	                             "expression": " 4+3/8-1",
	                             "save": "result",
	                             "transitions": {
		                               "next": 1,
		                               "error": 2
	                              },
	                             "context": {
		                               "value": 180
	                             }
                              }`
    var arithE Arithmetic
    json.Unmarshal([]byte(arithmeticExpression1), &arithE)
    post := InfixToPostfix(arithE.Expression, arithE.Context, 1)
    fmt.Println("Expression postfix,", post)
    tree := NewInterpTree(post)
    fmt.Println("Res: ", tree.Root.elem)
}
