package command_registry_test

import (
	. "github.com/cloudfoundry/cli/cf/command_registry/fake_command"

	. "github.com/cloudfoundry/cli/cf/command_registry"

	. "github.com/cloudfoundry/cli/cf/i18n"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CommandRegistry", func() {
	Context("i18n", func() {
		It("initialize i18n T() func", func() {
			Ω(T).ToNot(BeNil())
		})
	})

	Context("Register()", func() {
		It("registers a command into the Command Registry map", func() {
			Ω(Commands.CommandExists("fake-command2")).To(BeFalse())

			Register(FakeCommand2{})

			Ω(Commands.CommandExists("fake-command2")).To(BeTrue())
		})
	})

	Describe("Commands", func() {
		Context("CommandExists()", func() {
			It("returns true the command exists in the list", func() {
				Ω(Commands.CommandExists("fake-command")).To(BeTrue())
			})

			It("returns false if the command doesn't exists in the list", func() {
				Ω(Commands.CommandExists("non-exist-cmd")).To(BeFalse())
			})
		})

		Context("FindCommand()", func() {
			It("returns the command interface when found", func() {
				cmd := Commands.FindCommand("fake-command")
				Ω(cmd.MetaData().Usage).To(Equal("Usage of fake-command"))
				Ω(cmd.MetaData().Description).To(Equal("Description for fake-command"))
			})
		})

		Context("SetCommand()", func() {
			It("replaces the command in registry with command provided", func() {
				updatedCmd := FakeCommand1{Data: "This is new data"}
				oldCmd := Commands.FindCommand("fake-command")
				Ω(oldCmd).ToNot(Equal(updatedCmd))

				Commands.SetCommand(updatedCmd)
				oldCmd = Commands.FindCommand("fake-command")
				Ω(oldCmd).To(Equal(updatedCmd))
			})

		})
	})

})
