package functions

import (
	"reflect"
	"testing"
)

func TestRoutesApi_AppsAppRoutesPostByCode(t *testing.T) {
	type fields struct {
		Configuration *Configuration
	}
	type args struct {
		app  string
		body RouteEx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RouteWrapper
		want1   *APIResponse
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := RoutesApi{
				Configuration: tt.fields.Configuration,
			}
			got, got1, err := a.AppsAppRoutesPostByCode(tt.args.app, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoutesApi.AppsAppRoutesPostByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoutesApi.AppsAppRoutesPostByCode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RoutesApi.AppsAppRoutesPostByCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
