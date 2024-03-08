package globalTypes

type DbEntry struct {
	Url    string `json:"url"`
	Type   string `json:"type"`
	Date   string `json:"date"`
	UserId string `json:"userId"`
}

type BrokerEntry struct {
	FileName string
	FileType string
	UserId   string
}
