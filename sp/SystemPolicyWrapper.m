#import <dlfcn.h>
#import "SystemPolicyWrapper.h"
#import "SPExecutionPolicy.h"
#import "SPExecutionHistoryItem.h"
#import "SPKernelExtensionPolicy.h"
#import "SPKernelExtensionPolicyItem.h"

history_items
init_history_items(void) {
	history_items self;
	self.framework_handle = dlopen("/System/Library/PrivateFrameworks/SystemPolicy.framework/SystemPolicy", RTLD_LAZY);

	if (self.framework_handle != NULL) {
		self.pool = [[NSAutoreleasePool alloc] init];
		Class SPExecutionPolicyClass = NSClassFromString(@"SPExecutionPolicy");
		SPExecutionPolicy *execPolicy = [[[SPExecutionPolicyClass alloc] init] autorelease];
		self.historyItems = [execPolicy legacyExecutionHistory];
	}

	return self;
}

void
release_history_items(history_items self) {
	self.historyItems = nil;
	
	[self.pool release];
	self.pool = nil;

	dlclose(self.framework_handle);
	self.framework_handle = NULL;
}

unsigned long
history_items_length(history_items self) {
	return self.historyItems.count;
}

history_item
get_history_item(history_items self, unsigned long index) {
	history_item item;

	if (index < self.historyItems.count) {
		item.exec_path = [self.historyItems[index].execPath UTF8String];
		item.mmap_path = [self.historyItems[index].mmapPath UTF8String];
		item.signing_id = [self.historyItems[index].signingID UTF8String];
		item.team_id = [self.historyItems[index].teamID UTF8String];
		item.cd_hash = [self.historyItems[index].cdHash UTF8String];
		item.responsible_path = [self.historyItems[index].responsiblePath UTF8String];
		item.developer_name = [self.historyItems[index].developerName UTF8String];
		item.last_seen = [self.historyItems[index].lastSeen timeIntervalSince1970];
	}

	return item;
}

kext_items
init_kext_items(void) {
	kext_items self;
	self.framework_handle = dlopen("/System/Library/PrivateFrameworks/SystemPolicy.framework/SystemPolicy", RTLD_LAZY);

	if (self.framework_handle != NULL) {
		self.pool = [[NSAutoreleasePool alloc] init];
		Class SPKernelExtensionPolicyClass = NSClassFromString(@"SPKernelExtensionPolicy");
		SPKernelExtensionPolicy *kextPolicy = [[[SPKernelExtensionPolicyClass alloc] init] autorelease];
		self.kextItems = [kextPolicy currentPolicy];
	}

	return self;
}

void
release_kext_items(kext_items self) {	
	self.kextItems = nil;

	[self.pool release];
	self.pool = nil;
	
	dlclose(self.framework_handle);	
	self.framework_handle = NULL;
}

unsigned long
kext_items_length(kext_items self) {
	return self.kextItems.count;
}

kext_item
get_kext_item(kext_items self, unsigned long index) {
	kext_item item;

	if (index < self.kextItems.count) {
		item.developer_name = [self.kextItems[index].developerName UTF8String];
		item.application_name = [self.kextItems[index].applicationName UTF8String];
		item.application_path = [self.kextItems[index].applicationPath UTF8String];
		item.team_id = [self.kextItems[index].teamID UTF8String];
		item.bundle_id_count = self.kextItems[index].bundleIDs.count;
		item.allowed = [self.kextItems[index] isAllowed];
		item.reboot_required = [self.kextItems[index] isRebootRequired];
		item.modified = [self.kextItems[index] isModified];
	}

	return item;
}

const char *
get_kext_bundle_id(kext_items self, unsigned long kextIndex, unsigned long bundleIDIndex) {
	const char *bundle_id = NULL;

	if (kextIndex < self.kextItems.count) {
		SPKernelExtensionPolicyItem *item = self.kextItems[kextIndex];
		if (bundleIDIndex < item.bundleIDs.count) {
			bundle_id = [item.bundleIDs[bundleIDIndex] UTF8String];
		}
	}
	return bundle_id;
}
