package incomingdata

type Deck struct {
	Catagory string
	QuizInfo []Info
}

type Info struct {
	Id          string
	Type        string
	Information string
}
