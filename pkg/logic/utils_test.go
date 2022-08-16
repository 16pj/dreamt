package logic

import (
	"reflect"
	"testing"
)

func Test_MostFrequentWords(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				sentence: "Hello World World Hello Hello This is a test",
			},
			want: []string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostFrequentWords(tt.args.sentence, 2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostFrequentWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
