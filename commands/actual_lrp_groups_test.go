package commands_test

import (
	"code.cloudfoundry.org/bbs/fake_bbs"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/cfdot/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("ActualLRPGroups", func() {
	var (
		fakeBBSClient   *fake_bbs.FakeClient
		actualLRPGroups []*models.ActualLRPGroup
		returnedError   error
		stdout, stderr  *gbytes.Buffer
	)

	BeforeEach(func() {
		stdout = gbytes.NewBuffer()
		stderr = gbytes.NewBuffer()
		fakeBBSClient = &fake_bbs.FakeClient{}
	})

	JustBeforeEach(func() {
		fakeBBSClient.ActualLRPGroupsReturns(actualLRPGroups, returnedError)
	})

	Context("when the bbs responds with actual lrp groups", func() {
		BeforeEach(func() {
			actualLRPGroups = []*models.ActualLRPGroup{
				{
					Instance: &models.ActualLRP{
						State: "running",
					},
				},
			}
		})

		It("prints a json stream of all the actual lrp groups", func() {
			err := commands.ActualLRPGroups(stdout, stderr, fakeBBSClient, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(stdout).To(gbytes.Say(`"state":"running"`))
		})
	})

	Context("when the bbs errors", func() {
		BeforeEach(func() {
			returnedError = models.ErrUnknownError
		})

		It("fails with a relevant error", func() {
			err := commands.ActualLRPGroups(stdout, stderr, fakeBBSClient, nil)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(models.ErrUnknownError))
		})
	})
})
