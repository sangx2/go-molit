package molit

import (
	"fmt"
	"net/http"
)

type Molit struct {
	serviceKey string
}

func NewMolit(serviceKey string) *Molit {
	return &Molit{
		serviceKey: serviceKey,
	}
}

// GetTradeApts 아파트 매매 거래 내역 조회
func (m *Molit) GetTradeApts(areaCode string, yyyy, mm int) (string, []Apt, error) {
	req := NewRequest(RequestAptTrade, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]Apt{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetRentApts 아파트 전월세 거래 내역 조회
func (m *Molit) GetRentApts(areaCode string, yyyy, mm int) (string, []Apt, error) {
	req := NewRequest(RequestAptRent, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]Apt{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetTradeRHs 연립다세대 매매 거래 내역 조회
func (m *Molit) GetTradeRHs(areaCode string, yyyy, mm int) (string, []RH, error) {
	req := NewRequest(RequestRhTrade, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]RH{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetRentRHs 연립다세대 전월세 거래 내역 조회
func (m *Molit) GetRentRHs(areaCode string, yyyy, mm int) (string, []RH, error) {
	req := NewRequest(RequestRhRent, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]RH{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetTradeSHs 단독다가구 매매 거래 내역 조회
func (m *Molit) GetTradeSHs(areaCode string, yyyy, mm int) (string, []SH, error) {
	req := NewRequest(RequestShTrade, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]SH{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetRentSHs 단독다가구 전월세 거래 내역 조회
func (m *Molit) GetRentSHs(areaCode string, yyyy, mm int) (string, []SH, error) {
	req := NewRequest(RequestShRent, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()
	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]SH{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetTradeOffices 오피스텔 매매 거래 내역 조회
func (m *Molit) GetTradeOffices(areaCode string, yyyy, mm int) (string, []Office, error) {
	req := NewRequest(RequestOfficeTrade, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]Office{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetRentOffices 오피스텔 전월세 거래 내역 조회
func (m *Molit) GetRentOffices(areaCode string, yyyy, mm int) (string, []Office, error) {
	req := NewRequest(RequestOfficeRent, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]Office{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}

// GetTradeLands 토지 매매 거래 내역 조회
func (m *Molit) GetTradeLands(areaCode string, yyyy, mm int) (string, []Land, error) {
	req := NewRequest(RequestLandTrade, m.serviceKey, areaCode, yyyy, mm)
	if e := req.isValid(); e != nil {
		return "", nil, e
	}

	url := req.GetUrl()

	resp, e := http.Get(url)
	if e != nil {
		return url, nil, e
	}
	defer resp.Body.Close()

	response := ResponseFromXML([]Land{}, resp.Body)
	if response.Header.ResultCode != 0 {
		return url, nil, fmt.Errorf("code: %d, msg: %s", response.Header.ResultCode,
			response.Header.ResultMsg)
	}
	return url, response.Body, nil
}
