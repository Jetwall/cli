package v7action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"io"
)

// Droplet represents a Cloud Controller droplet.
type Droplet struct {
	GUID       string
	State      constant.DropletState
	CreatedAt  string
	Stack      string
	Image      string
	Buildpacks []DropletBuildpack
}

type DropletBuildpack ccv3.DropletBuildpack

// CreateApplicationDroplet creates a new droplet without a package for the app with
// guid appGUID.
func (actor Actor) CreateApplicationDroplet(appGUID string) (Droplet, Warnings, error) {
	ccDroplet, warnings, err := actor.CloudControllerClient.CreateDroplet(appGUID)

	return actor.convertCCToActorDroplet(ccDroplet), Warnings(warnings), err
}

// SetApplicationDropletByApplicationNameAndSpace sets the droplet for an application.
func (actor Actor) SetApplicationDropletByApplicationNameAndSpace(appName string, spaceGUID string, dropletGUID string) (Warnings, error) {
	allWarnings := Warnings{}
	application, warnings, err := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return allWarnings, err
	}
	_, apiWarnings, err := actor.CloudControllerClient.SetApplicationDroplet(application.GUID, dropletGUID)
	actorWarnings := Warnings(apiWarnings)
	allWarnings = append(allWarnings, actorWarnings...)

	if newErr, ok := err.(ccerror.UnprocessableEntityError); ok {
		return allWarnings, actionerror.AssignDropletError{Message: newErr.Message}
	}

	return allWarnings, err
}

func (actor Actor) SetApplicationDroplet(appGUID string, dropletGUID string) (Warnings, error) {
	_, warnings, err := actor.CloudControllerClient.SetApplicationDroplet(appGUID, dropletGUID)

	if newErr, ok := err.(ccerror.UnprocessableEntityError); ok {
		return Warnings(warnings), actionerror.AssignDropletError{Message: newErr.Message}
	}

	return Warnings(warnings), err
}

// GetApplicationDroplets returns the list of droplets that belong to application.
func (actor Actor) GetApplicationDroplets(appName string, spaceGUID string) ([]Droplet, Warnings, error) {
	allWarnings := Warnings{}
	application, warnings, err := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return nil, allWarnings, err
	}

	ccv3Droplets, apiWarnings, err := actor.CloudControllerClient.GetDroplets(
		ccv3.Query{Key: ccv3.AppGUIDFilter, Values: []string{application.GUID}},
	)
	actorWarnings := Warnings(apiWarnings)
	allWarnings = append(allWarnings, actorWarnings...)
	if err != nil {
		return nil, allWarnings, err
	}

	var droplets []Droplet
	for _, ccv3Droplet := range ccv3Droplets {
		droplets = append(droplets, actor.convertCCToActorDroplet(ccv3Droplet))
	}

	return droplets, allWarnings, err
}

func (actor Actor) GetCurrentDropletByApplication(appGUID string) (Droplet, Warnings, error) {
	droplet, warnings, err := actor.CloudControllerClient.GetApplicationDropletCurrent(appGUID)
	switch err.(type) {
	case ccerror.ApplicationNotFoundError:
		return Droplet{}, Warnings(warnings), actionerror.ApplicationNotFoundError{GUID: appGUID}
	case ccerror.DropletNotFoundError:
		return Droplet{}, Warnings(warnings), actionerror.DropletNotFoundError{AppGUID: appGUID}
	}
	return actor.convertCCToActorDroplet(droplet), Warnings(warnings), err
}

func (actor Actor) UploadDroplet(dropletGUID string, dropletPath string, progressReader io.Reader, size int64) (Warnings, error) {
	var allWarnings Warnings

	jobURL, uploadWarnings, err := actor.CloudControllerClient.UploadDropletBits(dropletGUID, dropletPath, progressReader, size)
	allWarnings = append(allWarnings, uploadWarnings...)
	if err != nil {
		return allWarnings, err
	}

	jobWarnings, jobErr := actor.CloudControllerClient.PollJob(jobURL)
	allWarnings = append(allWarnings, jobWarnings...)
	if jobErr != nil {
		return allWarnings, jobErr
	}

	return allWarnings, nil
}

func (actor Actor) convertCCToActorDroplet(ccDroplet ccv3.Droplet) Droplet {
	var buildpacks []DropletBuildpack
	for _, ccBuildpack := range ccDroplet.Buildpacks {
		buildpacks = append(buildpacks, DropletBuildpack(ccBuildpack))
	}

	return Droplet{
		GUID:       ccDroplet.GUID,
		State:      ccDroplet.State,
		CreatedAt:  ccDroplet.CreatedAt,
		Stack:      ccDroplet.Stack,
		Buildpacks: buildpacks,
		Image:      ccDroplet.Image,
	}
}
