// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zitadel/oidc/v3/pkg/op (interfaces: HasRedirectGlobs)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	oidc "github.com/zitadel/oidc/v3/pkg/oidc"
	op "github.com/zitadel/oidc/v3/pkg/op"
)

// MockHasRedirectGlobs is a mock of HasRedirectGlobs interface.
type MockHasRedirectGlobs struct {
	ctrl     *gomock.Controller
	recorder *MockHasRedirectGlobsMockRecorder
}

// MockHasRedirectGlobsMockRecorder is the mock recorder for MockHasRedirectGlobs.
type MockHasRedirectGlobsMockRecorder struct {
	mock *MockHasRedirectGlobs
}

// NewMockHasRedirectGlobs creates a new mock instance.
func NewMockHasRedirectGlobs(ctrl *gomock.Controller) *MockHasRedirectGlobs {
	mock := &MockHasRedirectGlobs{ctrl: ctrl}
	mock.recorder = &MockHasRedirectGlobsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasRedirectGlobs) EXPECT() *MockHasRedirectGlobsMockRecorder {
	return m.recorder
}

// AccessTokenType mocks base method.
func (m *MockHasRedirectGlobs) AccessTokenType() op.AccessTokenType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenType")
	ret0, _ := ret[0].(op.AccessTokenType)
	return ret0
}

// AccessTokenType indicates an expected call of AccessTokenType.
func (mr *MockHasRedirectGlobsMockRecorder) AccessTokenType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenType", reflect.TypeOf((*MockHasRedirectGlobs)(nil).AccessTokenType))
}

// ApplicationType mocks base method.
func (m *MockHasRedirectGlobs) ApplicationType() op.ApplicationType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationType")
	ret0, _ := ret[0].(op.ApplicationType)
	return ret0
}

// ApplicationType indicates an expected call of ApplicationType.
func (mr *MockHasRedirectGlobsMockRecorder) ApplicationType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationType", reflect.TypeOf((*MockHasRedirectGlobs)(nil).ApplicationType))
}

// AuthMethod mocks base method.
func (m *MockHasRedirectGlobs) AuthMethod() oidc.AuthMethod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMethod")
	ret0, _ := ret[0].(oidc.AuthMethod)
	return ret0
}

// AuthMethod indicates an expected call of AuthMethod.
func (mr *MockHasRedirectGlobsMockRecorder) AuthMethod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMethod", reflect.TypeOf((*MockHasRedirectGlobs)(nil).AuthMethod))
}

// ClockSkew mocks base method.
func (m *MockHasRedirectGlobs) ClockSkew() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClockSkew")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// ClockSkew indicates an expected call of ClockSkew.
func (mr *MockHasRedirectGlobsMockRecorder) ClockSkew() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClockSkew", reflect.TypeOf((*MockHasRedirectGlobs)(nil).ClockSkew))
}

// DevMode mocks base method.
func (m *MockHasRedirectGlobs) DevMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DevMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DevMode indicates an expected call of DevMode.
func (mr *MockHasRedirectGlobsMockRecorder) DevMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DevMode", reflect.TypeOf((*MockHasRedirectGlobs)(nil).DevMode))
}

// GetID mocks base method.
func (m *MockHasRedirectGlobs) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockHasRedirectGlobsMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockHasRedirectGlobs)(nil).GetID))
}

// GrantTypes mocks base method.
func (m *MockHasRedirectGlobs) GrantTypes() []oidc.GrantType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantTypes")
	ret0, _ := ret[0].([]oidc.GrantType)
	return ret0
}

// GrantTypes indicates an expected call of GrantTypes.
func (mr *MockHasRedirectGlobsMockRecorder) GrantTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantTypes", reflect.TypeOf((*MockHasRedirectGlobs)(nil).GrantTypes))
}

// IDTokenLifetime mocks base method.
func (m *MockHasRedirectGlobs) IDTokenLifetime() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDTokenLifetime")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// IDTokenLifetime indicates an expected call of IDTokenLifetime.
func (mr *MockHasRedirectGlobsMockRecorder) IDTokenLifetime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDTokenLifetime", reflect.TypeOf((*MockHasRedirectGlobs)(nil).IDTokenLifetime))
}

// IDTokenUserinfoClaimsAssertion mocks base method.
func (m *MockHasRedirectGlobs) IDTokenUserinfoClaimsAssertion() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDTokenUserinfoClaimsAssertion")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IDTokenUserinfoClaimsAssertion indicates an expected call of IDTokenUserinfoClaimsAssertion.
func (mr *MockHasRedirectGlobsMockRecorder) IDTokenUserinfoClaimsAssertion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDTokenUserinfoClaimsAssertion", reflect.TypeOf((*MockHasRedirectGlobs)(nil).IDTokenUserinfoClaimsAssertion))
}

