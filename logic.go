package main

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

const LogicBinaryOp string = "> >= < <= == || &&"//!=
