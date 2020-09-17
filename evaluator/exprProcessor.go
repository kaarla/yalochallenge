package evaluator

import(
  "fmt"
  "strconv"
  "strings"
)

//public method to process a string to beused by the shunting yard algorithm
func ProcessString(expr string, context map[string]int, kind int) (string, error){
  cleanExpr := subsContext(expr, context)
  cleanExpr = preconditions(cleanExpr)
  valid, symbol := validateSymbols(cleanExpr, kind)
  if valid {
    return cleanExpr, nil
  }
  return "", fmt.Errorf("Uncaught ReferenceError: %s is not defined", symbol)
}

//function to add the context to the expression
func subsContext(expr string, context map[string]int) string{
  for s, i := range context{
    expr = strings.Replace(expr, s, strconv.Itoa(i), -1)
  }
  return expr
}

//helper to have a cleaner string expresion
func preconditions(expr string) string{
  infix := strings.Replace(expr, " ", "", -1)
  infix = strings.Replace(infix, "**", "^", -1)
  infix = strings.Replace(infix, ">=", "G", -1)
  infix = strings.Replace(infix, "<=", "m", -1)
  infix = strings.Replace(infix, "!=", "d", -1)
  aux := ""
  symbols := []string{"(", ")", "+", "-", "/", "^", "!", ">", "G", "<", "m", "==", "d", "||", "&&"}
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
        return false, s
      }
    }
  }
  return true, ""
}
