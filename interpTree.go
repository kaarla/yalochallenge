package main

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

func (node *node)hasRightSon() bool{
  if node.rightSon != nil{
    return true
  }
  return false
}

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

func NewInterpTree(postfix []string) *InterpTree{
  return buildtree(postfix[:len(postfix)-1])
}

func buildtree(postfix []string) *InterpTree{
  // fmt.Println("Len init: ", len(postfix), postfix[4])
  tree := InterpTree{}
  var stack []*node
  token := ""
  for i:= 0; i < len(postfix); i++ {
    token = postfix[i]
    nod := newNode(token)
    fmt.Println("nodo, tipo, unorbin:", nod.elem, nod.kind, nod.unOrBin)

    if nod.kind == 0{
      if nod.unOrBin == 2{
        nod.rightSon = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        nod.rightSon.father = nod

        fmt.Println("len: ", len(stack))
        nod.leftSon = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        nod.leftSon.father = nod

        nod.elem = fmt.Sprintf("%f", eval(nod))
        fmt.Println("Elem eval:", nod.elem)

        remove(nod.rightSon)
        remove(nod.leftSon)

        stack = append(stack, nod)
      }else{
        nod.rightSon = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        nod.elem = fmt.Sprintf("%f", eval(nod))
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

func remove(nod *node){
  if nod.father.hasRightSon() && nod.father.rightSon.equals(nod){
    nod.father.rightSon = nil
  }else{
    nod.father.leftSon = nil
  }
}

func eval(nod *node) float64{
  y, _ := strconv.ParseFloat(nod.rightSon.elem, 32)
  if nod.unOrBin == 1{
    return y
  }
  x, _ := strconv.ParseFloat(nod.leftSon.elem, 32)
  fmt.Println("x, y", x, y)
  switch nod.elem {
  case "+":
    return x + y
  case "-":
    return x - y
  case "/":
    return x / y
  case "^":
    return math.Pow(x, y)
  default:
    return x
  }
}
