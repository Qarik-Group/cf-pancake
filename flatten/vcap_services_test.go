package flatten_test

import (
	"github.com/cloudfoundry-community/go-cfenv"
	. "github.com/starkandwayne/cf-pancake/flatten"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flatten VCAP Services", func() {
	BeforeEach(func() {
	})
	It("does nothing for nil $VCAP_SERVICES", func() {
		Expect(VCAPServices(nil)).To(Equal(EnvVars{}))
	})
	It("does nothing for empty $VCAP_SERVICES", func() {
		Expect(VCAPServices(&cfenv.Services{})).To(Equal(EnvVars{}))
	})
	It("parse $VCAP_SERVICES with a service", func() {
		services := cfenv.Services{
			"elephantsql": []cfenv.Service{
				cfenv.Service{
					Name: "mydb",
					Credentials: map[string]interface{}{
						"uri": "postgres://hostname:port/dbname",
					},
				},
			},
		}
		flattened := VCAPServices(&services)
		Expect(flattened).ToNot(Equal(EnvVars{}))
		Expect(flattened["ELEPHANTSQL_URI"]).To(Equal("postgres://hostname:port/dbname"))
		Expect(flattened["MYDB_URI"]).To(Equal("postgres://hostname:port/dbname"))
	})
})
