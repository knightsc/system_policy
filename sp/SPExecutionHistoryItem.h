#ifndef SPExecutionHistoryItem_h
#define SPExecutionHistoryItem_h

@interface SPExecutionHistoryItem : NSObject <NSCoding>

@property (readonly, nonatomic) NSString *execPath;
@property (readonly, nonatomic) NSString *mmapPath;
@property (readonly, nonatomic) NSString *signingID;
@property (readonly, nonatomic) NSString *teamID;
@property (readonly, nonatomic) NSString *cdHash;
@property (readonly, nonatomic) NSString *responsiblePath;
@property (readonly, nonatomic) NSString *developerName;
@property (readonly, nonatomic) NSDate *lastSeen;

@end

#endif /* SPExecutionHistoryItem_h */
