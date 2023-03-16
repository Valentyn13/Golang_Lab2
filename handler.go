package lab2

import (
	"bytes"
	"fmt"
	"io"
)

type ComputeHandler struct {
  Input  io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {

  bufRead := make([]byte, 128)
  _, err := ch.Input.Read(bufRead)
  if err != nil {
    return err
  }
  bufRead = bytes.Trim(bufRead, "\x00")

  expression := string(bufRead)

  res, err := PrefixSolver(expression)
  if err != nil {
    // fmt.Println("Error: ", err)
    return err
  }
  resToOutput := fmt.Sprintf("%v", res)
  ch.Output.Write([]byte(string(resToOutput)))
  return nil
}