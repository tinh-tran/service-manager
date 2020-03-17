package osb

import (
	"github.com/Peripli/service-manager/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"log"
	"net/url"
)

var _ = Describe("OSB Controller test", func() {

	var clientKey = `-----BEGIN PRIVATE KEY-----
  MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDaO3W1LP5M20sF
  fnPI+s3pqVRPbnHe5TiepguuMLqcM4HS6eJz5/IimmILUexCLZ83WZYOcAGFqRNR
  zLUrhbOH62RK+U8JvaB4JA/rFzXOQ698RDzAVo7ZFhiHGO3o1Y27icdfF2ps2MZX
  CY6UxK1x7P1ZYXds4gefJQaqiZrIcuwfb97+hIlgVwYh6k3AkBQMqL0gb/vZhH+I
  2BdHZCIDDTbenSwegP0IIyneg75IQDVQJydzR4i5JswXclgofpLi5A7s/WXCIhg/
  B5ODlKpbU+ziV1lrSnmEPtPU6UKh9iJGaukHbJMHFo0z3B0MfGOUxpPEaV3n6Rpg
  Zp8Kkw9tAgMBAAECggEADmKc/7RXjvllmJcdSsI9kIl45UOCfg7eDJclbfYIVwOO
  Kzj/lGRVsbI7hEOCL1qShDODkLARaZ4bh+jWiGfnza3WjpqgeyPk0AaQhg6hnVcY
  2jglSQhroiOyujUKea6aCSKr4bjJayNe753RqDzOshPNH3ctSCAeIH9wUQ2BBnVt
  r+kbDbrXr834/XkOB73r85CX+d690THYu1CdqT6OAju74u+19gC1UhCn95F/Sa/v
  ej0TEC8EqUOcsRpfjUEDx6Ywwr5RrblzEaS997IX0LZb21/8g6qVUbq+oKZtvyOe
  P/cq17cuvr3pvKhrGi5uOlX1JPan969jQqe7xxJkgQKBgQDxsRHAP4y1tgjMaaoM
  dkFA4WoQ+ilevIVtfxO5beb89ooANcjlX7lnGcAL1WAjoa5lr05No2pHudE1ivdt
  eJgER13/BtgVjz65E5gzYEjDqcle2K/3LVIOslvFk0M9umflZdjlz0fl1IlutcMd
  4lZ3cS8zJPVCR+D2CyDH+IxDuwKBgQDnJtqgPLjtlRxmjXaQ4zMcFpcdTqXr0Ddf
  AaDA0iC+mihNfmxmpMjPmpiROO8MtrvvQy9Kzdg53ptl6e2alzcWoJ8p6vDnMFUa
  PuIf3/lw9PfLkImGDiRyn6GcwGaWluHUXlphJvTi5A2Ql5HmW9XHlwRx/ZGur7bm
  lRhsAB7C9wKBgQDgzC0SfwlFSdbNKcp8ZNE0o3Sf7c3ky7veqD+UTOB3kGey4lPE
  5E/x0UWKvB/7hDpNYcyW8dO8etxXzLVuIKhj8m0+8wKwqtdQFSWPQ5LqSlV93lVs
  tb6I5OPu1JXKKELSXvRqa20YG6LoUi708LwzxBZ+n3Vu/KQEtTz8QfVUWQKBgQCA
  hZrzky+jcc/7uVYeUyU8zdaxxeP9TKUs3wPZkjwAnkggZlWxcJfyzltcC5Lmt8eg
  zfNCnVdHPd2becjRtpg7rY0xyl6tvLLkx+gEnwzbYGlStwewELb1QIqkVFn2Cuh/
  owKPmBB7AyADsDLAKXmg4vfmxX016p9Ab8/HZP21mwKBgC8MRd0XwV3sCR1PmW3O
  vrlZ+SswoN/7pwRTRWX/S0AHjYBJ+Bn25p3v4R1PcaESpYzYnuDWAgIl2uqncdeA
  IlSzIMiKfxuJIOzpJHEdwhVmFriEIIrwdAA0jPMLeXGFTmI/vWCmcG9F5XGwlRRJ
  gB9ceWwBvf0HhZYfJ3XCJZXM
-----END PRIVATE KEY-----`

	var cret = `-----BEGIN CERTIFICATE-----
MIICrDCCAZQCCQCziU7at44ipjANBgkqhkiG9w0BAQUFADAYMRYwFAYDVQQDDA1G
aXJzdCBNLiBMYXN0MB4XDTIwMDIyMjEyNDMxMloXDTIwMDMyMzEyNDMxMlowGDEW
MBQGA1UEAwwNRmlyc3QgTS4gTGFzdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBANo7dbUs/kzbSwV+c8j6zempVE9ucd7lOJ6mC64wupwzgdLp4nPn8iKa
YgtR7EItnzdZlg5wAYWpE1HMtSuFs4frZEr5Twm9oHgkD+sXNc5Dr3xEPMBWjtkW
GIcY7ejVjbuJx18XamzYxlcJjpTErXHs/Vlhd2ziB58lBqqJmshy7B9v3v6EiWBX
BiHqTcCQFAyovSBv+9mEf4jYF0dkIgMNNt6dLB6A/QgjKd6DvkhANVAnJ3NHiLkm
zBdyWCh+kuLkDuz9ZcIiGD8Hk4OUqltT7OJXWWtKeYQ+09TpQqH2IkZq6QdskwcW
jTPcHQx8Y5TGk8RpXefpGmBmnwqTD20CAwEAATANBgkqhkiG9w0BAQUFAAOCAQEA
QcLPaEwZ6EYoY7aa4sOzkV4AENEkdLcz/DQOFns5LisFtCUbGoPufzs4ozn9Bngy
fTSrUqV/I5l7bQV18vhWH86OqBYiDrxZMaTIgySuzN3aXJCpsw4JP0rHZjrRjFDx
hpL8qoDDR9vDvjvqE2jlqXMPAe0DZEljRzG+EARODnaCEFFpzEkosQLlPSXyn51I
3ffwNHcPQQeZCknqJ9BI8a4JdEP1cZDdl6TPu1rsakFfCHSKCwrKa6blCZRxVvpd
qYxHGtKZSU5BCswd7c3r8SL5qzmAscmu6orqwzGsvLHAx3Y9OcF+7weDZdz2OB3p
OOzY8kGVInUs83tZOfMVjQ==
-----END CERTIFICATE-----`

	var brokerTLS types.ServiceBroker

	BeforeEach(func() {
		brokerTLS = types.ServiceBroker{
			Base: types.Base{
				ID:     "123",
				Labels: map[string][]string{},
				Ready:  true,
			},
			Name:      "tls-broker",
			BrokerURL: "url",
			Credentials: &types.Credentials{
				Basic: &types.Basic{
					Username: "user",
					Password: "pass",
				},
				TLS: &types.TLS{
					Certificate: cret,
					Key:         clientKey,
				},
			},
		}
	})

	Describe("test osb create proxy", func() {
		logger := logrus.Entry{}
		targetBrokerURL, err := url.Parse("http://example.com/proxy/")
		if err != nil {
			log.Fatal(err)
		}

		It("create proxy with tls should return a new reverse proxy with its own tls setting", func() {
			reverseProxy, _ := buildProxy(targetBrokerURL, &logger, &brokerTLS)
			Expect(reverseProxy).NotTo(Equal(nil))
			reverseProxy2, _ := buildProxy(targetBrokerURL, &logger, &brokerTLS)
			Expect(reverseProxy2.Transport == reverseProxy.Transport).To(Equal(false))
		})
	})
})
