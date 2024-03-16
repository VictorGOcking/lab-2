package lab2

import (
	check "gopkg.in/check.v1"
	"os"
	"strings"
)

var _ = check.Suite(&MySuite{})

var expectedHandler = []TestCase{
	{
		task:     "Valid expression convert",
		postfix:  "4 2 - 3 * 5 +",
		prefix:   "+ * - 4 2 3 5",
		errorMsg: "",
	},
	{
		task:     "Invalid expression convert",
		postfix:  "4 2 - 3 * 5 $",
		prefix:   "+ * - 4 2 3 5",
		errorMsg: "incorrect syntax",
	},
}

func (s *MySuite) TestCompute(c *check.C) {
	for _, exp := range expectedHandler {
		writer := strings.Builder{}
		handler := ComputeHandler{
			Reader: strings.NewReader(exp.postfix),
			Writer: &writer,
		}

		err := handler.Compute()
		res := writer.String()

		if exp.errorMsg == "" {
			c.Assert(err, check.IsNil)
		} else {
			c.Assert(err.Error(), check.Equals, exp.errorMsg)
		}

		if err == nil {
			c.Assert(res, check.Equals, exp.prefix)
		}
	}
}

func ExampleCompute() {
	handler := ComputeHandler{
		Reader: strings.NewReader("4 2 - 3 * 5 +"),
		Writer: os.Stdout,
	}

	handler.Compute()

	// Output:
	// + * - 4 2 3 5
}
