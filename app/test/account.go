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

// CreateAccountCreated Create runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func CreateAccountCreated(t *testing.T, ctrl app.AccountController, payload *app.CreateAccountPayload) http.ResponseWriter {
	return CreateAccountCreatedWithContext(t, context.Background(), ctrl, payload)
}

// CreateAccountCreatedWithContext Create runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func CreateAccountCreatedWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, payload *app.CreateAccountPayload) http.ResponseWriter {
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
	req, err := http.NewRequest("POST", fmt.Sprintf("/cellar/accounts"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	createCtx, err := app.NewCreateAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	createCtx.Payload = payload

	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

	return rw
}

// DeleteAccountNoContent Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteAccountNoContent(t *testing.T, ctrl app.AccountController, accountID int) http.ResponseWriter {
	return DeleteAccountNoContentWithContext(t, context.Background(), ctrl, accountID)
}

// DeleteAccountNoContentWithContext Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteAccountNoContentWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteAccountContext(goaCtx, service)
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

	return rw
}

// DeleteAccountNotFound Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteAccountNotFound(t *testing.T, ctrl app.AccountController, accountID int) http.ResponseWriter {
	return DeleteAccountNotFoundWithContext(t, context.Background(), ctrl, accountID)
}

// DeleteAccountNotFoundWithContext Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func DeleteAccountNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteAccountContext(goaCtx, service)
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

	return rw
}

// ListAccountNotFound List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ListAccountNotFound(t *testing.T, ctrl app.AccountController) http.ResponseWriter {
	return ListAccountNotFoundWithContext(t, context.Background(), ctrl)
}

// ListAccountNotFoundWithContext List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ListAccountNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	listCtx, err := app.NewListAccountContext(goaCtx, service)
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

	return rw
}

// ListAccountOK List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListAccountOK(t *testing.T, ctrl app.AccountController) (http.ResponseWriter, *app.AccountCollection) {
	return ListAccountOKWithContext(t, context.Background(), ctrl)
}

// ListAccountOKWithContext List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListAccountOKWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController) (http.ResponseWriter, *app.AccountCollection) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	listCtx, err := app.NewListAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AccountCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AccountCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.AccountCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	return rw, mt
}

// ListAccountOKLink List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListAccountOKLink(t *testing.T, ctrl app.AccountController) (http.ResponseWriter, *app.AccountLinkCollection) {
	return ListAccountOKLinkWithContext(t, context.Background(), ctrl)
}

// ListAccountOKLinkWithContext List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListAccountOKLinkWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController) (http.ResponseWriter, *app.AccountLinkCollection) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	listCtx, err := app.NewListAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AccountLinkCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AccountLinkCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.AccountLinkCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	return rw, mt
}

// ListAccountOKTiny List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListAccountOKTiny(t *testing.T, ctrl app.AccountController) (http.ResponseWriter, *app.AccountTinyCollection) {
	return ListAccountOKTinyWithContext(t, context.Background(), ctrl)
}

// ListAccountOKTinyWithContext List runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ListAccountOKTinyWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController) (http.ResponseWriter, *app.AccountTinyCollection) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	listCtx, err := app.NewListAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.List(listCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AccountTinyCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AccountTinyCollection)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.AccountTinyCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	return rw, mt
}

// ShowAccountNotFound Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ShowAccountNotFound(t *testing.T, ctrl app.AccountController, accountID int) http.ResponseWriter {
	return ShowAccountNotFoundWithContext(t, context.Background(), ctrl, accountID)
}

// ShowAccountNotFoundWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
func ShowAccountNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) http.ResponseWriter {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountContext(goaCtx, service)
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

	return rw
}

// ShowAccountOK Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowAccountOK(t *testing.T, ctrl app.AccountController, accountID int) (http.ResponseWriter, *app.Account) {
	return ShowAccountOKWithContext(t, context.Background(), ctrl, accountID)
}

// ShowAccountOKWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowAccountOKWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) (http.ResponseWriter, *app.Account) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Account
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Account)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.Account", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	return rw, mt
}

// ShowAccountOKLink Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowAccountOKLink(t *testing.T, ctrl app.AccountController, accountID int) (http.ResponseWriter, *app.AccountLink) {
	return ShowAccountOKLinkWithContext(t, context.Background(), ctrl, accountID)
}

// ShowAccountOKLinkWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowAccountOKLinkWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) (http.ResponseWriter, *app.AccountLink) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AccountLink
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AccountLink)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.AccountLink", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	return rw, mt
}

// ShowAccountOKTiny Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowAccountOKTiny(t *testing.T, ctrl app.AccountController, accountID int) (http.ResponseWriter, *app.AccountTiny) {
	return ShowAccountOKTinyWithContext(t, context.Background(), ctrl, accountID)
}

// ShowAccountOKTinyWithContext Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
func ShowAccountOKTinyWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) (http.ResponseWriter, *app.AccountTiny) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AccountTiny
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AccountTiny)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.AccountTiny", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	return rw, mt
}

// UpdateAccountNoContent Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateAccountNoContent(t *testing.T, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) http.ResponseWriter {
	return UpdateAccountNoContentWithContext(t, context.Background(), ctrl, accountID, payload)
}

// UpdateAccountNoContentWithContext Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateAccountNoContentWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) http.ResponseWriter {
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
	req, err := http.NewRequest("PUT", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

	return rw
}

// UpdateAccountNotFound Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateAccountNotFound(t *testing.T, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) http.ResponseWriter {
	return UpdateAccountNotFoundWithContext(t, context.Background(), ctrl, accountID, payload)
}

// UpdateAccountNotFoundWithContext Update runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
func UpdateAccountNotFoundWithContext(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) http.ResponseWriter {
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
	req, err := http.NewRequest("PUT", fmt.Sprintf("/cellar/accounts/%v", accountID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountID"] = []string{fmt.Sprintf("%v", accountID)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateAccountContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	return rw
}
