package spike

import "testing"

func TestEmbedded(t *testing.T) {
	if len(Data) == 0 {
		t.Fatal("embedded file is empty")
	}
}
