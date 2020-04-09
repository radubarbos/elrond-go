package mock

import (
	"github.com/ElrondNetwork/elrond-go/data"
)

// StorageManagerStub --
type StorageManagerStub struct {
	DatabaseCalled              func() data.DBWriteCacher
	TakeSnapshotCalled          func([]byte)
	SetCheckpointCalled         func([]byte)
	PruneCalled                 func([]byte)
	CancelPruneCalled           func([]byte)
	MarkForEvictionCalled       func([]byte, data.ModifiedHashes) error
	GetDbThatContainsHashCalled func([]byte) data.DBWriteCacher
	IsPruningEnabledCalled      func() bool
	EnterSnapshotModeCalled     func()
	ExitSnapshotModeCalled      func()
	IsInterfaceNilCalled        func() bool
}

// Database --
func (sms *StorageManagerStub) Database() data.DBWriteCacher {
	if sms.DatabaseCalled != nil {
		return sms.DatabaseCalled()
	}
	return nil
}

// TakeSnapshot --
func (sms *StorageManagerStub) TakeSnapshot([]byte) {

}

// SetCheckpoint --
func (sms *StorageManagerStub) SetCheckpoint([]byte) {

}

// Prune --
func (sms *StorageManagerStub) Prune([]byte) {

}

// CancelPrune --
func (sms *StorageManagerStub) CancelPrune([]byte) {

}

// MarkForEviction --
func (sms *StorageManagerStub) MarkForEviction(d []byte, m data.ModifiedHashes) error {
	if sms.MarkForEvictionCalled != nil {
		return sms.MarkForEvictionCalled(d, m)
	}
	return nil
}

// GetDbThatContainsHash --
func (sms *StorageManagerStub) GetDbThatContainsHash(d []byte) data.DBWriteCacher {
	if sms.GetDbThatContainsHashCalled != nil {
		return sms.GetDbThatContainsHashCalled(d)
	}

	return nil
}

// IsPruningEnabled --
func (sms *StorageManagerStub) IsPruningEnabled() bool {
	if sms.IsPruningEnabledCalled != nil {
		return sms.IsPruningEnabledCalled()
	}
	return false
}

// EnterSnapshotMode --
func (sms *StorageManagerStub) EnterSnapshotMode() {
	if sms.EnterSnapshotModeCalled != nil {
		sms.EnterSnapshotModeCalled()
	}
}

// ExitSnapshotMode --
func (sms *StorageManagerStub) ExitSnapshotMode() {
	if sms.ExitSnapshotModeCalled != nil {
		sms.ExitSnapshotModeCalled()
	}
}

// IsInterfaceNil --
func (sms *StorageManagerStub) IsInterfaceNil() bool {
	return sms == nil
}
