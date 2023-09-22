package controller

import (
	mo "challengeeasybroker/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestContenidoContactShow(t *testing.T) {

	testsexpected := mo.Contact{
		Pagination: mo.Pagination{
			Limit:     2,
			Page:      1,
			Total:     833,
			Next_Page: "https://api.stagingeb.com/v1/contact_requests?limit=2&page=2",
		},
		Content: []mo.Content{
			{
				Name:        "Omar Juarez",
				Phone:       "7771567090",
				Email:       "omar+homeseeker@easybroker.com",
				Property_id: "EB-C6501",
				Message:     "Me interesa mucho esta propiedad y quiero recibir más información.\r\n¡Gracias!",
				Source:      "stagingeb.com",
				Happened_at: "2023-08-02T20:02:05-06:00",
			},
			{
				Name:        "Omar Juarez",
				Phone:       "7771567090",
				Email:       "omar+homeseeker@easybroker.com",
				Property_id: "EB-C6501",
				Message:     "Me interesa mucho esta propiedad y quiero recibir más información.\r\n¡Gracias!",
				Source:      "stagingeb.com",
				Happened_at: "2023-08-02T20:02:05-06:00",
			},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/contact?page=1&limit=2", nil)
	w := httptest.NewRecorder()
	ContenidoContactShow(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	var actual mo.Contact
	if err := json.Unmarshal(data, &actual); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(actual, testsexpected) {
		t.Errorf("Expected Hello john but got %v", string(data))
		t.Errorf("Expected Teste john but got %v", (testsexpected))

	}
}
