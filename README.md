Open Questions
===

1. Should we delete the folder when deleting the PVC?
   NOTE: this might be hard, we'd have to mount it in the controller pod, currently only mounting from the node.
2. Do we want the controller to read the FS info?
   This allows easier setup, the StorageClass only needs an ID,  but requires the controller to auth and use AWS APIs.
   Otherwise, it could work like static provisioning and not hit the AWS API from the controller.
3. Do we need to worry about separating customers?
   Can one customer re-mount this volume to get access to other volumes?
   While mostly a case of blocking network access like anything else, we should be aware of it. The lustre FS IP+name is leaked by mount.
4. How do you want folder names to look / how should they be configured?
   How can they be configured, what data do we have.
   Currently: 
   - /pvc-33a152db-2915-4f63-9a6a-8892854ec9f7/
   - /pvc-720e34cc-6ca1-4ac8-8e14-affe52a2ffb7/
   ....
