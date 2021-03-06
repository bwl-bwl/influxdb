package tenant

import (
	"fmt"

	"github.com/influxdata/influxdb"
)

var (
	// ErrInvalidURMID is used when the service was provided
	// an invalid ID format.
	ErrInvalidURMID = &influxdb.Error{
		Code: influxdb.EInvalid,
		Msg:  "provided user resource mapping ID has invalid format",
	}

	// ErrURMNotFound is used when the user resource mapping is not found.
	ErrURMNotFound = &influxdb.Error{
		Msg:  "user to resource mapping not found",
		Code: influxdb.ENotFound,
	}
)

// UnavailableURMServiceError is used if we aren't able to interact with the
// store, it means the store is not available at the moment (e.g. network).
func UnavailableURMServiceError(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EInternal,
		Msg:  fmt.Sprintf("Unable to connect to resource mapping service. Please try again; Err: %v", err),
		Op:   "kv/userResourceMapping",
	}
}

// CorruptURMError is used when the config cannot be unmarshalled from the
// bytes stored in the kv.
func CorruptURMError(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EInternal,
		Msg:  fmt.Sprintf("Unknown internal user resource mapping data error; Err: %v", err),
		Op:   "kv/userResourceMapping",
	}
}

// ErrUnprocessableMapping is used when a user resource mapping  is not able to be converted to JSON.
func ErrUnprocessableMapping(err error) *influxdb.Error {
	return &influxdb.Error{
		Code: influxdb.EUnprocessableEntity,
		Msg:  fmt.Sprintf("unable to convert mapping of user to resource into JSON; Err %v", err),
	}
}

// NonUniqueMappingError is an internal error when a user already has
// been mapped to a resource
func NonUniqueMappingError(userID influxdb.ID) error {
	return &influxdb.Error{
		Code: influxdb.EInternal,
		Msg:  fmt.Sprintf("Unexpected error when assigning user to a resource: mapping for user %s already exists", userID.String()),
	}
}
