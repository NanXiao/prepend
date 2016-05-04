package prepend

import "testing"

func TestPrepend(t *testing.T)  {
	s := []interface{}{1, 2, 3}
	for i, v := range Prepend(1, []interface{}{2, 3}) {
		if s[i] != v {
			t.Error("Prepend failed!")
		}
	}
}
