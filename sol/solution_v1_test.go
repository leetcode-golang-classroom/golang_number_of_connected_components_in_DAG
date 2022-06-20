package sol

import "testing"

func BenchmarkTestV1(b *testing.B) {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	for idx := 0; idx < b.N; idx++ {
		countNumberOfConnectedComponentsV1(n, edges)
	}
}
func Test_countNumberOfConnectedComponentsV1(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 5 and edges = [[0, 1], [1, 2], [3, 4]] ",
			args: args{n: 5, edges: [][]int{{0, 1}, {1, 2}, {3, 4}}},
			want: 2,
		},
		{
			name: "n = 5 and edges = [[0, 1], [1, 2], [2, 3], [3, 4]]",
			args: args{n: 5, edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumberOfConnectedComponentsV1(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countNumberOfConnectedComponentsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
