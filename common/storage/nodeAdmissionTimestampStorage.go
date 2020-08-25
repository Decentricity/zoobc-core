package storage

import (
	"sync"

	"github.com/zoobc/zoobc-core/common/blocker"
	"github.com/zoobc/zoobc-core/common/model"
)

type (
	// NodeAdmissionTimesatmpStorage store next node admission timestamp
	NodeAdmissionTimestampStorage struct {
		sync.RWMutex
		nextNodeAdmissionTimestamp model.NodeAdmissionTimestamp
	}
)

// NewNodeAdmissionTimestampStorage returns new NodeAdmissionTimesatmpStorage instance
func NewNodeAdmissionTimestampStorage() *NodeAdmissionTimestampStorage {
	return &NodeAdmissionTimestampStorage{
		nextNodeAdmissionTimestamp: model.NodeAdmissionTimestamp{},
	}
}

// SetItem setter of NodeAdmissionTimestampStorage
// Note: lastChange curretly unused
func (ns *NodeAdmissionTimestampStorage) SetItem(lastChange, item interface{}) error {
	ns.Lock()
	defer ns.Unlock()
	var (
		ok                         bool
		newNodeAdmissionTimesatamp model.NodeAdmissionTimestamp
	)
	// covert type
	newNodeAdmissionTimesatamp, ok = (item).(model.NodeAdmissionTimestamp)
	if !ok {
		return blocker.NewBlocker(blocker.ValidationErr, "WrongType item. Expected model.NodeAdmissionTimestamp")
	}
	ns.nextNodeAdmissionTimestamp = newNodeAdmissionTimesatamp
	return nil
}

// GetItem getter of NodeAdmissionTimestampStorage
func (ns *NodeAdmissionTimestampStorage) GetItem(lastChange, item interface{}) error {
	ns.Lock()
	defer ns.Unlock()

	if ns.nextNodeAdmissionTimestamp.Timestamp == 0 {
		return blocker.NewBlocker(blocker.ValidationErr, "EmptyNodeAdmissionTimestampStorage")
	}
	var (
		nodeAdmissionTimesatampCopy, ok = item.(*model.NodeAdmissionTimestamp)
	)
	if !ok {
		return blocker.NewBlocker(blocker.ValidationErr, "WrongType item. Expected *model.NodeAdmissionTimestamp")
	}
	// copy stored NodeAdmissionTimestamp value into reference variable requester
	*nodeAdmissionTimesatampCopy = ns.nextNodeAdmissionTimestamp
	return nil
}

// GetSize return the size of NodeAdmissionTimestampStorage
func (ns *NodeAdmissionTimestampStorage) GetSize() int64 {
	return int64(ns.nextNodeAdmissionTimestamp.XXX_Size())
}

// ClearCache cleaner of NodeAdmissionTimestampStorage
func (ns *NodeAdmissionTimestampStorage) ClearCache() error {
	ns.nextNodeAdmissionTimestamp = model.NodeAdmissionTimestamp{}
	return nil
}