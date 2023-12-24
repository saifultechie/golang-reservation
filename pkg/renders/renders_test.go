package renders

import (
	models "github/saifultechie/booking/pkg/templateData"
	"net/http"
	"testing"
)

func TestAddTemplateData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := addTemplateData(&td, r)
	if result.Flash != "123" {
		t.Error("failed")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-yhis", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}
