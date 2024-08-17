package main

import (
	"fmt"

	bexpr "github.com/hashicorp/go-bexpr"
)

type Internal struct {
	Name string

	// Use an alternative name for referencing this field in expressions
	Values []int `bexpr:"fields"`

	// Hides this field so it cannot be used in expression evaluation
	Hidden int `bexpr:"-"`

	// Unexported fields are not available for use by the evaluator
	unexported int
}

type Matchable struct {
	Jwt_claims map[string]interface{}
}

var data Matchable = Matchable{
	Jwt_claims: map[string]interface{}{
		"sub": "organization:zain-stacks-testing:project:zzz:stack:zst:deployment:production:operation:plan",
	},
}

var expressions []string = []string{
	"Jwt_claims.sub matches `organization:zain-stacks-testing:project:zzz:stack:zst:deployment:production:operation:*`",
}

func main() {
	for _, expression := range expressions {
		eval, err := bexpr.CreateEvaluator(expression)

		if err != nil {
			fmt.Printf("Failed to create evaluator for expression %q: %v\n", expression, err)
			continue
		}

		result, err := eval.Evaluate(data)
		if err != nil {
			fmt.Printf("Failed to run evaluation of expression %q: %v\n", expression, err)
			continue
		}

		fmt.Printf("Result of expression %q evaluation: %t\n", expression, result)
	}
}
