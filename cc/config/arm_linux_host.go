// Copyright 2016 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"strings"

	"android/soong/android"
)

var (
	// Extended cflags
	linuxArmCflags = []string{
		//		"-msse3",
		//		"-mfpmath=sse",
		"-m32",
		//		"-march=prescott",
		"-D_FILE_OFFSET_BITS=64",
		"-D_LARGEFILE_SOURCE=1",
	}

	linuxAarch64Cflags = []string{
		"-m64",
	}

	linuxArmLdflags = []string{
		"-m32",
	}

	linuxAarch64Ldflags = []string{
		"-m64",
	}

	linuxArmClangLdflags = append(ClangFilterUnknownCflags(linuxArmLdflags), []string{}...)

	linuxAarch64ClangLdflags = append(ClangFilterUnknownCflags(linuxAarch64Ldflags), []string{}...)

	linuxArmClangCppflags = []string{}

	linuxAarch64ClangCppflags = []string{}
)

func init() {

	// Extended cflags
	pctx.StaticVariable("LinuxArmCflags", strings.Join(linuxArmCflags, " "))
	pctx.StaticVariable("LinuxAarch64Cflags", strings.Join(linuxAarch64Cflags, " "))
	pctx.StaticVariable("LinuxArmLdflags", strings.Join(linuxArmLdflags, " "))
	pctx.StaticVariable("LinuxAarch64Ldflags", strings.Join(linuxAarch64Ldflags, " "))

	pctx.StaticVariable("LinuxArmClangCflags",
		strings.Join(ClangFilterUnknownCflags(linuxArmCflags), " "))
	pctx.StaticVariable("LinuxAarch64ClangCflags",
		strings.Join(ClangFilterUnknownCflags(linuxAarch64Cflags), " "))
	pctx.StaticVariable("LinuxArmClangLdflags", strings.Join(linuxArmClangLdflags, " "))
	pctx.StaticVariable("LinuxAarch64ClangLdflags", strings.Join(linuxAarch64ClangLdflags, " "))
	pctx.StaticVariable("LinuxArmClangCppflags", strings.Join(linuxArmClangCppflags, " "))
	pctx.StaticVariable("LinuxAarch64ClangCppflags", strings.Join(linuxAarch64ClangCppflags, " "))
	// Yasm flags
	pctx.StaticVariable("LinuxArmYasmFlags", "-f elf32 -m x86")
	pctx.StaticVariable("LinuxAarch64YasmFlags", "-f elf64 -m amd64")
}

type toolchainLinuxArm struct {
	toolchain32Bit
	toolchainLinux
}

type toolchainLinuxAarch64 struct {
	toolchain64Bit
	toolchainLinux
}

func (t *toolchainLinuxArm) Name() string {
	return "arm"
}

func (t *toolchainLinuxAarch64) Name() string {
	return "aarch64"
}

func (t *toolchainLinuxArm) Cflags() string {
	return "${config.LinuxCflags} ${config.LinuxArmCflags}"
}

func (t *toolchainLinuxAarch64) Cflags() string {
	return "${config.LinuxCflags} ${config.LinuxAarch64Cflags}"
}

func (t *toolchainLinuxArm) Ldflags() string {
	return "${config.LinuxLdflags} ${config.LinuxArmLdflags}"
}

func (t *toolchainLinuxAarch64) Ldflags() string {
	return "${config.LinuxLdflags} ${config.LinuxAarch64Ldflags}"
}

func (t *toolchainLinuxArm) ClangTriple() string {
	return "aarch64-unknown-linux-gnu"
}

func (t *toolchainLinuxArm) ClangCflags() string {
	return "${config.LinuxClangCflags} ${config.LinuxArmClangCflags}"
}

func (t *toolchainLinuxArm) ClangCppflags() string {
	return "${config.LinuxClangCppflags} ${config.LinuxArmClangCppflags}"
}

func (t *toolchainLinuxAarch64) ClangTriple() string {
	return "aarch64-unknown-linux-gnu"
}

func (t *toolchainLinuxAarch64) ClangCflags() string {
	return "${config.LinuxClangCflags} ${config.LinuxAarch64ClangCflags}"
}

func (t *toolchainLinuxAarch64) ClangCppflags() string {
	return "${config.LinuxClangCppflags} ${config.LinuxAarch64ClangCppflags}"
}

func (t *toolchainLinuxArm) ClangLdflags() string {
	return "${config.LinuxClangLdflags} ${config.LinuxArmClangLdflags}"
}

func (t *toolchainLinuxAarch64) ClangLdflags() string {
	return "${config.LinuxClangLdflags} ${config.LinuxAarch64ClangLdflags}"
}

func (t *toolchainLinuxArm) YasmFlags() string {
	return "${config.LinuxArmYasmFlags}"
}

func (t *toolchainLinuxAarch64) YasmFlags() string {
	return "${config.LinuxAarch64YasmFlags}"
}

func (t *toolchainLinuxArm) InstructionSetFlags(isa string) (string, error) {
    switch isa {
    case "arm":
        return "-O2 -fomit-frame-pointer -fstrict-aliasing -funswitch-loops", nil
    case "thumb":
        return "-mthumb -Os -fomit-frame-pointer -fno-strict-aliasing", nil
    default:
        return t.toolchainBase.InstructionSetFlags(isa)
    }
}

func (t *toolchainLinuxArm) ClangInstructionSetFlags(isa string) (string, error) {
    switch isa {
    case "arm":
        return "-O2 -fomit-frame-pointer -fstrict-aliasing -funswitch-loops", nil
    case "thumb":
        return "-mthumb -Os -fomit-frame-pointer -fno-strict-aliasing", nil
    default:
        return t.toolchainBase.ClangInstructionSetFlags(isa)
    }
}

func (t *toolchainLinuxAarch64) InstructionSetFlags(isa string) (string, error) {
    switch isa {
    case "arm":
        return "-O2 -fomit-frame-pointer -fstrict-aliasing -funswitch-loops", nil
    case "thumb":
        return "-mthumb -Os -fomit-frame-pointer -fno-strict-aliasing", nil
    default:
        return t.toolchainBase.InstructionSetFlags(isa)
    }
}

func (t *toolchainLinuxAarch64) ClangInstructionSetFlags(isa string) (string, error) {
    switch isa {
    case "arm":
        return "-O2 -fomit-frame-pointer -fstrict-aliasing -funswitch-loops", nil
    case "thumb":
        return "-mthumb -Os -fomit-frame-pointer -fno-strict-aliasing", nil
    default:
        return t.toolchainBase.ClangInstructionSetFlags(isa)
    }
}

var toolchainLinuxArmSingleton Toolchain = &toolchainLinuxArm{}
var toolchainLinuxAarch64Singleton Toolchain = &toolchainLinuxAarch64{}

func linuxArmToolchainFactory(arch android.Arch) Toolchain {
	return toolchainLinuxArmSingleton
}

func linuxAarch64ToolchainFactory(arch android.Arch) Toolchain {
	return toolchainLinuxAarch64Singleton
}

func init() {
	registerToolchainFactory(android.Linux, android.Arm, linuxArmToolchainFactory)
	registerToolchainFactory(android.Linux, android.Aarch64, linuxAarch64ToolchainFactory)
}
