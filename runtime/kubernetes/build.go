// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package kubernetes

import (
	"context"
	"fmt"

	"github.com/go-vela/types/pipeline"

	"github.com/buildkite/yaml"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InspectBuild displays details about the pod for the init step.
func (c *client) InspectBuild(ctx context.Context, b *pipeline.Build) ([]byte, error) {
	c.Logger.Tracef("inspecting build pod for pipeline %s", b.ID)

	output := []byte(fmt.Sprintf("> Inspecting pod for pipeline %s", b.ID))

	// TODO: The environment gets populated in AssembleBuild, after InspectBuild runs.
	//       But, we should make sure that secrets can't be leaked here anyway.
	buildOutput, err := yaml.Marshal(c.Pod)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to serialize pod: %w", err)
	}

	output = append(output, buildOutput...)

	// TODO: make other k8s Inspect* funcs no-ops (prefer this method):
	// 	     InspectVolume, InspectImage, InspectNetwork
	return output, nil
}

// SetupBuild prepares the pod metadata for the pipeline build.
func (c *client) SetupBuild(ctx context.Context, b *pipeline.Build) error {
	c.Logger.Tracef("setting up for build %s", b.ID)

	// TODO: name from c.config and options. Allow a filename instead to load from filesystem.
	podsTemplateResponse, err := c.VelaKubernetes.VelaV1alpha1().PipelinePodsTemplates(c.config.Namespace).Get(
		context.Background(), "asdf", metav1.GetOptions{},
	)
	if err != nil {
		return err
	}

	// save the podTemplate to use later in SetupContainer and other Setup methods
	c.podTemplate = &podsTemplateResponse.Spec.Template

	// These labels will be used to call k8s watch APIs.
	labels := map[string]string{
		"pipeline":        b.ID,
		"worker":          c.config.WorkerHostname,
		"worker-flavor":   b.Worker.Flavor,
		"worker-platform": b.Worker.Platform,
	}

	if c.podTemplate.Meta.Labels != nil {
		// merge the template labels into the worker-defined labels.
		for k, v := range c.podTemplate.Meta.Labels {
			// do not allow overwriting any of the worker-defined labels.
			if _, ok := labels[k]; ok {
				continue
			}
			labels[k] = v
		}
	}

	// create the object metadata for the pod
	//
	// https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1?tab=doc#ObjectMeta
	c.Pod.ObjectMeta = metav1.ObjectMeta{
		Name:        b.ID,
		Labels:      labels,
		Annotations: c.podTemplate.Meta.Annotations,
	}

	// TODO: Vela admin defined worker-specific: AutomountServiceAccountToken

	if c.podTemplate.Spec.NodeSelector != nil {
		c.Pod.Spec.NodeSelector = c.podTemplate.Spec.NodeSelector
	}
	if c.podTemplate.Spec.Tolerations != nil {
		c.Pod.Spec.Tolerations = c.podTemplate.Spec.Tolerations
	}
	if c.podTemplate.Spec.Affinity != nil {
		c.Pod.Spec.Affinity = c.podTemplate.Spec.Affinity
	}

	// create the restart policy for the pod
	//
	// https://pkg.go.dev/k8s.io/api/core/v1?tab=doc#RestartPolicy
	c.Pod.Spec.RestartPolicy = v1.RestartPolicyNever

	if c.podTemplate.Spec.DNSPolicy != nil {
		c.Pod.Spec.DNSPolicy = *c.podTemplate.Spec.DNSPolicy
	}
	if c.podTemplate.Spec.DNSConfig != nil {
		c.Pod.Spec.DNSConfig = c.podTemplate.Spec.DNSConfig
	}

	if c.podTemplate.Spec.SecurityContext != nil {
		if c.Pod.Spec.SecurityContext == nil {
			c.Pod.Spec.SecurityContext = &v1.PodSecurityContext{}
		}
		if c.podTemplate.Spec.SecurityContext.RunAsNonRoot != nil {
			c.Pod.Spec.SecurityContext.RunAsNonRoot = c.podTemplate.Spec.SecurityContext.RunAsNonRoot
		}
		if c.podTemplate.Spec.SecurityContext.Sysctls != nil {
			c.Pod.Spec.SecurityContext.Sysctls = c.podTemplate.Spec.SecurityContext.Sysctls
		}
	}

	return nil
}

