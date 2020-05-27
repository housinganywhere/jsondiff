package jsondiff

import (
	"testing"
)

var cases = []struct {
	a      string
	b      string
	result Difference
}{
	{`{"a": 5}`, `["a"]`, NoMatch},
	{`{"a": 5}`, `{"a": 6}`, NoMatch},
	{`{"a": 5}`, `{"a": true}`, NoMatch},
	{`{"a": 5}`, `{"a": 5}`, FullMatch},
	{`{"a": 5}`, `{"a": 5, "b": 6}`, NoMatch},
	{`{"a": 5, "b": 6}`, `{"a": 5}`, SupersetMatch},
	{`{"a": 5, "b": 6}`, `{"b": 6}`, SupersetMatch},
	{`{"a": null}`, `{"a": 1}`, NoMatch},
	{`{"a": null}`, `{"a": null}`, FullMatch},
	{`{"a": "null"}`, `{"a": null}`, NoMatch},
	{`{"a": 3.1415}`, `{"a": 3.14156}`, NoMatch},
	{`{"a": 3.1415}`, `{"a": 3.1415}`, FullMatch},
	{`{"a": 4213123123}`, `{"a": "4213123123"}`, NoMatch},
	{`{"a": 4213123123}`, `{"a": 4213123123}`, FullMatch},
	{`"dns"`, `"dns.1239102398"`, FullMatch},       // regex for literal strings
	{`{"a":"dns"}`, `{"a":"dnsutils"}`, FullMatch}, // regex for literal string values
	{`{"name": "dns"}`, `{"name": "dnsutils"}`, FullMatch},
	{`{"dnsutils": 4213123123}`, `{"dns[a-zA-Z]{5}": 4213123123}`, SupersetMatch},
	{`{"dnsutils": 4213123123}`, `{"dns": 4213123123, "b": "lessQQ"}`, NoMatch},
	{`{"dnsutils": 4213123123, "existsNot": "morePewPew"}`, `{"dns": 4213123124}`, SupersetMatch},
	{`"metadata": {"name": "monitoring-dnsutils","namespace": "monitoring"}`, `"metadata": {"name": "monitoring-dns","namespace": "monitoring"}`, FullMatch},
	{`{"metadata": {"name": "monitoring-dnsutils","namespace": "monitoring"}}`, `{"metadata": {"name": "monitoring-dnsutils.90239123","namespace": "monitoring"}}`, FullMatch},
}

func TestCompare(t *testing.T) {
	opts := DefaultConsoleOptions()
	opts.PrintTypes = false
	for i, c := range cases {
		result, _ := Compare([]byte(c.a), []byte(c.b), &opts)
		if result != c.result {
			t.Errorf("case %d failed, got: %s, expected: %s", i, result, c.result)
		}
	}
}
