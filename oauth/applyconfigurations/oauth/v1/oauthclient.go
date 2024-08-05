// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	oauthv1 "github.com/openshift/api/oauth/v1"
	internal "github.com/openshift/client-go/oauth/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// OAuthClientApplyConfiguration represents a declarative configuration of the OAuthClient type for use
// with apply.
type OAuthClientApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration       `json:",inline"`
	*v1.ObjectMetaApplyConfiguration    `json:"metadata,omitempty"`
	Secret                              *string                              `json:"secret,omitempty"`
	AdditionalSecrets                   []string                             `json:"additionalSecrets,omitempty"`
	RespondWithChallenges               *bool                                `json:"respondWithChallenges,omitempty"`
	RedirectURIs                        []string                             `json:"redirectURIs,omitempty"`
	GrantMethod                         *oauthv1.GrantHandlerType            `json:"grantMethod,omitempty"`
	ScopeRestrictions                   []ScopeRestrictionApplyConfiguration `json:"scopeRestrictions,omitempty"`
	AccessTokenMaxAgeSeconds            *int32                               `json:"accessTokenMaxAgeSeconds,omitempty"`
	AccessTokenInactivityTimeoutSeconds *int32                               `json:"accessTokenInactivityTimeoutSeconds,omitempty"`
}

// OAuthClient constructs a declarative configuration of the OAuthClient type for use with
// apply.
func OAuthClient(name string) *OAuthClientApplyConfiguration {
	b := &OAuthClientApplyConfiguration{}
	b.WithName(name)
	b.WithKind("OAuthClient")
	b.WithAPIVersion("oauth.openshift.io/v1")
	return b
}

// ExtractOAuthClient extracts the applied configuration owned by fieldManager from
// oAuthClient. If no managedFields are found in oAuthClient for fieldManager, a
// OAuthClientApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// oAuthClient must be a unmodified OAuthClient API object that was retrieved from the Kubernetes API.
// ExtractOAuthClient provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractOAuthClient(oAuthClient *oauthv1.OAuthClient, fieldManager string) (*OAuthClientApplyConfiguration, error) {
	return extractOAuthClient(oAuthClient, fieldManager, "")
}

// ExtractOAuthClientStatus is the same as ExtractOAuthClient except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractOAuthClientStatus(oAuthClient *oauthv1.OAuthClient, fieldManager string) (*OAuthClientApplyConfiguration, error) {
	return extractOAuthClient(oAuthClient, fieldManager, "status")
}

func extractOAuthClient(oAuthClient *oauthv1.OAuthClient, fieldManager string, subresource string) (*OAuthClientApplyConfiguration, error) {
	b := &OAuthClientApplyConfiguration{}
	err := managedfields.ExtractInto(oAuthClient, internal.Parser().Type("com.github.openshift.api.oauth.v1.OAuthClient"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(oAuthClient.Name)

	b.WithKind("OAuthClient")
	b.WithAPIVersion("oauth.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithKind(value string) *OAuthClientApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithAPIVersion(value string) *OAuthClientApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithName(value string) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithGenerateName(value string) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithNamespace(value string) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithUID(value types.UID) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithResourceVersion(value string) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithGeneration(value int64) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithCreationTimestamp(value metav1.Time) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *OAuthClientApplyConfiguration) WithLabels(entries map[string]string) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *OAuthClientApplyConfiguration) WithAnnotations(entries map[string]string) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *OAuthClientApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *OAuthClientApplyConfiguration) WithFinalizers(values ...string) *OAuthClientApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *OAuthClientApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSecret sets the Secret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Secret field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithSecret(value string) *OAuthClientApplyConfiguration {
	b.Secret = &value
	return b
}

// WithAdditionalSecrets adds the given value to the AdditionalSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalSecrets field.
func (b *OAuthClientApplyConfiguration) WithAdditionalSecrets(values ...string) *OAuthClientApplyConfiguration {
	for i := range values {
		b.AdditionalSecrets = append(b.AdditionalSecrets, values[i])
	}
	return b
}

// WithRespondWithChallenges sets the RespondWithChallenges field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RespondWithChallenges field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithRespondWithChallenges(value bool) *OAuthClientApplyConfiguration {
	b.RespondWithChallenges = &value
	return b
}

// WithRedirectURIs adds the given value to the RedirectURIs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RedirectURIs field.
func (b *OAuthClientApplyConfiguration) WithRedirectURIs(values ...string) *OAuthClientApplyConfiguration {
	for i := range values {
		b.RedirectURIs = append(b.RedirectURIs, values[i])
	}
	return b
}

// WithGrantMethod sets the GrantMethod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GrantMethod field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithGrantMethod(value oauthv1.GrantHandlerType) *OAuthClientApplyConfiguration {
	b.GrantMethod = &value
	return b
}

// WithScopeRestrictions adds the given value to the ScopeRestrictions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ScopeRestrictions field.
func (b *OAuthClientApplyConfiguration) WithScopeRestrictions(values ...*ScopeRestrictionApplyConfiguration) *OAuthClientApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithScopeRestrictions")
		}
		b.ScopeRestrictions = append(b.ScopeRestrictions, *values[i])
	}
	return b
}

// WithAccessTokenMaxAgeSeconds sets the AccessTokenMaxAgeSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AccessTokenMaxAgeSeconds field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithAccessTokenMaxAgeSeconds(value int32) *OAuthClientApplyConfiguration {
	b.AccessTokenMaxAgeSeconds = &value
	return b
}

// WithAccessTokenInactivityTimeoutSeconds sets the AccessTokenInactivityTimeoutSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AccessTokenInactivityTimeoutSeconds field is set to the value of the last call.
func (b *OAuthClientApplyConfiguration) WithAccessTokenInactivityTimeoutSeconds(value int32) *OAuthClientApplyConfiguration {
	b.AccessTokenInactivityTimeoutSeconds = &value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *OAuthClientApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
