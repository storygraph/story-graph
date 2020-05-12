package storygraph_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/storygraph/story-graph/pkg/storygraph"
)

var _ = Describe("Weenie", func() {
	var weenie *Weenie

	BeforeEach(func() {
		weenie = NewWeenie(
			"Aragorn",
			"wounded",
			map[string]string{
				"faction": "fellowship of the ring",
			},
		)
	})

	Describe("Checking weenie equality", func() {
		Context("With different weenie", func() {
			It("should not be equal", func() {
				otherWeenie := NewWeenie(
					"saruman",
					"healthy",
					map[string]string{
						"faction":   "evil orc organization",
						"expertise": "dark magic",
					},
				)

				Expect(weenie.IsEqualTo(otherWeenie)).To(Equal(false))
			})
		})

		Context("With equal weenie", func() {
			It("should be equal", func() {
				otherWeenie := NewWeenie(
					"Aragorn",
					"wounded",
					map[string]string{
						"faction": "fellowship of the ring",
					},
				)

				Expect(weenie.IsEqualTo(otherWeenie)).To(Equal(true))
			})
		})
	})
})
