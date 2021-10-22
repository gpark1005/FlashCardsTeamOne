package incomingdata

type test interface{}

//Matching start
type Matching struct {
	Id       string                 `json:"id"`
	Type     string                 `json:"type"`
	Category string                 `json:"category"`
	Quesions map[string]interface{} `json:"questions"` //map to string/interface
	Options  map[string]interface{} `json:"options"`
	Answers  map[string]interface{} `json:"answers"`
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

//Matching end

type Info struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Information string `json:"information"`
}

type QNA struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Information string `json:"information"`
	Quesion     string `json:"question"`
	Answer      string `json:"answer"`
}

type TrueOrFalse struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Quesion  string `json:"question"`
	Answer   string `json:"answer"`
}

//MC start
type MultipleChoice struct {
	Id       string          `json:"id"`
	Type     string          `json:"type"`
	Category string          `json:"category"`
	Quesion  string          `json:"question"`
	Options  MultipleChoiceO `json:"options"`
	Answers  int             `json:"answers"`
}

type MultipleChoiceO struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
}

//MC end
