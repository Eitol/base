package couchdb

import (
	"reflect"
	"testing"
)

type testStruct struct {
	A string
}

func Test_ensureTypeField(t *testing.T) {
	doc := testStruct{
		A: "Test",
	}
	type args struct {
		doc interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "",
			args: args{
				doc: doc,
			},
			want: map[string]interface{}{
				"A":    "Test",
				"type": "testStruct",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ensureTypeField(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ensureTypeField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ensureTypeField() = %v, want %v", got, tt.want)
			}
		})
	}
}
