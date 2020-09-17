package main

import(
  "fmt"
  "strconv"
  "strings"
)

func ProcessString(expr string, context map[string]int, kind int) (string, error){
  cleanExpr := subsContext(expr, context)
  cleanExpr = preconditions(cleanExpr)
  valid, symbol := validateSymbols(cleanExpr, kind)
  if valid {
    return cleanExpr, nil
  }
  return "", fmt.Errorf("Uncaught ReferenceError: %g is not defined", symbol)
}

func subsContext(expr string, context map[string]int) string{
  for s, _ := range context{
    expr = strings.Replace(expr, s, strconv.Itoa(context[s]), -1)
  }
  return expr
}

func preconditions(expr string) string{
  infix := strings.Replace(expr, " ", "", -1)
  infix = strings.Replace(expr, "**", "^", -1)
  infix = strings.Replace(expr, ">=", "GT", -1)
  infix = strings.Replace(expr, "<=", "mt", -1)
  infix = strings.Replace(expr, "!=", "dif", -1)
  // infix = strings.Replace(expr, "true", "T", -1)
  // infix = strings.Replace(expr, "false", "F", -1)
  aux := ""
  symbols := []string{"(", ")", "+", "-", "/", "^", "!", ">", "GT", "<", "mt", "==", "dif", "||", "&&"}
  for _, s := range symbols{
    aux = " " + s + " "
    infix = strings.Replace(infix, s, aux, -1)
  }
  return infix
}

//kind: 0 for logic expressions, 1 for arithmetic
func validateSymbols(expr string, kind int) (bool, string){
  symbols := ArithmeticValidSymbols
  if kind == 0{
    symbols += LogicBinaryOp
  }
  aux := strings.Split(expr, " ")
  for _, s := range aux{
    if !strings.Contains(symbols, s){ //check if it's a valid symbol
      if _, err := strconv.Atoi(s); err != nil{ //check if it's a number
        fmt.Println("Validate, error with symbol: ", s)
        return false, s
      }
    }
  }
  return true, ""
}
