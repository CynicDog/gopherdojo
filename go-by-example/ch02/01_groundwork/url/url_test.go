package url

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	const url = "https://github.com/cynicdog"

	got, err := Parse(url)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", url, err)
	}
	want := &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "cynicdog",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot %#v\nwant %#v", url, got, want)
	}
}

func TestURLString(t *testing.T) {
	u := &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "cynicdog",
	}

	got := u.String()
	want := "https://github.com/cynicdog"
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestParseWithoutPath(t *testing.T) {
	const url = "https://github.com"

	got, err := Parse(url)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", url, err)
	}

	want := &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Parse(%q) mismatch (-want +got):\n%s", url, diff)
	}
}

// a table of test cases
var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{ // test case meant to fail
		name: "with_data_scheme",
		uri:  "data:text/plain;base64,R28gYnkgRXhhbXBsZQ==",
		want: &URL{Scheme: "data"},
	},
	{
		name: "full",
		uri:  "https://github.com/cynicdog",
		want: &URL{
			Scheme: "https",
			Host:   "github.com",
			Path:   "cynicdog",
		},
	},
	{
		name: "without_path",
		uri:  "https://github.com",
		want: &URL{
			Scheme: "https",
			Host:   "github.com",
			Path:   "",
		},
	},
}

func TestParseTable(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.uri)
			if err != nil {
				t.Fatalf("Parse(%q) err = %v, want <nil>", tt.uri, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Parse(%q) mismatch (-want +got):\n%s", tt.uri, diff)
			}
		})
	}
}
