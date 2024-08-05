// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// OpenShiftAPIServerStatusApplyConfiguration represents a declarative configuration of the OpenShiftAPIServerStatus type for use
// with apply.
type OpenShiftAPIServerStatusApplyConfiguration struct {
	OperatorStatusApplyConfiguration `json:",inline"`
	LatestAvailableRevision          *int32 `json:"latestAvailableRevision,omitempty"`
}

// OpenShiftAPIServerStatusApplyConfiguration constructs a declarative configuration of the OpenShiftAPIServerStatus type for use with
// apply.
func OpenShiftAPIServerStatus() *OpenShiftAPIServerStatusApplyConfiguration {
	return &OpenShiftAPIServerStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *OpenShiftAPIServerStatusApplyConfiguration) WithObservedGeneration(value int64) *OpenShiftAPIServerStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *OpenShiftAPIServerStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *OpenShiftAPIServerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *OpenShiftAPIServerStatusApplyConfiguration) WithVersion(value string) *OpenShiftAPIServerStatusApplyConfiguration {
	b.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *OpenShiftAPIServerStatusApplyConfiguration) WithReadyReplicas(value int32) *OpenShiftAPIServerStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *OpenShiftAPIServerStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *OpenShiftAPIServerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.Generations = append(b.Generations, *values[i])
	}
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *OpenShiftAPIServerStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *OpenShiftAPIServerStatusApplyConfiguration {
	b.LatestAvailableRevision = &value
	return b
}
