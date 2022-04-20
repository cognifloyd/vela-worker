// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:operation POST /api/v1/shutdown system Shutdown
//
// Perform a soft shutdown of the worker
//
// ---
// produces:
// - application/json
// security:
//   - ApiKeyAuth: []
// responses:
//   '501':
//     description: Endpoint is not yet implemented
//     schema:
//       type: string

// Shutdown represents the API handler to shut down all
// executors currently running on a worker.
//
// This function performs a soft shutdown of a worker.
// Any build running during this time will safely complete, then
// the worker will safely shut itself down.
func Shutdown(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "This endpoint is not yet implemented")
}
