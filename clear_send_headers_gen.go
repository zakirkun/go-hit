// +build !generate

package hit

import errortrace "github.com/Eun/go-hit/errortrace"

// ⚠️⚠️⚠️ This file was autogenerated by generators/clear/clear ⚠️⚠️⚠️ //

// IClearSendHeaders provides methods to clear steps.
type IClearSendHeaders interface {
	IStep
	// Add clears all matching Add steps
	Add(value ...interface{}) IStep
}
type clearSendHeaders struct {
	cp callPath
	tr *errortrace.ErrorTrace
}

func newClearSendHeaders(cp callPath) IClearSendHeaders {
	return &clearSendHeaders{cp: cp, tr: ett.Prepare()}
}
func (v *clearSendHeaders) trace() *errortrace.ErrorTrace {
	return v.tr
}
func (*clearSendHeaders) when() StepTime {
	return CleanStep
}
func (v *clearSendHeaders) callPath() callPath {
	return v.cp
}
func (v *clearSendHeaders) exec(hit *hitImpl) error {
	if err := removeSteps(hit, v.callPath()); err != nil {
		return err
	}
	return nil
}
func (v *clearSendHeaders) Add(value ...interface{}) IStep {
	return removeStep(v.callPath().Push("Add", value))
}
