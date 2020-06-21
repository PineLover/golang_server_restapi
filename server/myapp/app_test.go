package myapp

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func TestIndex(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK,resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))

}

func TestUsers(t * testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK,resp.StatusCode)
	data , _ := ioutil.ReadAll(resp.Body)

	//assert.Equal("Hello World",string(data))
	assert.Contains(string(data),"Get UserInfo")

}