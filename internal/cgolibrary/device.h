#pragma once

#include <stdint.h>

#ifdef __cplusplus
#include <torch/torch.h>
extern "C" {
#endif

#ifdef __cplusplus
typedef torch::Device *Device;
#else
typedef void *Device;
#endif

const char *Torch_Device(const char *device_type, Device *device);
void SetNumThreads(int32_t n);

#ifdef __cplusplus
}
#endif