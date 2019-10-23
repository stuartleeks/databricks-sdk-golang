// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package models

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AclItem) DeepCopyInto(out *AclItem) {
	*out = *in
	if in.Permission != nil {
		in, out := &in.Permission, &out.Permission
		*out = new(AclPermission)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AclItem.
func (in *AclItem) DeepCopy() *AclItem {
	if in == nil {
		return nil
	}
	out := new(AclItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScale) DeepCopyInto(out *AutoScale) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScale.
func (in *AutoScale) DeepCopy() *AutoScale {
	if in == nil {
		return nil
	}
	out := new(AutoScale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAttributes) DeepCopyInto(out *ClusterAttributes) {
	*out = *in
	if in.SparkConf != nil {
		in, out := &in.SparkConf, &out.SparkConf
		*out = new(SparkConfPair)
		**out = **in
	}
	if in.SSHPublicKeys != nil {
		in, out := &in.SSHPublicKeys, &out.SSHPublicKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomTags != nil {
		in, out := &in.CustomTags, &out.CustomTags
		*out = make([]ClusterTag, len(*in))
		copy(*out, *in)
	}
	if in.ClusterLogConf != nil {
		in, out := &in.ClusterLogConf, &out.ClusterLogConf
		*out = new(ClusterLogConf)
		(*in).DeepCopyInto(*out)
	}
	if in.InitScripts != nil {
		in, out := &in.InitScripts, &out.InitScripts
		*out = make([]InitScriptInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SparkEnvVars != nil {
		in, out := &in.SparkEnvVars, &out.SparkEnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ClusterSource != nil {
		in, out := &in.ClusterSource, &out.ClusterSource
		*out = new(ClusterSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAttributes.
func (in *ClusterAttributes) DeepCopy() *ClusterAttributes {
	if in == nil {
		return nil
	}
	out := new(ClusterAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCloudProviderNodeInfo) DeepCopyInto(out *ClusterCloudProviderNodeInfo) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ClusterCloudProviderNodeStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCloudProviderNodeInfo.
func (in *ClusterCloudProviderNodeInfo) DeepCopy() *ClusterCloudProviderNodeInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterCloudProviderNodeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEvent) DeepCopyInto(out *ClusterEvent) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ClusterEventType)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = new(EventDetails)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEvent.
func (in *ClusterEvent) DeepCopy() *ClusterEvent {
	if in == nil {
		return nil
	}
	out := new(ClusterEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	if in.AutoScale != nil {
		in, out := &in.AutoScale, &out.AutoScale
		*out = new(AutoScale)
		**out = **in
	}
	if in.Driver != nil {
		in, out := &in.Driver, &out.Driver
		*out = new(SparkNode)
		**out = **in
	}
	if in.Executors != nil {
		in, out := &in.Executors, &out.Executors
		*out = make([]SparkNode, len(*in))
		copy(*out, *in)
	}
	if in.SparkConf != nil {
		in, out := &in.SparkConf, &out.SparkConf
		*out = new(SparkConfPair)
		**out = **in
	}
	if in.ClusterLogConf != nil {
		in, out := &in.ClusterLogConf, &out.ClusterLogConf
		*out = new(ClusterLogConf)
		(*in).DeepCopyInto(*out)
	}
	if in.InitScripts != nil {
		in, out := &in.InitScripts, &out.InitScripts
		*out = make([]InitScriptInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SparkEnvVars != nil {
		in, out := &in.SparkEnvVars, &out.SparkEnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterState)
		**out = **in
	}
	if in.DefaultTags != nil {
		in, out := &in.DefaultTags, &out.DefaultTags
		*out = make([]ClusterTag, len(*in))
		copy(*out, *in)
	}
	if in.ClusterLogStatus != nil {
		in, out := &in.ClusterLogStatus, &out.ClusterLogStatus
		*out = new(LogSyncStatus)
		**out = **in
	}
	if in.TerminationReason != nil {
		in, out := &in.TerminationReason, &out.TerminationReason
		*out = new(TerminationReason)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInstance) DeepCopyInto(out *ClusterInstance) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInstance.
func (in *ClusterInstance) DeepCopy() *ClusterInstance {
	if in == nil {
		return nil
	}
	out := new(ClusterInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLibraryStatuses) DeepCopyInto(out *ClusterLibraryStatuses) {
	*out = *in
	if in.LibraryStatuses != nil {
		in, out := &in.LibraryStatuses, &out.LibraryStatuses
		*out = make([]LibraryFullStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLibraryStatuses.
func (in *ClusterLibraryStatuses) DeepCopy() *ClusterLibraryStatuses {
	if in == nil {
		return nil
	}
	out := new(ClusterLibraryStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogConf) DeepCopyInto(out *ClusterLogConf) {
	*out = *in
	if in.Dbfs != nil {
		in, out := &in.Dbfs, &out.Dbfs
		*out = new(DbfsStorageInfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogConf.
func (in *ClusterLogConf) DeepCopy() *ClusterLogConf {
	if in == nil {
		return nil
	}
	out := new(ClusterLogConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSize) DeepCopyInto(out *ClusterSize) {
	*out = *in
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(AutoScale)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSize.
func (in *ClusterSize) DeepCopy() *ClusterSize {
	if in == nil {
		return nil
	}
	out := new(ClusterSize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.NewCluster != nil {
		in, out := &in.NewCluster, &out.NewCluster
		*out = new(NewCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.Libraries != nil {
		in, out := &in.Libraries, &out.Libraries
		*out = make([]Library, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTag) DeepCopyInto(out *ClusterTag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTag.
func (in *ClusterTag) DeepCopy() *ClusterTag {
	if in == nil {
		return nil
	}
	out := new(ClusterTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronSchedule) DeepCopyInto(out *CronSchedule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronSchedule.
func (in *CronSchedule) DeepCopy() *CronSchedule {
	if in == nil {
		return nil
	}
	out := new(CronSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbfsStorageInfo) DeepCopyInto(out *DbfsStorageInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbfsStorageInfo.
func (in *DbfsStorageInfo) DeepCopy() *DbfsStorageInfo {
	if in == nil {
		return nil
	}
	out := new(DbfsStorageInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpec) DeepCopyInto(out *DiskSpec) {
	*out = *in
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(DiskType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpec.
func (in *DiskSpec) DeepCopy() *DiskSpec {
	if in == nil {
		return nil
	}
	out := new(DiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskType) DeepCopyInto(out *DiskType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskType.
func (in *DiskType) DeepCopy() *DiskType {
	if in == nil {
		return nil
	}
	out := new(DiskType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDetails) DeepCopyInto(out *EventDetails) {
	*out = *in
	if in.PreviousAttributes != nil {
		in, out := &in.PreviousAttributes, &out.PreviousAttributes
		*out = new(ClusterAttributes)
		(*in).DeepCopyInto(*out)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = new(ClusterAttributes)
		(*in).DeepCopyInto(*out)
	}
	if in.PreviousClusterSize != nil {
		in, out := &in.PreviousClusterSize, &out.PreviousClusterSize
		*out = new(ClusterSize)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSize != nil {
		in, out := &in.ClusterSize, &out.ClusterSize
		*out = new(ClusterSize)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDetails.
func (in *EventDetails) DeepCopy() *EventDetails {
	if in == nil {
		return nil
	}
	out := new(EventDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileInfo) DeepCopyInto(out *FileInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileInfo.
func (in *FileInfo) DeepCopy() *FileInfo {
	if in == nil {
		return nil
	}
	out := new(FileInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitScriptInfo) DeepCopyInto(out *InitScriptInfo) {
	*out = *in
	if in.Dbfs != nil {
		in, out := &in.Dbfs, &out.Dbfs
		*out = new(DbfsStorageInfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitScriptInfo.
func (in *InitScriptInfo) DeepCopy() *InitScriptInfo {
	if in == nil {
		return nil
	}
	out := new(InitScriptInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePoolAndStats) DeepCopyInto(out *InstancePoolAndStats) {
	*out = *in
	if in.CustomTags != nil {
		in, out := &in.CustomTags, &out.CustomTags
		*out = make([]ClusterTag, len(*in))
		copy(*out, *in)
	}
	in.DiskSpec.DeepCopyInto(&out.DiskSpec)
	if in.PreloadedSparkVersions != nil {
		in, out := &in.PreloadedSparkVersions, &out.PreloadedSparkVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultTags != nil {
		in, out := &in.DefaultTags, &out.DefaultTags
		*out = make([]ClusterTag, len(*in))
		copy(*out, *in)
	}
	out.Stats = in.Stats
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePoolAndStats.
func (in *InstancePoolAndStats) DeepCopy() *InstancePoolAndStats {
	if in == nil {
		return nil
	}
	out := new(InstancePoolAndStats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstancePoolStats) DeepCopyInto(out *InstancePoolStats) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstancePoolStats.
func (in *InstancePoolStats) DeepCopy() *InstancePoolStats {
	if in == nil {
		return nil
	}
	out := new(InstancePoolStats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(JobSettings)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobEmailNotifications) DeepCopyInto(out *JobEmailNotifications) {
	*out = *in
	if in.OnStart != nil {
		in, out := &in.OnStart, &out.OnStart
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OnSuccess != nil {
		in, out := &in.OnSuccess, &out.OnSuccess
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobEmailNotifications.
func (in *JobEmailNotifications) DeepCopy() *JobEmailNotifications {
	if in == nil {
		return nil
	}
	out := new(JobEmailNotifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSettings) DeepCopyInto(out *JobSettings) {
	*out = *in
	if in.NewCluster != nil {
		in, out := &in.NewCluster, &out.NewCluster
		*out = new(NewCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.NotebookTask != nil {
		in, out := &in.NotebookTask, &out.NotebookTask
		*out = new(NotebookTask)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkJarTask != nil {
		in, out := &in.SparkJarTask, &out.SparkJarTask
		*out = new(SparkJarTask)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkPythonTask != nil {
		in, out := &in.SparkPythonTask, &out.SparkPythonTask
		*out = new(SparkPythonTask)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkSubmitTask != nil {
		in, out := &in.SparkSubmitTask, &out.SparkSubmitTask
		*out = new(SparkSubmitTask)
		(*in).DeepCopyInto(*out)
	}
	if in.Libraries != nil {
		in, out := &in.Libraries, &out.Libraries
		*out = make([]Library, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EmailNotifications != nil {
		in, out := &in.EmailNotifications, &out.EmailNotifications
		*out = new(JobEmailNotifications)
		(*in).DeepCopyInto(*out)
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(CronSchedule)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSettings.
func (in *JobSettings) DeepCopy() *JobSettings {
	if in == nil {
		return nil
	}
	out := new(JobSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTask) DeepCopyInto(out *JobTask) {
	*out = *in
	if in.NotebookTask != nil {
		in, out := &in.NotebookTask, &out.NotebookTask
		*out = new(NotebookTask)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkJarTask != nil {
		in, out := &in.SparkJarTask, &out.SparkJarTask
		*out = new(SparkJarTask)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkPythonTask != nil {
		in, out := &in.SparkPythonTask, &out.SparkPythonTask
		*out = new(SparkPythonTask)
		(*in).DeepCopyInto(*out)
	}
	if in.SparkSubmitTask != nil {
		in, out := &in.SparkSubmitTask, &out.SparkSubmitTask
		*out = new(SparkSubmitTask)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTask.
func (in *JobTask) DeepCopy() *JobTask {
	if in == nil {
		return nil
	}
	out := new(JobTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Library) DeepCopyInto(out *Library) {
	*out = *in
	if in.Pypi != nil {
		in, out := &in.Pypi, &out.Pypi
		*out = new(PythonPyPiLibrary)
		**out = **in
	}
	if in.Maven != nil {
		in, out := &in.Maven, &out.Maven
		*out = new(MavenLibrary)
		(*in).DeepCopyInto(*out)
	}
	if in.Cran != nil {
		in, out := &in.Cran, &out.Cran
		*out = new(RCranLibrary)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Library.
func (in *Library) DeepCopy() *Library {
	if in == nil {
		return nil
	}
	out := new(Library)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryFullStatus) DeepCopyInto(out *LibraryFullStatus) {
	*out = *in
	if in.Library != nil {
		in, out := &in.Library, &out.Library
		*out = new(Library)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(LibraryInstallStatus)
		**out = **in
	}
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryFullStatus.
func (in *LibraryFullStatus) DeepCopy() *LibraryFullStatus {
	if in == nil {
		return nil
	}
	out := new(LibraryFullStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSyncStatus) DeepCopyInto(out *LogSyncStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSyncStatus.
func (in *LogSyncStatus) DeepCopy() *LogSyncStatus {
	if in == nil {
		return nil
	}
	out := new(LogSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MavenLibrary) DeepCopyInto(out *MavenLibrary) {
	*out = *in
	if in.Exclusions != nil {
		in, out := &in.Exclusions, &out.Exclusions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MavenLibrary.
func (in *MavenLibrary) DeepCopy() *MavenLibrary {
	if in == nil {
		return nil
	}
	out := new(MavenLibrary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewCluster) DeepCopyInto(out *NewCluster) {
	*out = *in
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(AutoScale)
		**out = **in
	}
	if in.SparkConf != nil {
		in, out := &in.SparkConf, &out.SparkConf
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CustomTags != nil {
		in, out := &in.CustomTags, &out.CustomTags
		*out = make([]ClusterTag, len(*in))
		copy(*out, *in)
	}
	if in.ClusterLogConf != nil {
		in, out := &in.ClusterLogConf, &out.ClusterLogConf
		*out = new(ClusterLogConf)
		(*in).DeepCopyInto(*out)
	}
	if in.InitScripts != nil {
		in, out := &in.InitScripts, &out.InitScripts
		*out = make([]InitScriptInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SparkEnvVars != nil {
		in, out := &in.SparkEnvVars, &out.SparkEnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewCluster.
func (in *NewCluster) DeepCopy() *NewCluster {
	if in == nil {
		return nil
	}
	out := new(NewCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeType) DeepCopyInto(out *NodeType) {
	*out = *in
	if in.NodeInfo != nil {
		in, out := &in.NodeInfo, &out.NodeInfo
		*out = new(ClusterCloudProviderNodeInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeType.
func (in *NodeType) DeepCopy() *NodeType {
	if in == nil {
		return nil
	}
	out := new(NodeType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookOutput) DeepCopyInto(out *NotebookOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookOutput.
func (in *NotebookOutput) DeepCopy() *NotebookOutput {
	if in == nil {
		return nil
	}
	out := new(NotebookOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookTask) DeepCopyInto(out *NotebookTask) {
	*out = *in
	if in.BaseParameters != nil {
		in, out := &in.BaseParameters, &out.BaseParameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookTask.
func (in *NotebookTask) DeepCopy() *NotebookTask {
	if in == nil {
		return nil
	}
	out := new(NotebookTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectInfo) DeepCopyInto(out *ObjectInfo) {
	*out = *in
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(ObjectType)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(Language)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectInfo.
func (in *ObjectInfo) DeepCopy() *ObjectInfo {
	if in == nil {
		return nil
	}
	out := new(ObjectInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParamPair) DeepCopyInto(out *ParamPair) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParamPair.
func (in *ParamPair) DeepCopy() *ParamPair {
	if in == nil {
		return nil
	}
	out := new(ParamPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterPair) DeepCopyInto(out *ParameterPair) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterPair.
func (in *ParameterPair) DeepCopy() *ParameterPair {
	if in == nil {
		return nil
	}
	out := new(ParameterPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalName) DeepCopyInto(out *PrincipalName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalName.
func (in *PrincipalName) DeepCopy() *PrincipalName {
	if in == nil {
		return nil
	}
	out := new(PrincipalName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicTokenInfo) DeepCopyInto(out *PublicTokenInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicTokenInfo.
func (in *PublicTokenInfo) DeepCopy() *PublicTokenInfo {
	if in == nil {
		return nil
	}
	out := new(PublicTokenInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PythonPyPiLibrary) DeepCopyInto(out *PythonPyPiLibrary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PythonPyPiLibrary.
func (in *PythonPyPiLibrary) DeepCopy() *PythonPyPiLibrary {
	if in == nil {
		return nil
	}
	out := new(PythonPyPiLibrary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RCranLibrary) DeepCopyInto(out *RCranLibrary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RCranLibrary.
func (in *RCranLibrary) DeepCopy() *RCranLibrary {
	if in == nil {
		return nil
	}
	out := new(RCranLibrary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Run) DeepCopyInto(out *Run) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RunState)
		(*in).DeepCopyInto(*out)
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(CronSchedule)
		**out = **in
	}
	if in.Task != nil {
		in, out := &in.Task, &out.Task
		*out = new(JobTask)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSpec != nil {
		in, out := &in.ClusterSpec, &out.ClusterSpec
		*out = new(ClusterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterInstance != nil {
		in, out := &in.ClusterInstance, &out.ClusterInstance
		*out = new(ClusterInstance)
		**out = **in
	}
	if in.OverridingParameters != nil {
		in, out := &in.OverridingParameters, &out.OverridingParameters
		*out = new(RunParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Trigger != nil {
		in, out := &in.Trigger, &out.Trigger
		*out = new(TriggerType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Run.
func (in *Run) DeepCopy() *Run {
	if in == nil {
		return nil
	}
	out := new(Run)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunParameters) DeepCopyInto(out *RunParameters) {
	*out = *in
	if in.JarParams != nil {
		in, out := &in.JarParams, &out.JarParams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotebookParams != nil {
		in, out := &in.NotebookParams, &out.NotebookParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PythonParams != nil {
		in, out := &in.PythonParams, &out.PythonParams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SparkSubmitParams != nil {
		in, out := &in.SparkSubmitParams, &out.SparkSubmitParams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunParameters.
func (in *RunParameters) DeepCopy() *RunParameters {
	if in == nil {
		return nil
	}
	out := new(RunParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunState) DeepCopyInto(out *RunState) {
	*out = *in
	if in.LifeCycleState != nil {
		in, out := &in.LifeCycleState, &out.LifeCycleState
		*out = new(RunLifeCycleState)
		**out = **in
	}
	if in.ResultState != nil {
		in, out := &in.ResultState, &out.ResultState
		*out = new(RunResultState)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunState.
func (in *RunState) DeepCopy() *RunState {
	if in == nil {
		return nil
	}
	out := new(RunState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMetadata) DeepCopyInto(out *SecretMetadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMetadata.
func (in *SecretMetadata) DeepCopy() *SecretMetadata {
	if in == nil {
		return nil
	}
	out := new(SecretMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretScope) DeepCopyInto(out *SecretScope) {
	*out = *in
	if in.BackendType != nil {
		in, out := &in.BackendType, &out.BackendType
		*out = new(ScopeBackendType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretScope.
func (in *SecretScope) DeepCopy() *SecretScope {
	if in == nil {
		return nil
	}
	out := new(SecretScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkConfPair) DeepCopyInto(out *SparkConfPair) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkConfPair.
func (in *SparkConfPair) DeepCopy() *SparkConfPair {
	if in == nil {
		return nil
	}
	out := new(SparkConfPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkEnvPair) DeepCopyInto(out *SparkEnvPair) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkEnvPair.
func (in *SparkEnvPair) DeepCopy() *SparkEnvPair {
	if in == nil {
		return nil
	}
	out := new(SparkEnvPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkJarTask) DeepCopyInto(out *SparkJarTask) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkJarTask.
func (in *SparkJarTask) DeepCopy() *SparkJarTask {
	if in == nil {
		return nil
	}
	out := new(SparkJarTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkNode) DeepCopyInto(out *SparkNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkNode.
func (in *SparkNode) DeepCopy() *SparkNode {
	if in == nil {
		return nil
	}
	out := new(SparkNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkPythonTask) DeepCopyInto(out *SparkPythonTask) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkPythonTask.
func (in *SparkPythonTask) DeepCopy() *SparkPythonTask {
	if in == nil {
		return nil
	}
	out := new(SparkPythonTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkSubmitTask) DeepCopyInto(out *SparkSubmitTask) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkSubmitTask.
func (in *SparkSubmitTask) DeepCopy() *SparkSubmitTask {
	if in == nil {
		return nil
	}
	out := new(SparkSubmitTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkVersion) DeepCopyInto(out *SparkVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkVersion.
func (in *SparkVersion) DeepCopy() *SparkVersion {
	if in == nil {
		return nil
	}
	out := new(SparkVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminationReason) DeepCopyInto(out *TerminationReason) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(TerminationCode)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParameterPair, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminationReason.
func (in *TerminationReason) DeepCopy() *TerminationReason {
	if in == nil {
		return nil
	}
	out := new(TerminationReason)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewItem) DeepCopyInto(out *ViewItem) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ViewType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewItem.
func (in *ViewItem) DeepCopy() *ViewItem {
	if in == nil {
		return nil
	}
	out := new(ViewItem)
	in.DeepCopyInto(out)
	return out
}
