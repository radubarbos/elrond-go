package builtInFunctions

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/core/vmcommon"
	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/process"
)

var _ process.BuiltinFunction = (*changeOwnerAddress)(nil)

type changeOwnerAddress struct {
	gasCost      uint64
	mutExecution sync.RWMutex
}

// NewChangeOwnerAddressFunc create a new change owner built in function
func NewChangeOwnerAddressFunc(gasCost uint64) *changeOwnerAddress {
	return &changeOwnerAddress{gasCost: gasCost}
}

// SetNewGasConfig is called whenever gas cost is changed
func (c *changeOwnerAddress) SetNewGasConfig(gasCost *process.GasCost) {
	if gasCost == nil {
		return
	}

	c.mutExecution.Lock()
	c.gasCost = gasCost.BuiltInCost.ChangeOwnerAddress
	c.mutExecution.Unlock()
}

// ProcessBuiltinFunction processes simple protocol built-in function
func (c *changeOwnerAddress) ProcessBuiltinFunction(
	acntSnd, acntDst state.UserAccountHandler,
	vmInput *vmcommon.ContractCallInput,
) (*vmcommon.VMOutput, error) {
	c.mutExecution.RLock()
	defer c.mutExecution.RUnlock()

	if vmInput == nil {
		return nil, process.ErrNilVmInput
	}
	if len(vmInput.Arguments) == 0 {
		return nil, process.ErrInvalidArguments
	}
	if vmInput.CallValue.Cmp(zero) != 0 {
		return nil, process.ErrBuiltInFunctionCalledWithValue
	}
	if len(vmInput.Arguments[0]) != len(vmInput.CallerAddr) {
		return nil, process.ErrInvalidAddressLength
	}
	if vmInput.GasProvided < c.gasCost {
		return nil, process.ErrNotEnoughGas
	}
	gasRemaining := computeGasRemaining(acntSnd, vmInput.GasProvided, c.gasCost)
	if check.IfNil(acntDst) {
		// cross-shard call, in sender shard only the gas is taken out
		return &vmcommon.VMOutput{ReturnCode: vmcommon.Ok, GasRemaining: gasRemaining}, nil
	}

	if !bytes.Equal(vmInput.CallerAddr, acntDst.GetOwnerAddress()) {
		return nil, fmt.Errorf("%w not the owner of the account", process.ErrOperationNotPermitted)
	}

	err := acntDst.ChangeOwnerAddress(vmInput.CallerAddr, vmInput.Arguments[0])
	if err != nil {
		return nil, err
	}

	return &vmcommon.VMOutput{GasRemaining: gasRemaining, ReturnCode: vmcommon.Ok}, nil
}

func computeGasRemaining(snd state.UserAccountHandler, gasProvided uint64, gasToUse uint64) uint64 {
	if gasProvided < gasToUse {
		return 0
	}
	// in case of built in functions - gas is consumed in sender shard, returned already in sender shard
	// thus we must return with 0 here
	if check.IfNil(snd) {
		return 0
	}

	return gasProvided - gasToUse
}

// IsInterfaceNil returns true if underlying object in nil
func (c *changeOwnerAddress) IsInterfaceNil() bool {
	return c == nil
}
