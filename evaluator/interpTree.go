package evaluator

import(
  "fmt"
  "math"
  "strconv"
  "strings"
)

//internal private struct for tree nodes
type node struct{
  elem string
  father *node
  leftSon *node
  rightSon *node
  kind int
  unOrBin int
}

//node constructor
func newNode(elem string) *node{
  node := node{}
  node.elem = elem
  node.kind = 1 //1 is for numbers
  node.unOrBin = 0 //is for numbers
  if strings.Index(OperatorsPrecedence, elem) > -1 {
    node.kind = 0 //0 is for operators
  }
  if strings.Index(ArithmeticUnaryOp, elem) > -1{
    node.unOrBin = 1 //1 for unary
  }
  if strings.Index(ArithmeticBinaryOp, elem) > -1 || strings.Index(LogicBinaryOp, elem) > -1{
    node.unOrBin = 2 //2 for binary
  }
  return &node
}

//to check if the node has rigth son
func (node *node)hasRightSon() bool{
  if node.rightSon != nil{
    return true
  }
  return false
}

//to check if 2 nodes are the same
func (node *node)equals(n *node) bool{
  if n.elem == ""{
    return false
  }
  return node.elem == n.elem
}

//struct for the actual tree
type InterpTree struct{
  Root *node
}

//Public function to create an interp tree
func NewInterpTree(postfix []string) *InterpTree{
  return buildtree(postfix[:len(postfix)-1])
}

//private function to build an autoevaluated tree
func buildtree(postfix []string) *InterpTree{
    tree := InterpTree{}
    var stack []*node
    token := ""
    for i:= 0; i < len(postfix); i++ {
        token = postfix[i]
        nod := newNode(token)

        if nod.kind == 0{
            if nod.unOrBin == 2{
                nod.rightSon = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                nod.rightSon.father = nod

                nod.leftSon = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                nod.leftSon.father = nod

                ar, log, k := eval(nod)
                if k == 0{
                    nod.elem = fmt.Sprintf("%f", ar)
                }else{
                    nod.elem = strconv.FormatBool(log)
                }

                remove(nod.rightSon)
                remove(nod.leftSon)

                stack = append(stack, nod)
            }else{
                nod.rightSon = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                ar, log, k := eval(nod)
                if k == 0{
                    nod.elem = fmt.Sprintf("%f", ar)
                }else{
                    nod.elem = strconv.FormatBool(log)
                }
                stack = append(stack, nod)
            }
            tree.Root = nod
          }else{
              stack = append(stack, nod)
          }
    }
    tree.Root = stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    return &tree
}

//function to remove nodes that we don't need anymore
func remove(nod *node){
  if nod.father.hasRightSon() && nod.father.rightSon.equals(nod){
    nod.father.rightSon = nil
  }else{
    nod.father.leftSon = nil
  }
}

//function to eval an expression
func eval(nod *node) (float64, bool, int){
  xL, yL := true, true
  y, errY := strconv.ParseFloat(nod.rightSon.elem, 32)
  if nod.unOrBin == 1{
    return factorial(y), true, 0
  }
  x, errX := strconv.ParseFloat(nod.leftSon.elem, 32)
  if errY != nil || errX != nil{
    yL, errY = strconv.ParseBool(nod.rightSon.elem)
    xL, errX = strconv.ParseBool(nod.leftSon.elem)
    if errY != nil || errX != nil{
      return 0.0, false, -1
    }
  }
  switch nod.elem {
    case "+":
      return x + y, true, 0
    case "-":
      return x - y, true, 0
    case "/":
      return x / y, true, 0
    case "^":
      return math.Pow(x, y), true, 0
    case ">":
      return 0.0, y > x, 1
    case "G":
      return 0.0, x >= y, 1
    case "<":
      return 0.0, x < y, 1
    case "m":
      return 0.0, x <= y, 1
    case "=":
      return 0.0, x == y, 1
    case "d":
      return 0.0, x != y, 1
    case "|":
      return 0.0, (xL || yL), 1
    case "&":
      return 0.0, xL && yL, 1
    default:
      return 0.0, false, -1
  }
}

//function to calculate factorial and avoid importing another package
func factorial(x float64) float64{
  if x <= 1{
    return 1
  }
  return x * factorial(x - 1)
}
