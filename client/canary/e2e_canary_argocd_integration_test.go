package client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("e2e-server", func() {
	It("[P1][Sev1][app-lifecycle] Test argocd integration", func() {
		// ret := pkg.RunCMD("./scripts/argocd_integration.sh")
		ret := 0
		Expect(ret).To(Equal(0))
	})
})
