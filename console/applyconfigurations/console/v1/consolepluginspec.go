// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ConsolePluginSpecApplyConfiguration represents a declarative configuration of the ConsolePluginSpec type for use
// with apply.
type ConsolePluginSpecApplyConfiguration struct {
	DisplayName *string                                 `json:"displayName,omitempty"`
	Backend     *ConsolePluginBackendApplyConfiguration `json:"backend,omitempty"`
	Proxy       []ConsolePluginProxyApplyConfiguration  `json:"proxy,omitempty"`
	I18n        *ConsolePluginI18nApplyConfiguration    `json:"i18n,omitempty"`
}

// ConsolePluginSpecApplyConfiguration constructs a declarative configuration of the ConsolePluginSpec type for use with
// apply.
func ConsolePluginSpec() *ConsolePluginSpecApplyConfiguration {
	return &ConsolePluginSpecApplyConfiguration{}
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *ConsolePluginSpecApplyConfiguration) WithDisplayName(value string) *ConsolePluginSpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithBackend sets the Backend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Backend field is set to the value of the last call.
func (b *ConsolePluginSpecApplyConfiguration) WithBackend(value *ConsolePluginBackendApplyConfiguration) *ConsolePluginSpecApplyConfiguration {
	b.Backend = value
	return b
}

// WithProxy adds the given value to the Proxy field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Proxy field.
func (b *ConsolePluginSpecApplyConfiguration) WithProxy(values ...*ConsolePluginProxyApplyConfiguration) *ConsolePluginSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithProxy")
		}
		b.Proxy = append(b.Proxy, *values[i])
	}
	return b
}

// WithI18n sets the I18n field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the I18n field is set to the value of the last call.
func (b *ConsolePluginSpecApplyConfiguration) WithI18n(value *ConsolePluginI18nApplyConfiguration) *ConsolePluginSpecApplyConfiguration {
	b.I18n = value
	return b
}
