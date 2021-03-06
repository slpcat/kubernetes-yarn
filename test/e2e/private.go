/*
Copyright 2015 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/golang/glog"
)

// A basic test to check the deployment of the
// container.cloud.google.com/_b_k8s_test/serve_hostname image
// with the TestBasicImage test. This test is only supported
// for the providers GCE and GKE.
func TestPrivate(c *client.Client) bool {
	if testContext.provider != "gce" && testContext.provider != "gke" {
		glog.Infof("Skipping test private which is only supported for providers gce and gke (not %s)", testContext.provider)
		return true
	}
	glog.Info("Calling out to TestBasic")
	return TestBasicImage(c, "container.cloud.google.com/_b_k8s_test/serve_hostname")
}
