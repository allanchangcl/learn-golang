package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_Handler(t *testing.T) {
	testCases := []struct {
		desc           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			desc:           "good",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   message,
			// expectedBody: "wrong message",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := NewHandlers(nil)
			h.Home(tC.out, tC.in)
			if tC.out.Code != tC.expectedStatus {
				// t.Errorf("expected status %d, got %d", tC.expectedStatus, tC.out.Code)
				t.Logf("expected: %d\ngot: %d\n", tC.expectedStatus, tC.out.Code)
				t.Fail()
			}

			body := tC.out.Body.String()
			// fmt.Println(body)
			if body != tC.expectedBody {
				// t.Errorf("expected body %s, got %s", tC.expectedBody, body)
				t.Logf("expected: %s\ngot: %s\n", tC.expectedBody, body)
			}
		})
	}
}
