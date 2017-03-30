package functions

import "testing"

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
		want    *RouteEx
		want1   *APIResponse
		wantErr bool
	}{
		// Add test cases.
		{
			name:    "go",
			wantErr: false,
			args: args{
				app: "tang-client",
				body: RouteEx{
					Route: Route{
						Path:  "/clienttest",
						Image: "dcatalog.hnaresearch.com/tangqian/godemo:0.1",
					},
					AppName: "tang-client",
					Runtime: "golang",
					Code: `package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
}

func main() {
	p := &Person{Name: "World"}
	json.NewDecoder(os.Stdin).Decode(p)
	fmt.Printf("Hello %v!", p.Name)
}`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewRoutesApiWithBasePathAndSteelServer("http://192.168.56.6:8080/v1", "http://localhost:8081/v1")
			_, _, err := a.AppsAppRoutesPostByCode(tt.args.app, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoutesApi.AppsAppRoutesPostByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("RoutesApi.AppsAppRoutesPostByCode() got = %v, want %v", got, tt.want)
			// }
			// if !reflect.DeepEqual(got1, tt.want1) {
			// 	t.Errorf("RoutesApi.AppsAppRoutesPostByCode() got1 = %v, want %v", got1, tt.want1)
			// }
		})
	}
}
