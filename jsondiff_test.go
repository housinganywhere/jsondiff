package jsondiff

import (
	"testing"
)

var cases = []struct {
	a      string // event
	b      string // match template
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

	// {`{"a": 3.1415}`, `{"a": 3.14156}`, NoMatch},
	// {`{"a": 3.1415}`, `{"a": 3.1415}`, FullMatch},
	// {`{"a": 4213123123}`, `{"a": "4213123123"}`, NoMatch},
	// {`{"a": 4213123123}`, `{"a": 4213123123}`, FullMatch},
	// {`"dns"`, `"dns.1239102398"`, FullMatch},       // regex for literal strings
	// {`{"a":"dns"}`, `{"a":"dnsutils"}`, FullMatch}, // regex for literal string values
	// {`{"name": "dns"}`, `{"name": "dnsutils"}`, FullMatch},
	// {`{"dnsutils": 4213123123}`, `{"dns[a-zA-Z]{5}": 4213123123}`, SupersetMatch},
	// {`{"dnsutils": 4213123123}`, `{"dns": 4213123123, "b": "lessQQ"}`, NoMatch},
	// {`{"dnsutils": 4213123123, "existsNot": "morePewPew"}`, `{"dns": 4213123124}`, SupersetMatch},
	// {`{"kind":"Pod","namespace":"stage","name":"stage-storybook"}`, `{"kind":"Pod","namespace":"stage","name":"stage-storybook-5bd85966c4-c6k5p"}`, FullMatch},
	// {`{"namespace":"stage","name":"stage-storybook"}`, `{"name":"stage-storybook"}`, SupersetMatch},
	// {`{"namespace":"stage","name":"stage-osiris", "kind":"Pod"}`, `{"namespace":"stage","name":"stage-osiris-sheets"}`, SupersetMatch},

	{`{"namespace":"stage","name":"stage-osiris"}`, `{"namespace":"stage","name":"stage-osiris-sheets","kind":"Pod"}`, NoMatch},
	{`{"kind":"Pod","namespace":"stage","name":"stage-storybook-5bd85966c4-c6k5p"}`, `{"kind":"Pod","namespace":"stage","name":"stage-storybook.*"}`, FullMatch},
	// literal cases for BigBrother
	// SawCompletedJob
	{`{"apiVersion": "batch/v1beta1","kind": "CronJob","name": "production-osiris-sheets","namespace": "production","resourceVersion": "405900509","uid": "9c1b28d0-5f07-11ea-9f38-42010a840052"}`, `{"name": "production-osiris-sheets","namespace": "production"}`, SupersetMatch},
	// Pod Start
	{`{"kind":"Pod","name":"production-osiris-sheets-1594287000-p2r4c","namespace":"production","resourceVersion":"405935284","uid":"c6da6414-c1c6-11ea-9ab1-42010a8401a6"}`, `{"kind":"Pod","name":"production-osiris","namespace":"production"}`, SupersetMatch},
	// Pod Killed
	{`{"apiVersion": "v1","kind": "Pod","name": "production-imaginary-586d4bc49d-vmn2q","namespace": "production","resourceVersion": "405899391"}`, `{"kind": "Pod","name": "production-imaginary","namespace": "production"}`, SupersetMatch},
	// HelmRelease Warnings
	// {``, ``, SupersetMatch},
	//
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
