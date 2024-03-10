package globalTypes

type DbEntry struct {
	Url    string `json:"url"`
	Type   string `json:"type"`
	Date   string `json:"date"`
	UserId string `json:"userId"`
}

type BrokerEntry struct {
	VideoId  string `json:"VideoId"`
	FileName string `json:"FileName"`
	FileType string `json:"FileType"`
	UserId   string `json:"UserId"`
}

type DynamoEntry struct {
	FilePath string `json:"FilePath"`
	BrokerEntry
	Date int64 `json:"Date"`
}
