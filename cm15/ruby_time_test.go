package cm15_test

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc/cm15"
)

var _ = Describe("RubyTime", func() {
	Describe("MarshalJSON", func() {
		It("marshals to JSON", func() {
			rt := cm15.RubyTime{time.Date(2019, 2, 4, 23, 18, 40, 0, time.UTC)}
			Expect(json.Marshal(&rt)).To(BeEquivalentTo(`"2019/02/04 23:18:40 +0000"`))
		})
	})

	Describe("UnmarshalJSON", func() {
		It("unmarshals from JSON", func() {
			var rt cm15.RubyTime
			err := json.Unmarshal([]byte(`"2019/02/04 23:18:40 +0000"`), &rt)
			Expect(err).NotTo(HaveOccurred())
			Expect(rt.Time.In(time.UTC)).To(Equal(time.Date(2019, 2, 4, 23, 18, 40, 0, time.UTC)))
		})
	})
})
