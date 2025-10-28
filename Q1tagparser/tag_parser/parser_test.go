package tagparser

import (
	"reflect"
	"testing"
)

func TestParseTags(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"golang", []string{"golang"}},
		{"golang, python", []string{"golang", "python"}},
		{"java, C#, python", []string{"java", "C#", "python"}},
		{" csharp", []string{"csharp"}},
		{",csharp", []string{"csharp"}},
		{"csharp,", []string{"csharp"}},
		{"C# byte code, python", []string{"C# byte code", "python"}},
		{"csharp,,python", []string{"csharp", "python"}},
		{" , csharp , , python ", []string{"csharp", "python"}},
	}

	for _, test := range tests {
		got := ParseTags(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ParseTags(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}
