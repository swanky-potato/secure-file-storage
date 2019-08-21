package sfm

import "testing"

func TestCreateChecksum(t *testing.T) {
	dd := []byte("some dummy data")
	cs, err := createChecksum(dd)
	if err != nil {
		t.Error(err)
	}
	if string(cs) == "hWm9XbsjOKnX0GsHWGDP/3dYMjyxmcFu1czQHj/0mmk= " {
		t.Error("Checksum does not match with what is expected")
	}
}