// AssembleBuild finalizes the pipeline build setup.
// This creates the pod in kubernetes for the pipeline build.
// After creation, image is the only container field we can edit in kubernetes.
// So, all environment, volume, and other container metadata must be setup
// before running AssembleBuild.
func (c *client) AssembleBuild(ctx context.Context, b *pipeline.Build) error {
	c.Logger.Tracef("assembling build %s", b.ID)

	var err error

	// last minute Environment setup
	for _, _service := range b.Services {
		err = c.setupContainerEnvironment(_service)
		if err != nil {
			return err
		}
	}

	for _, _stage := range b.Stages {
		// TODO: remove hardcoded reference
		if _stage.Name == "init" {
			continue
		}

		for _, _step := range _stage.Steps {
			err = c.setupContainerEnvironment(_step)
			if err != nil {
				return err
			}
		}
	}

	for _, _step := range b.Steps {
		// TODO: remove hardcoded reference
		if _step.Name == "init" {
			continue
		}

		err = c.setupContainerEnvironment(_step)
		if err != nil {
			return err
		}
	}

	for _, _secret := range b.Secrets {
		if _secret.Origin.Empty() {
			continue
		}

		err = c.setupContainerEnvironment(_secret.Origin)
		if err != nil {
			return err
		}
	}

	// If the api call to create the pod fails, the pod might
	// partially exist. So, set this first to make sure all
	// remnants get deleted.
	c.createdPod = true

	// TODO: create pipeline pod event watcher
	// RunContainer (image error watch): c.Kubernetes.CoreV1().Events(c.config.Namespace).Watch(context.Background(), opts)
	// WaitContainer: c.Kubernetes.CoreV1().Pods(c.config.Namespace).Watch(context.Background(), opts)
	// TailContainer: c.Kubernetes.CoreV1().Pods(c.config.Namespace).GetLogs(c.Pod.ObjectMeta.Name, opts).Stream(context.Background())

	c.Logger.Infof("creating pod %s", c.Pod.ObjectMeta.Name)
	// send API call to create the pod
	//
	// https://pkg.go.dev/k8s.io/client-go/kubernetes/typed/core/v1?tab=doc#PodInterface
	// nolint: contextcheck // ignore non-inherited new context
	_, err = c.Kubernetes.CoreV1().
		Pods(c.config.Namespace).
		Create(context.Background(), c.Pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

// RemoveBuild deletes (kill, remove) the pipeline build metadata.
// This deletes the kubernetes pod.
func (c *client) RemoveBuild(ctx context.Context, b *pipeline.Build) error {
	c.Logger.Tracef("removing build %s", b.ID)

	if !c.createdPod {
		// nothing to do
		return nil
	}

	// create variables for the delete options
	//
	// This is necessary because the delete options
	// expect all values to be passed by reference.
	var (
		period = int64(0)
		policy = metav1.DeletePropagationForeground
	)

	// create options for removing the pod
	//
	// https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1?tab=doc#DeleteOptions
	opts := metav1.DeleteOptions{
		GracePeriodSeconds: &period,
		// https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1?tab=doc#DeletionPropagation
		PropagationPolicy: &policy,
	}

	c.Logger.Infof("removing pod %s", c.Pod.ObjectMeta.Name)
	// send API call to delete the pod
	// nolint: contextcheck // ignore non-inherited new context
	err := c.Kubernetes.CoreV1().
		Pods(c.config.Namespace).
		Delete(context.Background(), c.Pod.ObjectMeta.Name, opts)
	if err != nil {
		return err
	}

	c.Pod = &v1.Pod{}
	c.createdPod = false

	return nil
}
