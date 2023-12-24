package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm__Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if !form.Valid() {
		t.Error("got invalid when should have valid")
	}
}

func TestForm__Required(t *testing.T) {
	// create a request using httptest
	r := httptest.NewRequest("POST", "/whatever", nil)
	// create form using blank object r.PostForm = {}
	form := New(r.PostForm)
	// then required filed a,b,c
	form.Required("a", "b", "c")
	// return false
	if form.Valid() {
		t.Error("got invalid when should have valid")
	}

	// another time check ready data in type url.Values{} and add a , b, c
	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	// create a request using httptest
	r = httptest.NewRequest("POST", "/whatever", nil)
	// then requst body override by postData
	r.PostForm = postData
	// // create form using postedData
	form = New(r.PostForm)
	// provide required field
	form.Required("a", "b", "c")

	// return true
	if !form.Valid() {
		t.Error("does not shows required field when it does ")
	}

}

func TestForm__MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("first_name", "saiful")

	r.PostForm = postedData

	form := New(r.PostForm)
	valid := form.MinLength("first_name", 4)

	if !valid {
		t.Error("first name field length less then 4")
	}

}

func TestForm__Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("first_name", "saiful")

	r.PostForm = postedData
	form := New(r.PostForm)

	isExist := form.Has("first_name")
	if !isExist {
		t.Error("first name does not blabk")
	}
}
