package globalTypes

type DbEntry struct {
	Url    string `json:"url"`
	Type   string `json:"type"`
	Date   string `json:"date"`
	UserId string `json:"userId"`
}

type BrokerEntry struct {
	VideoId  string
	FileName string
	FileType string
	UserId   string
}

type DynamoEntry struct {
	FilePath string
	BrokerEntry
	Date int64
}
