package forms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOffer(t *testing.T) {

	var form NewOfferForm

	form = NewOfferForm{Name: "aa"} // name too short
	assert.Error(t, form.Curate())

	form = NewOfferForm{Name: "aaa"} // name ok
	assert.NoError(t, form.Curate())

	form = NewOfferForm{Name: "aaaaaaaaaaaaaaaaaaaaa"} // name too long
	assert.Error(t, form.Curate())

}

func TestManyNewOffers(t *testing.T) {

	var tests_OK = []NewOfferForm{
		{Name: "aaa", Price: 1},
		{Name: "aaa", Price: 9999999999999999999999999999999999999},
	}
	for _, tt := range tests_OK {
		testname := fmt.Sprintf("%s", tt.Name)
		t.Run(testname, func(t *testing.T) {
			err := tt.Curate()
			if err != nil {
				t.Errorf("got %s %s", tt.Name, err.Error())
			}
		})
	}

	var tests_FAIL = []NewOfferForm{
		{Name: "aaa", Price: -1},
		{Name: "aaa", Price: 99999999999999999999999999999999999999},
	}
	for _, tt := range tests_FAIL {
		testname := fmt.Sprintf("%s", tt.Name)
		t.Run(testname, func(t *testing.T) {
			err := tt.Curate()
			if err == nil {
				t.Errorf("got %s %s", tt.Name, err.Error())
			}
		})
	}
}
