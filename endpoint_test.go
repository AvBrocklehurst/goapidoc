package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_endpoint_String(t *testing.T) {
	tests := []struct {
		name string
		ep   endpoint
		want string
	}{
		{
			name: "Valid endpoint string test",
			ep: endpoint{
				Route:   "/test/route",
				Returns: "user",
			},
			want: fmt.Sprintf("%s\n%s", "/test/route", "user"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ep.String(); got != tt.want {
				t.Errorf("%s: endpoint.String() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_endpoint_parseLine(t *testing.T) {
	type args struct {
		line string
		fp   *fileParser
	}
	tests := []struct {
		name   string
		ep     *endpoint
		args   args
		wantEp endpoint
	}{
		{
			name: "Valid endpoint parse line: Route test",
			ep:   &endpoint{},
			args: args{
				line: "route: /route/name",
				fp:   &fileParser{},
			},
			wantEp: endpoint{
				Route: "/route/name",
			},
		},
		{
			name: "Valid endpoint parse line: Method test",
			ep:   &endpoint{},
			args: args{
				line: "method: GET",
				fp:   &fileParser{},
			},
			wantEp: endpoint{
				Method: "GET",
			},
		},
		{
			name: "Valid endpoint parse line: Description test",
			ep:   &endpoint{},
			args: args{
				line: "description: some description explaining some things",
				fp:   &fileParser{},
			},
			wantEp: endpoint{
				Description: "some description explaining some things",
			},
		},
		{
			name: "Valid endpoint parse line: Name test",
			ep:   &endpoint{},
			args: args{
				line: "name: Method Name",
				fp:   &fileParser{},
			},
			wantEp: endpoint{
				Description: "Method Name",
			},
		},
		{
			name: "Valid endpoint parse line: Params test",
			ep:   &endpoint{},
			args: args{
				line: "params: id int query",
				fp:   &fileParser{},
			},
			wantEp: endpoint{
				Params: []param{
					param{
						Name:     "id",
						Type:     "int",
						Location: "query",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ep.parseLine(tt.args.line, tt.args.fp)
			if !reflect.DeepEqual(*tt.ep, tt.wantEp) {
				t.Errorf("%s: endpoint.parseLine(). Endpoint is %v, want %v", tt.name, *tt.ep, tt.wantEp)
			}
		})
	}
}

func Test_endpoint_parseParam(t *testing.T) {
	type args struct {
		parts []string
	}
	tests := []struct {
		name string
		ep   *endpoint
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ep.parseParam(tt.args.parts)
		})
	}
}

func Test_endpoint_parseReturns(t *testing.T) {
	type args struct {
		parts []string
		fp    *fileParser
	}
	tests := []struct {
		name string
		ep   *endpoint
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ep.parseReturns(tt.args.parts, tt.args.fp)
		})
	}
}
