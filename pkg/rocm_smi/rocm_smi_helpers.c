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


#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>

#include "rocm_smi/rocm_smi_helpers.h"
#include "rocm_smi/rocm_smi.h"
#include "rocm_smi/kfd_ioctl.h"



// Basically a copy from the lookup code in rocm_smi.h (Supported Functions)
rsmi_status_t rsmi_helper_func_variants_get(uint32_t deviceIndex, rsmi_helper_function_t* functions)
{
    rsmi_func_id_iter_handle_t iter_handle, var_iter, sub_var_iter;
    rsmi_func_id_value_t value;
    rsmi_status_t err;
    rsmi_status_t last_err;

    int num_funcs = 0;

    err = rsmi_dev_supported_func_iterator_open(deviceIndex, &iter_handle);
    if (err == RSMI_STATUS_SUCCESS)
    {
        while (1) {
            err = rsmi_func_iter_value_get(iter_handle, &value);
            if (err == RSMI_STATUS_SUCCESS)
            {
                num_funcs++;
            }
            err = rsmi_func_iter_next(iter_handle);
            if (err == RSMI_STATUS_NO_DATA) {
                break;
            }
        }
        err = rsmi_dev_supported_func_iterator_close(&iter_handle);
    } else {
        return err;
    }
    if (num_funcs == 0)
    {
        return RSMI_STATUS_NO_DATA;
    }
    last_err = err;

    err = rsmi_dev_supported_func_iterator_open(deviceIndex, &iter_handle);

    while (1) {
        err = rsmi_func_iter_value_get(iter_handle, &value);
        if (err != RSMI_STATUS_SUCCESS)
        {
            err = rsmi_func_iter_next(iter_handle);
            if (err == RSMI_STATUS_NO_DATA) {
                last_err = err;
                break;
            }
            continue;
        }
        rsmi_helper_function_variants_t* func = &functions->functionlist[functions->num_functions];
        func->num_sensors = 0;

        memset(func->name, '\0', RSMI_HELPER_MAX_NAME_LEN*sizeof(char));
        int ret = snprintf(func->name, RSMI_HELPER_MAX_NAME_LEN-1, "%s", value.name);
        if (ret >= 0)
        {
            func->name[ret] = '\0';
        }

        err = rsmi_dev_supported_variant_iterator_open(iter_handle, &var_iter);
        if (err != RSMI_STATUS_NO_DATA) {
            while (1) {
                err = rsmi_func_iter_value_get(var_iter, &value);
                
                rsmi_helper_variants_t* variants = &func->sensorlist[func->num_sensors];
                variants->num_variants = 0;
                if (value.id == RSMI_DEFAULT_VARIANT) {
                    variants->variant_id = RSMI_DEFAULT_VARIANT;
                } else {
                    variants->variant_id = value.id;
                }

                err = rsmi_dev_supported_variant_iterator_open(var_iter, &sub_var_iter);
                if (err != RSMI_STATUS_NO_DATA) {

                    while (1) {
                        err = rsmi_func_iter_value_get(sub_var_iter, &value);

                        variants->variantlist[variants->num_variants] = value.id;
                        variants->num_variants++;

                        err = rsmi_func_iter_next(sub_var_iter);
                        if (err == RSMI_STATUS_NO_DATA) {
                            break;
                        }
                    }
                    err = rsmi_dev_supported_func_iterator_close(&sub_var_iter);
                }
                func->num_sensors++;
                err = rsmi_func_iter_next(var_iter);
                if (err == RSMI_STATUS_NO_DATA) {
                    break;
                }
            }
            err = rsmi_dev_supported_func_iterator_close(&var_iter);
        }
        functions->num_functions++;
        err = rsmi_func_iter_next(iter_handle);
        if (err == RSMI_STATUS_NO_DATA) {
            break;
        }
    }
    err = rsmi_dev_supported_func_iterator_close(&iter_handle);
    return (int)last_err;
}
