package lab2

import (
	"fmt"
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})

var expected = []struct {
	task     string
	postfix  string
	prefix   string
	errorMsg string
}{
	{
		task:     "Converts simple expression",
		postfix:  "4 2 - 3 * 5 +",
		prefix:   "+ * - 4 2 3 5",
		errorMsg: "",
	},
	{
		task:     "Converts simple expression",
		postfix:  "4 2 3 ^ 6 / -",
		prefix:   "- 4 / ^ 2 3 6",
		errorMsg: "",
	},
	{
		task:     "Converts complex expression",
		postfix:  "5 3 + 8 2 ^ * 4 / 6 - 7 * 9 +",
		prefix:   "+ * - / * + 5 3 ^ 8 2 4 6 7 9",
		errorMsg: "",
	},
	{
		task:     "Converts complex expression",
		postfix:  "7 2 ^ 4 3 * + 5 / 9 - 6 * 8 2 / +",
		prefix:   "+ * - / + ^ 7 2 * 4 3 5 9 6 / 8 2",
		errorMsg: "",
	},
	{
		task:     "Converts complex expression",
		postfix:  "8 4 / 2 * 5 3 ^ + 7 - 6 * 9 / 4 2 ^ -",
		prefix:   "- / * - + * / 8 4 2 ^ 5 3 7 6 9 ^ 4 2",
		errorMsg: "",
	},
	{
		task:     "Considers empty expression invalid",
		postfix:  "",
		prefix:   "",
		errorMsg: "empty expression",
	},
	{
		task:     "Considers non-postfix syntax invalid",
		postfix:  "2 + 10 * 6 / 3",
		prefix:   "",
		errorMsg: "incorrect syntax",
	},
	{
		task:     "Considers wrong number of operators invalid",
		postfix:  "2 3 + * /",
		prefix:   "",
		errorMsg: "incorrect number of operators",
	},
	{
		task:     "Considers wrongly spaced operators invalid",
		postfix:  "2 3 7* +",
		prefix:   "",
		errorMsg: "incorrect syntax",
	},
	{
		task:     "Considers wrongly spaced operators invalid",
		postfix:  "2 3 * 4 ++",
		prefix:   "",
		errorMsg: "incorrect operators separation",
	},
	{
		task:     "Considers wrong symbols invalid",
		postfix:  "2 3 * 4 ~",
		prefix:   "",
		errorMsg: "incorrect syntax",
	},
	{
		task:     "Considers letters in expression invalid",
		postfix:  "4 2 - b * 5 +",
		prefix:   "",
		errorMsg: "incorrect syntax",
	},
}

func (s *MySuite) TestPostfixToPrefix(c *check.C) {

	for _, exp := range expected {
		res, err := PostfixToPrefix(exp.postfix)
		if exp.errorMsg == "" {
			c.Assert(err, check.IsNil)
		} else {
			c.Assert(err.Error(), check.Equals, exp.errorMsg)
		}

		if res != exp.prefix {
			c.Errorf("Task: %s\nExpected: %s\nActual: %s", exp.task, exp.prefix, res)
		}
	}
}

func ExamplePostfixToPrefix() {
	res, err := PostfixToPrefix("4 2 - 3 * 5 +")

	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(res)

	// Output:
	// + * - 4 2 3 5
}
