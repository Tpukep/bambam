package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestBasicTypesInStruct(t *testing.T) {

	cv.Convey("Given a parsable golang source file with a simple struct", t, func() {
		cv.Convey("then we can extract the struct and convert the basic types to capnp", func() {

			in1 := `
type In1 struct {
	Str string
	N   int
    D   float64
}`
			expect1 := `struct In1Capn { str @0: Text; n @1: Int64; d @2: Float64; } `
			// List(Float64) next

			cv.So(ExtractString2String(in1), ShouldMatchModuloSpaces, expect1)
		})
	})
}

func TestCaseConversion(t *testing.T) {

	cv.Convey("Given a parsable golang source file with a simple struct", t, func() {
		cv.Convey("then the capnp generated struct has uppercase struct name and lowercase field name, as per capnp requirements.", func() {

			in1 := `
type in1 struct {
	Str string
	N   int
    D   float64
}`
			expect1 := `struct In1Capn { str @0: Text; n @1: Int64; d @2: Float64;}`
			// List(Float64) next

			cv.So(ExtractString2String(in1), ShouldMatchModuloSpaces, expect1)
		})
	})
}

func TestNonStructTypeDefsAreIgnored(t *testing.T) {

	cv.Convey("Given a go type definition 'type SyncMsg int32' that isn't a struct definition,", t, func() {
		cv.Convey("then bambam should ignore it and not define a new SyncMsgCapn type", func() {

			in1 := `type SyncMsg int32`
			expect1 := ``

			cv.So(ExtractString2String(in1), ShouldMatchModuloSpaces, expect1)
		})
	})
}
