package util

import (
	"reflect"
	"testing"
)

func TestMapSliceMixed2MapSliceString(t *testing.T) {
	params := map[string]interface{}{
		"keyInt":         10,
		"keyFloat":       3.2,
		"keyStr":         "valueStr",
		"keySliceMixed":  []interface{}{10, 3.2, "valueStr", "3.2"},
		"keySliceInt":    []int{10, 32},
		"keySliceFloat":  []float64{10, 3.2},
		"keySliceString": []string{"valueStr", "3.2"},
	}
	dataCorrect := map[string][]string{
		"keyInt":         []string{"10"},
		"keyFloat":       []string{"3.2"},
		"keyStr":         []string{"valueStr"},
		"keySliceMixed":  []string{"10", "3.2", "valueStr", "3.2"},
		"keySliceInt":    []string{"10", "32"},
		"keySliceFloat":  []string{"10", "3.2"},
		"keySliceString": []string{"valueStr", "3.2"},
	}

	data := MapSliceMixed2MapSliceString(params)
	t.Logf("origin:%#v\n", params)
	t.Logf("data:%#v\n", data)

	if !reflect.DeepEqual(data, dataCorrect) {
		t.Logf(`!reflect.DeepEqual(data, dataCorrect)`)
	}
}
