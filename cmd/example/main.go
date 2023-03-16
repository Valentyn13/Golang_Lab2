package main

import (
	"flag"
	"io"
	"os"
	"strings"
)

var (
  inputExpression = flag.String("e", "", "Expression to compute")
  inputFromFile   = flag.String("f", "", "Input file")
  resultOutput    = flag.String("o", "", "File for output")
)

func main() {
  flag.Parse()

  var input io.Reader = nil
  var output = os.Stdout

  if *inputExpression != "" {
    input = strings.NewReader(*inputExpression)
  }

  if *inputFromFile != "" {
    f, err := os.Open(*inputFromFile)
    if err != nil {
      os.Stderr.WriteString("Error opening file")
    }
    defer f.Close()
    input = f
  }

  if *resultOutput != "" {
    file, err := os.Create(*resultOutput)
    if err != nil {
      os.Stderr.WriteString("Error")
    }
    defer file.Close()
    output = file
  }

  if input == nil {
    os.Stderr.WriteString("No input provided")
    return
  }

//   handler := &lab2.ComputeHandler{
//     Input:  input,
//     Output: output,
//   }

//   err := handler.Compute()
//   if err != nil {
//     println(err)
//   }
}