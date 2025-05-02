package judge

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"go-code-platform/problems"
)

type Submission struct {
	ProblemID string `json:"problem_id"`
	Code      string `json:"code"`
}

func HandleSubmission(c *gin.Context) {
	var sub Submission
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Find problem
	var prob problems.Problem
	found := false
	for _, p := range problems.ProblemSet {
		if p.ID == sub.ProblemID {
			prob = p
			found = true
			break
		}
	}
	if !found {
		c.JSON(404, gin.H{"error": "Problem not found"})
		return
	}

	// Write code to temp file
	tmpFile, _ := ioutil.TempFile("", "*.go")
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte(sub.Code))
	tmpFile.Close()

	// Compile
	exeFile := tmpFile.Name() + ".out"
	cmd := exec.Command("go", "build", "-o", exeFile, tmpFile.Name())
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		c.JSON(400, gin.H{"error": "Compilation failed", "details": stderr.String()})
		return
	}
	defer os.Remove(exeFile)

	// Test cases
	results := []string{}
	for _, tc := range prob.TestCases {
		cmd := exec.Command(exeFile)
		cmd.Stdin = bytes.NewBufferString(tc.Input)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		err := runWithTimeout(cmd, 2*time.Second)
		if err != nil {
			results = append(results, fmt.Sprintf("❌ Timeout or error: %s", err.Error()))
			continue
		}

		actual := out.String()
		if trim(actual) == trim(tc.Output) {
			results = append(results, "✅ Passed")
		} else {
			results = append(results, fmt.Sprintf("❌ Failed: got '%s', expected '%s'", trim(actual), trim(tc.Output)))
		}
	}

	c.JSON(200, gin.H{
		"results": results,
	})
}

func trim(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}

func runWithTimeout(cmd *exec.Cmd, timeout time.Duration) error {
	done := make(chan error)
	go func() { done <- cmd.Run() }()
	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		cmd.Process.Kill()
		return fmt.Errorf("execution timeout")
	}
}
