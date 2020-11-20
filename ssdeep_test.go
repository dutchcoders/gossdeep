package ssdeep

import (
	"io/ioutil"
	"math/rand"
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

/*
Obtain the fuzzy hash from the state.
*/
func TestHashBytes(t *testing.T) {
	gotData, err := HashBytes([]byte("test"))
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

func BenchmarkSSDeepString_1k(b *testing.B) {
	runSSDeepBenchmarkString(1024, b)
}

func BenchmarkSSDeepString_10k(b *testing.B) {
	runSSDeepBenchmarkString(10*1024, b)
}

func BenchmarkSSDeepString_100k(b *testing.B) {
	runSSDeepBenchmarkString(100*1024, b)
}

func BenchmarkSSDeepString_250k(b *testing.B) {
	runSSDeepBenchmarkString(250*1024, b)
}

func BenchmarkSSDeepString_500k(b *testing.B) {
	runSSDeepBenchmarkString(500*1024, b)
}

func BenchmarkSSDeepConvertString_1k(b *testing.B) {
	runSSDeepBenchmarkBytes(true, 1024, b)
}

func BenchmarkSSDeepConvertString_10k(b *testing.B) {
	runSSDeepBenchmarkBytes(true, 10*1024, b)
}

func BenchmarkSSDeepConvertString_100k(b *testing.B) {
	runSSDeepBenchmarkBytes(true, 100*1024, b)
}

func BenchmarkSSDeepConvertString_250k(b *testing.B) {
	runSSDeepBenchmarkBytes(true, 250*1024, b)
}

func BenchmarkSSDeepConvertString_500k(b *testing.B) {
	runSSDeepBenchmarkBytes(true, 500*1024, b)
}

func BenchmarkSSDeepBytes_1k(b *testing.B) {
	runSSDeepBenchmarkBytes(false, 1024, b)
}

func BenchmarkSSDeepBytes_10k(b *testing.B) {
	runSSDeepBenchmarkBytes(false, 10*1024, b)
}

func BenchmarkSSDeepBytes_100k(b *testing.B) {
	runSSDeepBenchmarkBytes(false, 100*1024, b)
}

func BenchmarkSSDeepBytes_250k(b *testing.B) {
	runSSDeepBenchmarkBytes(false, 250*1024, b)
}

func BenchmarkSSDeepBytes_500k(b *testing.B) {
	runSSDeepBenchmarkBytes(false, 500*1024, b)
}

func runSSDeepBenchmarkBytes(convertString bool, i int, b *testing.B) {
	bs := randomBytes(i, b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if convertString {
			HashString(string(bs))
		}
		HashBytes(bs)
	}

}

func randomBytes(i int, b *testing.B) []byte {
	bs := make([]byte, i)
	_, err := rand.Read(bs)
	if err != nil {
		b.Fatal(err)
	}
	return bs
}

func runSSDeepBenchmarkString(i int, b *testing.B) {
	s := string(randomBytes(i, b))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashString(s)
	}

}
