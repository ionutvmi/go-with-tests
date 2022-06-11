package maps

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		dictionary map[string]string
		query      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should find the key in the map",
			args: args{
				dictionary: map[string]string{"test": "this is just a test"},
				query:      "test",
			},
			want: "this is just a test",
		},
		{
			name: "should return an empty string if key doesn't exist",
			args: args{
				dictionary: map[string]string{"test": "this is just a test"},
				query:      "not exist",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Search(tt.args.dictionary, tt.args.query); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
