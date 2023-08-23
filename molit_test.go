package molit

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	// fix me
	serviceKey = ""
	areaCode   = "41117"
	yyyy       = 2023
	mm         = 7
)

func TestMolit(t *testing.T) {
	m := NewMolit(serviceKey)

	// 아파트
	url, tradeApts, e := m.GetTradeApts(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetTradeApts: %s", e.Error())
	}
	fmt.Printf("tradeApts: url: %s\n", url)
	bTradeApts, _ := json.Marshal(tradeApts)
	fmt.Printf("tradeApts: data: %s\n", bTradeApts)

	url, rentApts, e := m.GetRentApts(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetRentApts: %s", e.Error())
	}
	fmt.Printf("rentApts: url: %s\n", url)
	bRentApts, _ := json.Marshal(rentApts)
	fmt.Printf("rentApts: data: %s\n", bRentApts)

	// 오피스텔
	url, tradeOffices, e := m.GetTradeOffices(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetTradeOffices: %s", e.Error())
	}
	fmt.Printf("tradeOffice: url: %s\n", url)
	bTradeOffices, _ := json.Marshal(tradeOffices)
	fmt.Printf("tradeOffice: data: %s\n", bTradeOffices)

	url, rentOffices, e := m.GetRentOffices(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetRentOffices: %s", e.Error())
	}
	fmt.Printf("rentOffice: url: %s\n", url)
	bRentOffices, _ := json.Marshal(rentOffices)
	fmt.Printf("rentOffice: data: %s\n", bRentOffices)

	// 단독다가구
	url, tradeSHs, e := m.GetTradeSHs(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetTradeSHs: %s", e.Error())
	}
	fmt.Printf("tradeSHs: url: %s\n", url)
	bTradeSHs, _ := json.Marshal(tradeSHs)
	fmt.Printf("tradeSHs: data: %s\n", bTradeSHs)

	url, rentSHs, e := m.GetRentSHs(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetRentSHs: %s", e.Error())
	}
	fmt.Printf("rentSHs: url: %s\n", url)
	bRentSHs, _ := json.Marshal(rentSHs)
	fmt.Printf("rentSHs: data: %s\n", bRentSHs)

	// 연립다세대
	url, tradeRHs, e := m.GetTradeRHs(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetTradeRHs: %s", e.Error())
	}
	fmt.Printf("tradeRHs: url: %s\n", url)
	bTradeRHs, _ := json.Marshal(tradeRHs)
	fmt.Printf("tradeRHs: data: %s\n", bTradeRHs)

	url, rentRHs, e := m.GetRentRHs(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetRentRHs: %s", e.Error())
	}
	fmt.Printf("rentRHs: url: %s\n", url)
	bRentRHs, _ := json.Marshal(rentRHs)
	fmt.Printf("rentRHs: data: %s\n", bRentRHs)

	// 토지
	url, tradeLands, e := m.GetTradeLands(areaCode, yyyy, mm)
	if e != nil {
		t.Fatalf("GetTradeLands: %s", e.Error())
	}
	fmt.Printf("tradeLands: url: %s\n", url)
	bTradeLands, _ := json.Marshal(tradeLands)
	fmt.Printf("tradeLands: data: %s\n", bTradeLands)
}
