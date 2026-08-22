// HTTP 层集成测试(≈ MockMvc):验证完整契约
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestAPICreateGetPay(t *testing.T) {
	// 创建
	req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(`{"item":"coffee","quantity":2}`))
	w := httptest.NewRecorder()
	handleOrders(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("create: got %d, want 201", w.Code)
	}
	var o Order
	if err := json.Unmarshal(w.Body.Bytes(), &o); err != nil || o.ID == 0 {
		t.Fatalf("bad order response: %s", w.Body.String())
	}

	// 查询
	req = httptest.NewRequest(http.MethodGet, "/orders/"+strconv.FormatInt(o.ID, 10), nil)
	w = httptest.NewRecorder()
	handleOrder(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("get: got %d, want 200", w.Code)
	}

	// 支付
	req = httptest.NewRequest(http.MethodPost, "/orders/"+strconv.FormatInt(o.ID, 10)+"/pay", nil)
	w = httptest.NewRecorder()
	handleOrder(w, req)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `"paid":true`) {
		t.Fatalf("pay: got %d %s", w.Code, w.Body.String())
	}
}

func TestAPINotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/orders/999", nil)
	w := httptest.NewRecorder()
	handleOrder(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("got %d, want 404", w.Code)
	}
}
