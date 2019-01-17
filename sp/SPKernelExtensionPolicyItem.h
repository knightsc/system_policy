#ifndef SPKernelExtensionPolicyItem_h
#define SPKernelExtensionPolicyItem_h

@interface SPKernelExtensionPolicyItem : NSObject <NSCoding>

@property (readonly, nonatomic) NSString *developerName;
@property (readonly, nonatomic) NSString *applicationName;
@property (readonly, nonatomic) NSString *applicationPath;
@property (readonly, nonatomic) NSString *teamID;
@property (readonly, nonatomic) NSArray *bundleIDs;
@property (nonatomic, getter=isAllowed) char allowed;
@property (readonly, nonatomic, getter=isRebootRequired) char rebootRequired;
@property (readonly, nonatomic, getter=isModified) char modified;

@end

#endif /* SPKernelExtensionPolicyItem_h */
