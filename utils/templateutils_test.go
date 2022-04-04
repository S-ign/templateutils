package templateutils_test

import (
	"reflect"
	"testing"

	templateutils "github.com/S-ign/templateutils/utils"
)

func TestAllSubstrInBetweenTwoStrs(t *testing.T) {
	tests := []struct {
		name string
		got  map[string]string
		want map[string]string
	}{
		{
			name: "two words",
			got:  templateutils.GetAllSubstrInBetweenTwoStrs("{{.hello}} whats up {{.blah}}", "{{.", "}}"),
			want: map[string]string{"hello": "", "blah": ""},
		},
		{
			name: "many many words",
			got: templateutils.GetAllSubstrInBetweenTwoStrs(
				"welcome to, {{.company}}, {{.employee}}. Todays task is to: {{.task1}} {{.task2}} {{.task3}} {{.task4}}",
				"{{.", "}}"),
			want: map[string]string{"company": "", "employee": "", "task1": "", "task2": "", "task3": "", "task4": ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got: %v\nwant: %v", tt.got, tt.want)
			}
		})
	}
}
