#include "cuda.h"

#ifdef WITH_CUDA
#include "c10/cuda/CUDAStream.h"
#endif

#include <torch/torch.h>

bool IsCUDAAvailable() {
  return torch::cuda::is_available();
}
