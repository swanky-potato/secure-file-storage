package sfm

import "testing"

func TestCreateHash(t *testing.T) {
	hash := createHash("passphrase")
	if hash != "d3eb8e78cad217d10feca1080bfb61dc" {
		t.Error("hash is not matching expected value")
	}
}

func TestEncryptDecrypt(t *testing.T) {
	bd := []byte("data to encrypt")
	enc := encrypt(bd, createHash("passphrase"))
	d := decrypt(enc, createHash("passphrase"))
	if string(d) != string(bd) {
		t.Errorf("encrypted data mismatched decrypt \n encrypted: %s \n decrypted: %s", string(bd), string(d))
	}
}
