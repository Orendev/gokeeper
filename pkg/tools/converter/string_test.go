package converter

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestStringToUUID(t *testing.T) {
	type args struct {
		str string
	}
	id := uuid.New()
	tests := []struct {
		name string
		args args
		want uuid.UUID
	}{
		{
			name: "positive test #1 StringToUUID",
			args: args{str: id.String()},
			want: id,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUUID(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStringToUUID(b *testing.B) {
	id := uuid.New()
	for i := 0; i < b.N; i++ {
		StringToUUID(id.String())
	}
}
