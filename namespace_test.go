package sfm

import (
	"errors"
	"testing"
)

var testStorage = "/Users/steven/Projects/Golang/src/github.com/SionX/secure-file-storage/files"

func TestAddNamespace(t *testing.T) {
	StorageLocation = testStorage
	nsin := Namespace{Path: "nsx", GID: 1}
	if err := nsin.Create(); err != nil {
		t.Fatal(err)
	}
}

func TestGetNamespace(t *testing.T) {
	StorageLocation = testStorage
	nscheck := Namespace{Path: "nsx", GID: 1}
	ns, err := GetNamespace("nsx")
	if err != nil {
		t.Fatal(err)
	}
	if ns != nscheck {
		t.Fatal(errors.New("returned namespace not as expected"))
	}
}

func TestUpdateNamespace(t *testing.T) {
	StorageLocation = testStorage
	nscheck := Namespace{Path: "nsx", GID: 2}
	if err := nscheck.Update(); err != nil {
		t.Fatal(err)
	}
	ns, err := GetNamespace("nsx")
	if err != nil {
		t.Fatal(err)
	}
	if ns != nscheck {
		t.Fatal(errors.New("returned namespace not as expected"))
	}
}

func TestRemoveNamespace(t *testing.T) {
	StorageLocation = testStorage
	ns, err := GetNamespace("nsx")
	if err != nil {
		t.Fatal(err)
	}
	if err := ns.Delete(); err != nil {
		t.Fatal(err)
	}
}

func TestNamespaceMap(t *testing.T) {
	StorageLocation = testStorage
	spacescheck := []Namespace{}

	// add test namespace structure
	nsadd := []Namespace{}
	nsadd = append(nsadd, Namespace{Path: "nsx", GID: 1})
	nsadd = append(nsadd, Namespace{Path: "nsx/a", GID: 1})
	nsadd = append(nsadd, Namespace{Path: "nsx/b", GID: 1})
	nsadd = append(nsadd, Namespace{Path: "nsx/c/b", GID: 1})
	nsadd = append(nsadd, Namespace{Path: "nsx/b/a", GID: 1})

	for _, ns := range nsadd {
		if err := ns.Create(); err != nil {
			t.Error(err)
		}
		spacescheck = append(spacescheck, ns)
	}

	// map Namespace test structure
	spaces, err := GetNamespaceMapBelow("nsx")
	if err != nil {
		t.Error(err)
	}
	if len(spaces) != len(spacescheck) {
		t.Error(errors.New("Length of []spaces is not as long as expected"))

	}

	// remove namespace test structure
	ns, err := GetNamespace("nsx")
	if err != nil {
		t.Error(err)
	}
	if err := ns.Delete(); err != nil {
		t.Error(err)
	}
}
