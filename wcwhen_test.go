package main_test

import (
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("[INTEGRATION] wcwhen", func() {
	It("should successfully run end to end", func() {
		command := exec.Command(pathToBin)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		Eventually(session, "1m").Should(gexec.Exit(0))

		expectedName := "wcwhen - A CLI for rugby world cup 2019"
		expectedToHaveVenue := "wcwhen [global options] command [command options] [arguments...]"
		expectedCommands := "team     team -name <teamName>"
		Expect(string(session.Out.Contents())).ToNot(BeEmpty())
		Expect(string(session.Out.Contents())).To(ContainSubstring(expectedName))
		Expect(string(session.Out.Contents())).To(ContainSubstring(expectedToHaveVenue))
		Expect(string(session.Out.Contents())).To(ContainSubstring(expectedCommands))
	})
})
