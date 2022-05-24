package arrays

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldMapItemsAndReturnNewTransformedArray(t *testing.T) {
	s := Of([]string{"1", "2", "3"})
	res := s.Map(func(i string) string { return fmt.Sprintf("%s-test", i) })

	assert.True(t, res.Every(func(i string) bool { return strings.Contains(i, "-test") }))
}

func TestItShouldFilterAndReturnOnlyElementsThatPredicateReturnedTrue(t *testing.T) {
	s := Of([]int{1, 2, 3, 4, 5, 6, 7, 9, 11})
	res := s.Filter(func(i int) bool { return i%2 == 0 })

	assert.Len(t, res.Value, 3)
}
