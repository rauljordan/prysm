package bls

/*
#cgo bn384 CFLAGS:-UBLS_MAX_OP_UNIT_SIZE -DBLS_MAX_OP_UNIT_SIZE=6
#include "external/bls/bls.h"
*/
import "C"

func figureOut() {
	res := C.blsInit(C.int(1), C.BLS_MAX_OP_UNIT_SIZE)
	fmt.Println(res)
}