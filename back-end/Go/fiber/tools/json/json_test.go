package json

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	segJson "github.com/segmentio/encoding/json"
)

type JSONData struct {
	Name string
	Age  uint8
}

var (
	jsoniterParser = jsoniter.ConfigCompatibleWithStandardLibrary
	s              = []byte(`{"Name":"John","Age":20}`)
)

func BenchmarkJsoniter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := JSONData{
			Name: "John",
			Age:  20,
		}
		_, _ = jsoniterParser.Marshal(&data)
	}
}

func BenchmarkJson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := JSONData{
			Name: "John",
			Age:  20,
		}
		_, _ = json.Marshal(&data)
	}
}

func BenchmarkSegJson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := JSONData{
			Name: "John",
			Age:  20,
		}
		_, _ = segJson.Marshal(&data)
	}
}

func BenchmarkJsoniterDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var data JSONData
		_ = jsoniterParser.Unmarshal(s, &data)
	}
}

func BenchmarkStdJsonDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var data JSONData
		_ = json.Unmarshal(s, &data)
	}
}

func BenchmarkSegJsonDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var data JSONData
		_ = segJson.Unmarshal(s, &data)
	}
}
