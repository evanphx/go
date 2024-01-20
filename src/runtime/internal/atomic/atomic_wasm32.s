// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//go:build wasm32

#include "textflag.h"

TEXT ·StorepNoWB(SB), NOSPLIT, $0-16
	MOVW ptr+0(FP), R0
	MOVW val+8(FP), 0(R0)
	RET
