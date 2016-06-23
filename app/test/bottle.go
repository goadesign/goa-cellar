package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa-cellar/app"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// CreateBottleCreated Create runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func CreateBottleCreated(t *testing.T, ctrl app.BottleController, accountID string, payload *app.CreateBottlePayload) http.ResponseWriter {
	return CreateBottleCreatedWithContext(t, context.Background(), ctrl, accountID, payload)
}

// CreateBottleCreatedWithContext Create runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func CreateBottleCreatedWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, payload *app.CreateBottlePayload) http.ResponseWriter {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 201 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	createCtx, err := app.NewCreateBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	createCtx.Payload = payload

	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// CreateBottleNotFound Create runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func CreateBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, payload *app.CreateBottlePayload) http.ResponseWriter {
	return CreateBottleNotFoundWithContext(t, context.Background(), ctrl, accountID, payload)
}

// CreateBottleNotFoundWithContext Create runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func CreateBottleNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, payload *app.CreateBottlePayload) http.ResponseWriter {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 404 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	createCtx, err := app.NewCreateBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	createCtx.Payload = payload

	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// DeleteBottleNoContent Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteBottleNoContent(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) http.ResponseWriter {
	return DeleteBottleNoContentWithContext(t, context.Background(), ctrl, accountID, bottleID)
}

// DeleteBottleNoContentWithContext Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteBottleNoContentWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// DeleteBottleNotFound Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) http.ResponseWriter {
	return DeleteBottleNotFoundWithContext(t, context.Background(), ctrl, accountID, bottleID)
}

// DeleteBottleNotFoundWithContext Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteBottleNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// ListBottleNotFound List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ListBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string) http.ResponseWriter {
	return ListBottleNotFoundWithContext(t, context.Background(), ctrl, accountID)
}

// ListBottleNotFoundWithContext List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ListBottleNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	listCtx, err := app.NewListBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// ListBottleOK List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListBottleOK(t *testing.T, ctrl app.BottleController, accountID string) (http.ResponseWriter, *app.BottleCollection) {
	return ListBottleOKWithContext(t, context.Background(), ctrl, accountID)
}

// ListBottleOKWithContext List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListBottleOKWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string) (http.ResponseWriter, *app.BottleCollection) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	listCtx, err := app.NewListBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	var mt *app.BottleCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.BottleCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.BottleCollection", resp)
		}
	}

	return rw, mt
}

// ListBottleOKTiny List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListBottleOKTiny(t *testing.T, ctrl app.BottleController, accountID string) (http.ResponseWriter, *app.BottleTinyCollection) {
	return ListBottleOKTinyWithContext(t, context.Background(), ctrl, accountID)
}

// ListBottleOKTinyWithContext List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListBottleOKTinyWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string) (http.ResponseWriter, *app.BottleTinyCollection) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	listCtx, err := app.NewListBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	var mt *app.BottleTinyCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.BottleTinyCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.BottleTinyCollection", resp)
		}
	}

	return rw, mt
}

// RateBottleNoContent Rate runs the method Rate of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func RateBottleNoContent(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) http.ResponseWriter {
	return RateBottleNoContentWithContext(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// RateBottleNoContentWithContext Rate runs the method Rate of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func RateBottleNoContentWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) http.ResponseWriter {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 204 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/cellar/accounts/%v/bottles/%v/actions/rate", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	rateCtx, err := app.NewRateBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	rateCtx.Payload = payload

	err = ctrl.Rate(rateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// RateBottleNotFound Rate runs the method Rate of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func RateBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) http.ResponseWriter {
	return RateBottleNotFoundWithContext(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// RateBottleNotFoundWithContext Rate runs the method Rate of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func RateBottleNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) http.ResponseWriter {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 404 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/cellar/accounts/%v/bottles/%v/actions/rate", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	rateCtx, err := app.NewRateBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	rateCtx.Payload = payload

	err = ctrl.Rate(rateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// ShowBottleNotFound Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ShowBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) http.ResponseWriter {
	return ShowBottleNotFoundWithContext(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleNotFoundWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ShowBottleNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// ShowBottleOK Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowBottleOK(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) (http.ResponseWriter, *app.Bottle) {
	return ShowBottleOKWithContext(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleOKWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowBottleOKWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) (http.ResponseWriter, *app.Bottle) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	var mt *app.Bottle
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Bottle)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.Bottle", resp)
		}
	}

	return rw, mt
}

// ShowBottleOKFull Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowBottleOKFull(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) (http.ResponseWriter, *app.BottleFull) {
	return ShowBottleOKFullWithContext(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleOKFullWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowBottleOKFullWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) (http.ResponseWriter, *app.BottleFull) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	var mt *app.BottleFull
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.BottleFull)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.BottleFull", resp)
		}
	}

	return rw, mt
}

// ShowBottleOKTiny Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowBottleOKTiny(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) (http.ResponseWriter, *app.BottleTiny) {
	return ShowBottleOKTinyWithContext(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleOKTinyWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowBottleOKTinyWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) (http.ResponseWriter, *app.BottleTiny) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	var mt *app.BottleTiny
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.BottleTiny)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.BottleTiny", resp)
		}
	}

	return rw, mt
}

// UpdateBottleNoContent Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateBottleNoContent(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) http.ResponseWriter {
	return UpdateBottleNoContentWithContext(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// UpdateBottleNoContentWithContext Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateBottleNoContentWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) http.ResponseWriter {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 204 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}

// UpdateBottleNotFound Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) http.ResponseWriter {
	return UpdateBottleNotFoundWithContext(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// UpdateBottleNotFoundWithContext Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateBottleNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) http.ResponseWriter {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 404 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}
	prms["bottleID"] = []string{fmt.Sprintf("%v", bottleID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	return rw
}
