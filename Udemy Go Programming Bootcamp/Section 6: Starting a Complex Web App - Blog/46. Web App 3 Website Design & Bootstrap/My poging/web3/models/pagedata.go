package models

type PageData struct{
	StrMap  map[string]string
	IntMap  map[string]int
	FltMap  map[string]float32
	DataMap map[string]interface{} // may contain any type of data
	CSRFToken string
	Warning   string
	Error     string
}