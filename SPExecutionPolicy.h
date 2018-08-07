#ifndef SPExecutionPolicy_h
#define SPExecutionPolicy_h

@class SPExecutionHistoryItem;

@interface SPExecutionPolicy : NSObject

- (instancetype)init;
- (NSArray<SPExecutionHistoryItem *> *)legacyExecutionHistory;

@end

#endif /* SPExecutionPolicy_h */
