package main

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"reflect"
)

type TestCase struct {
	Input  []byte
	Output []byte
}

type Acceptor struct {
	cmd *exec.Cmd
}

func (a *Acceptor) Accept(tc TestCase) error {
	buf := bytes.NewBuffer(make([]byte, 0, len(tc.Output)))
	r, w := io.Pipe()
	a.cmd.Stdin, a.cmd.Stdout = r, buf

	n, err := io.Copy(w, bytes.NewReader(tc.Input))
	if err != nil {
		return err
	}
	if n != int64(len(tc.Input)) {
		return errors.New("invalid copy")
	}

	if err := a.cmd.Run(); err != nil {
		return err
	}

	if !reflect.DeepEqual(tc.Output, buf.Bytes()) {
		return errors.New("not accept")
	}
	return nil
}

func main() {
}
