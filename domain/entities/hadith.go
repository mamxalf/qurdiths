package entities

type Book struct {
	ID     string
	Name   string
	Amount int32
}

type HadithJson struct {
	Number     int32  `json:"number"`
	Arabian    string `json:"arab"`
	Indonesian string `json:"id"`
}

type Hadith struct {
	Book       string
	Number     int32
	Arabian    string
	Indonesian string
}

type InsertHadith struct {
	Name string
	File string
}
