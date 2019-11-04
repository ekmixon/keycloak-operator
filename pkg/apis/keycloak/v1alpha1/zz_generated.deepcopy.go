// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationExecutionInfo) DeepCopyInto(out *AuthenticationExecutionInfo) {
	*out = *in
	if in.RequirementChoices != nil {
		in, out := &in.RequirementChoices, &out.RequirementChoices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationExecutionInfo.
func (in *AuthenticationExecutionInfo) DeepCopy() *AuthenticationExecutionInfo {
	if in == nil {
		return nil
	}
	out := new(AuthenticationExecutionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticatorConfig) DeepCopyInto(out *AuthenticatorConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticatorConfig.
func (in *AuthenticatorConfig) DeepCopy() *AuthenticatorConfig {
	if in == nil {
		return nil
	}
	out := new(AuthenticatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedIdentity) DeepCopyInto(out *FederatedIdentity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedIdentity.
func (in *FederatedIdentity) DeepCopy() *FederatedIdentity {
	if in == nil {
		return nil
	}
	out := new(FederatedIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keycloak) DeepCopyInto(out *Keycloak) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keycloak.
func (in *Keycloak) DeepCopy() *Keycloak {
	if in == nil {
		return nil
	}
	out := new(Keycloak)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Keycloak) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIClient) DeepCopyInto(out *KeycloakAPIClient) {
	*out = *in
	if in.DefaultRoles != nil {
		in, out := &in.DefaultRoles, &out.DefaultRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RedirectUris != nil {
		in, out := &in.RedirectUris, &out.RedirectUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WebOrigins != nil {
		in, out := &in.WebOrigins, &out.WebOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProtocolMappers != nil {
		in, out := &in.ProtocolMappers, &out.ProtocolMappers
		*out = make([]KeycloakProtocolMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIClient.
func (in *KeycloakAPIClient) DeepCopy() *KeycloakAPIClient {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIPasswordReset) DeepCopyInto(out *KeycloakAPIPasswordReset) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIPasswordReset.
func (in *KeycloakAPIPasswordReset) DeepCopy() *KeycloakAPIPasswordReset {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIPasswordReset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIRealm) DeepCopyInto(out *KeycloakAPIRealm) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]*KeycloakAPIUser, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakAPIUser)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Clients != nil {
		in, out := &in.Clients, &out.Clients
		*out = make([]*KeycloakAPIClient, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakAPIClient)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.IdentityProviders != nil {
		in, out := &in.IdentityProviders, &out.IdentityProviders
		*out = make([]*KeycloakIdentityProvider, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakIdentityProvider)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.EventsListeners != nil {
		in, out := &in.EventsListeners, &out.EventsListeners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIRealm.
func (in *KeycloakAPIRealm) DeepCopy() *KeycloakAPIRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakAPIUser) DeepCopyInto(out *KeycloakAPIUser) {
	*out = *in
	if in.RealmRoles != nil {
		in, out := &in.RealmRoles, &out.RealmRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClientRoles != nil {
		in, out := &in.ClientRoles, &out.ClientRoles
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.RequiredActions != nil {
		in, out := &in.RequiredActions, &out.RequiredActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FederatedIdentities != nil {
		in, out := &in.FederatedIdentities, &out.FederatedIdentities
		*out = make([]FederatedIdentity, len(*in))
		copy(*out, *in)
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]KeycloakCredential, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakAPIUser.
func (in *KeycloakAPIUser) DeepCopy() *KeycloakAPIUser {
	if in == nil {
		return nil
	}
	out := new(KeycloakAPIUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakCredential) DeepCopyInto(out *KeycloakCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakCredential.
func (in *KeycloakCredential) DeepCopy() *KeycloakCredential {
	if in == nil {
		return nil
	}
	out := new(KeycloakCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakExternalAccess) DeepCopyInto(out *KeycloakExternalAccess) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakExternalAccess.
func (in *KeycloakExternalAccess) DeepCopy() *KeycloakExternalAccess {
	if in == nil {
		return nil
	}
	out := new(KeycloakExternalAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakIdentityProvider) DeepCopyInto(out *KeycloakIdentityProvider) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakIdentityProvider.
func (in *KeycloakIdentityProvider) DeepCopy() *KeycloakIdentityProvider {
	if in == nil {
		return nil
	}
	out := new(KeycloakIdentityProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakList) DeepCopyInto(out *KeycloakList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Keycloak, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakList.
func (in *KeycloakList) DeepCopy() *KeycloakList {
	if in == nil {
		return nil
	}
	out := new(KeycloakList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakProtocolMapper) DeepCopyInto(out *KeycloakProtocolMapper) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakProtocolMapper.
func (in *KeycloakProtocolMapper) DeepCopy() *KeycloakProtocolMapper {
	if in == nil {
		return nil
	}
	out := new(KeycloakProtocolMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealm) DeepCopyInto(out *KeycloakRealm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealm.
func (in *KeycloakRealm) DeepCopy() *KeycloakRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmList) DeepCopyInto(out *KeycloakRealmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakRealm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmList.
func (in *KeycloakRealmList) DeepCopy() *KeycloakRealmList {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmSpec) DeepCopyInto(out *KeycloakRealmSpec) {
	*out = *in
	if in.InstanceSelector != nil {
		in, out := &in.InstanceSelector, &out.InstanceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(KeycloakAPIRealm)
		(*in).DeepCopyInto(*out)
	}
	if in.RealmOverrides != nil {
		in, out := &in.RealmOverrides, &out.RealmOverrides
		*out = make([]*RedirectorIdentityProviderOverride, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RedirectorIdentityProviderOverride)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmSpec.
func (in *KeycloakRealmSpec) DeepCopy() *KeycloakRealmSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmStatus) DeepCopyInto(out *KeycloakRealmStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmStatus.
func (in *KeycloakRealmStatus) DeepCopy() *KeycloakRealmStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakSpec) DeepCopyInto(out *KeycloakSpec) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ExternalAccess = in.ExternalAccess
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakSpec.
func (in *KeycloakSpec) DeepCopy() *KeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakStatus) DeepCopyInto(out *KeycloakStatus) {
	*out = *in
	if in.SecondaryResources != nil {
		in, out := &in.SecondaryResources, &out.SecondaryResources
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakStatus.
func (in *KeycloakStatus) DeepCopy() *KeycloakStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUserRole) DeepCopyInto(out *KeycloakUserRole) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUserRole.
func (in *KeycloakUserRole) DeepCopy() *KeycloakUserRole {
	if in == nil {
		return nil
	}
	out := new(KeycloakUserRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedirectorIdentityProviderOverride) DeepCopyInto(out *RedirectorIdentityProviderOverride) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedirectorIdentityProviderOverride.
func (in *RedirectorIdentityProviderOverride) DeepCopy() *RedirectorIdentityProviderOverride {
	if in == nil {
		return nil
	}
	out := new(RedirectorIdentityProviderOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenResponse) DeepCopyInto(out *TokenResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenResponse.
func (in *TokenResponse) DeepCopy() *TokenResponse {
	if in == nil {
		return nil
	}
	out := new(TokenResponse)
	in.DeepCopyInto(out)
	return out
}
