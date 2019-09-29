package parser

import "testing"

func TestCasesForParser(t *testing.T) {
	cases := map[string]int{
		"1+1":         2,
		"1+(1+4)":     6,
		"2+4":         6,
		"2*3+1":       7,
		"2 * (3 + 1)": 8,
	}

	for k, v := range cases {
		if result, err := ExecSe(k); err != nil {
			t.Errorf("exec simple expr [%s] expect: %v but error: %s", result, v, err)
			t.FailNow()
		} else {
			if result != v {
				t.Errorf("exec simple expr [%s] expect: %v but got: %v", result, v, result)
				t.FailNow()
			}
		}
	}
}
