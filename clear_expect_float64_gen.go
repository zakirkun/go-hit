// +build !generate

package hit

import errortrace "github.com/Eun/go-hit/errortrace"

// ⚠️⚠️⚠️ This file was autogenerated by generators/clear/clear ⚠️⚠️⚠️ //

// IClearExpectFloat64 provides methods to clear steps.
type IClearExpectFloat64 interface {
	IStep
	// Between clears all matching Between steps
	Between(value ...float64) IStep
	// Equal clears all matching Equal steps
	Equal(value ...float64) IStep
	// GreaterOrEqualThan clears all matching GreaterOrEqualThan steps
	GreaterOrEqualThan(value ...float64) IStep
	// GreaterThan clears all matching GreaterThan steps
	GreaterThan(value ...float64) IStep
	// LessOrEqualThan clears all matching LessOrEqualThan steps
	LessOrEqualThan(value ...float64) IStep
	// LessThan clears all matching LessThan steps
	LessThan(value ...float64) IStep
	// NotBetween clears all matching NotBetween steps
	NotBetween(value ...float64) IStep
	// NotEqual clears all matching NotEqual steps
	NotEqual(value ...float64) IStep
	// NotOneOf clears all matching NotOneOf steps
	NotOneOf(value ...float64) IStep
	// OneOf clears all matching OneOf steps
	OneOf(value ...float64) IStep
}
type clearExpectFloat64 struct {
	cp callPath
	tr *errortrace.ErrorTrace
}

func newClearExpectFloat64(cp callPath) IClearExpectFloat64 {
	return &clearExpectFloat64{cp: cp, tr: ett.Prepare()}
}
func (v *clearExpectFloat64) trace() *errortrace.ErrorTrace {
	return v.tr
}
func (*clearExpectFloat64) when() StepTime {
	return CleanStep
}
func (v *clearExpectFloat64) callPath() callPath {
	return v.cp
}
func (v *clearExpectFloat64) exec(hit *hitImpl) error {
	if err := removeSteps(hit, v.callPath()); err != nil {
		return err
	}
	return nil
}
func (v *clearExpectFloat64) Between(value ...float64) IStep {
	return removeStep(v.callPath().Push("Between", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) Equal(value ...float64) IStep {
	return removeStep(v.callPath().Push("Equal", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) GreaterOrEqualThan(value ...float64) IStep {
	return removeStep(v.callPath().Push("GreaterOrEqualThan", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) GreaterThan(value ...float64) IStep {
	return removeStep(v.callPath().Push("GreaterThan", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) LessOrEqualThan(value ...float64) IStep {
	return removeStep(v.callPath().Push("LessOrEqualThan", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) LessThan(value ...float64) IStep {
	return removeStep(v.callPath().Push("LessThan", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) NotBetween(value ...float64) IStep {
	return removeStep(v.callPath().Push("NotBetween", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) NotEqual(value ...float64) IStep {
	return removeStep(v.callPath().Push("NotEqual", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) NotOneOf(value ...float64) IStep {
	return removeStep(v.callPath().Push("NotOneOf", float64SliceToInterfaceSlice(value)))
}
func (v *clearExpectFloat64) OneOf(value ...float64) IStep {
	return removeStep(v.callPath().Push("OneOf", float64SliceToInterfaceSlice(value)))
}
