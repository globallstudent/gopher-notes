package problems

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Problem struct {
	ID          string
	Title       string
	Description string
	TestCases   []TestCase
}

type TestCase struct {
	Input  string
	Output string
}

var ProblemSet = []Problem{
	{
		ID:          "sum",
		Title:       "Sum Two Numbers",
		Description: "Read two integers and return their sum.",
		TestCases: []TestCase{
			{Input: "2 3", Output: "5"},
			{Input: "10 15", Output: "25"},
		},
	},
}

func ListProblems(c *gin.Context) {
	c.JSON(http.StatusOK, ProblemSet)
}
