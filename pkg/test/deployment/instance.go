//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package deployment

import (
	"github.com/hashicorp/go-multierror"

	"istio.io/istio/pkg/test/framework/scopes"
	"istio.io/istio/pkg/test/kube"

	kubeApiCore "k8s.io/api/core/v1"
)

// Instance represents an Istio deployment instance that has been performed by this test code.
type Instance struct {
	// The deployment namespace.
	namespace string

	// Path to the yaml file that is generated from the template.
	yamlFilePath string
}

// Deploy this deployment instance.
func (i *Instance) Deploy(a *kube.Accessor, wait bool) (err error) {
	scopes.CI.Infof("Applying Yaml file: %s", i.yamlFilePath)
	if err = a.Apply(i.namespace, i.yamlFilePath); err != nil {
		return multierror.Prefix(err, "kube apply of generated yaml filed:")
	}

	if wait {
		fetchFn := func() ([]kubeApiCore.Pod, error) {
			return a.GetPods(i.namespace)
		}
		_, err := fetchFn()
		if err != nil {
			scopes.CI.Errorf("Error retrieving pods in namespace: %s: %v", i.namespace, err)
			return err
		}
		if err := a.WaitUntilPodsAreReady(fetchFn); err != nil {
			scopes.CI.Errorf("Wait for Istio pods failed: %v", err)
			return err
		}
	}

	return nil
}

// Delete this deployment instance.
func (i *Instance) Delete(a *kube.Accessor, wait bool) (err error) {

	if err = a.Delete(i.namespace, i.yamlFilePath); err != nil {
		scopes.CI.Warnf("Error deleting deployment: %v", err)
	}

	if wait {
		// TODO: Just for waiting for deployment namespace deletion may not be enough. There are CRDs
		// and roles/rolebindings in other parts of the system as well. We should also wait for deletion of them.
		if e := a.WaitForNamespaceDeletion(i.namespace); e != nil {
			scopes.CI.Warnf("Error waiting for environment deletion: %v", e)
			err = multierror.Append(err, e)
		}
	}

	return
}
