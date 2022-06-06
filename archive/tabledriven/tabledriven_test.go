package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// t.Fatal("not implemented")
	type args struct {
		a int
		b int
	}
	test := []struct {
		name string
		args args
		want int
	}{
		{

			name: "2+2",
			args: args{
				a: 2,
				b: 2,
			},
			want: 4,
		},
		{

			name: "3+3",
			args: args{
				a: 3,
				b: 3,
			},
			want: 6,
		},
	}
	// t.Run(tt.name, func (t *testing.T))
	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			if got := add(v.args.a, v.args.b); got != v.want {
				t.Errorf( "add() = %v, want %v", got, v.want)
			}

		})

	}

}
