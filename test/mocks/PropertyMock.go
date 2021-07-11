package mocks

import (
	"github.com/jgrath/go-and-find-with-go/store"
	. "github.com/jgrath/go-and-find-with-go/util"
	"github.com/stretchr/testify/mock"
)

type MockPropertyStoreStruct struct {
	mock.Mock
}

func (mockedStore *MockPropertyStoreStruct) FindProperties() ([]*SystemProperty, error) {
	mockedObject := mockedStore.Called()
	return mockedObject.Get(0).([]*SystemProperty), mockedObject.Error(1)
}

func (mockedStore *MockPropertyStoreStruct) AddProperty(property *SystemProperty) (error){
	mockedObject := mockedStore.Called()
	return mockedObject.Error(0)
}

func InitPropertyMockStore() *MockPropertyStoreStruct {
	mock := new(MockPropertyStoreStruct)
	store.MainPropertyStore = mock
	return mock
}
