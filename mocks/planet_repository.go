// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/core/domains/planet/repositories/planet_repository.go

// Package mock is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/domains/planet/entities"
	bson "go.mongodb.org/mongo-driver/bson"
)

// MockIPlanetRepository is a mock of IPlanetRepository interface.
type MockIPlanetRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPlanetRepositoryMockRecorder
}

// MockIPlanetRepositoryMockRecorder is the mock recorder for MockIPlanetRepository.
type MockIPlanetRepositoryMockRecorder struct {
	mock *MockIPlanetRepository
}

// NewMockIPlanetRepository creates a new mock instance.
func NewMockIPlanetRepository(ctrl *gomock.Controller) *MockIPlanetRepository {
	mock := &MockIPlanetRepository{ctrl: ctrl}
	mock.recorder = &MockIPlanetRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPlanetRepository) EXPECT() *MockIPlanetRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIPlanetRepository) Create(planet entities.Planet) (entities.Planet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", planet)
	ret0, _ := ret[0].(entities.Planet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIPlanetRepositoryMockRecorder) Create(planet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIPlanetRepository)(nil).Create), planet)
}

// Delete mocks base method.
func (m *MockIPlanetRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIPlanetRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIPlanetRepository)(nil).Delete), id)
}

// GetById mocks base method.
func (m *MockIPlanetRepository) GetById(id string) (entities.Planet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(entities.Planet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIPlanetRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIPlanetRepository)(nil).GetById), id)
}

// GetPlanets mocks base method.
func (m *MockIPlanetRepository) GetPlanets(filter bson.M) (entities.Planets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlanets", filter)
	ret0, _ := ret[0].(entities.Planets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlanets indicates an expected call of GetPlanets.
func (mr *MockIPlanetRepositoryMockRecorder) GetPlanets(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanets", reflect.TypeOf((*MockIPlanetRepository)(nil).GetPlanets), filter)
}

// Update mocks base method.
func (m *MockIPlanetRepository) Update(id string, planet entities.Planet) (entities.Planet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, planet)
	ret0, _ := ret[0].(entities.Planet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIPlanetRepositoryMockRecorder) Update(id, planet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPlanetRepository)(nil).Update), id, planet)
}