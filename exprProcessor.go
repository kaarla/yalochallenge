package main

import(
  "fmt"
  "strconv"
  "strings"
)

func ProcessString(expr string, context map[string]int, kind int) string{
  cleanExpr := subsContext(expr, context)
  if validateSymbols(cleanExpr, kind) {
    return preconditions(cleanExpr)
  }
  return ""
}

func subsContext(expr string, context map[string]int) string{
  result := ""
  for s, _ := range context{
    result = strings.Replace(expr, s, strconv.Itoa(context[s]), -1)
  }
  return result
}

func preconditions(expr string) string{
  infix := strings.Replace(expr, " ", "", -1)
  infix = strings.Replace(expr, "**", "^", -1)
  aux := ""
  symbols := []string{"(", ")", "+", "-", "/", "^", "!", ">", ">=", "<", "<=", "==", "!=", "||", "&&"}
  for _, s := range symbols{
    aux = " " + s + " "
    //fmt.Println(s)
    infix = strings.Replace(infix, s, aux, -1)
    //fmt.Println(infix)
  }
  return infix
}

//kind: 0 for logic expressions, 1 for arithmetic
func validateSymbols(expr string, kind int) bool{
  symbols := ArithmeticValidSymbols
  if kind == 0{
    symbols += LogicBinaryOp
  }
  temp := strings.Replace(expr, " ", "", -1)
  aux := strings.Split(temp, "")
  for _, s := range aux{
    if !strings.Contains(symbols, s){ //check if it's a valid symbol
      if _, err := strconv.Atoi(s); err != nil{ //check if it's a number
        fmt.Println("Validate, error with symbol: ", s)
        return false
      }
    }
  }
  return true
}
