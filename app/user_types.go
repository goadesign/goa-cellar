//************************************************************************//
// API "cellar": Application User Types
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

// BottlePayload type
type BottlePayload struct {
	Color     *string
	Country   *string
	Name      *string
	Region    *string
	Review    *string
	Sweetness *int
	Varietal  *string
	Vineyard  *string
	Vintage   *int
}

// Validate validates the type instance.
func (ut *BottlePayload) Validate() (err error) {
	if ut.Color != nil {
		if !(*ut.Color == "red" || *ut.Color == "white" || *ut.Color == "rose" || *ut.Color == "yellow" || *ut.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`response.color`, *ut.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
	}
	if ut.Country != nil {
		if len(*ut.Country) < 2 {
			err = goa.InvalidLengthError(`response.country`, *ut.Country, len(*ut.Country), 2, true, err)
		}
	}
	if ut.Name != nil {
		if len(*ut.Name) < 2 {
			err = goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 2, true, err)
		}
	}
	if ut.Review != nil {
		if len(*ut.Review) < 3 {
			err = goa.InvalidLengthError(`response.review`, *ut.Review, len(*ut.Review), 3, true, err)
		}
	}
	if ut.Review != nil {
		if len(*ut.Review) > 300 {
			err = goa.InvalidLengthError(`response.review`, *ut.Review, len(*ut.Review), 300, false, err)
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness < 1 {
			err = goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 1, true, err)
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness > 5 {
			err = goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 5, false, err)
		}
	}
	if ut.Varietal != nil {
		if len(*ut.Varietal) < 4 {
			err = goa.InvalidLengthError(`response.varietal`, *ut.Varietal, len(*ut.Varietal), 4, true, err)
		}
	}
	if ut.Vineyard != nil {
		if len(*ut.Vineyard) < 2 {
			err = goa.InvalidLengthError(`response.vineyard`, *ut.Vineyard, len(*ut.Vineyard), 2, true, err)
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage < 1900 {
			err = goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 1900, true, err)
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage > 2020 {
			err = goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 2020, false, err)
		}
	}
	return
}

// MarshalBottlePayload validates and renders an instance of BottlePayload into a interface{}
func MarshalBottlePayload(source *BottlePayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if err2 := source.Validate(); err2 != nil {
		return nil, goa.ReportError(err, err2)
	}
	tmp66 := map[string]interface{}{
		"color":     source.Color,
		"country":   source.Country,
		"name":      source.Name,
		"region":    source.Region,
		"review":    source.Review,
		"sweetness": source.Sweetness,
		"varietal":  source.Varietal,
		"vineyard":  source.Vineyard,
		"vintage":   source.Vintage,
	}
	target = tmp66
	return
}

// UnmarshalBottlePayload unmarshals and validates a raw interface{} into an instance of BottlePayload
func UnmarshalBottlePayload(source interface{}, inErr error) (target *BottlePayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(BottlePayload)
		if v, ok := val["color"]; ok {
			var tmp67 string
			if val, ok := v.(string); ok {
				tmp67 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Color`, v, "string", err)
			}
			target.Color = &tmp67
		}
		if v, ok := val["country"]; ok {
			var tmp68 string
			if val, ok := v.(string); ok {
				tmp68 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			target.Country = &tmp68
		}
		if v, ok := val["name"]; ok {
			var tmp69 string
			if val, ok := v.(string); ok {
				tmp69 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			target.Name = &tmp69
		}
		if v, ok := val["region"]; ok {
			var tmp70 string
			if val, ok := v.(string); ok {
				tmp70 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Region`, v, "string", err)
			}
			target.Region = &tmp70
		}
		if v, ok := val["review"]; ok {
			var tmp71 string
			if val, ok := v.(string); ok {
				tmp71 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Review`, v, "string", err)
			}
			target.Review = &tmp71
		}
		if v, ok := val["sweetness"]; ok {
			var tmp72 int
			if f, ok := v.(float64); ok {
				tmp72 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Sweetness`, v, "int", err)
			}
			target.Sweetness = &tmp72
		}
		if v, ok := val["varietal"]; ok {
			var tmp73 string
			if val, ok := v.(string); ok {
				tmp73 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Varietal`, v, "string", err)
			}
			target.Varietal = &tmp73
		}
		if v, ok := val["vineyard"]; ok {
			var tmp74 string
			if val, ok := v.(string); ok {
				tmp74 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Vineyard`, v, "string", err)
			}
			target.Vineyard = &tmp74
		}
		if v, ok := val["vintage"]; ok {
			var tmp75 int
			if f, ok := v.(float64); ok {
				tmp75 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Vintage`, v, "int", err)
			}
			target.Vintage = &tmp75
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	if err2 := target.Validate(); err2 != nil {
		err = goa.ReportError(err, err2)
	}
	return
}
