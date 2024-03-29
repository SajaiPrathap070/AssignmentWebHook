package model

type InputData struct {
	Ev     string `json:"ev"`
	Et     string `json:"et"`
	Id     string `json:"id"`
	Uid    string `json:"uid"`
	Mid    string `json:"mid"`
	T      string `json:"t"`
	P      string `json:"p"`
	L      string `json:"l"`
	Sc     string `json:"sc"`
	Atrk1  string `json:"atrk1"`
	Atrv1  string `json:"atrv1"`
	Atrt1  string `json:"atrt1"`
	Atrk2  string `json:"atrk2"`
	Atrv2  string `json:"atrv2"`
	Atrt2  string `json:"atrt2"`
	Uatrk1 string `json:"uatrk1"`
	Uatrv1 string `json:"uatrv1"`
	Uatrt1 string `json:"uatrt1"`
	Uatrk2 string `json:"uatrk2"`
	Uatrv2 string `json:"uatrv2"`
	Uatrt2 string `json:"uatrt2"`
	Uatrk3 string `json:"uatrk3"`
	Uatrv3 int    `json:"uatrv3"`
	Uatrt3 string `json:"uatrt3"`
}

type OutputData struct {
	Event       string              `json:"event"`
	EventType   string              `json:"event_type"`
	AppId       string              `json:"app_id"`
	UserId      string              `json:"user_id"`
	MessageId   string              `json:"message_id"`
	PageTitle   string              `json:"page_title"`
	PageUrl     string              `json:"page_url"`
	BrowserLang string              `json:"browser_language"`
	ScreenSize  string              `json:"screen_size"`
	Attributes  map[string]AttrData `json:"attributes"`
	Traits      map[string]AttrData `json:"traits"`
}

type AttrData struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}
