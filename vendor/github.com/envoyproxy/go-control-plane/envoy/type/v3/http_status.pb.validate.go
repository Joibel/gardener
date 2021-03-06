// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/v3/http_status.proto

package envoy_type_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _http_status_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HttpStatus with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HttpStatus) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _HttpStatus_Code_NotInLookup[m.GetCode()]; ok {
		return HttpStatusValidationError{
			field:  "Code",
			reason: "value must not be in list [0]",
		}
	}

	if _, ok := StatusCode_name[int32(m.GetCode())]; !ok {
		return HttpStatusValidationError{
			field:  "Code",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// HttpStatusValidationError is the validation error returned by
// HttpStatus.Validate if the designated constraints aren't met.
type HttpStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpStatusValidationError) ErrorName() string { return "HttpStatusValidationError" }

// Error satisfies the builtin error interface
func (e HttpStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpStatusValidationError{}

var _HttpStatus_Code_NotInLookup = map[StatusCode]struct{}{
	0: {},
}
