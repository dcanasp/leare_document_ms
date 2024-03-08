package res

type Dto_S_V struct {
	Success bool        `json:"success"`
	Value   interface{} `json:"value"`
}

type Dto_S_E struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
}

type Dto_S struct {
	Success bool `json:"success"`
}
