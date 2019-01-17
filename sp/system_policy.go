// Package sp provides access to the SystemPolicy.framework on macOS.
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

// An ExecutionHistoryItem represents a 32-bit application that has been run.
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

// A KernelExtensionPolicyItem represents a KEXT that is either waiting for
// approval to load or is was approved and loaded.
type KernelExtensionPolicyItem struct {
	DeveloperName   string
	ApplicationName string
	ApplicationPath string
	TeamID          string
	BundleID        string
	Allowed         bool
	RebootRequired  bool
	Modified        bool
}

// LegacyExecutionHistory returns a list of 32-bit applications that have
// been executed on a macOS machine.
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

// CurrentKernelExtensionPolicy returns a list of items that represent whether
// a specific KEXT was approved and loaded or not.
func CurrentKernelExtensionPolicy() []KernelExtensionPolicyItem {
	self := C.init_kext_items()
	defer C.release_kext_items(self)

	policy := make([]KernelExtensionPolicyItem, 0)

	length := uint64(C.kext_items_length(self))
	for i := uint64(0); i < length; i++ {
		item := C.get_kext_item(self, C.ulong(i))

		for j := uint64(0); j < uint64(item.bundle_id_count); j++ {
			ki := KernelExtensionPolicyItem{}
			ki.DeveloperName = C.GoString(item.developer_name)
			ki.ApplicationName = C.GoString(item.application_name)
			ki.ApplicationPath = C.GoString(item.application_path)
			ki.TeamID = C.GoString(item.team_id)
			ki.BundleID = C.GoString(C.get_kext_bundle_id(self, C.ulong(i), C.ulong(j)))
			ki.Allowed = item.allowed == '\x01'
			ki.RebootRequired = item.reboot_required == '\x01'
			ki.Modified = item.modified == '\x01'

			policy = append(policy, ki)
		}
	}

	return policy
}
