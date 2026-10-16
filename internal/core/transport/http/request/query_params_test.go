package core_http_request

import (
	"errors"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

func TestGetTimeQueryParam(t *testing.T) {
	for _, tc := range []struct {
		name, input, want string
		invalid           bool
	}{
		{name: "missing"},
		{name: "legacy date", input: "2026-10-16", want: "2026-10-16T00:00:00Z"},
		{name: "local midnight", input: "2026-10-16T00:00:00+05:00", want: "2026-10-15T19:00:00Z"},
		{name: "UTC", input: "2026-10-16T19:00:00.000Z", want: "2026-10-16T19:00:00Z"},
		{name: "invalid", input: "yesterday", invalid: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/statistics?from="+url.QueryEscape(tc.input), nil)
			got, err := GetTimeQueryParam(request, "from")
			if tc.invalid {
				if !errors.Is(err, core_errors.ErrInvalidArgument) {
					t.Fatalf("expected invalid argument, got %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if tc.want == "" {
				if got != nil {
					t.Fatal("missing parameter must remain nil")
				}
				return
			}
			if got == nil || got.UTC().Format(time.RFC3339) != tc.want {
				t.Fatalf("got %v, want %s", got, tc.want)
			}
		})
	}
}
