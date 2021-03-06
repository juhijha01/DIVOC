package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"vaccination-module/models"
)

func TestAuthenticate(t *testing.T) {
	type args struct {
		request string
		want int
	}
	//ctx1 := &gin.Context{}
	//ctx1.Request, _ = http.NewRequest("POST", "/auth", strings.NewReader("{}"))
	//w := httptest.NewRecorder()


	tests := []struct {
		name string
		args args
	}{
		{
			name: "shouldAuthenticate",
			args: args{request:`{"mobile":"98765443210", "password":"1234"}`, want:200},
		},{
			name: "shouldReturnBadRequestForInvalidInput",
			args: args{request:`{"mobileNumber":"a", "password":"1234"}`, want:400},
		},{
			name: "shouldFail",
			args: args{request:`{"mobile":"0000000000", "password":"1234"}`, want:401},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//OperatorLogin(ctx) //need to mock writer and assert.
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request, _ =  http.NewRequest("POST", "/auth", strings.NewReader(tt.args.request))
			OperatorLogin(ctx)
			//ctx.

			//var got gin.H
			//err := json.Unmarshal(w.Body.Bytes(), &got)
			//if err != nil {
			//	t.Fatal(err)
			//}
			//fmt.Printf("%v", got)
			assert.Equal(t, tt.args.want, w.Result().StatusCode) // w
		})
	}
}

func TestOperatorConfiguration(t *testing.T) {
	type args struct {
		body string
		want int
		DailyLimit int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ShouldGetConfiguration",
			args: args{ body : "" , want: 200, DailyLimit: 10},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request, _ =  http.NewRequest("POST", "/configuration", strings.NewReader("{}"))
			OperatorConfiguration(ctx)
			assert.Equal(t, tt.args.want, w.Result().StatusCode)
			var res = models.OperatorConfiguration{}
			_ = json.Unmarshal([]byte(w.Body.String()), &res)
			assert.Equal(t, tt.args.DailyLimit ,res.DailyLimit)
		})
	}
}
