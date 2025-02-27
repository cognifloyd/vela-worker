// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"strings"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/pipeline"
)

// Skip creates the ruledata from the build and repository
// information and returns true if the data does not match
// the ruleset for the given container.
func Skip(c *pipeline.Container, b *library.Build, r *library.Repo) bool {
	// check if the container provided is empty
	if c == nil {
		return true
	}

	event := b.GetEvent()
	action := b.GetEventAction()

	// if the build has an event action, concatenate event and event action for matching
	if !strings.EqualFold(action, "") {
		event = event + ":" + action
	}

	// create ruledata from build and repository information
	//
	// https://pkg.go.dev/github.com/go-vela/types/pipeline#RuleData
	ruledata := &pipeline.RuleData{
		Branch:   b.GetBranch(),
		Event:    event,
		Repo:     r.GetFullName(),
		Status:   b.GetStatus(),
		Parallel: c.Ruleset.If.Parallel,
	}

	// check if the build event is tag
	if strings.EqualFold(b.GetEvent(), constants.EventTag) {
		// add tag information to ruledata with refs/tags prefix removed
		ruledata.Tag = strings.TrimPrefix(b.GetRef(), "refs/tags/")
	}

	// check if the build event is deployment
	if strings.EqualFold(b.GetEvent(), constants.EventDeploy) {
		// handle when deployment event is for a tag
		if strings.HasPrefix(b.GetRef(), "refs/tags/") {
			// add tag information to ruledata with refs/tags prefix removed
			ruledata.Tag = strings.TrimPrefix(b.GetRef(), "refs/tags/")
		}
		// add deployment target information to ruledata
		ruledata.Target = b.GetDeploy()
	}

	// return the inverse of container execute
	//
	// https://pkg.go.dev/github.com/go-vela/types/pipeline#Container.Execute
	return !c.Execute(ruledata)
}
