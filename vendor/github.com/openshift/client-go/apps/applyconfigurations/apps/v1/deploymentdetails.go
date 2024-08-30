// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// DeploymentDetailsApplyConfiguration represents an declarative configuration of the DeploymentDetails type for use
// with apply.
type DeploymentDetailsApplyConfiguration struct {
	Message *string                             `json:"message,omitempty"`
	Causes  []DeploymentCauseApplyConfiguration `json:"causes,omitempty"`
}

// DeploymentDetailsApplyConfiguration constructs an declarative configuration of the DeploymentDetails type for use with
// apply.
func DeploymentDetails() *DeploymentDetailsApplyConfiguration {
	return &DeploymentDetailsApplyConfiguration{}
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *DeploymentDetailsApplyConfiguration) WithMessage(value string) *DeploymentDetailsApplyConfiguration {
	b.Message = &value
	return b
}

// WithCauses adds the given value to the Causes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Causes field.
func (b *DeploymentDetailsApplyConfiguration) WithCauses(values ...*DeploymentCauseApplyConfiguration) *DeploymentDetailsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCauses")
		}
		b.Causes = append(b.Causes, *values[i])
	}
	return b
}