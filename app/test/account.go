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

// CreateAccountCreated test setup
func CreateAccountCreated(t *testing.T, ctrl app.AccountController, payload *app.CreateAccountPayload) {
	CreateAccountCreatedCtx(t, context.Background(), ctrl, payload)
}

// CreateAccountCreatedCtx test setup
func CreateAccountCreatedCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, payload *app.CreateAccountPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 201 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
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

}

// DeleteAccountNoContent test setup
func DeleteAccountNoContent(t *testing.T, ctrl app.AccountController, accountID int) {
	DeleteAccountNoContentCtx(t, context.Background(), ctrl, accountID)
}

// DeleteAccountNoContentCtx test setup
func DeleteAccountNoContentCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) {
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

}

// DeleteAccountNotFound test setup
func DeleteAccountNotFound(t *testing.T, ctrl app.AccountController, accountID int) {
	DeleteAccountNotFoundCtx(t, context.Background(), ctrl, accountID)
}

// DeleteAccountNotFoundCtx test setup
func DeleteAccountNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) {
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

}

// ListAccountNotFound test setup
func ListAccountNotFound(t *testing.T, ctrl app.AccountController) {
	ListAccountNotFoundCtx(t, context.Background(), ctrl)
}

// ListAccountNotFoundCtx test setup
func ListAccountNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AccountController) {
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

}

// ListAccountOK test setup
func ListAccountOK(t *testing.T, ctrl app.AccountController) *app.AccountCollection {
	return ListAccountOKCtx(t, context.Background(), ctrl)
}

// ListAccountOKCtx test setup
func ListAccountOKCtx(t *testing.T, ctx context.Context, ctrl app.AccountController) *app.AccountCollection {
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

	a, ok := resp.(*app.AccountCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.AccountCollection", resp)
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

// ListAccountOKLink test setup
func ListAccountOKLink(t *testing.T, ctrl app.AccountController) *app.AccountLinkCollection {
	return ListAccountOKLinkCtx(t, context.Background(), ctrl)
}

// ListAccountOKLinkCtx test setup
func ListAccountOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.AccountController) *app.AccountLinkCollection {
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

	a, ok := resp.(*app.AccountLinkCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.AccountLinkCollection", resp)
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

// ListAccountOKTiny test setup
func ListAccountOKTiny(t *testing.T, ctrl app.AccountController) *app.AccountTinyCollection {
	return ListAccountOKTinyCtx(t, context.Background(), ctrl)
}

// ListAccountOKTinyCtx test setup
func ListAccountOKTinyCtx(t *testing.T, ctx context.Context, ctrl app.AccountController) *app.AccountTinyCollection {
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

	a, ok := resp.(*app.AccountTinyCollection)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.AccountTinyCollection", resp)
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

// ShowAccountNotFound test setup
func ShowAccountNotFound(t *testing.T, ctrl app.AccountController, accountID int) {
	ShowAccountNotFoundCtx(t, context.Background(), ctrl, accountID)
}

// ShowAccountNotFoundCtx test setup
func ShowAccountNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) {
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

}

// ShowAccountOK test setup
func ShowAccountOK(t *testing.T, ctrl app.AccountController, accountID int) *app.Account {
	return ShowAccountOKCtx(t, context.Background(), ctrl, accountID)
}

// ShowAccountOKCtx test setup
func ShowAccountOKCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) *app.Account {
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

	a, ok := resp.(*app.Account)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Account", resp)
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

// ShowAccountOKLink test setup
func ShowAccountOKLink(t *testing.T, ctrl app.AccountController, accountID int) *app.AccountLink {
	return ShowAccountOKLinkCtx(t, context.Background(), ctrl, accountID)
}

// ShowAccountOKLinkCtx test setup
func ShowAccountOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) *app.AccountLink {
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

	a, ok := resp.(*app.AccountLink)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.AccountLink", resp)
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

// ShowAccountOKTiny test setup
func ShowAccountOKTiny(t *testing.T, ctrl app.AccountController, accountID int) *app.AccountTiny {
	return ShowAccountOKTinyCtx(t, context.Background(), ctrl, accountID)
}

// ShowAccountOKTinyCtx test setup
func ShowAccountOKTinyCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int) *app.AccountTiny {
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

	a, ok := resp.(*app.AccountTiny)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.AccountTiny", resp)
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

// UpdateAccountNoContent test setup
func UpdateAccountNoContent(t *testing.T, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) {
	UpdateAccountNoContentCtx(t, context.Background(), ctrl, accountID, payload)
}

// UpdateAccountNoContentCtx test setup
func UpdateAccountNoContentCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 204 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
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

}

// UpdateAccountNotFound test setup
func UpdateAccountNotFound(t *testing.T, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) {
	UpdateAccountNotFoundCtx(t, context.Background(), ctrl, accountID, payload)
}

// UpdateAccountNotFoundCtx test setup
func UpdateAccountNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.AccountController, accountID int, payload *app.UpdateAccountPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 404 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
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

}
