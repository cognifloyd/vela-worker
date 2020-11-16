// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"net/http"

	"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"
)

// register is a helper function to register
// the worker in the database, updating the item
// if the worker already exists
func (w *Worker) register(config *library.Worker) error {
	// check to see if the worker already exists in the database
	_, resp, err := w.VelaClient.Worker.Get(config.GetHostname())
	if err != nil {
		// check to see if the response was nil
		if resp == nil {
			return fmt.Errorf("unable to retrieve worker %s from the server: %v", config.GetHostname(), err)
		}
		// check to see if the worker was not found and if we need to add it
		if resp.StatusCode == http.StatusNotFound {
			logrus.Infof("registering worker %s with the server", config.GetHostname())
			_, _, err := w.VelaClient.Worker.Add(config)
			if err != nil {
				// log the error instead of returning so the operation doesn't block worker deployment
				return fmt.Errorf("unable to register worker %s with the server: %v", config.GetHostname(), err)
			}

			// successfully added the worker so return nil
			return nil
		}

		return fmt.Errorf("unable to retrieve worker %s from the server: %v", config.GetHostname(), err)
	}

	// the worker exists in the db, update it with the new config
	logrus.Infof("worker %s previously registered with server, updating information", config.GetHostname())
	_, _, err = w.VelaClient.Worker.Update(config.GetHostname(), config)
	if err != nil {
		// log the error instead of returning so the operation doesn't block worker deployment
		return fmt.Errorf("unable to update worker %s on the server: %v", config.GetHostname(), err)
	}

	return nil
}