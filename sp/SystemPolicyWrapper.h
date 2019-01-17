#import <Foundation/Foundation.h>

@class SPExecutionHistoryItem;
@class SPKernelExtensionPolicyItem;

typedef struct {
	void *framework_handle;
	NSAutoreleasePool *pool;
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

history_items init_history_items(void);
void release_history_items(history_items self);
unsigned long history_items_length(history_items self);
history_item get_history_item(history_items self, unsigned long index);

typedef struct {
	void *framework_handle;
	NSAutoreleasePool *pool;
	NSArray<SPKernelExtensionPolicyItem *> *kextItems;
} kext_items;

typedef struct {
	const char *developer_name;
	const char *application_name;
	const char *application_path;
	const char *team_id;
	unsigned long bundle_id_count;
	char allowed;
	char reboot_required;
	char modified;
} kext_item;

kext_items init_kext_items(void);
void release_kext_items(kext_items self);
unsigned long kext_items_length(kext_items self);
kext_item get_kext_item(kext_items self, unsigned long index);
const char *get_kext_bundle_id(kext_items self, unsigned long kextIndex, unsigned long bundleIDIndex);
