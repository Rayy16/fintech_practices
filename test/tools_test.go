package test

import (
	"bytes"
	Init "fintechpractices/init"
	"fintechpractices/tools"
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestMd5(t *testing.T) {
	m1 := tools.GenMD5WithSalt("hello world", tools.Salt)
	m2 := tools.GenMD5("hello worldCC-fintech-practices")
	if m1 == m2 {
		fmt.Println(m1, len(m1))
	} else {
		t.Errorf("%s not equal to %s", m1, m2)
	}
}

func TestEncrypt(t *testing.T) {
	tools.TestEncrypt()
}

func TestParseToken(t *testing.T) {
	token, err := tools.GenToken("admin")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(token)

	time.Sleep(time.Second)
	claims, err := tools.ParseToken(token)
	if err != nil {
		t.Error(err.Error())
	}
	if claims.UserAccount != "admin" {
		t.Fail()
	} else {
		fmt.Println(claims)
	}
}

func TestSplitCmd(t *testing.T) {
	testCases := []struct {
		cmdStr string
		name   string
		args   []string
	}{
		{
			cmdStr: `python inference.py --driven_audio chinese_news.wav --source_image WDA_KatieHill_000.mp4 --result_dir examples --enhancer gfpgan --file_name testdemo`,
			name:   "python",
			args:   []string{"inference.py", "--driven_audio", "chinese_news.wav", "--source_image", "WDA_KatieHill_000.mp4", "--result_dir", "examples", "--enhancer", "gfpgan", "--file_name", "testdemo"},
		},
		{
			cmdStr: "ls -l",
			name:   "ls",
			args:   []string{"-l"},
		},
		{
			cmdStr: "",
			name:   "",
			args:   nil,
		},
	}
	for _, testCase := range testCases {
		name, args := tools.SplitCmd(testCase.cmdStr)
		if name != testCase.name {
			t.Errorf("expect %s, got %s", testCase.name, name)
		}

		if args == nil && testCase.args == nil {
			continue
		}
		if args == nil || testCase.args == nil || len(args) != len(testCase.args) {
			t.Errorf("expect %v, got %v", testCase.args, args)
			continue
		}
		for i := range testCase.args {
			if testCase.args[i] != args[i] {
				t.Errorf("expect %v, got %v", testCase.args, args)
			}
		}
	}
}

func TestRunCmd(t *testing.T) {
	testCases := []struct {
		cmdStr string
		std    [2]string
		err    error
	}{
		{
			cmdStr: `go version`,
			std: [2]string{
				`go version go1.20.6 windows/amd64`,
				"",
			},
			err: nil,
		},
		{
			cmdStr: `Python --version`,
			std: [2]string{
				`Python 3.9.7`,
				"",
			},
			err: nil,
		},
	}
	for _, testCase := range testCases {
		std, err := tools.RunCmd(tools.SplitCmd(testCase.cmdStr))
		if !strings.Contains(std[0], testCase.std[0]) {
			t.Errorf("expected %v, got %v", testCase.std[0], std[0])
		}
		if !strings.Contains(std[1], testCase.std[1]) {
			t.Errorf("expected %v, got %v", testCase.std[1], std[1])
		}
		if err != testCase.err {
			t.Errorf("expected %v, got %v", testCase.err, err)
		}
	}
}

func TestExec(t *testing.T) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("python", `--version`)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout.String(), stderr.String())
}

func TestExecTool(t *testing.T) {
	std, err := tools.RunCmd("activate", []string{"so-vits-svc_cc"})
	if err != nil {
		fmt.Printf("err: %s\nstdout: %s\nstderr: %s\n", err.Error(), std[0], std[1])
		t.Fail()
	}
	fmt.Printf("stdout: %s\nstderr: %s\n", std[0], std[1])

	std, err = tools.RunCmd("activate", []string{"base"})
	if err != nil {
		fmt.Printf("err: %s\nstdout: %s\nstderr: %s\n", err.Error(), std[0], std[1])
		t.Fail()
	}
	fmt.Printf("stdout: %s\nstderr: %s\n", std[0], std[1])
}

func TestExtractVedio(t *testing.T) {
	Init.Initialization()

	src := `C:\Users\Administrator\Desktop\CC\12-8_clone\fintech_practices\files\dp\96860df8c3aa4859da7283764e4c204c.mp4.mp4`
	dest := `C:\Users\Administrator\Desktop\CC\12-8_clone\fintech_practices\files\cover_image\dbc732a1fbf712bb3a07e62f86800cd6`
	err := tools.ExtractVedioToImage(src, dest)
	if err != nil {
		t.Error(err.Error())
	}
}
