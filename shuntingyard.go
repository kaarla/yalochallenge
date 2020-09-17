package main

import(
  "strings"
)

var OperatorsPrecedence string = "-+/^!>GT<mt==dif||&&"

func InfixToPostfix(originalExpr string, context map[string]int, kind int) ([]string, error){
  infix, err := ProcessString(originalExpr, context, kind)
  aux := ""
  var stack []int
  if err != nil{
    return nil, err
  }

  tokens := strings.Split(infix, " ")
  for _, token := range tokens{
    if len(token) == 0{
      continue;
    }
    c := string(token[0])
    priority := strings.Index(OperatorsPrecedence, c)

    if priority != -1{
      if(len(stack) == 0){
        stack = append(stack, priority)
      }else{
        for len(stack) != 0{
          precedence2 := stack[len(stack)-1] / 2
          precedence1 := priority / 2
          if precedence2 > precedence1 || (precedence2 == precedence1 && c != "^"){
            aux += string(OperatorsPrecedence[stack[len(stack)-1]]) + " "
            stack = stack[:len(stack)-1]
          }else{
            break
          }
        }
        stack = append(stack, priority)
      }
    }else if c == "("{
      stack = append(stack, -2)
    }else if c == ")"{
      for stack[len(stack)-1] != -2{
        aux += string(OperatorsPrecedence[stack[len(stack)-1]]) + " "
        stack = stack[:len(stack)-1]
      }
      stack = stack[:len(stack)-1]
    }else{
      aux += token + " "
    }
  }
  for len(stack) != 0{
    aux += string(OperatorsPrecedence[stack[len(stack)-1]]) + " "
    if len(stack) == 1{
      break
    }
    stack = stack[:len(stack)-1]
  }

  var result []string
  for _, t := range strings.Split(aux, " "){
    result = append(result, t)
  }
  return result, nil
}
