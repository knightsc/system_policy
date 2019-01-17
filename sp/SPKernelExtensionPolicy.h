#ifndef SPKernelExtensionPolicy_h
#define SPKernelExtensionPolicy_h

@class SPKernelExtensionPolicyItem;

@interface SPKernelExtensionPolicy : NSObject

- (instancetype)init;
- (NSArray<SPKernelExtensionPolicyItem *> *)currentPolicy;

@end

#endif /* SPKernelExtensionPolicy_h */
