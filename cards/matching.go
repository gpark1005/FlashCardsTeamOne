package cards

type Matching struct {
	Id        string                 `json:"id"`
	Type      string                 `json:"type"`
	Category  string                 `json:"category"`
	Questions map[string]interface{} `json:"questions"` //map to string/interface
	Options   map[string]interface{} `json:"options"`
	Answers   map[string]interface{} `json:"answers"`
}

type MatchingQ struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

type MatchingO struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
}

type MatchingA struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}
