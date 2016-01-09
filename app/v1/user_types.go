//************************************************************************//
// API "cellar" version 1.0: Application User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-cellar
// --design=github.com/raphael/testd/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package v1

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
		if len(*ut.Review) < 10 {
			err = goa.InvalidLengthError(`response.review`, *ut.Review, len(*ut.Review), 10, true, err)
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
	if source.Color != nil {
		if !(*source.Color == "red" || *source.Color == "white" || *source.Color == "rose" || *source.Color == "yellow" || *source.Color == "sparkling") {
			err = goa.InvalidEnumValueError(`.color`, *source.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
		}
	}
	if source.Country != nil {
		if len(*source.Country) < 2 {
			err = goa.InvalidLengthError(`.country`, *source.Country, len(*source.Country), 2, true, err)
		}
	}
	if source.Name != nil {
		if len(*source.Name) < 2 {
			err = goa.InvalidLengthError(`.name`, *source.Name, len(*source.Name), 2, true, err)
		}
	}
	if source.Review != nil {
		if len(*source.Review) < 10 {
			err = goa.InvalidLengthError(`.review`, *source.Review, len(*source.Review), 10, true, err)
		}
	}
	if source.Review != nil {
		if len(*source.Review) > 300 {
			err = goa.InvalidLengthError(`.review`, *source.Review, len(*source.Review), 300, false, err)
		}
	}
	if source.Sweetness != nil {
		if *source.Sweetness < 1 {
			err = goa.InvalidRangeError(`.sweetness`, *source.Sweetness, 1, true, err)
		}
	}
	if source.Sweetness != nil {
		if *source.Sweetness > 5 {
			err = goa.InvalidRangeError(`.sweetness`, *source.Sweetness, 5, false, err)
		}
	}
	if source.Varietal != nil {
		if len(*source.Varietal) < 4 {
			err = goa.InvalidLengthError(`.varietal`, *source.Varietal, len(*source.Varietal), 4, true, err)
		}
	}
	if source.Vineyard != nil {
		if len(*source.Vineyard) < 2 {
			err = goa.InvalidLengthError(`.vineyard`, *source.Vineyard, len(*source.Vineyard), 2, true, err)
		}
	}
	if source.Vintage != nil {
		if *source.Vintage < 1900 {
			err = goa.InvalidRangeError(`.vintage`, *source.Vintage, 1900, true, err)
		}
	}
	if source.Vintage != nil {
		if *source.Vintage > 2020 {
			err = goa.InvalidRangeError(`.vintage`, *source.Vintage, 2020, false, err)
		}
	}
	tmp79 := map[string]interface{}{
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
	target = tmp79
	return
}

// UnmarshalBottlePayload unmarshals and validates a raw interface{} into an instance of BottlePayload
func UnmarshalBottlePayload(source interface{}, inErr error) (target *BottlePayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(BottlePayload)
		if v, ok := val["color"]; ok {
			var tmp80 string
			if val, ok := v.(string); ok {
				tmp80 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Color`, v, "string", err)
			}
			if err == nil {
				if !(tmp80 == "red" || tmp80 == "white" || tmp80 == "rose" || tmp80 == "yellow" || tmp80 == "sparkling") {
					err = goa.InvalidEnumValueError(`load.Color`, tmp80, []interface{}{"red", "white", "rose", "yellow", "sparkling"}, err)
				}
			}
			target.Color = &tmp80
		}
		if v, ok := val["country"]; ok {
			var tmp81 string
			if val, ok := v.(string); ok {
				tmp81 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Country`, v, "string", err)
			}
			if err == nil {
				if len(tmp81) < 2 {
					err = goa.InvalidLengthError(`load.Country`, tmp81, len(tmp81), 2, true, err)
				}
			}
			target.Country = &tmp81
		}
		if v, ok := val["name"]; ok {
			var tmp82 string
			if val, ok := v.(string); ok {
				tmp82 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if len(tmp82) < 2 {
					err = goa.InvalidLengthError(`load.Name`, tmp82, len(tmp82), 2, true, err)
				}
			}
			target.Name = &tmp82
		}
		if v, ok := val["region"]; ok {
			var tmp83 string
			if val, ok := v.(string); ok {
				tmp83 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Region`, v, "string", err)
			}
			target.Region = &tmp83
		}
		if v, ok := val["review"]; ok {
			var tmp84 string
			if val, ok := v.(string); ok {
				tmp84 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Review`, v, "string", err)
			}
			if err == nil {
				if len(tmp84) < 10 {
					err = goa.InvalidLengthError(`load.Review`, tmp84, len(tmp84), 10, true, err)
				}
				if len(tmp84) > 300 {
					err = goa.InvalidLengthError(`load.Review`, tmp84, len(tmp84), 300, false, err)
				}
			}
			target.Review = &tmp84
		}
		if v, ok := val["sweetness"]; ok {
			var tmp85 int
			if f, ok := v.(float64); ok {
				tmp85 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Sweetness`, v, "int", err)
			}
			if err == nil {
				if tmp85 < 1 {
					err = goa.InvalidRangeError(`load.Sweetness`, tmp85, 1, true, err)
				}
				if tmp85 > 5 {
					err = goa.InvalidRangeError(`load.Sweetness`, tmp85, 5, false, err)
				}
			}
			target.Sweetness = &tmp85
		}
		if v, ok := val["varietal"]; ok {
			var tmp86 string
			if val, ok := v.(string); ok {
				tmp86 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Varietal`, v, "string", err)
			}
			if err == nil {
				if len(tmp86) < 4 {
					err = goa.InvalidLengthError(`load.Varietal`, tmp86, len(tmp86), 4, true, err)
				}
			}
			target.Varietal = &tmp86
		}
		if v, ok := val["vineyard"]; ok {
			var tmp87 string
			if val, ok := v.(string); ok {
				tmp87 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Vineyard`, v, "string", err)
			}
			if err == nil {
				if len(tmp87) < 2 {
					err = goa.InvalidLengthError(`load.Vineyard`, tmp87, len(tmp87), 2, true, err)
				}
			}
			target.Vineyard = &tmp87
		}
		if v, ok := val["vintage"]; ok {
			var tmp88 int
			if f, ok := v.(float64); ok {
				tmp88 = int(f)
			} else {
				err = goa.InvalidAttributeTypeError(`load.Vintage`, v, "int", err)
			}
			if err == nil {
				if tmp88 < 1900 {
					err = goa.InvalidRangeError(`load.Vintage`, tmp88, 1900, true, err)
				}
				if tmp88 > 2020 {
					err = goa.InvalidRangeError(`load.Vintage`, tmp88, 2020, false, err)
				}
			}
			target.Vintage = &tmp88
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
