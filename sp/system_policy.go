package sp

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import "SystemPolicyWrapper.h"
*/
import "C"
import (
	"math"
	"time"
)

type ExecutionHistoryItem struct {
	ExecPath        string
	MmapPath        string
	SigningID       string
	TeamID          string
	CDHash          string
	ResponsiblePath string
	DeveloperName   string
	LastSeen        time.Time
}

func LegacyExecutionHistory() []ExecutionHistoryItem {
	self := C.init_history_items()
	defer C.release_history_items(self)

	history := make([]ExecutionHistoryItem, 0)

	length := uint64(C.history_items_length(self))
	for i := uint64(0); i < length; i++ {
		item := C.get_history_item(self, C.ulong(i))

		ehi := ExecutionHistoryItem{}
		ehi.ExecPath = C.GoString(item.exec_path)
		ehi.MmapPath = C.GoString(item.mmap_path)
		ehi.SigningID = C.GoString(item.signing_id)
		ehi.TeamID = C.GoString(item.team_id)
		ehi.CDHash = C.GoString(item.cd_hash)
		ehi.ResponsiblePath = C.GoString(item.responsible_path)
		ehi.DeveloperName = C.GoString(item.developer_name)
		sec, dec := math.Modf(float64(item.last_seen))
		ehi.LastSeen = time.Unix(int64(sec), int64(dec*(1e9)))

		history = append(history, ehi)
	}

	return history
}
