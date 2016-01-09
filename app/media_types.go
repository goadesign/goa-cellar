//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-cellar
// --design=github.com/raphael/testd/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/raphael/goa"

// A tenant account
// Identifier: application/vnd.account+json
type Account struct {
	// Date of creation
	CreatedAt *string
	// Email of account owner
	CreatedBy *string
	// API href of account
	Href *string
	// ID of account
	ID *int
	// Name of account
	Name *string
}

// Account views
type AccountViewEnum string

const (
	// Account default view
	AccountDefaultView AccountViewEnum = "default"
	// Account link view
	AccountLinkView AccountViewEnum = "link"
)

// LoadAccount loads raw data into an instance of Account
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAccount(raw interface{}) (res *Account, err error) {
	res, err = UnmarshalAccount(raw, err)
	return
}

// Dump produces raw data from an instance of Account running all the
// validations. See LoadAccount for the definition of raw data.
func (mt *Account) Dump(view AccountViewEnum) (res map[string]interface{}, err error) {
	if view == AccountDefaultView {
		res, err = MarshalAccount(mt, err)
	}
	if view == AccountLinkView {
		res, err = MarshalAccountLink(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Account) Validate() (err error) {
	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.created_by`, *mt.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	return
}

// MarshalAccount validates and renders an instance of Account into a interface{}
// using view "default".
func MarshalAccount(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *source.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`.created_at`, *source.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if source.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *source.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`.created_by`, *source.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	tmp9 := map[string]interface{}{
		"created_at": source.CreatedAt,
		"created_by": source.CreatedBy,
		"href":       source.Href,
		"id":         source.ID,
		"name":       source.Name,
	}
	target = tmp9
	return
}

// MarshalAccountLink validates and renders an instance of Account into a interface{}
// using view "link".
func MarshalAccountLink(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp10 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp10
	return
}

// UnmarshalAccount unmarshals and validates a raw interface{} into an instance of Account
func UnmarshalAccount(source interface{}, inErr error) (target *Account, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Account)
		if v, ok := val["created_at"]; ok {
			var tmp11 string
			if val, ok := v.(string); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedAt`, v, "string", err)
			}
			if err == nil {
				if err2 := goa.ValidateFormat(goa.FormatDateTime, tmp11); err2 != nil {
					err = goa.InvalidFormatError(`load.CreatedAt`, tmp11, goa.FormatDateTime, err2, err)
				}
			}
			target.CreatedAt = &tmp11
		}
		if v, ok := val["created_by"]; ok {
			var tmp12 string
			if val, ok := v.(string); ok {
				tmp12 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedBy`, v, "string", err)
			}
			if err == nil {
				if err2 := goa.ValidateFormat(goa.FormatEmail, tmp12); err2 != nil {
					err = goa.InvalidFormatError(`load.CreatedBy`, tmp12, goa.FormatEmail, err2, err)
				}
			}
			target.CreatedBy = &tmp12
		}
		if v, ok := val["href"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = &tmp13
		}
		if v, ok := val["id"]; ok {
			var tmp14 int
			if f, ok := v.(float64); ok {
				tmp14 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = &tmp14
		}
		if v, ok := val["name"]; ok {
			var tmp15 string
			if val, ok := v.(string); ok {
				tmp15 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			target.Name = &tmp15
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
