package dns

import (
	"fmt"
)

type configDNSError interface {
	error
	Network() bool
	NotFound() bool
	FailedToSave() bool
	ValidationFailed() bool
}

type ZoneError struct {
	zoneName         string
	httpErrorMessage string
	err              error
}

func (e *ZoneError) Network() bool {
	if httpErrorMessage != "" {
		return true
	}
	return false
}

func (e *ZoneError) NotFound() bool {
	if e.err == nil && httpErrorMessage == "" {
		return true
	}
	return false
}

func (e *ZoneError) FailedToSave() bool {
	return false
}

func (e *ZoneError) ValidationFailed() bool {
	return false
}

func (e *ZoneError) Error() string {
	if e == nil {
		return "<nil>"
	}

	if e.Network() {
		return fmt.Sprintf("Zone \"%s\" network error: [%s]", e.zoneName, e.httpErrorMessage)
	}

	if e.NotFound() {
		return fmt.Sprintf("Zone \"%s\" not found.", e.zoneName)
	}

	if e.FailedToSave() {
		return fmt.Sprintf("Zone \"%s\" failed to save: [%s]", e.zoneName, e.err.Error())
	}

	if e.ValidationFailed() {
		return fmt.Sprintf("Zone \"%s\" validation failed: [%s]", e.zoneName, e.err.Error())
	}
}

type RecordError struct {
	fieldName        string
	httpErrorMessage string
	err              error
}

func (e *ZoneError) Network() bool {
	if httpErrorMessage != "" {
		return true
	}
	return false
}

func (e *RecordError) NotFound() bool {
	return false
}

func (e *RecordError) FailedToSave() bool {
	if fieldName == "" {
		return true
	}
	return false
}

func (e *RecordError) ValidationFailed() bool {
	if fieldName != "" {
		return true
	}
	return false
}

func (e *RecordError) Error() string {
	if e == nil {
		return "<nil>"
	}

	if e.Network() {
		return fmt.Sprintf("Record network error: [%s]", e.httpErrorMessage)
	}

	if e.NotFound() {
		return fmt.Sprintf("Record not found.")
	}

	if e.FailedToSave() {
		return fmt.Sprintf("Record failed to save: [%s]", e.err.Error())
	}

	if e.ValidationFailed() {
		return fmt.Sprintf("Record validation failed for field [%s]: [%s]", e.fieldName, e.err.Error())
	}
}
