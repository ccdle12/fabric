/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by mockery v1.0.0
package mocks

import chaincode "github.com/hyperledger/fabric/common/chaincode"
import mock "github.com/stretchr/testify/mock"

// LifeCycleChangeListener is an autogenerated mock type for the LifeCycleChangeListener type
type LifeCycleChangeListener struct {
	mock.Mock
}

// LifeCycleChangeListener provides a mock function with given fields: channel, chaincodes
func (_m *LifeCycleChangeListener) LifeCycleChangeListener(channel string, chaincodes chaincode.MetadataSet) {
	_m.Called(channel, chaincodes)
}