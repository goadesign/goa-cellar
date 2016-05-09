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
	"testing"
)

// CreateBottleCreated test setup
func CreateBottleCreated(t *testing.T, ctrl app.BottleController, accountID string, payload *app.CreateBottlePayload) {
	CreateBottleCreatedCtx(t, context.Background(), ctrl, accountID, payload)
}

// CreateBottleCreatedCtx test setup
func CreateBottleCreatedCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, payload *app.CreateBottlePayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
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
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	createCtx, err := app.NewCreateBottleContext(goaCtx, service)
	createCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

}

// DeleteBottleNoContent test setup
func DeleteBottleNoContent(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) {
	DeleteBottleNoContentCtx(t, context.Background(), ctrl, accountID, bottleID)
}

// DeleteBottleNoContentCtx test setup
func DeleteBottleNoContentCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// DeleteBottleNotFound test setup
func DeleteBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) {
	DeleteBottleNotFoundCtx(t, context.Background(), ctrl, accountID, bottleID)
}

// DeleteBottleNotFoundCtx test setup
func DeleteBottleNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ListBottleNotFound test setup
func ListBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string) {
	ListBottleNotFoundCtx(t, context.Background(), ctrl, accountID)
}

// ListBottleNotFoundCtx test setup
func ListBottleNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	listCtx, err := app.NewListBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ListBottleOK test setup
func ListBottleOK(t *testing.T, ctrl app.BottleController, accountID string) *app.BottleCollection {
	return ListBottleOKCtx(t, context.Background(), ctrl, accountID)
}

// ListBottleOKCtx test setup
func ListBottleOKCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string) *app.BottleCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	listCtx, err := app.NewListBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.BottleCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.BottleCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ListBottleOKTiny test setup
func ListBottleOKTiny(t *testing.T, ctrl app.BottleController, accountID string) *app.BottleTinyCollection {
	return ListBottleOKTinyCtx(t, context.Background(), ctrl, accountID)
}

// ListBottleOKTinyCtx test setup
func ListBottleOKTinyCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string) *app.BottleTinyCollection {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	listCtx, err := app.NewListBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.BottleTinyCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.BottleTinyCollection", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// RateBottleNoContent test setup
func RateBottleNoContent(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) {
	RateBottleNoContentCtx(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// RateBottleNoContentCtx test setup
func RateBottleNoContentCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
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
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	rateCtx, err := app.NewRateBottleContext(goaCtx, service)
	rateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Rate(rateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// RateBottleNotFound test setup
func RateBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) {
	RateBottleNotFoundCtx(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// RateBottleNotFoundCtx test setup
func RateBottleNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.RateBottlePayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
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
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	rateCtx, err := app.NewRateBottleContext(goaCtx, service)
	rateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Rate(rateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ShowBottleNotFound test setup
func ShowBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) {
	ShowBottleNotFoundCtx(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleNotFoundCtx test setup
func ShowBottleNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ShowBottleOK test setup
func ShowBottleOK(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) *app.Bottle {
	return ShowBottleOKCtx(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleOKCtx test setup
func ShowBottleOKCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) *app.Bottle {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Bottle)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Bottle", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowBottleOKFull test setup
func ShowBottleOKFull(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) *app.BottleFull {
	return ShowBottleOKFullCtx(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleOKFullCtx test setup
func ShowBottleOKFullCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) *app.BottleFull {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.BottleFull)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.BottleFull", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowBottleOKTiny test setup
func ShowBottleOKTiny(t *testing.T, ctrl app.BottleController, accountID string, bottleID int) *app.BottleTiny {
	return ShowBottleOKTinyCtx(t, context.Background(), ctrl, accountID, bottleID)
}

// ShowBottleOKTinyCtx test setup
func ShowBottleOKTinyCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int) *app.BottleTiny {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v/bottles/%v", accountID, bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.BottleTiny)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.BottleTiny", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// UpdateBottleNoContent test setup
func UpdateBottleNoContent(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) {
	UpdateBottleNoContentCtx(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// UpdateBottleNoContentCtx test setup
func UpdateBottleNoContentCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
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
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateBottleContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// UpdateBottleNotFound test setup
func UpdateBottleNotFound(t *testing.T, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) {
	UpdateBottleNotFoundCtx(t, context.Background(), ctrl, accountID, bottleID, payload)
}

// UpdateBottleNotFoundCtx test setup
func UpdateBottleNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, accountID string, bottleID int, payload *app.UpdateBottlePayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
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
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateBottleContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}
