package main

import "testing"

func Test_englishGreet_Greet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "english",
			want: "hello\n",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &englishGreet{}
			if got := e.Greet(); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
