// Copyright (c) 2019 RRZE, University Erlangen-Nuremberg

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
