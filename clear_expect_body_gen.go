// +build !generate

package hit

import errortrace "github.com/Eun/go-hit/errortrace"

// ⚠️⚠️⚠️ This file was autogenerated by generators/clear/clear ⚠️⚠️⚠️ //

// IClearExpectBody provides methods to clear steps.
type IClearExpectBody interface {
	IStep
	// Bytes clears all matching Bytes steps
	Bytes() IClearExpectBytes
	// Float32 clears all matching Float32 steps
	Float32() IClearExpectFloat32
	// Float64 clears all matching Float64 steps
	Float64() IClearExpectFloat64
	// FormValues clears all matching FormValues steps
	FormValues(value ...string) IClearExpectFormValues
	// Int clears all matching Int steps
	Int() IClearExpectInt
	// Int16 clears all matching Int16 steps
	Int16() IClearExpectInt16
	// Int32 clears all matching Int32 steps
	Int32() IClearExpectInt32
	// Int64 clears all matching Int64 steps
	Int64() IClearExpectInt64
	// Int8 clears all matching Int8 steps
	Int8() IClearExpectInt8
	// JSON clears all matching JSON steps
	JSON() IClearExpectBodyJSON
	// String clears all matching String steps
	String() IClearExpectString
	// Uint clears all matching Uint steps
	Uint() IClearExpectUint
	// Uint16 clears all matching Uint16 steps
	Uint16() IClearExpectUint16
	// Uint32 clears all matching Uint32 steps
	Uint32() IClearExpectUint32
	// Uint64 clears all matching Uint64 steps
	Uint64() IClearExpectUint64
	// Uint8 clears all matching Uint8 steps
	Uint8() IClearExpectUint8
}
type clearExpectBody struct {
	cp callPath
	tr *errortrace.ErrorTrace
}

func newClearExpectBody(cp callPath) IClearExpectBody {
	return &clearExpectBody{cp: cp, tr: ett.Prepare()}
}
func (v *clearExpectBody) trace() *errortrace.ErrorTrace {
	return v.tr
}
func (*clearExpectBody) when() StepTime {
	return CleanStep
}
func (v *clearExpectBody) callPath() callPath {
	return v.cp
}
func (v *clearExpectBody) exec(hit *hitImpl) error {
	if err := removeSteps(hit, v.callPath()); err != nil {
		return err
	}
	return nil
}
func (v *clearExpectBody) Bytes() IClearExpectBytes {
	return newClearExpectBytes(v.callPath().Push("Bytes", nil))
}
func (v *clearExpectBody) Float32() IClearExpectFloat32 {
	return newClearExpectFloat32(v.callPath().Push("Float32", nil))
}
func (v *clearExpectBody) Float64() IClearExpectFloat64 {
	return newClearExpectFloat64(v.callPath().Push("Float64", nil))
}
func (v *clearExpectBody) FormValues(value ...string) IClearExpectFormValues {
	return newClearExpectFormValues(v.callPath().Push("FormValues", stringSliceToInterfaceSlice(value)))
}
func (v *clearExpectBody) Int() IClearExpectInt {
	return newClearExpectInt(v.callPath().Push("Int", nil))
}
func (v *clearExpectBody) Int16() IClearExpectInt16 {
	return newClearExpectInt16(v.callPath().Push("Int16", nil))
}
func (v *clearExpectBody) Int32() IClearExpectInt32 {
	return newClearExpectInt32(v.callPath().Push("Int32", nil))
}
func (v *clearExpectBody) Int64() IClearExpectInt64 {
	return newClearExpectInt64(v.callPath().Push("Int64", nil))
}
func (v *clearExpectBody) Int8() IClearExpectInt8 {
	return newClearExpectInt8(v.callPath().Push("Int8", nil))
}
func (v *clearExpectBody) JSON() IClearExpectBodyJSON {
	return newClearExpectBodyJSON(v.callPath().Push("JSON", nil))
}
func (v *clearExpectBody) String() IClearExpectString {
	return newClearExpectString(v.callPath().Push("String", nil))
}
func (v *clearExpectBody) Uint() IClearExpectUint {
	return newClearExpectUint(v.callPath().Push("Uint", nil))
}
func (v *clearExpectBody) Uint16() IClearExpectUint16 {
	return newClearExpectUint16(v.callPath().Push("Uint16", nil))
}
func (v *clearExpectBody) Uint32() IClearExpectUint32 {
	return newClearExpectUint32(v.callPath().Push("Uint32", nil))
}
func (v *clearExpectBody) Uint64() IClearExpectUint64 {
	return newClearExpectUint64(v.callPath().Push("Uint64", nil))
}
func (v *clearExpectBody) Uint8() IClearExpectUint8 {
	return newClearExpectUint8(v.callPath().Push("Uint8", nil))
}
