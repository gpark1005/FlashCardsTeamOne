package incomingdata

type NewDeck struct {
	Catagory string
	QuizInfo Info
}

type Info struct {
	Id          string
	Type        string
	Information string
}
