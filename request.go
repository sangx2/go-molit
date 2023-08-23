package molit

import "fmt"

type RequestType int

const (
	RequestAptTrade = iota
	RequestAptRent
	RequestOfficeTrade
	RequestOfficeRent
	RequestShTrade
	RequestShRent
	RequestRhTrade
	RequestRhRent
	RequestLandTrade
)

const (
	urlBase        = "http://openapi.molit.go.kr"
	urlAptTrade    = urlBase + ":8081/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcAptTrade?"
	urlAptRent     = urlBase + ":8081/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcAptRent?"
	urlRhTrade     = urlBase + ":8081/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcRHTrade?"
	urlRhRent      = urlBase + ":8081/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcRHRent?"
	urlShTrade     = urlBase + ":8081/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcSHTrade?"
	urlShRent      = urlBase + ":8081/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcSHRent?"
	urlOfficeTrade = urlBase + "/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcOffiTrade?"
	urlOfficeRent  = urlBase + "/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcOffiRent?"
	urlLandTrade   = urlBase + "/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcLandTrade?"

	urlLawdCode   = "LAWD_CD="
	urlDeal       = "DEAL_YMD="
	urlServiceKey = "serviceKey="
)

type Request struct {
	reqType RequestType

	key, areaCode string

	yyyy, mm int
}

func NewRequest(requestType RequestType, key, areaCode string, yyyy, mm int) *Request {
	return &Request{
		reqType:  requestType,
		key:      key,
		areaCode: areaCode,
		yyyy:     yyyy,
		mm:       mm,
	}
}

func (r *Request) isValid() error {
	switch r.reqType {
	case RequestAptTrade, RequestOfficeTrade, RequestShTrade, RequestRhTrade, RequestLandTrade:
		if r.yyyy < 2006 {
			return fmt.Errorf("invalid year: %d < 2006", r.yyyy)
		}
	case RequestAptRent, RequestOfficeRent, RequestShRent, RequestRhRent:
		if r.yyyy < 2011 {
			return fmt.Errorf("invalid year: %d < 2011", r.yyyy)
		}
	default:
		return fmt.Errorf("invalid request type")
	}

	if r.mm < 1 || r.mm > 12 {
		return fmt.Errorf("invalid month: %d", r.mm)
	}

	return nil
}

func (r *Request) GetUrl() string {
	var url string

	switch r.reqType {
	case RequestAptTrade:
		url = urlAptTrade
	case RequestAptRent:
		url = urlAptRent
	case RequestOfficeTrade:
		url = urlOfficeTrade
	case RequestOfficeRent:
		url = urlOfficeRent
	case RequestShTrade:
		url = urlShTrade
	case RequestShRent:
		url = urlShRent
	case RequestRhTrade:
		url = urlRhTrade
	case RequestRhRent:
		url = urlRhRent
	case RequestLandTrade:
		url = urlLandTrade
	}

	url += "&" + urlLawdCode + r.areaCode
	url += "&" + urlDeal + fmt.Sprintf("%04d%02d", r.yyyy, r.mm)
	url += "&" + urlServiceKey + r.key

	return url
}
