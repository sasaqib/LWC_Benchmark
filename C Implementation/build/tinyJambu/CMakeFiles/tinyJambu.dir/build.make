# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.23

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/local/Cellar/cmake/3.23.2/bin/cmake

# The command to remove a file.
RM = /usr/local/Cellar/cmake/3.23.2/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = "/Users/lino/Research/tiny jambu/C Implementation/v2"

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = "/Users/lino/Research/tiny jambu/C Implementation/v2/build"

# Include any dependencies generated for this target.
include tinyJambu/CMakeFiles/tinyJambu.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include tinyJambu/CMakeFiles/tinyJambu.dir/compiler_depend.make

# Include the progress variables for this target.
include tinyJambu/CMakeFiles/tinyJambu.dir/progress.make

# Include the compile flags for this target's objects.
include tinyJambu/CMakeFiles/tinyJambu.dir/flags.make

tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.obj: tinyJambu/CMakeFiles/tinyJambu.dir/flags.make
tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.obj: ../tinyJambu/encrypt.c
tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.obj: tinyJambu/CMakeFiles/tinyJambu.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir="/Users/lino/Research/tiny jambu/C Implementation/v2/build/CMakeFiles" --progress-num=$(CMAKE_PROGRESS_1) "Building C object tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.obj"
	cd "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu" && /usr/local/bin/arm-none-eabi-gcc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.obj -MF CMakeFiles/tinyJambu.dir/encrypt.c.obj.d -o CMakeFiles/tinyJambu.dir/encrypt.c.obj -c "/Users/lino/Research/tiny jambu/C Implementation/v2/tinyJambu/encrypt.c"

tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/tinyJambu.dir/encrypt.c.i"
	cd "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu" && /usr/local/bin/arm-none-eabi-gcc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E "/Users/lino/Research/tiny jambu/C Implementation/v2/tinyJambu/encrypt.c" > CMakeFiles/tinyJambu.dir/encrypt.c.i

tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/tinyJambu.dir/encrypt.c.s"
	cd "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu" && /usr/local/bin/arm-none-eabi-gcc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S "/Users/lino/Research/tiny jambu/C Implementation/v2/tinyJambu/encrypt.c" -o CMakeFiles/tinyJambu.dir/encrypt.c.s

# Object files for target tinyJambu
tinyJambu_OBJECTS = \
"CMakeFiles/tinyJambu.dir/encrypt.c.obj"

# External object files for target tinyJambu
tinyJambu_EXTERNAL_OBJECTS =

tinyJambu/libtinyJambu.a: tinyJambu/CMakeFiles/tinyJambu.dir/encrypt.c.obj
tinyJambu/libtinyJambu.a: tinyJambu/CMakeFiles/tinyJambu.dir/build.make
tinyJambu/libtinyJambu.a: tinyJambu/CMakeFiles/tinyJambu.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir="/Users/lino/Research/tiny jambu/C Implementation/v2/build/CMakeFiles" --progress-num=$(CMAKE_PROGRESS_2) "Linking C static library libtinyJambu.a"
	cd "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu" && $(CMAKE_COMMAND) -P CMakeFiles/tinyJambu.dir/cmake_clean_target.cmake
	cd "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu" && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/tinyJambu.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
tinyJambu/CMakeFiles/tinyJambu.dir/build: tinyJambu/libtinyJambu.a
.PHONY : tinyJambu/CMakeFiles/tinyJambu.dir/build

tinyJambu/CMakeFiles/tinyJambu.dir/clean:
	cd "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu" && $(CMAKE_COMMAND) -P CMakeFiles/tinyJambu.dir/cmake_clean.cmake
.PHONY : tinyJambu/CMakeFiles/tinyJambu.dir/clean

tinyJambu/CMakeFiles/tinyJambu.dir/depend:
	cd "/Users/lino/Research/tiny jambu/C Implementation/v2/build" && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" "/Users/lino/Research/tiny jambu/C Implementation/v2" "/Users/lino/Research/tiny jambu/C Implementation/v2/tinyJambu" "/Users/lino/Research/tiny jambu/C Implementation/v2/build" "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu" "/Users/lino/Research/tiny jambu/C Implementation/v2/build/tinyJambu/CMakeFiles/tinyJambu.dir/DependInfo.cmake" --color=$(COLOR)
.PHONY : tinyJambu/CMakeFiles/tinyJambu.dir/depend

