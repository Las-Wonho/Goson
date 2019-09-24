package goson_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Las-Wonho/Goson/goson"
)

func TestSample(t *testing.T) {

	simpleJson := goson.New().SetElement(".data1", "value1").SetElement(".data2", "value2")

	assert.Equal(t, simpleJson.GetEliment("data1"), "value1", "json structure should accessible data")
}
