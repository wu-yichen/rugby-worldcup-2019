package api_test

import (
	"errors"

	"github.com/followme1987/wcwhen/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/followme1987/wcwhen/fakes"
)

var _ = Describe("Api", func() {
	It("should retrieve data from API server", func() {
		expected := []byte("result")
		fake := fakes.FakeScraper{GetDataFunc: func() ([]byte, error) {
			result := expected
			return result, nil
		}}

		actual, err := api.GetMatches(fake)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	It("should return error with error message if fail to retrieve data from API server", func() {
		fake := fakes.FakeScraper{GetDataFunc: func() ([]byte, error) {
			return nil, errors.New("server is down")
		}}
		_, err := api.GetMatches(fake)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("server is down"))
	})
})
