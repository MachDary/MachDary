package api

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/MachDary/MachDary/consensus/segwit"
	"github.com/MachDary/MachDary/protocol/vm"
)

// DecodeProgResp is response for decode program
type DecodeProgResp struct {
	Instructions string `json:"instructions"`
}

func (a *API) decodeProgram(ctx context.Context, ins struct {
	Program string `json:"program"`
}) Response {
	prog, err := hex.DecodeString(ins.Program)
	if err != nil {
		return NewErrorResponse(err)
	}

	// if program is P2SH script, convert it into actual executed program
	if segwit.IsP2WSHScript(prog) {
		if witnessProg, err := segwit.ConvertP2SHProgram(prog); err == nil {
			prog = witnessProg
		}
	}

	insts, err := vm.ParseProgram(prog)
	if err != nil {
		return NewErrorResponse(err)
	}

	var result string
	for _, inst := range insts {
		result += fmt.Sprintf("%s %s\n", inst.Op, hex.EncodeToString(inst.Data))
	}
	return NewSuccessResponse(DecodeProgResp{Instructions: result})
}
