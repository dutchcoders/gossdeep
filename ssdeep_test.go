package ssdeep

import (
	"io/ioutil"
	"testing"
)

func TestHashString(t *testing.T) {
	gotData, err := HashString("test")
	if err != nil {
		t.Fatal(err)
	}

	data := "3:Hn:Hn"
	if gotData != data {
		t.Errorf("result mismatch:\nwant: %v\n got: %v", data, gotData)
	}
}

func TestFilename(t *testing.T) {
	f, err := ioutil.TempFile("", "fuzzy")
	if err != nil {
		t.Fatal(err)
	}

	f.Write([]byte("test"))
	f.Close()

	gotData, err := HashFilename(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	data := "3:Hn:Hn"
	if gotData != data {
		t.Errorf("hash mismatch:\nwant: %v\n got: %v", data, gotData)
	}
}

func TestClone(t *testing.T) {
	fs, err := New()
	if err != nil {
		t.Fatal(err)
	}

	fs.Update("test")
	defer fs.Free()

	fs2, err := fs.Clone()
	if err != nil {
		t.Fatal(err)
	}

	defer fs2.Free()

	err = fs.Update("test")
	if err != nil {
		t.Fatal(err)
	}

	gotData, err := fs2.Digest()
	if err != nil {
		t.Fatal(err)
	}

	data := "3:Hn:Hn"
	if gotData != data {
		t.Errorf("clone mismatch:\nwant: %v\n got: %v", data, gotData)
	}

	gotData, err = fs.Digest()
	if err != nil {
		t.Fatal(err)
	}

	data = "3:HHn:HH"
	if gotData != data {
		t.Errorf("clone mismatch:\nwant: %v\n got: %v", data, gotData)
	}

}

func TestScore(t *testing.T) {
	gotData, err := Compare("3:Hn:Hn", "3:Hn:Hn")
	if err != nil {
		t.Fatal(err)
	}

	data := 100
	if gotData != data {
		t.Errorf("score mismatch:\nwant: %v\n got: %v", data, gotData)
	}
}
