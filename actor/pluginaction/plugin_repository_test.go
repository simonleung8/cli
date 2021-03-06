package pluginaction_test

import (
	"errors"
	"strings"

	. "code.cloudfoundry.org/cli/actor/pluginaction"
	"code.cloudfoundry.org/cli/actor/pluginaction/pluginactionfakes"
	"code.cloudfoundry.org/cli/api/plugin"
	"code.cloudfoundry.org/cli/util/configv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plugin Repository Actions", func() {
	var (
		actor            Actor
		fakeConfig       *pluginactionfakes.FakeConfig
		fakePluginClient *pluginactionfakes.FakePluginClient
	)

	BeforeEach(func() {
		fakeConfig = new(pluginactionfakes.FakeConfig)
		fakePluginClient = new(pluginactionfakes.FakePluginClient)
		actor = NewActor(fakeConfig, fakePluginClient)
	})

	Describe("AddPluginRepository", func() {
		var err error

		JustBeforeEach(func() {
			err = actor.AddPluginRepository("some-repo", "some-URL")
		})

		Context("when passed a url without a scheme", func() {
			It("prepends https://", func() {
				_ = actor.AddPluginRepository("some-repo2", "some-URL")
				url := fakePluginClient.GetPluginRepositoryArgsForCall(1)
				Expect(strings.HasPrefix(url, "https://")).To(BeTrue())
			})
		})

		Context("when passed an IP address with a port is passed without a scheme", func() {
			It("prepends https://", func() {
				_ = actor.AddPluginRepository("some-repo2", "127.0.0.1:5000")
				url := fakePluginClient.GetPluginRepositoryArgsForCall(1)
				Expect(strings.HasPrefix(url, "https://")).To(BeTrue())
			})
		})

		Context("when url ends with a trailing slash", func() {
			It("removes the trailing slash", func() {
				_ = actor.AddPluginRepository("some-repo2", "some-URL/")
				url := fakePluginClient.GetPluginRepositoryArgsForCall(1)
				Expect(strings.HasSuffix(url, "/")).To(BeFalse())
			})
		})

		Context("when the repository name is taken", func() {
			BeforeEach(func() {
				fakeConfig.PluginRepositoriesReturns([]configv3.PluginRepository{
					{
						Name: "repo-1",
						URL:  "https://URL-1",
					},
					{
						Name: "some-repo",
						URL:  "https://www.com",
					},
				})
			})

			It("returns the RepositoryNameTakenError", func() {
				Expect(err).To(MatchError(RepositoryNameTakenError{Name: "some-repo"}))
			})
		})

		Context("when the repository URL is taken", func() {
			BeforeEach(func() {
				fakeConfig.PluginRepositoriesReturns([]configv3.PluginRepository{
					{
						Name: "repo-1",
						URL:  "https://URL-1",
					},
					{
						Name: "repo-2",
						URL:  "https://some-URL",
					},
				})
			})

			It("returns the RepositoryURLTakenError", func() {
				Expect(err).To(MatchError(RepositoryURLTakenError{Name: "repo-2", URL: "https://some-URL"}))
			})
		})

		Context("when getting the repository errors", func() {
			BeforeEach(func() {
				fakePluginClient.GetPluginRepositoryReturns(plugin.PluginRepository{}, errors.New("generic-error"))
			})

			It("returns a 'AddPluginRepositoryError", func() {
				Expect(err).To(MatchError(AddPluginRepositoryError{
					Name:    "some-repo",
					URL:     "https://some-URL",
					Message: "generic-error",
				}))
			})
		})

		Context("when no errors occur", func() {
			BeforeEach(func() {
				fakePluginClient.GetPluginRepositoryReturns(plugin.PluginRepository{}, nil)
			})

			It("adds the repo to the config and returns nil", func() {
				Expect(err).ToNot(HaveOccurred())

				Expect(fakePluginClient.GetPluginRepositoryCallCount()).To(Equal(1))
				Expect(fakePluginClient.GetPluginRepositoryArgsForCall(0)).To(Equal("https://some-URL"))

				Expect(fakeConfig.AddPluginRepositoryCallCount()).To(Equal(1))
				repoName, repoURL := fakeConfig.AddPluginRepositoryArgsForCall(0)
				Expect(repoName).To(Equal("some-repo"))
				Expect(repoURL).To(Equal("https://some-URL"))
			})
		})
	})
})
