package ssdeep

/*
#cgo linux LDFLAGS:-L/usr/local/lib/ -lfuzzy -ldl -I/usr/local/include/
#include <stdlib.h>
#include <fuzzy.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type FuzzyState struct {
	cstate *_Ctype_struct_fuzzy_state
}

/*
Construct a fuzzy_state object and return it.
*/
func New() (*FuzzyState, error) {
	var cstate *_Ctype_struct_fuzzy_state
	if cstate = C.fuzzy_new(); cstate == nil {
		return nil, errors.New("")
	}
	return &FuzzyState{cstate}, nil
}

/*
Create a copy of a fuzzy_state object and return it.
*/
func (fs *FuzzyState) Clone() (*FuzzyState, error) {
	var cstate *_Ctype_struct_fuzzy_state
	if cstate = C.fuzzy_clone(fs.cstate); cstate == nil {
		return nil, errors.New("")
	}
	return &FuzzyState{cstate}, nil
}

/*
Dispose a fuzzy state.
*/
func (fs *FuzzyState) Free() {
	C.fuzzy_free(fs.cstate)
}

/*
Feed the data contained in the given buffer to the state.<F37>
*/
func (fs *FuzzyState) Update(str string) error {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	length := C.size_t(len(str))
	if C.fuzzy_update(fs.cstate, (*C.uchar)(unsafe.Pointer(cstr)), length) != 0 {
		return errors.New("")
	}

	return nil

}

/*
Obtain the fuzzy hash from the state.
*/
func (fs *FuzzyState) Digest() (string, error) {
	buf := (*C.char)(C.calloc(C.FUZZY_MAX_RESULT, 1))
	defer C.free(unsafe.Pointer(buf))

	if C.fuzzy_digest(fs.cstate, buf, C.uint(0)) != 0 {
		return "", errors.New("")
	}

	return C.GoString(buf), nil

}

/*
Computes the match score between two fuzzy hash signatures.
*/
func Compare(str1, str2 string) (int, error) {
	cstr1 := C.CString(str1)
	defer C.free(unsafe.Pointer(cstr1))

	cstr2 := C.CString(str2)
	defer C.free(unsafe.Pointer(cstr2))

	if score := C.fuzzy_compare(cstr1, cstr2); score >= 0 {
		return int(score), nil
	}

	return -1, errors.New("")
}

/*
Obtain the fuzzy hash from the state.
*/
func HashString(str string) (string, error) {
	buf := (*C.char)(C.calloc(C.FUZZY_MAX_RESULT, 1))
	defer C.free(unsafe.Pointer(buf))
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	length := C.uint32_t(len(str))
	if C.fuzzy_hash_buf((*C.uchar)(unsafe.Pointer(cstr)), length, buf) != 0 {
		return "", errors.New("")
	}

	return C.GoString(buf), nil
}

/*
Compute the fuzzy hash of a file.
*/
func HashFilename(filename string) (string, error) {
	buf := (*C.char)(C.calloc(C.FUZZY_MAX_RESULT, 1))
	defer C.free(unsafe.Pointer(buf))
	cf := C.CString(filename)
	defer C.free(unsafe.Pointer(cf))

	if C.fuzzy_hash_filename(cf, buf) != 0 {
		return "", errors.New("")
	}

	return C.GoString(buf), nil
}
