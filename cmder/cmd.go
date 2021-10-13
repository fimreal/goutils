package cmder

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"runtime"
)

var CMDArgs [2]string

func init() {
	if runtime.GOOS == "windows" {
		CMDArgs[0] = "cmd"
		CMDArgs[1] = "/c"
	} else {
		CMDArgs[0] = "/bin/sh"
		CMDArgs[1] = "-c"
	}
}

// EX. cmdStrArray := "ls -l"
// no command output, just return error/nil
func EZCMD(cmdStr string) error {
	cmd := exec.Command(CMDArgs[0], CMDArgs[1], cmdStr)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// EX. cmdStrArray := "ls -l"
// return out AND error/nil
func CMD(cmdStr string) (string, error) {
	cmd := exec.Command(CMDArgs[0], CMDArgs[1], cmdStr)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// EX. cmdStrArray := "ls -l"
// return stdout, stderr, error/nil
func CombinedCMD(cmdStr string) (string, string, error) {
	cmd := exec.Command(CMDArgs[0], CMDArgs[1], cmdStr)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr, err
}

// 逐行输出命令执行结果
func SuperCMD(cmdStr string) error {
	cmd := exec.Command(CMDArgs[0], CMDArgs[1], cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()
	return nil
}
