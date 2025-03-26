package bun_test

import (
	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tydanny/foodwheel/internal/app/foodwheel/cuisines/cuisineadapter/bun"
)

var _ = Describe("Config", Label("unit", "config"), func() {
	var config bun.Config

	BeforeEach(func() {
		config = bun.Config{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			Name:     "foodwheel",
			Params: map[string]string{
				"sslmode": "disable",
			},
		}
	})

	It("Should output a valid DSN", func() {
		config, err := pgxpool.ParseConfig(config.DSN())
		Expect(err).NotTo(HaveOccurred())
		Expect(config.ConnConfig.Host).To(Equal("localhost"))
		Expect(config.ConnConfig.Port).To(Equal(uint16(5432)))
		Expect(config.ConnConfig.User).To(Equal("postgres"))
		Expect(config.ConnConfig.Password).To(Equal("postgres"))
		Expect(config.ConnConfig.Database).To(Equal("foodwheel"))
		Expect(config.ConnConfig.TLSConfig).To(BeNil())
	})

	DescribeTable("Invalid Validate",
		func(expectedErr error, configMutator func(*bun.Config)) {
			configMutator(&config)
			Expect(config.Validate()).To(MatchError(expectedErr))
		},
		Entry("missing host", bun.ErrMissingHost, func(c *bun.Config) { c.Host = "" }),
		Entry("missing port", bun.ErrMissingPort, func(c *bun.Config) { c.Port = "" }),
		Entry("missing user", bun.ErrMissingUser, func(c *bun.Config) { c.User = "" }),
		Entry("missing name", bun.ErrMissingName, func(c *bun.Config) { c.Name = "" }),
	)
})
