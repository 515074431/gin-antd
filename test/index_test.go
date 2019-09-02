package test

import (
	"github.com/515074431/gin-antd/routers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexGetRouter(t *testing.T)  {
	r := routers.SetupRouter()
	w:=httptest.NewRecorder()
	req,_:= http.NewRequest(http.MethodGet,"/",nil)
	r.ServeHTTP(w,req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t,"Hello gin", w.Body.String())
}

