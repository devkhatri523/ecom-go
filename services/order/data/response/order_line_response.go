package response

import "encoding/json"

type OrderLineResponse struct {
	Id       int32   `json:"id"`
	Quantity float64 `json:"quantity"`
}

func (r OrderLineResponse) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
