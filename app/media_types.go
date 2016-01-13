//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-cellar
// --design=github.com/raphael/goa-cellar/design
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
	Href string
	// ID of account
	ID int
	// Name of account
	Name string
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

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

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
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp36 := map[string]interface{}{
		"created_at": source.CreatedAt,
		"created_by": source.CreatedBy,
		"href":       source.Href,
		"id":         source.ID,
		"name":       source.Name,
	}
	target = tmp36
	return
}

// MarshalAccountLink validates and renders an instance of Account into a interface{}
// using view "link".
func MarshalAccountLink(source *Account, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp37 := map[string]interface{}{
		"href": source.Href,
		"id":   source.ID,
		"name": source.Name,
	}
	target = tmp37
	return
}

// UnmarshalAccount unmarshals and validates a raw interface{} into an instance of Account
func UnmarshalAccount(source interface{}, inErr error) (target *Account, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Account)
		if v, ok := val["created_at"]; ok {
			var tmp38 string
			if val, ok := v.(string); ok {
				tmp38 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp38
		}
		if v, ok := val["created_by"]; ok {
			var tmp39 string
			if val, ok := v.(string); ok {
				tmp39 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedBy`, v, "string", err)
			}
			target.CreatedBy = &tmp39
		}
		if v, ok := val["href"]; ok {
			var tmp40 string
			if val, ok := v.(string); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp40
		} else {
			err = goa.MissingAttributeError(`load`, "href", err)
		}
		if v, ok := val["id"]; ok {
			var tmp41 int
			if f, ok := v.(float64); ok {
				tmp41 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp41
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["name"]; ok {
			var tmp42 string
			if val, ok := v.(string); ok {
				tmp42 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			target.Name = tmp42
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}

// A bottle of wine
// Identifier: application/vnd.bottle+json
type Bottle struct {
	// Account that owns bottle
	Account *Account
	Color   string
	Country *string
	// Date of creation
	CreatedAt *string
	// API href of bottle
	Href string
	// ID of bottle
	ID   int
	Name string
	// Rating of bottle between 1 and 5
	Rating    *int
	Region    *string
	Review    *string
	Sweetness *int
	// Date of last update
	UpdatedAt *string
	Varietal  string
	Vineyard  string
	Vintage   int
}

// Bottle views
type BottleViewEnum string

const (
	// Bottle default view
	BottleDefaultView BottleViewEnum = "default"
	// Bottle full view
	BottleFullView BottleViewEnum = "full"
	// Bottle tiny view
	BottleTinyView BottleViewEnum = "tiny"
)

// LoadBottle loads raw data into an instance of Bottle
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadBottle(raw interface{}) (res *Bottle, err error) {
	res, err = UnmarshalBottle(raw, err)
	return
}

// Dump produces raw data from an instance of Bottle running all the
// validations. See LoadBottle for the definition of raw data.
func (mt *Bottle) Dump(view BottleViewEnum) (res map[string]interface{}, err error) {
	if view == BottleDefaultView {
		res, err = MarshalBottle(mt, err)
	}
	if view == BottleFullView {
		res, err = MarshalBottleFull(mt, err)
	}
	if view == BottleTinyView {
		res, err = MarshalBottleTiny(mt, err)
	}
	return
}

// Validate validates the media type instance.
func (mt *Bottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MissingAttributeError(`response`, "href", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Vineyard == "" {
		err = goa.MissingAttributeError(`response`, "vineyard", err)
	}
	if mt.Varietal == "" {
		err = goa.MissingAttributeError(`response`, "varietal", err)
	}

	if mt.Color == "" {
		err = goa.MissingAttributeError(`response`, "color", err)
	}

	if mt.Account.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.Account.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.account.created_at`, *mt.Account.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if mt.Account.CreatedBy != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Account.CreatedBy); err2 != nil {
			err = goa.InvalidFormatError(`response.account.created_by`, *mt.Account.CreatedBy, goa.FormatEmail, err2, err)
		}
	}
	if !(mt.Color == "red" || mt.Color == "white" || mt.Color == "rose" || mt.Color == "yellow" || mt.Color == "sparkling") {
		err = goa.InvalidEnumValueError(`response.color`, mt.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
	}
	if mt.Country != nil {
		if len(*mt.Country) < 2 {
			err = goa.InvalidLengthError(`response.country`, *mt.Country, len(*mt.Country), 2, true, err)
		}
	}
	if mt.CreatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.CreatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.created_at`, *mt.CreatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if len(mt.Name) < 2 {
		err = goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 2, true, err)
	}
	if mt.Rating != nil {
		if *mt.Rating < 1 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 1, true, err)
		}
	}
	if mt.Rating != nil {
		if *mt.Rating > 5 {
			err = goa.InvalidRangeError(`response.rating`, *mt.Rating, 5, false, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) < 3 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 3, true, err)
		}
	}
	if mt.Review != nil {
		if len(*mt.Review) > 300 {
			err = goa.InvalidLengthError(`response.review`, *mt.Review, len(*mt.Review), 300, false, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness < 1 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 1, true, err)
		}
	}
	if mt.Sweetness != nil {
		if *mt.Sweetness > 5 {
			err = goa.InvalidRangeError(`response.sweetness`, *mt.Sweetness, 5, false, err)
		}
	}
	if mt.UpdatedAt != nil {
		if err2 := goa.ValidateFormat(goa.FormatDateTime, *mt.UpdatedAt); err2 != nil {
			err = goa.InvalidFormatError(`response.updated_at`, *mt.UpdatedAt, goa.FormatDateTime, err2, err)
		}
	}
	if len(mt.Varietal) < 4 {
		err = goa.InvalidLengthError(`response.varietal`, mt.Varietal, len(mt.Varietal), 4, true, err)
	}
	if len(mt.Vineyard) < 2 {
		err = goa.InvalidLengthError(`response.vineyard`, mt.Vineyard, len(mt.Vineyard), 2, true, err)
	}
	if mt.Vintage < 1900 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true, err)
	}
	if mt.Vintage > 2020 {
		err = goa.InvalidRangeError(`response.vintage`, mt.Vintage, 2020, false, err)
	}
	return
}

// MarshalBottle validates and renders an instance of Bottle into a interface{}
// using view "default".
func MarshalBottle(source *Bottle, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp43 := map[string]interface{}{
		"href":     source.Href,
		"id":       source.ID,
		"name":     source.Name,
		"rating":   source.Rating,
		"varietal": source.Varietal,
		"vineyard": source.Vineyard,
		"vintage":  source.Vintage,
	}
	target = tmp43
	if err == nil {
		links := make(map[string]interface{})
		links["account"], err = MarshalAccountLink(source.Account, err)
		target["links"] = links
	}
	return
}

// MarshalBottleFull validates and renders an instance of Bottle into a interface{}
// using view "full".
func MarshalBottleFull(source *Bottle, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp44 := map[string]interface{}{
		"color":      source.Color,
		"country":    source.Country,
		"created_at": source.CreatedAt,
		"href":       source.Href,
		"id":         source.ID,
		"name":       source.Name,
		"rating":     source.Rating,
		"region":     source.Region,
		"review":     source.Review,
		"sweetness":  source.Sweetness,
		"updated_at": source.UpdatedAt,
		"varietal":   source.Varietal,
		"vineyard":   source.Vineyard,
		"vintage":    source.Vintage,
	}
	if source.Account != nil {
		tmp44["account"], err = MarshalAccount(source.Account, err)
	}
	target = tmp44
	if err == nil {
		links := make(map[string]interface{})
		links["account"], err = MarshalAccountLink(source.Account, err)
		target["links"] = links
	}
	return
}

// MarshalBottleTiny validates and renders an instance of Bottle into a interface{}
// using view "tiny".
func MarshalBottleTiny(source *Bottle, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp45 := map[string]interface{}{
		"href":   source.Href,
		"id":     source.ID,
		"name":   source.Name,
		"rating": source.Rating,
	}
	target = tmp45
	if err == nil {
		links := make(map[string]interface{})
		links["account"], err = MarshalAccountLink(source.Account, err)
		target["links"] = links
	}
	return
}

// UnmarshalBottle unmarshals and validates a raw interface{} into an instance of Bottle
func UnmarshalBottle(source interface{}, inErr error) (target *Bottle, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Bottle)
		if v, ok := val["account"]; ok {
			var tmp46 *Account
			tmp46, err = UnmarshalAccount(v, err)
			target.Account = tmp46
		}
		if v, ok := val["color"]; ok {
			var tmp47 string
			if val, ok := v.(string); ok {
				tmp47 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Color`, v, "string", err)
			}
			target.Color = tmp47
		} else {
			err = goa.MissingAttributeError(`load`, "color", err)
		}
		if v, ok := val["country"]; ok {
			var tmp48 string
			if val, ok := v.(string); ok {
				tmp48 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp48
		}
		if v, ok := val["created_at"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.CreatedAt`, v, "string", err)
			}
			target.CreatedAt = &tmp49
		}
		if v, ok := val["href"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Href`, v, "string", err)
			}
			target.Href = tmp50
		} else {
			err = goa.MissingAttributeError(`load`, "href", err)
		}
		if v, ok := val["id"]; ok {
			var tmp51 int
			if f, ok := v.(float64); ok {
				tmp51 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.ID`, v, "int", err)
			}
			target.ID = tmp51
		} else {
			err = goa.MissingAttributeError(`load`, "id", err)
		}
		if v, ok := val["name"]; ok {
			var tmp52 string
			if val, ok := v.(string); ok {
				tmp52 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			target.Name = tmp52
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["rating"]; ok {
			var tmp53 int
			if f, ok := v.(float64); ok {
				tmp53 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Rating`, v, "int", err)
			}
			target.Rating = &tmp53
		}
		if v, ok := val["region"]; ok {
			var tmp54 string
			if val, ok := v.(string); ok {
				tmp54 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Region`, v, "string", err)
			}
			target.Region = &tmp54
		}
		if v, ok := val["review"]; ok {
			var tmp55 string
			if val, ok := v.(string); ok {
				tmp55 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Review`, v, "string", err)
			}
			target.Review = &tmp55
		}
		if v, ok := val["sweetness"]; ok {
			var tmp56 int
			if f, ok := v.(float64); ok {
				tmp56 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Sweetness`, v, "int", err)
			}
			target.Sweetness = &tmp56
		}
		if v, ok := val["updated_at"]; ok {
			var tmp57 string
			if val, ok := v.(string); ok {
				tmp57 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.UpdatedAt`, v, "string", err)
			}
			target.UpdatedAt = &tmp57
		}
		if v, ok := val["varietal"]; ok {
			var tmp58 string
			if val, ok := v.(string); ok {
				tmp58 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Varietal`, v, "string", err)
			}
			target.Varietal = tmp58
		} else {
			err = goa.MissingAttributeError(`load`, "varietal", err)
		}
		if v, ok := val["vineyard"]; ok {
			var tmp59 string
			if val, ok := v.(string); ok {
				tmp59 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Vineyard`, v, "string", err)
			}
			target.Vineyard = tmp59
		} else {
			err = goa.MissingAttributeError(`load`, "vineyard", err)
		}
		if v, ok := val["vintage"]; ok {
			var tmp60 int
			if f, ok := v.(float64); ok {
				tmp60 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Vintage`, v, "int", err)
			}
			target.Vintage = tmp60
		} else {
			err = goa.MissingAttributeError(`load`, "vintage", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}

// BottleCollection media type
// Identifier: application/vnd.bottle+json; type=collection
type BottleCollection []*Bottle

// BottleCollection views
type BottleCollectionViewEnum string

const (
	// BottleCollection default view
	BottleCollectionDefaultView BottleCollectionViewEnum = "default"
	// BottleCollection tiny view
	BottleCollectionTinyView BottleCollectionViewEnum = "tiny"
)

// LoadBottleCollection loads raw data into an instance of BottleCollection
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadBottleCollection(raw interface{}) (res BottleCollection, err error) {
	res, err = UnmarshalBottleCollection(raw, err)
	return
}

// Dump produces raw data from an instance of BottleCollection running all the
// validations. See LoadBottleCollection for the definition of raw data.
func (mt BottleCollection) Dump(view BottleCollectionViewEnum) (res []map[string]interface{}, err error) {
	if view == BottleCollectionDefaultView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp61 := range mt {
			var tmp62 map[string]interface{}
			tmp62, err = MarshalBottle(tmp61, err)
			res[i] = tmp62
		}
	}
	if view == BottleCollectionTinyView {
		res = make([]map[string]interface{}, len(mt))
		for i, tmp63 := range mt {
			var tmp64 map[string]interface{}
			tmp64, err = MarshalBottleTiny(tmp63, err)
			res[i] = tmp64
		}
	}
	return
}

// Validate validates the media type instance.
func (mt BottleCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Account.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.Account.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].account.created_at`, *e.Account.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if e.Account.CreatedBy != nil {
			if err2 := goa.ValidateFormat(goa.FormatEmail, *e.Account.CreatedBy); err2 != nil {
				err = goa.InvalidFormatError(`response[*].account.created_by`, *e.Account.CreatedBy, goa.FormatEmail, err2, err)
			}
		}
		if !(e.Color == "red" || e.Color == "white" || e.Color == "rose" || e.Color == "yellow" || e.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`response[*].color`, e.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
		if e.Country != nil {
			if len(*e.Country) < 2 {
				err = goa.InvalidLengthError(`response[*].country`, *e.Country, len(*e.Country), 2, true, err)
			}
		}
		if e.CreatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.CreatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].created_at`, *e.CreatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if len(e.Name) < 2 {
			err = goa.InvalidLengthError(`response[*].name`, e.Name, len(e.Name), 2, true, err)
		}
		if e.Rating != nil {
			if *e.Rating < 1 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 1, true, err)
			}
		}
		if e.Rating != nil {
			if *e.Rating > 5 {
				err = goa.InvalidRangeError(`response[*].rating`, *e.Rating, 5, false, err)
			}
		}
		if e.Review != nil {
			if len(*e.Review) < 3 {
				err = goa.InvalidLengthError(`response[*].review`, *e.Review, len(*e.Review), 3, true, err)
			}
		}
		if e.Review != nil {
			if len(*e.Review) > 300 {
				err = goa.InvalidLengthError(`response[*].review`, *e.Review, len(*e.Review), 300, false, err)
			}
		}
		if e.Sweetness != nil {
			if *e.Sweetness < 1 {
				err = goa.InvalidRangeError(`response[*].sweetness`, *e.Sweetness, 1, true, err)
			}
		}
		if e.Sweetness != nil {
			if *e.Sweetness > 5 {
				err = goa.InvalidRangeError(`response[*].sweetness`, *e.Sweetness, 5, false, err)
			}
		}
		if e.UpdatedAt != nil {
			if err2 := goa.ValidateFormat(goa.FormatDateTime, *e.UpdatedAt); err2 != nil {
				err = goa.InvalidFormatError(`response[*].updated_at`, *e.UpdatedAt, goa.FormatDateTime, err2, err)
			}
		}
		if len(e.Varietal) < 4 {
			err = goa.InvalidLengthError(`response[*].varietal`, e.Varietal, len(e.Varietal), 4, true, err)
		}
		if len(e.Vineyard) < 2 {
			err = goa.InvalidLengthError(`response[*].vineyard`, e.Vineyard, len(e.Vineyard), 2, true, err)
		}
		if e.Vintage < 1900 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 1900, true, err)
		}
		if e.Vintage > 2020 {
			err = goa.InvalidRangeError(`response[*].vintage`, e.Vintage, 2020, false, err)
		}
	}
	return
}

// MarshalBottleCollection validates and renders an instance of BottleCollection into a interface{}
// using view "default".
func MarshalBottleCollection(source BottleCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalBottle(res, err)
	}
	return
}

// MarshalBottleCollectionTiny validates and renders an instance of BottleCollection into a interface{}
// using view "tiny".
func MarshalBottleCollectionTiny(source BottleCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalBottleTiny(res, err)
	}
	return
}

// UnmarshalBottleCollection unmarshals and validates a raw interface{} into an instance of BottleCollection
func UnmarshalBottleCollection(source interface{}, inErr error) (target BottleCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Bottle, len(val))
		for tmp65, v := range val {
			target[tmp65], err = UnmarshalBottle(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}
