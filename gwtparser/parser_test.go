package gwtparser

import (
	"gwtcodegen/model"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []model.Given
	}{
		{
			name: "single given, single when, single then",
			args: args{
				input:
					"// given this test input\n" +
					"// when this test condition\n" +
					"// then this test result",
			},
			want: []model.Given{
				{
					Text: "Given this test input",
					When: []model.When{
						{
							Text: "When this test condition",
							Then: []model.Then{
								{
									Text: "Then this test result",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "single Given, single When, single Then",
			args: args{
				input:
				"// Given this test input\n" +
					"// When this test condition\n" +
					"// Then this test result",
			},
			want: []model.Given{
				{
					Text: "Given this test input",
					When: []model.When{
						{
							Text: "When this test condition",
							Then: []model.Then{
								{
									Text: "Then this test result",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
