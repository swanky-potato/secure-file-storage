package sfm

import (
	"errors"
	"testing"
)

func TestCreateChecksum(t *testing.T) {
	dd := []byte("some dummy data")
	cs := createChecksum(dd)
	if string(cs) == "hWm9XbsjOKnX0GsHWGDP/3dYMjyxmcFu1czQHj/0mmk= " {
		t.Error("Checksum does not match with what is expected")
	}
}

func TestCRUDOnFile(t *testing.T) {
	StorageLocation = testStorage
	ns := Namespace{Path: "nsf", GID: 1}
	if err := ns.Create(); err != nil {
		t.Error(err)
	}

	content := []byte("hello world, unit testing file here")
	pass := "password"

	// create File struct to store
	f1 := File{
		Name:      "unit",
		Extention: ".txt",
		Namespace: ns.Path,
		Content:   content,
		UID:       1,
		Extra:     "",
	}

	// store file on disk
	f1.Store(ns, pass)
	// read stored file
	f2, err := ReadFile(ns, f1.Name, pass)
	if err != nil {
		t.Error(err)
	}

	// Update a file
	f1.Name = "applepi"
	// store file on disk
	f1.Update(pass)
	// read stored file
	f2, err = ReadFile(ns, f1.Name, pass)
	if err != nil {
		t.Error(err)
	}
	// Compair return with expected values
	if f1.Name != f2.Name {
		t.Error(errors.New("Stored " + f1.Name + " not equal to input  " + f2.Name))
	}
	if f1.Namespace != f2.Namespace {
		t.Error(errors.New("Stored " + f1.Namespace + " not equal to input " + f2.Namespace))
	}

	// test removing a file
	if err := f1.Delete(); err != nil {
		t.Error(err)
	}
	// cleanup namespace
	if err := ns.Delete(); err != nil {
		t.Fatal(err)
	}
}
