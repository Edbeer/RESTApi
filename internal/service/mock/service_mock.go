// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/Edbeer/restapi/internal/entity"
	utils "github.com/Edbeer/restapi/pkg/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAuth) Delete(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthMockRecorder) Delete(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuth)(nil).Delete), ctx, userID)
}

// FindUsersByName mocks base method.
func (m *MockAuth) FindUsersByName(ctx context.Context, name string, pq *utils.PaginationQuery) (*entity.UsersList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUsersByName", ctx, name, pq)
	ret0, _ := ret[0].(*entity.UsersList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUsersByName indicates an expected call of FindUsersByName.
func (mr *MockAuthMockRecorder) FindUsersByName(ctx, name, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUsersByName", reflect.TypeOf((*MockAuth)(nil).FindUsersByName), ctx, name, pq)
}

// GetUserByID mocks base method.
func (m *MockAuth) GetUserByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockAuthMockRecorder) GetUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAuth)(nil).GetUserByID), ctx, userID)
}

// GetUsers mocks base method.
func (m *MockAuth) GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*entity.UsersList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, pq)
	ret0, _ := ret[0].(*entity.UsersList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockAuthMockRecorder) GetUsers(ctx, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockAuth)(nil).GetUsers), ctx, pq)
}

// Login mocks base method.
func (m *MockAuth) Login(ctx context.Context, user *entity.User) (*entity.UserWithToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, user)
	ret0, _ := ret[0].(*entity.UserWithToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), ctx, user)
}

// Register mocks base method.
func (m *MockAuth) Register(ctx context.Context, user *entity.User) (*entity.UserWithToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(*entity.UserWithToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuth)(nil).Register), ctx, user)
}

// Update mocks base method.
func (m *MockAuth) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAuthMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuth)(nil).Update), ctx, user)
}

// MockNews is a mock of News interface.
type MockNews struct {
	ctrl     *gomock.Controller
	recorder *MockNewsMockRecorder
}

// MockNewsMockRecorder is the mock recorder for MockNews.
type MockNewsMockRecorder struct {
	mock *MockNews
}

// NewMockNews creates a new mock instance.
func NewMockNews(ctrl *gomock.Controller) *MockNews {
	mock := &MockNews{ctrl: ctrl}
	mock.recorder = &MockNewsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNews) EXPECT() *MockNewsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNews) Create(ctx context.Context, news *entity.News) (*entity.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, news)
	ret0, _ := ret[0].(*entity.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockNewsMockRecorder) Create(ctx, news interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNews)(nil).Create), ctx, news)
}

// Delete mocks base method.
func (m *MockNews) Delete(ctx context.Context, newsID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, newsID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNewsMockRecorder) Delete(ctx, newsID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNews)(nil).Delete), ctx, newsID)
}

// GetNews mocks base method.
func (m *MockNews) GetNews(ctx context.Context, pq *utils.PaginationQuery) (*entity.NewsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNews", ctx, pq)
	ret0, _ := ret[0].(*entity.NewsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNews indicates an expected call of GetNews.
func (mr *MockNewsMockRecorder) GetNews(ctx, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNews", reflect.TypeOf((*MockNews)(nil).GetNews), ctx, pq)
}

// GetNewsByID mocks base method.
func (m *MockNews) GetNewsByID(ctx context.Context, newsID uuid.UUID) (*entity.NewsBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewsByID", ctx, newsID)
	ret0, _ := ret[0].(*entity.NewsBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewsByID indicates an expected call of GetNewsByID.
func (mr *MockNewsMockRecorder) GetNewsByID(ctx, newsID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewsByID", reflect.TypeOf((*MockNews)(nil).GetNewsByID), ctx, newsID)
}

// SearchNews mocks base method.
func (m *MockNews) SearchNews(ctx context.Context, title string, pq *utils.PaginationQuery) (*entity.NewsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNews", ctx, title, pq)
	ret0, _ := ret[0].(*entity.NewsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNews indicates an expected call of SearchNews.
func (mr *MockNewsMockRecorder) SearchNews(ctx, title, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNews", reflect.TypeOf((*MockNews)(nil).SearchNews), ctx, title, pq)
}

// Update mocks base method.
func (m *MockNews) Update(ctx context.Context, news *entity.News) (*entity.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, news)
	ret0, _ := ret[0].(*entity.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockNewsMockRecorder) Update(ctx, news interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNews)(nil).Update), ctx, news)
}

// MockComments is a mock of Comments interface.
type MockComments struct {
	ctrl     *gomock.Controller
	recorder *MockCommentsMockRecorder
}

// MockCommentsMockRecorder is the mock recorder for MockComments.
type MockCommentsMockRecorder struct {
	mock *MockComments
}

// NewMockComments creates a new mock instance.
func NewMockComments(ctrl *gomock.Controller) *MockComments {
	mock := &MockComments{ctrl: ctrl}
	mock.recorder = &MockCommentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComments) EXPECT() *MockCommentsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockComments) Create(ctx context.Context, comments *entity.Comment) (*entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, comments)
	ret0, _ := ret[0].(*entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommentsMockRecorder) Create(ctx, comments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockComments)(nil).Create), ctx, comments)
}

// Delete mocks base method.
func (m *MockComments) Delete(ctx context.Context, commentID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, commentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentsMockRecorder) Delete(ctx, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockComments)(nil).Delete), ctx, commentID)
}

// GetAllByNewsID mocks base method.
func (m *MockComments) GetAllByNewsID(ctx context.Context, newsID uuid.UUID, pq *utils.PaginationQuery) (*entity.CommentsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByNewsID", ctx, newsID, pq)
	ret0, _ := ret[0].(*entity.CommentsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByNewsID indicates an expected call of GetAllByNewsID.
func (mr *MockCommentsMockRecorder) GetAllByNewsID(ctx, newsID, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByNewsID", reflect.TypeOf((*MockComments)(nil).GetAllByNewsID), ctx, newsID, pq)
}

// GetByID mocks base method.
func (m *MockComments) GetByID(ctx context.Context, commentID uuid.UUID) (*entity.CommentBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, commentID)
	ret0, _ := ret[0].(*entity.CommentBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCommentsMockRecorder) GetByID(ctx, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockComments)(nil).GetByID), ctx, commentID)
}

// Update mocks base method.
func (m *MockComments) Update(ctx context.Context, comments *entity.Comment) (*entity.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, comments)
	ret0, _ := ret[0].(*entity.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommentsMockRecorder) Update(ctx, comments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockComments)(nil).Update), ctx, comments)
}

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockSession) CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, session, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionMockRecorder) CreateSession(ctx, session, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSession)(nil).CreateSession), ctx, session, expire)
}

// DeleteSessionByID mocks base method.
func (m *MockSession) DeleteSessionByID(ctx context.Context, sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionByID", ctx, sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionByID indicates an expected call of DeleteSessionByID.
func (mr *MockSessionMockRecorder) DeleteSessionByID(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionByID", reflect.TypeOf((*MockSession)(nil).DeleteSessionByID), ctx, sessionID)
}

// GetSessionByID mocks base method.
func (m *MockSession) GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByID", ctx, sessionID)
	ret0, _ := ret[0].(*entity.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByID indicates an expected call of GetSessionByID.
func (mr *MockSessionMockRecorder) GetSessionByID(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByID", reflect.TypeOf((*MockSession)(nil).GetSessionByID), ctx, sessionID)
}
