# Distributed under the OSI-approved BSD 3-Clause License.  See accompanying
# file Copyright.txt or https://cmake.org/licensing for details.

cmake_minimum_required(VERSION 3.5)

file(MAKE_DIRECTORY
  "/Users/lino/Research/pico-sdk/tools/pioasm"
  "/Users/lino/Research/tiny jambu/C Implementation/v2/build/pioasm"
  "/Users/lino/Research/tiny jambu/C Implementation/v2/build/pico-sdk/src/rp2_common/cyw43_driver/pioasm"
  "/Users/lino/Research/tiny jambu/C Implementation/v2/build/pico-sdk/src/rp2_common/cyw43_driver/pioasm/tmp"
  "/Users/lino/Research/tiny jambu/C Implementation/v2/build/pico-sdk/src/rp2_common/cyw43_driver/pioasm/src/PioasmBuild-stamp"
  "/Users/lino/Research/tiny jambu/C Implementation/v2/build/pico-sdk/src/rp2_common/cyw43_driver/pioasm/src"
  "/Users/lino/Research/tiny jambu/C Implementation/v2/build/pico-sdk/src/rp2_common/cyw43_driver/pioasm/src/PioasmBuild-stamp"
)

set(configSubDirs )
foreach(subDir IN LISTS configSubDirs)
    file(MAKE_DIRECTORY "/Users/lino/Research/tiny jambu/C Implementation/v2/build/pico-sdk/src/rp2_common/cyw43_driver/pioasm/src/PioasmBuild-stamp/${subDir}")
endforeach()
