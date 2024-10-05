package test

import (
	"os/exec"
	"testing"
)

func TestTryGetzhFrompy(t *testing.T) {
	//zh_question := "你好"
	en_question := "hello"

	cmd1 := exec.Command("python", "../AI/AIpic.py", en_question)
	out1, _ := cmd1.CombinedOutput()

	t.Log(string(out1))
}
