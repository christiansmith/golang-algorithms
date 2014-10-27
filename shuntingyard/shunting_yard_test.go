package shuntingyard

import "testing"

func TestShuntingYard(t *testing.T) {
	exprs := map[string]int64{
		"(5+(3*5))":  int64(20),
		"((64/8)-5)": int64(3),
	}

	for expr, expecting := range exprs {
		result := EvaluateInfix(expr)
		if result != expecting {
			t.Error("Expected", expecting, "got", result)
		}
	}
}
