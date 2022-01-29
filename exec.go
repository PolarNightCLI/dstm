package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func bash(cmd *exec.Cmd) error {
	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	if err := cmd.Start(); err != nil {
		return err
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}
	if errStdout != nil {
		return errStdout
	}
	if errStderr != nil {
		return errStderr
	}
	return nil
}

func bash2(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	// 这句话好像没啥用，待测试
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func bash3(cmd1 *exec.Cmd, cmd2 *exec.Cmd) error {
	cmdout1, err := cmd1.StdoutPipe()
	if err != nil {
		return err
	}

	cmd2.Stdin = cmdout1
	cmd2.Stdout = os.Stdout

	err = cmd1.Start()
	if err != nil {
		return err
	}
	err = cmd2.Start()
	if err != nil {
		return err
	}
	cmd1.Wait()
	cmd2.Wait()

	// return errors.New(fmt.Sprint(cmd2.Stdout))
	return nil
}

func bash4(cmd *exec.Cmd) string {
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
