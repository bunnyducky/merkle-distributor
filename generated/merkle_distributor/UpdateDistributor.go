// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package merkle_distributor

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateDistributor is the `updateDistributor` instruction.
type UpdateDistributor struct {
	Root          *[32]uint8
	MaxTotalClaim *uint64
	MaxNumNodes   *uint64

	// [0] = [SIGNER] adminAuth
	//
	// [1] = [WRITE] distributor
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateDistributorInstructionBuilder creates a new `UpdateDistributor` instruction builder.
func NewUpdateDistributorInstructionBuilder() *UpdateDistributor {
	nd := &UpdateDistributor{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetRoot sets the "root" parameter.
func (inst *UpdateDistributor) SetRoot(root [32]uint8) *UpdateDistributor {
	inst.Root = &root
	return inst
}

// SetMaxTotalClaim sets the "maxTotalClaim" parameter.
func (inst *UpdateDistributor) SetMaxTotalClaim(maxTotalClaim uint64) *UpdateDistributor {
	inst.MaxTotalClaim = &maxTotalClaim
	return inst
}

// SetMaxNumNodes sets the "maxNumNodes" parameter.
func (inst *UpdateDistributor) SetMaxNumNodes(maxNumNodes uint64) *UpdateDistributor {
	inst.MaxNumNodes = &maxNumNodes
	return inst
}

// SetAdminAuthAccount sets the "adminAuth" account.
func (inst *UpdateDistributor) SetAdminAuthAccount(adminAuth ag_solanago.PublicKey) *UpdateDistributor {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(adminAuth).SIGNER()
	return inst
}

// GetAdminAuthAccount gets the "adminAuth" account.
func (inst *UpdateDistributor) GetAdminAuthAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetDistributorAccount sets the "distributor" account.
func (inst *UpdateDistributor) SetDistributorAccount(distributor ag_solanago.PublicKey) *UpdateDistributor {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(distributor).WRITE()
	return inst
}

// GetDistributorAccount gets the "distributor" account.
func (inst *UpdateDistributor) GetDistributorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst UpdateDistributor) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateDistributor,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateDistributor) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateDistributor) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Root == nil {
			return errors.New("Root parameter is not set")
		}
		if inst.MaxTotalClaim == nil {
			return errors.New("MaxTotalClaim parameter is not set")
		}
		if inst.MaxNumNodes == nil {
			return errors.New("MaxNumNodes parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AdminAuth is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Distributor is not set")
		}
	}
	return nil
}

func (inst *UpdateDistributor) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateDistributor")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("         Root", *inst.Root))
						paramsBranch.Child(ag_format.Param("MaxTotalClaim", *inst.MaxTotalClaim))
						paramsBranch.Child(ag_format.Param("  MaxNumNodes", *inst.MaxNumNodes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  adminAuth", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("distributor", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj UpdateDistributor) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	// Serialize `MaxTotalClaim` param:
	err = encoder.Encode(obj.MaxTotalClaim)
	if err != nil {
		return err
	}
	// Serialize `MaxNumNodes` param:
	err = encoder.Encode(obj.MaxNumNodes)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateDistributor) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	// Deserialize `MaxTotalClaim`:
	err = decoder.Decode(&obj.MaxTotalClaim)
	if err != nil {
		return err
	}
	// Deserialize `MaxNumNodes`:
	err = decoder.Decode(&obj.MaxNumNodes)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateDistributorInstruction declares a new UpdateDistributor instruction with the provided parameters and accounts.
func NewUpdateDistributorInstruction(
	// Parameters:
	root [32]uint8,
	maxTotalClaim uint64,
	maxNumNodes uint64,
	// Accounts:
	adminAuth ag_solanago.PublicKey,
	distributor ag_solanago.PublicKey) *UpdateDistributor {
	return NewUpdateDistributorInstructionBuilder().
		SetRoot(root).
		SetMaxTotalClaim(maxTotalClaim).
		SetMaxNumNodes(maxNumNodes).
		SetAdminAuthAccount(adminAuth).
		SetDistributorAccount(distributor)
}