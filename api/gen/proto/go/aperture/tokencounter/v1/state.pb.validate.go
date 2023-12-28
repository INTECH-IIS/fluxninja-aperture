// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: aperture/tokencounter/v1/state.proto

package tokencounterv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on State with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *State) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on State with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StateMultiError, or nil if none found.
func (m *State) ValidateAll() error {
	return m.validate(true)
}

func (m *State) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRequestsQueued() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StateValidationError{
						field:  fmt.Sprintf("RequestsQueued[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StateValidationError{
						field:  fmt.Sprintf("RequestsQueued[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StateValidationError{
					field:  fmt.Sprintf("RequestsQueued[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRequestsInflight() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StateValidationError{
						field:  fmt.Sprintf("RequestsInflight[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StateValidationError{
						field:  fmt.Sprintf("RequestsInflight[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StateValidationError{
					field:  fmt.Sprintf("RequestsInflight[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetTokenWindow()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StateValidationError{
					field:  "TokenWindow",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StateValidationError{
					field:  "TokenWindow",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokenWindow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StateValidationError{
				field:  "TokenWindow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TokenRate

	if len(errors) > 0 {
		return StateMultiError(errors)
	}

	return nil
}

// StateMultiError is an error wrapping multiple validation errors returned by
// State.ValidateAll() if the designated constraints aren't met.
type StateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StateMultiError) AllErrors() []error { return m }

// StateValidationError is the validation error returned by State.Validate if
// the designated constraints aren't met.
type StateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StateValidationError) ErrorName() string { return "StateValidationError" }

// Error satisfies the builtin error interface
func (e StateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StateValidationError{}

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RequestMultiError, or nil if none found.
func (m *Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	if all {
		switch v := interface{}(m.GetExpiresAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "ExpiresAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiresAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestValidationError{
				field:  "ExpiresAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Tokens

	if all {
		switch v := interface{}(m.GetWaitFor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "WaitFor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "WaitFor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWaitFor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestValidationError{
				field:  "WaitFor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for NumRetries

	if len(errors) > 0 {
		return RequestMultiError(errors)
	}

	return nil
}

// RequestMultiError is an error wrapping multiple validation errors returned
// by Request.ValidateAll() if the designated constraints aren't met.
type RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMultiError) AllErrors() []error { return m }

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on TakeNRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TakeNRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TakeNRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TakeNRequestMultiError, or
// nil if none found.
func (m *TakeNRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TakeNRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	if all {
		switch v := interface{}(m.GetDeadline()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TakeNRequestValidationError{
					field:  "Deadline",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TakeNRequestValidationError{
					field:  "Deadline",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeadline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TakeNRequestValidationError{
				field:  "Deadline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Tokens

	// no validation rules for CanWait

	if len(errors) > 0 {
		return TakeNRequestMultiError(errors)
	}

	return nil
}

// TakeNRequestMultiError is an error wrapping multiple validation errors
// returned by TakeNRequest.ValidateAll() if the designated constraints aren't met.
type TakeNRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TakeNRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TakeNRequestMultiError) AllErrors() []error { return m }

// TakeNRequestValidationError is the validation error returned by
// TakeNRequest.Validate if the designated constraints aren't met.
type TakeNRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TakeNRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TakeNRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TakeNRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TakeNRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TakeNRequestValidationError) ErrorName() string { return "TakeNRequestValidationError" }

// Error satisfies the builtin error interface
func (e TakeNRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTakeNRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TakeNRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TakeNRequestValidationError{}

// Validate checks the field values on TakeNResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TakeNResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TakeNResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TakeNResponseMultiError, or
// nil if none found.
func (m *TakeNResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TakeNResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AvailableNow

	// no validation rules for Current

	// no validation rules for Remaining

	if all {
		switch v := interface{}(m.GetCheckBackAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TakeNResponseValidationError{
					field:  "CheckBackAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TakeNResponseValidationError{
					field:  "CheckBackAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCheckBackAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TakeNResponseValidationError{
				field:  "CheckBackAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Ok

	if len(errors) > 0 {
		return TakeNResponseMultiError(errors)
	}

	return nil
}

// TakeNResponseMultiError is an error wrapping multiple validation errors
// returned by TakeNResponse.ValidateAll() if the designated constraints
// aren't met.
type TakeNResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TakeNResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TakeNResponseMultiError) AllErrors() []error { return m }

// TakeNResponseValidationError is the validation error returned by
// TakeNResponse.Validate if the designated constraints aren't met.
type TakeNResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TakeNResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TakeNResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TakeNResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TakeNResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TakeNResponseValidationError) ErrorName() string { return "TakeNResponseValidationError" }

// Error satisfies the builtin error interface
func (e TakeNResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTakeNResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TakeNResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TakeNResponseValidationError{}

// Validate checks the field values on ReturnNRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReturnNRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReturnNRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReturnNRequestMultiError,
// or nil if none found.
func (m *ReturnNRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReturnNRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	// no validation rules for Tokens

	if len(errors) > 0 {
		return ReturnNRequestMultiError(errors)
	}

	return nil
}

// ReturnNRequestMultiError is an error wrapping multiple validation errors
// returned by ReturnNRequest.ValidateAll() if the designated constraints
// aren't met.
type ReturnNRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReturnNRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReturnNRequestMultiError) AllErrors() []error { return m }

// ReturnNRequestValidationError is the validation error returned by
// ReturnNRequest.Validate if the designated constraints aren't met.
type ReturnNRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReturnNRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReturnNRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReturnNRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReturnNRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReturnNRequestValidationError) ErrorName() string { return "ReturnNRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReturnNRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReturnNRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReturnNRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReturnNRequestValidationError{}

// Validate checks the field values on ReturnNResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReturnNResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReturnNResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReturnNResponseMultiError, or nil if none found.
func (m *ReturnNResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReturnNResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return ReturnNResponseMultiError(errors)
	}

	return nil
}

// ReturnNResponseMultiError is an error wrapping multiple validation errors
// returned by ReturnNResponse.ValidateAll() if the designated constraints
// aren't met.
type ReturnNResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReturnNResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReturnNResponseMultiError) AllErrors() []error { return m }

// ReturnNResponseValidationError is the validation error returned by
// ReturnNResponse.Validate if the designated constraints aren't met.
type ReturnNResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReturnNResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReturnNResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReturnNResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReturnNResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReturnNResponseValidationError) ErrorName() string { return "ReturnNResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReturnNResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReturnNResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReturnNResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReturnNResponseValidationError{}

// Validate checks the field values on CancelQueuedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CancelQueuedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CancelQueuedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CancelQueuedRequestMultiError, or nil if none found.
func (m *CancelQueuedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CancelQueuedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	if len(errors) > 0 {
		return CancelQueuedRequestMultiError(errors)
	}

	return nil
}

// CancelQueuedRequestMultiError is an error wrapping multiple validation
// errors returned by CancelQueuedRequest.ValidateAll() if the designated
// constraints aren't met.
type CancelQueuedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CancelQueuedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CancelQueuedRequestMultiError) AllErrors() []error { return m }

// CancelQueuedRequestValidationError is the validation error returned by
// CancelQueuedRequest.Validate if the designated constraints aren't met.
type CancelQueuedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelQueuedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelQueuedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelQueuedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelQueuedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelQueuedRequestValidationError) ErrorName() string {
	return "CancelQueuedRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CancelQueuedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelQueuedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelQueuedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelQueuedRequestValidationError{}

// Validate checks the field values on CancelQueuedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CancelQueuedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CancelQueuedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CancelQueuedResponseMultiError, or nil if none found.
func (m *CancelQueuedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CancelQueuedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return CancelQueuedResponseMultiError(errors)
	}

	return nil
}

// CancelQueuedResponseMultiError is an error wrapping multiple validation
// errors returned by CancelQueuedResponse.ValidateAll() if the designated
// constraints aren't met.
type CancelQueuedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CancelQueuedResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CancelQueuedResponseMultiError) AllErrors() []error { return m }

// CancelQueuedResponseValidationError is the validation error returned by
// CancelQueuedResponse.Validate if the designated constraints aren't met.
type CancelQueuedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelQueuedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelQueuedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelQueuedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelQueuedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelQueuedResponseValidationError) ErrorName() string {
	return "CancelQueuedResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CancelQueuedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelQueuedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelQueuedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelQueuedResponseValidationError{}

// Validate checks the field values on CancelInflightRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CancelInflightRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CancelInflightRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CancelInflightRequestMultiError, or nil if none found.
func (m *CancelInflightRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CancelInflightRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestId

	if len(errors) > 0 {
		return CancelInflightRequestMultiError(errors)
	}

	return nil
}

// CancelInflightRequestMultiError is an error wrapping multiple validation
// errors returned by CancelInflightRequest.ValidateAll() if the designated
// constraints aren't met.
type CancelInflightRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CancelInflightRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CancelInflightRequestMultiError) AllErrors() []error { return m }

// CancelInflightRequestValidationError is the validation error returned by
// CancelInflightRequest.Validate if the designated constraints aren't met.
type CancelInflightRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelInflightRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelInflightRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelInflightRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelInflightRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelInflightRequestValidationError) ErrorName() string {
	return "CancelInflightRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CancelInflightRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelInflightRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelInflightRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelInflightRequestValidationError{}

// Validate checks the field values on CancelInflightResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CancelInflightResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CancelInflightResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CancelInflightResponseMultiError, or nil if none found.
func (m *CancelInflightResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CancelInflightResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return CancelInflightResponseMultiError(errors)
	}

	return nil
}

// CancelInflightResponseMultiError is an error wrapping multiple validation
// errors returned by CancelInflightResponse.ValidateAll() if the designated
// constraints aren't met.
type CancelInflightResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CancelInflightResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CancelInflightResponseMultiError) AllErrors() []error { return m }

// CancelInflightResponseValidationError is the validation error returned by
// CancelInflightResponse.Validate if the designated constraints aren't met.
type CancelInflightResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelInflightResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelInflightResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelInflightResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelInflightResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelInflightResponseValidationError) ErrorName() string {
	return "CancelInflightResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CancelInflightResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelInflightResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelInflightResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelInflightResponseValidationError{}

// Validate checks the field values on State_TokenWindow with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *State_TokenWindow) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on State_TokenWindow with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// State_TokenWindowMultiError, or nil if none found.
func (m *State_TokenWindow) ValidateAll() error {
	return m.validate(true)
}

func (m *State_TokenWindow) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStart()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, State_TokenWindowValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, State_TokenWindowValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return State_TokenWindowValidationError{
				field:  "Start",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEnd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, State_TokenWindowValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, State_TokenWindowValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return State_TokenWindowValidationError{
				field:  "End",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Sum

	// no validation rules for Count

	if len(errors) > 0 {
		return State_TokenWindowMultiError(errors)
	}

	return nil
}

// State_TokenWindowMultiError is an error wrapping multiple validation errors
// returned by State_TokenWindow.ValidateAll() if the designated constraints
// aren't met.
type State_TokenWindowMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m State_TokenWindowMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m State_TokenWindowMultiError) AllErrors() []error { return m }

// State_TokenWindowValidationError is the validation error returned by
// State_TokenWindow.Validate if the designated constraints aren't met.
type State_TokenWindowValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e State_TokenWindowValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e State_TokenWindowValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e State_TokenWindowValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e State_TokenWindowValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e State_TokenWindowValidationError) ErrorName() string {
	return "State_TokenWindowValidationError"
}

// Error satisfies the builtin error interface
func (e State_TokenWindowValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sState_TokenWindow.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = State_TokenWindowValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = State_TokenWindowValidationError{}