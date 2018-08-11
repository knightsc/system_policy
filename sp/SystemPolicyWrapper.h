#import <Foundation/Foundation.h>

@class SPExecutionHistoryItem;

typedef struct {
	void *framework_handle;
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
