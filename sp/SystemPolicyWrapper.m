#import <dlfcn.h>
#import "SystemPolicyWrapper.h"
#import "SPExecutionPolicy.h"
#import "SPExecutionHistoryItem.h"

history_items
init_history_items(void) {
	history_items self;
	self.framework_handle = dlopen("/System/Library/PrivateFrameworks/SystemPolicy.framework/SystemPolicy", RTLD_LAZY);

	if (self.framework_handle != NULL) {
		Class SPExecutionPolicyClass = NSClassFromString(@"SPExecutionPolicy");
		SPExecutionPolicy *execPolicy = [[SPExecutionPolicyClass alloc] init];
		self.historyItems = [execPolicy legacyExecutionHistory];
	}

	return self;
}

void
release_history_items(history_items self) {
	self.historyItems = nil;
	dlclose(self.framework_handle);
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
