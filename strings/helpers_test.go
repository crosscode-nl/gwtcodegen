package strings

import "testing"

func TestAlphaNumCamelCase(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name:       "normal operation",
			args:       args{
				line: "this is a string ( which should be ). converted to 1 camel case based string",
			},
			wantResult: "ThisIsAStringWhichShouldBeConvertedTo1CamelCaseBasedString",
		},
		{
			name:       "single word",
			args:       args{
				line: "this",
			},
			wantResult: "This",
		},
		{
			name:       "already camel cased words",
			args:       args{
				line: "this contains camelCased words",
			},
			wantResult: "ThisContainsCamelCasedWords",
		},
		{
			name:       "single character",
			args:       args{
				line: "t w d",
			},
			wantResult: "TWD",
		},
		{
			name:       "empty string",
			args:       args{
				line: "",
			},
			wantResult: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AlphaNumCamelCase(tt.args.line); gotResult != tt.wantResult {
				t.Errorf("CamelCase() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
