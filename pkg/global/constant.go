package global

const (
	HadoopComponentNamesCmd ="jps | egrep 'NameNode|DataNode|JournalNode|DFSZKFailoverController|ResourceManager|NodeManager|HMaster' | awk '{print $2}'| sort |tr '\n' ' ' "
)
