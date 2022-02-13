package main

import (
	"reflect"
	"testing"
)

func Test_findStructMethods(t *testing.T) {
	type args struct {
		path string
		id   string
	}
	tests := []struct {
		name    string
		args    args
		want    []Method
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				path: "td",
				id:   "Struct1",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, _, err := findStructAndMethods(tt.args.path, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("findStructAndMethods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findStructAndMethods() got = %v, want %v", got, tt.want)
			}
		})
	}
}
