package main

import (
	"github.com/amanelis/skynet/model"
	ws "github.com/gorilla/websocket"
)

type Helper struct {
}

func (h Helper) gdaxConnectWss() *ws.Conn {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)
	if err != nil {
		println(err.Error())
	}

	return wsConn
}

func (h Helper) gdaxSubscribeParams(t string, p string) map[string]string {
	return map[string]string{
		"type":       t,
		"product_id": p,
	}
}

func (h Helper) avrMinMax(l []*model.Order) map[string]float64 {
	var total float64
	var min = l[0].Price
	var max = l[0].Price

	for _, value := range l {
		total += value.Price

		if value.Price > max {
			max = value.Price
		}

		if value.Price < min {
			min = value.Price
		}
	}

	avg := total / float64(len(l))

	return map[string]float64{
		"avg": avg,
		"min": min,
		"max": max,
	}
}

func LoadHelper() Helper {
	return Helper{}
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
