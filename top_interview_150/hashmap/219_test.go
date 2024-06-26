package hashmap

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 3",
			args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2},
			want: false,
		},
		{
			name: "case 56",
			args: args{nums: []int{4, 1, 2, 3, 1, 5}, k: 3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
