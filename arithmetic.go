package main

type Arithmetic struct{
  Expression string           `json:"expression"`
  Save string                 `json:"save"`
  Transitions TransitionsA    `json:"transitions"`
  Context map[string]int      `json:"context"`
}

type TransitionsA struct{
  Next int    `json:"next"`
  Error int   `json:"error"`
}

type OutputA struct{
  Result float32
  Transition int
}
