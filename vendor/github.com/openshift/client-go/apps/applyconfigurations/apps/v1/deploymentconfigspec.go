// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	appsv1 "github.com/openshift/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// DeploymentConfigSpecApplyConfiguration represents an declarative configuration of the DeploymentConfigSpec type for use
// with apply.
type DeploymentConfigSpecApplyConfiguration struct {
	Strategy             *DeploymentStrategyApplyConfiguration `json:"strategy,omitempty"`
	MinReadySeconds      *int32                                `json:"minReadySeconds,omitempty"`
	Triggers             *appsv1.DeploymentTriggerPolicies     `json:"triggers,omitempty"`
	Replicas             *int32                                `json:"replicas,omitempty"`
	RevisionHistoryLimit *int32                                `json:"revisionHistoryLimit,omitempty"`
	Test                 *bool                                 `json:"test,omitempty"`
	Paused               *bool                                 `json:"paused,omitempty"`
	Selector             map[string]string                     `json:"selector,omitempty"`
	Template             *corev1.PodTemplateSpec               `json:"template,omitempty"`
}

// DeploymentConfigSpecApplyConfiguration constructs an declarative configuration of the DeploymentConfigSpec type for use with
// apply.
func DeploymentConfigSpec() *DeploymentConfigSpecApplyConfiguration {
	return &DeploymentConfigSpecApplyConfiguration{}
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithStrategy(value *DeploymentStrategyApplyConfiguration) *DeploymentConfigSpecApplyConfiguration {
	b.Strategy = value
	return b
}

// WithMinReadySeconds sets the MinReadySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReadySeconds field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithMinReadySeconds(value int32) *DeploymentConfigSpecApplyConfiguration {
	b.MinReadySeconds = &value
	return b
}

// WithTriggers sets the Triggers field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Triggers field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithTriggers(value appsv1.DeploymentTriggerPolicies) *DeploymentConfigSpecApplyConfiguration {
	b.Triggers = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithReplicas(value int32) *DeploymentConfigSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithRevisionHistoryLimit sets the RevisionHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RevisionHistoryLimit field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithRevisionHistoryLimit(value int32) *DeploymentConfigSpecApplyConfiguration {
	b.RevisionHistoryLimit = &value
	return b
}

// WithTest sets the Test field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Test field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithTest(value bool) *DeploymentConfigSpecApplyConfiguration {
	b.Test = &value
	return b
}

// WithPaused sets the Paused field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Paused field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithPaused(value bool) *DeploymentConfigSpecApplyConfiguration {
	b.Paused = &value
	return b
}

// WithSelector puts the entries into the Selector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Selector field,
// overwriting an existing map entries in Selector field with the same key.
func (b *DeploymentConfigSpecApplyConfiguration) WithSelector(entries map[string]string) *DeploymentConfigSpecApplyConfiguration {
	if b.Selector == nil && len(entries) > 0 {
		b.Selector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Selector[k] = v
	}
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *DeploymentConfigSpecApplyConfiguration) WithTemplate(value corev1.PodTemplateSpec) *DeploymentConfigSpecApplyConfiguration {
	b.Template = &value
	return b
}
