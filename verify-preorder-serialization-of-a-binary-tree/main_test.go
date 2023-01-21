package verifypreorderserializationofabinarytree

import "testing"

func Test_isValidSerialization(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#"}, true}, //
		{"2", args{preorder: "1,#"}, false},                      //
		{"3", args{preorder: "9,#,#,1"}, false},                  //
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerialization(tt.args.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}
