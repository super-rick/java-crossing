// 订单服务入口:标准库 net/http(Go 标准库即生产级)
// 契约:
//   POST /orders            创建订单 {"item":"coffee","quantity":2} -> 201 {order}
//   GET  /orders/{id}       查询订单 -> 200 {order} / 404
//   POST /orders/{id}/pay   幂等支付(并发安全) -> 200 {"paid":true} / 409
//   GET  /healthz           健康检查 -> 200 ok
//   GET  /metrics           Prometheus 指标 -> 200 text
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var store = NewStore()

func main() {
	http.HandleFunc("/orders", handleOrders)
	http.HandleFunc("/orders/", handleOrder)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "# HELP orders_total 创建的订单总数\n# TYPE orders_total counter\norders_total %d\n", store.TotalOrders())
	})
	port := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Item     string `json:"item"`
		Quantity int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Item == "" || req.Quantity <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"error":"invalid request"}`)
		return
	}
	o, err := store.CreateOrder(req.Item, req.Quantity)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, `{"error":%q}`, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/orders/"), "/")
	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch {
	case r.Method == http.MethodGet:
		o, err := store.GetOrder(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, `{"error":"order not found"}`)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(o)
	case r.Method == http.MethodPost && len(parts) == 2 && parts[1] == "pay":
		if err := store.PayOrder(id); err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, `{"error":%q}`, err.Error())
			return
		}
		fmt.Fprintln(w, `{"paid":true}`)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
