#ifndef ROCM_SMI_HELPERS_H
#define ROCM_SMI_HELPERS_H

#include "rocm_smi.h"
#include "kfd_ioctl.h"

#define RSMI_HELPER_MAX_NAME_LEN 200
#define RSMI_HELPER_MAX_FUNCTIONS 300
#define RSMI_HELPER_MAX_SENSORS 30
#define RSMI_HELPER_MAX_VARIANTS 100



typedef struct RocmSmiHelperFunctionVariantStruct {
    int num_variants;
    unsigned long long variant_id;
    unsigned long long variantlist[RSMI_HELPER_MAX_VARIANTS];
} rsmi_helper_variants_t;

typedef struct RocmSmiHelperFunctionStruct {
    char name[RSMI_HELPER_MAX_NAME_LEN];
    int num_sensors;
    rsmi_helper_variants_t sensorlist[RSMI_HELPER_MAX_SENSORS];
} rsmi_helper_function_variants_t;

typedef struct {
    int num_functions;
    rsmi_helper_function_variants_t functionlist[RSMI_HELPER_MAX_FUNCTIONS];
} rsmi_helper_function_t;

rsmi_status_t rsmi_helper_func_variants_get(uint32_t deviceIndex, rsmi_helper_function_t* functions);

// #define RSMI_HELPER_MAX_BUILD_LEN 200
// typedef struct {
//     uint32_t major;
//     uint32_t minor;
//     uint32_t patch;
//     char build[RSMI_HELPER_MAX_BUILD_LEN];
// } rsmi_helper_version_t;
// rsmi_status_t rsmi_helper_version_get(rsmi_helper_version_t* version);

#endif