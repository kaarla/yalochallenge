package main

import(
  "fmt"
  "encoding/json"
)

func main(){
    // logicExpression1 := `{
	  //                       "expression": "(amount > min)",
	  //                       "save": "validAmount",
    //                       "transitions": {
		//                            "isTrue": 5,
		//                            "isFalse": 10,
		//                            "isError": 25
	  //                       },
    //                       "context": {
		//                           "amount": 150,
		//                           "min": 45
	  //                       }
    //                     }`
    // var logicE LogicEx
    //
    // json.Unmarshal([]byte(logicExpression1), &logicE)
    // //fmt.Printf("Expression: %s, Save: %s", logicE.Expression, logicE.Save)
    // fmt.Println(logicE)

    arithmeticExpression1 := `{
	                             "expression": "3!",
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
    //fmt.Printf("Expression: %s, Save: %s", arithE.Expression, arithE.Save)
    //fmt.Println(arithE)
    fmt.Println("Expression infix,", ProcessString(arithE.Expression, arithE.Context, 1))
    post := InfixToPostfix(arithE.Expression, arithE.Context, 1)
    fmt.Println("Expression postfix,", post)
    tree := NewInterpTree(post)
    fmt.Println("Res: ", tree.Root.elem)
}