// IsScopeAllowed mocks base method.
func (m *MockHasRedirectGlobs) IsScopeAllowed(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsScopeAllowed", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsScopeAllowed indicates an expected call of IsScopeAllowed.
func (mr *MockHasRedirectGlobsMockRecorder) IsScopeAllowed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsScopeAllowed", reflect.TypeOf((*MockHasRedirectGlobs)(nil).IsScopeAllowed), arg0)
}

// LoginURL mocks base method.
func (m *MockHasRedirectGlobs) LoginURL(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginURL", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// LoginURL indicates an expected call of LoginURL.
func (mr *MockHasRedirectGlobsMockRecorder) LoginURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginURL", reflect.TypeOf((*MockHasRedirectGlobs)(nil).LoginURL), arg0)
}

// PostLogoutRedirectURIGlobs mocks base method.
func (m *MockHasRedirectGlobs) PostLogoutRedirectURIGlobs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostLogoutRedirectURIGlobs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// PostLogoutRedirectURIGlobs indicates an expected call of PostLogoutRedirectURIGlobs.
func (mr *MockHasRedirectGlobsMockRecorder) PostLogoutRedirectURIGlobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostLogoutRedirectURIGlobs", reflect.TypeOf((*MockHasRedirectGlobs)(nil).PostLogoutRedirectURIGlobs))
}

// PostLogoutRedirectURIs mocks base method.
func (m *MockHasRedirectGlobs) PostLogoutRedirectURIs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostLogoutRedirectURIs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// PostLogoutRedirectURIs indicates an expected call of PostLogoutRedirectURIs.
func (mr *MockHasRedirectGlobsMockRecorder) PostLogoutRedirectURIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostLogoutRedirectURIs", reflect.TypeOf((*MockHasRedirectGlobs)(nil).PostLogoutRedirectURIs))
}

// RedirectURIGlobs mocks base method.
func (m *MockHasRedirectGlobs) RedirectURIGlobs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedirectURIGlobs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// RedirectURIGlobs indicates an expected call of RedirectURIGlobs.
func (mr *MockHasRedirectGlobsMockRecorder) RedirectURIGlobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedirectURIGlobs", reflect.TypeOf((*MockHasRedirectGlobs)(nil).RedirectURIGlobs))
}

// RedirectURIs mocks base method.
func (m *MockHasRedirectGlobs) RedirectURIs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedirectURIs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// RedirectURIs indicates an expected call of RedirectURIs.
func (mr *MockHasRedirectGlobsMockRecorder) RedirectURIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedirectURIs", reflect.TypeOf((*MockHasRedirectGlobs)(nil).RedirectURIs))
}

// ResponseTypes mocks base method.
func (m *MockHasRedirectGlobs) ResponseTypes() []oidc.ResponseType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseTypes")
	ret0, _ := ret[0].([]oidc.ResponseType)
	return ret0
}

// ResponseTypes indicates an expected call of ResponseTypes.
func (mr *MockHasRedirectGlobsMockRecorder) ResponseTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseTypes", reflect.TypeOf((*MockHasRedirectGlobs)(nil).ResponseTypes))
}

// RestrictAdditionalAccessTokenScopes mocks base method.
func (m *MockHasRedirectGlobs) RestrictAdditionalAccessTokenScopes() func([]string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestrictAdditionalAccessTokenScopes")
	ret0, _ := ret[0].(func([]string) []string)
	return ret0
}

// RestrictAdditionalAccessTokenScopes indicates an expected call of RestrictAdditionalAccessTokenScopes.
func (mr *MockHasRedirectGlobsMockRecorder) RestrictAdditionalAccessTokenScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestrictAdditionalAccessTokenScopes", reflect.TypeOf((*MockHasRedirectGlobs)(nil).RestrictAdditionalAccessTokenScopes))
}

// RestrictAdditionalIdTokenScopes mocks base method.
func (m *MockHasRedirectGlobs) RestrictAdditionalIdTokenScopes() func([]string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestrictAdditionalIdTokenScopes")
	ret0, _ := ret[0].(func([]string) []string)
	return ret0
}

// RestrictAdditionalIdTokenScopes indicates an expected call of RestrictAdditionalIdTokenScopes.
func (mr *MockHasRedirectGlobsMockRecorder) RestrictAdditionalIdTokenScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestrictAdditionalIdTokenScopes", reflect.TypeOf((*MockHasRedirectGlobs)(nil).RestrictAdditionalIdTokenScopes))
}
