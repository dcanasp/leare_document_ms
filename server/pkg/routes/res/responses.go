package res

import (
	"encoding/json"
	"global/globalTypes"
	logs "global/logging"
	"net/http"
)

// unique for batch find
func Response_S_Structure(w http.ResponseWriter, success bool, value map[string]globalTypes.DynamoEntry) {

	w.Header().Set("Content-Type", "application/json")

	responseObj := Dto_S_V{
		Success: success,
		Value:   value,
	}

	errJson := json.NewEncoder(w).Encode(responseObj)
	if errJson != nil {
		logs.E.Printf("Failed to encode JSON: %v", errJson)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// unique for find unique
func Response_S_Dynamo(w http.ResponseWriter, success bool, value globalTypes.DynamoEntry) {

	w.Header().Set("Content-Type", "application/json")

	responseObj := Dto_S_V{
		Success: success,
		Value:   value,
	}

	errJson := json.NewEncoder(w).Encode(responseObj)
	if errJson != nil {
		logs.E.Printf("Failed to encode JSON: %v", errJson)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// all worked? what is the value
func Response_S_V(w http.ResponseWriter, success bool, value string) {

	w.Header().Set("Content-Type", "application/json")

	responseObj := Dto_S_V{
		Success: success,
		Value:   value,
	}

	errJson := json.NewEncoder(w).Encode(responseObj)
	if errJson != nil {
		logs.E.Printf("Failed to encode JSON: %v", errJson)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// there was an error. Success always false
func Response_Error(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")

	responseObj := Dto_S_E{
		Success: false,
		Error:   err,
	}
	w.WriteHeader(http.StatusBadRequest)
	errJson := json.NewEncoder(w).Encode(responseObj)
	if errJson != nil {
		logs.E.Printf("Failed to encode JSON: %v", errJson)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// for those responses where all worked but there is no data
func Response_Success(w http.ResponseWriter) {
	// there was an error. Success always false
	w.Header().Set("Content-Type", "application/json")

	responseObj := Dto_S{
		Success: true,
	}
	w.WriteHeader(http.StatusCreated)
	errJson := json.NewEncoder(w).Encode(responseObj)
	if errJson != nil {
		logs.E.Printf("Failed to encode JSON: %v", errJson)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
