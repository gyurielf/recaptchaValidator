package recaptchaValidator

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/volatiletech/null"
)

const endpoint = "https://www.google.com/recaptcha/api/siteverify"

type Response struct {
	Success bool         `json:"success"`
	Score   null.Float64 `json:"score,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}

func Verify(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		errorResponse(w, err, "invalid request body")
		log.Print(err)
		return
	}

	q := r.URL.Query()
	q.Add("secret", "123123123123key123123123")
	req.URL.RawQuery = q.Encode()

	rawResp, err := http.Get(endpoint)

	if err != nil {
		log.Print(err.Error())
		return
	}
	defer rawResp.Body.Close()

	var respBody Response
	if err = json.NewDecoder(rawResp.Body).Decode(&respBody); err != nil {
		errorResponse(w, err, "invalid request body")
		log.Print("error decoding recaptcha body")
		return
	}
}

func errorResponse(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusBadRequest)
	raw, _ := json.Marshal(Error{
		Message: message,
	})
	w.Write(raw)
}
