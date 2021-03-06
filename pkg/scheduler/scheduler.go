/*
Copyright 2014 Google Inc. All rights reserved.

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

package scheduler

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

// Scheduler is an interface implemented by things that know how to schedule pods
// onto machines.
type Scheduler interface {
	Schedule(api.Pod, MinionLister) (selectedMachine string, err error)
}

//hack: creating a new type so that existing schedulers aren't broken.
//a better way to handle this is necessary.
//See : https://github.com/GoogleCloudPlatform/kubernetes/issues/1517
type StatefulScheduler interface {
	Scheduler
	Delete(id string) error
}
