package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -F/System/Library/PrivateFrameworks
#cgo LDFLAGS: -framework Foundation -framework SystemPolicy
#import <Foundation/Foundation.h>
#import "SPExecutionPolicy.h"
#import "SPExecutionHistoryItem.h"

typedef struct {
	NSArray<SPExecutionHistoryItem *> *historyItems;
} history_items;

typedef struct {
    const char *exec_path;
	const char *mmap_path;
	const char *signing_id;
	const char *team_id;
	const char *cd_hash;
	const char *responsible_path;
	const char *developer_name;
	double last_seen;
} history_item;

history_items
init_history_items(void) {
	history_items self;
	SPExecutionPolicy *execPolicy = [[SPExecutionPolicy alloc] init];
	self.historyItems = [execPolicy legacyExecutionHistory];
	return self;
}

void
release_history_items(history_items self) {
	self.historyItems = nil;
}

unsigned long
history_items_length(history_items self) {
	return self.historyItems.count;
}

history_item
get_history_item(history_items self, unsigned long index) {
	history_item item;
	item.exec_path = [self.historyItems[index].execPath UTF8String];
	item.mmap_path = [self.historyItems[index].mmapPath UTF8String];
	item.signing_id = [self.historyItems[index].signingID UTF8String];
	item.team_id = [self.historyItems[index].teamID UTF8String];
	item.cd_hash = [self.historyItems[index].cdHash UTF8String];
	item.responsible_path = [self.historyItems[index].responsiblePath UTF8String];
	item.developer_name = [self.historyItems[index].developerName UTF8String];
	item.last_seen = [self.historyItems[index].lastSeen timeIntervalSince1970];

	return item;
}
*/
import "C"
import (
	"math"
	"time"
)

type executionHistoryItem struct {
	execPath        string
	mmapPath        string
	signingID       string
	teamID          string
	cdHash          string
	responsiblePath string
	developerName   string
	lastSeen        time.Time
}

func legacyExecutionHistory() []executionHistoryItem {
	self := C.init_history_items()
	defer C.release_history_items(self)

	history := make([]executionHistoryItem, 0)

	length := uint64(C.history_items_length(self))
	for i := uint64(0); i < length; i++ {
		item := C.get_history_item(self, C.ulong(i))

		ehi := executionHistoryItem{}
		ehi.execPath = C.GoString(item.exec_path)
		ehi.mmapPath = C.GoString(item.mmap_path)
		ehi.signingID = C.GoString(item.signing_id)
		ehi.teamID = C.GoString(item.team_id)
		ehi.cdHash = C.GoString(item.cd_hash)
		ehi.responsiblePath = C.GoString(item.responsible_path)
		ehi.developerName = C.GoString(item.developer_name)
		sec, dec := math.Modf(float64(item.last_seen))
		ehi.lastSeen = time.Unix(int64(sec), int64(dec*(1e9)))

		history = append(history, ehi)
	}

	return history
}
