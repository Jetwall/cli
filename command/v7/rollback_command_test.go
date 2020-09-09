package v7_test

import (
	"errors"
	"fmt"
	"time"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/command/v7/shared/sharedfakes"
	"code.cloudfoundry.org/cli/command/v7/v7fakes"
	"code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = FDescribe("rollback Command", func() {
	var (
		app             string
		binaryName      string
		executeErr      error
		fakeActor       *v7fakes.FakeRollbackActor
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		input           *Buffer
		testUI          *ui.UI

		fakeAppStager *sharedfakes.FakeAppStager
		cmd           v7.RollbackCommand
	)

	BeforeEach(func() {
		app = "some-app"
		binaryName = "faceman"
		fakeActor = new(v7fakes.FakeRollbackActor)
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeAppStager = new(sharedfakes.FakeAppStager)
		input = NewBuffer()
		testUI = ui.NewTestUI(input, NewBuffer(), NewBuffer())

		revisions := []resources.Revision{
			resources.Revision{Version: 1},
			resources.Revision{Version: 2},
		}

		fakeActor.GetRevisionsByApplicationNameAndSpaceReturns(
			revisions, v7action.Warnings{"warning-2"}, nil,
		)

		fakeConfig.BinaryNameReturns(binaryName)

		fakeConfig.CurrentUserReturns(configv3.User{Name: "steve"}, nil)

		fakeConfig.TargetedOrganizationReturns(configv3.Organization{
			Name: "some-org",
			GUID: "some-org-guid",
		})

		fakeConfig.TargetedSpaceReturns(configv3.Space{
			Name: "some-space",
			GUID: "some-space-guid",
		})

		cmd = v7.RollbackCommand{
			RequiredArgs: flag.AppName{AppName: app},
			Actor:        fakeActor,
			BaseCommand: v7.BaseCommand{
				UI:          testUI,
				Config:      fakeConfig,
				SharedActor: fakeSharedActor,
			},
			Stager: fakeAppStager,
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	It("displays the experimental warning", func() {
		Expect(testUI.Err).To(Say("This command is in EXPERIMENTAL stage and may change without notice"))
	})

	When("checking target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(actionerror.NoOrganizationTargetedError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(actionerror.NoOrganizationTargetedError{BinaryName: binaryName}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))

			checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrg).To(BeTrue())
			Expect(checkTargetedSpace).To(BeTrue())
		})
	})

	When("the user is not logged in", func() {
		var expectedErr error

		BeforeEach(func() {
			expectedErr = errors.New("some current user error")
			fakeConfig.CurrentUserReturns(configv3.User{}, expectedErr)
		})

		It("returns an error", func() {
			Expect(executeErr).To(Equal(expectedErr))
		})
	})

	// TODO: Missing app #174107413
	// When("getting the applications returns an error", func() {
	// 	var expectedErr error

	// 	BeforeEach(func() {
	// 		expectedErr = ccerror.RequestError{}
	// 		fakeActor.GetAppSummariesForSpaceReturns([]v7action.ApplicationSummary{}, v7action.Warnings{"warning-1", "warning-2"}, expectedErr)
	// 	})

	// 	It("returns the error and prints warnings", func() {
	// 		Expect(executeErr).To(Equal(ccerror.RequestError{}))

	// 		Expect(testUI.Out).To(Say(`Getting apps in org some-org / space some-space as steve\.\.\.`))

	// 		Expect(testUI.Err).To(Say("warning-1"))
	// 		Expect(testUI.Err).To(Say("warning-2"))
	// 	})
	// })

	When("an app has no revisions", func() {
		BeforeEach(func() {
			fakeActor.GetRevisionsByApplicationNameAndSpaceReturns([]resources.Revision{}, nil, nil)
		})

		It("displays an error saying that there are no revisions, and does not rollback", func() {
			Expect(executeErr).NotTo(BeNil())
			Expect(executeErr).To(MatchError(fmt.Sprintf("No revisions for app %s", app)))
		})
	})

	When("the first revision is set as the rollback target", func() {
		BeforeEach(func() {
			cmd.Version = flag.PositiveInteger{Value: 1}
		})

		When("the app has at least one revision", func() {
			BeforeEach(func() {
				fakeActor.GetApplicationByNameAndSpaceReturns(
					resources.Application{GUID: "123"},
					v7action.Warnings{"warning-1"},
					nil,
				)

				fakeActor.GetRevisionByApplicationAndVersionReturns(
					resources.Revision{Version: 1, GUID: "some-1-guid"},
					v7action.Warnings{"warning-3"},
					nil,
				)

				fakeActor.CreateDeploymentByApplicationAndRevisionReturns(
					"deployment-guid",
					v7action.Warnings{"warning-4"},
					nil,
				)
			})

			When("the user passes the force flag", func() {
				BeforeEach(func() {
					cmd.Force = true
				})

				It("skips the prompt and executes the rollback", func() {
					Expect(fakeActor.GetApplicationByNameAndSpaceCallCount()).To(Equal(1), "GetApplicationByNameAndSpace call count")
					appName, spaceGUID := fakeActor.GetApplicationByNameAndSpaceArgsForCall(0)
					Expect(appName).To(Equal(app))
					Expect(spaceGUID).To(Equal("some-space-guid"))

					Expect(fakeActor.GetRevisionByApplicationAndVersionCallCount()).To(Equal(1), "GetRevisionByApplicationAndVersion call count")
					appGUID, version := fakeActor.GetRevisionByApplicationAndVersionArgsForCall(0)
					Expect(appGUID).To(Equal("123"))
					Expect(version).To(Equal(1))

					Expect(fakeActor.CreateDeploymentByApplicationAndRevisionCallCount()).To(Equal(1), "CreateDeploymentByApplicationAndRevision call count")
					appGUID, revisionGUID := fakeActor.CreateDeploymentByApplicationAndRevisionArgsForCall(0)
					Expect(appGUID).To(Equal("123"))
					Expect(revisionGUID).To(Equal("some-1-guid"))

					Expect(testUI.Out).ToNot(Say("Rolling '%s' back to revision '1' will create a new revision. The new revision '3' will use the settings from revision '1'.", app))
					Expect(testUI.Out).ToNot(Say("Are you sure you want to continue?"))

					Expect(testUI.Out).To(Say("Rolling back to revision 1 for app some-app in org some-org / space some-space as steve..."))
					Expect(testUI.Err).To(Say("warning-1"))
					Expect(testUI.Err).To(Say("warning-2"))
					Expect(testUI.Err).To(Say("warning-3"))
					Expect(testUI.Err).To(Say("warning-4"))
					Expect(testUI.Out).To(Say("OK"))
				})

				Describe("staging logs", func() {
					When("there are no logging errors", func() {
						BeforeEach(func() {
							fakeActor.GetStreamingLogsForApplicationByNameAndSpaceStub = ReturnLogs(
								[]LogEvent{
									{Log: sharedaction.NewLogMessage("log-message-1", "OUT", time.Now(), sharedaction.StagingLog, "source-instance")},
									{Log: sharedaction.NewLogMessage("log-message-2", "OUT", time.Now(), sharedaction.StagingLog, "source-instance")},
									{Log: sharedaction.NewLogMessage("log-message-3", "OUT", time.Now(), "potato", "source-instance")},
								},
								v7action.Warnings{"log-warning-1", "log-warning-2"},
								nil,
							)
						})

						It("displays the staging logs and warnings", func() {
							Expect(testUI.Out).To(Say("Staging app and tracing logs..."))

							Expect(testUI.Err).To(Say("log-warning-1"))
							Expect(testUI.Err).To(Say("log-warning-2"))

							Eventually(testUI.Out).Should(Say("log-message-1"))
							Eventually(testUI.Out).Should(Say("log-message-2"))
							Eventually(testUI.Out).ShouldNot(Say("log-message-3"))

							Expect(fakeActor.GetStreamingLogsForApplicationByNameAndSpaceCallCount()).To(Equal(1))
							passedAppName, spaceGUID, _ := fakeActor.GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(0)
							Expect(passedAppName).To(Equal(app))
							Expect(spaceGUID).To(Equal("some-space-guid"))
						})
					})

					When("there are logging errors", func() {
						BeforeEach(func() {
							fakeActor.GetStreamingLogsForApplicationByNameAndSpaceStub = ReturnLogs(
								[]LogEvent{
									{Error: errors.New("some-random-err")},
									{Error: actionerror.LogCacheTimeoutError{}},
									{Log: sharedaction.NewLogMessage("log-message-1", "OUT", time.Now(), sharedaction.StagingLog, "source-instance")},
								},
								v7action.Warnings{"log-warning-1", "log-warning-2"},
								nil,
							)
						})

						It("displays the errors as warnings", func() {
							Expect(testUI.Out).To(Say("Staging app and tracing logs..."))

							Expect(testUI.Err).To(Say("log-warning-1"))
							Expect(testUI.Err).To(Say("log-warning-2"))
							Eventually(testUI.Err).Should(Say("Failed to retrieve logs from Log Cache: some-random-err"))
							Eventually(testUI.Err).Should(Say("timeout connecting to log server, no log will be shown"))

							Eventually(testUI.Out).Should(Say("log-message-1"))
						})
					})
				})

				When("when getting the application summary succeeds", func() {
					BeforeEach(func() {
						summary := v7action.DetailedApplicationSummary{
							ApplicationSummary: v7action.ApplicationSummary{
								Application:      resources.Application{},
								ProcessSummaries: v7action.ProcessSummaries{},
							},
							CurrentDroplet: resources.Droplet{},
						}
						fakeActor.GetDetailedAppSummaryReturnsOnCall(0, summary, v7action.Warnings{"app-1-summary-warning-1", "app-1-summary-warning-2"}, nil)
						fakeActor.GetDetailedAppSummaryReturnsOnCall(1, summary, v7action.Warnings{"app-2-summary-warning-1", "app-2-summary-warning-2"}, nil)
					})

					// TODO: Don't test the shared.AppSummaryDisplayer.AppDisplay method.
					// Use DI to pass in a new AppSummaryDisplayer to the Command instead.
					It("displays the app summary", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(fakeActor.GetDetailedAppSummaryCallCount()).To(Equal(2))
					})
				})

				When("getting the application summary fails", func() {
					BeforeEach(func() {
						fakeActor.GetDetailedAppSummaryReturns(
							v7action.DetailedApplicationSummary{},
							v7action.Warnings{"get-application-summary-warnings"},
							errors.New("get-application-summary-error"),
						)
					})

					It("does not display the app summary", func() {
						Expect(testUI.Out).ToNot(Say(`requested state:`))
					})

					It("returns the error from GetDetailedAppSummary", func() {
						Expect(executeErr).To(MatchError("get-application-summary-error"))
					})

					It("prints the warnings", func() {
						Expect(testUI.Err).To(Say("get-application-summary-warnings"))
					})
				})
			})

			When("user says yes to prompt", func() {
				BeforeEach(func() {
					_, err := input.Write([]byte("y\n"))
					Expect(err).NotTo(HaveOccurred())
				})

				It("successfully executes the command and outputs warnings", func() {
					Expect(fakeActor.GetApplicationByNameAndSpaceCallCount()).To(Equal(1), "GetApplicationByNameAndSpace call count")
					appName, spaceGUID := fakeActor.GetApplicationByNameAndSpaceArgsForCall(0)
					Expect(appName).To(Equal(app))
					Expect(spaceGUID).To(Equal("some-space-guid"))

					Expect(fakeActor.GetRevisionByApplicationAndVersionCallCount()).To(Equal(1), "GetRevisionByApplicationAndVersion call count")
					appGUID, version := fakeActor.GetRevisionByApplicationAndVersionArgsForCall(0)
					Expect(appGUID).To(Equal("123"))
					Expect(version).To(Equal(1))

					Expect(fakeActor.CreateDeploymentByApplicationAndRevisionCallCount()).To(Equal(1), "CreateDeploymentByApplicationAndRevision call count")
					appGUID, revisionGUID := fakeActor.CreateDeploymentByApplicationAndRevisionArgsForCall(0)
					Expect(appGUID).To(Equal("123"))
					Expect(revisionGUID).To(Equal("some-1-guid"))

					Expect(testUI.Out).To(Say("Rolling '%s' back to revision '1' will create a new revision. The new revision '3' will use the settings from revision '1'.", app))
					Expect(testUI.Out).To(Say("Are you sure you want to continue?"))
					Expect(testUI.Out).To(Say("Rolling back to revision 1 for app some-app in org some-org / space some-space as steve..."))
					Expect(testUI.Err).To(Say("warning-1"))
					Expect(testUI.Err).To(Say("warning-2"))
					Expect(testUI.Err).To(Say("warning-3"))
					Expect(testUI.Err).To(Say("warning-4"))
					Expect(testUI.Out).To(Say("OK"))
				})
			})

			When("user says no to prompt", func() {
				BeforeEach(func() {
					_, err := input.Write([]byte("n\n"))
					Expect(err).NotTo(HaveOccurred())
				})

				It("does not execute the command and outputs warnings", func() {
					Expect(fakeActor.GetApplicationByNameAndSpaceCallCount()).To(Equal(1), "GetApplicationByNameAndSpace call count")
					appName, spaceGUID := fakeActor.GetApplicationByNameAndSpaceArgsForCall(0)
					Expect(appName).To(Equal(app))
					Expect(spaceGUID).To(Equal("some-space-guid"))

					Expect(fakeActor.GetRevisionByApplicationAndVersionCallCount()).To(Equal(1), "GetRevisionByApplicationAndVersion call count")
					appGUID, version := fakeActor.GetRevisionByApplicationAndVersionArgsForCall(0)
					Expect(appGUID).To(Equal("123"))
					Expect(version).To(Equal(1))

					Expect(fakeActor.CreateDeploymentByApplicationAndRevisionCallCount()).To(Equal(0), "CreateDeploymentByApplicationAndRevision call count")

					Expect(testUI.Out).To(Say("Rolling '%s' back to revision '1' will create a new revision. The new revision '3' will use the settings from revision '1'.", app))
					Expect(testUI.Out).To(Say("App '%s' has not been rolled back to revision '1'.", app))
					Expect(testUI.Err).To(Say("warning-1"))
					Expect(testUI.Err).To(Say("warning-2"))
					Expect(testUI.Err).To(Say("warning-3"))
				})
			})

			When("the user chooses the default", func() {
				BeforeEach(func() {
					_, err := input.Write([]byte("\n"))
					Expect(err).NotTo(HaveOccurred())
				})

				It("cancels the rollback", func() {
					Expect(executeErr).NotTo(HaveOccurred())

					Expect(testUI.Out).To(Say("App '%s' has not been rolled back to revision '1'.", app))
					Expect(fakeActor.CreateDeploymentByApplicationAndRevisionCallCount()).To(Equal(0))
				})
			})
		})
	})
})
