package enum

type MetrcisMappingCategory string

const (
	MetrcisMappingCategory_Cpu     MetrcisMappingCategory = "CPU"
	MetrcisMappingCategory_Gpu     MetrcisMappingCategory = "GPU"
	MetrcisMappingCategory_Memory  MetrcisMappingCategory = "MEMORY"
	MetrcisMappingCategory_Disk    MetrcisMappingCategory = "DISK"
	MetrcisMappingCategory_Network MetrcisMappingCategory = "NETWORK"
)
