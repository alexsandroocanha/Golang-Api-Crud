package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_JSONTags(t *testing.T) {
	p := Product{ID: 1, Name: "Mouse", Price: 99.90}
	b, err := json.Marshal(p)

	assert.NoError(t, err)
	assert.JSONEq(t, `{"id_product":1,"name":"Mouse","price":99.9}`, string(b))
}
