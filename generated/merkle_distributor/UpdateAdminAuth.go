// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package merkle_distributor

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateAdminAuth is the `updateAdminAuth` instruction.
type UpdateAdminAuth struct {

	// [0] = [SIGNER] newAdminAuth
	//
	// [1] = [SIGNER] adminAuth
	//
	// [2] = [WRITE] distributor
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateAdminAuthInstructionBuilder creates a new `UpdateAdminAuth` instruction builder.
func NewUpdateAdminAuthInstructionBuilder() *UpdateAdminAuth {
	nd := &UpdateAdminAuth{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetNewAdminAuthAccount sets the "newAdminAuth" account.
func (inst *UpdateAdminAuth) SetNewAdminAuthAccount(newAdminAuth ag_solanago.PublicKey) *UpdateAdminAuth {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(newAdminAuth).SIGNER()
	return inst
}

// GetNewAdminAuthAccount gets the "newAdminAuth" account.
func (inst *UpdateAdminAuth) GetNewAdminAuthAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAuthAccount sets the "adminAuth" account.
func (inst *UpdateAdminAuth) SetAdminAuthAccount(adminAuth ag_solanago.PublicKey) *UpdateAdminAuth {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(adminAuth).SIGNER()
	return inst
}

// GetAdminAuthAccount gets the "adminAuth" account.
func (inst *UpdateAdminAuth) GetAdminAuthAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDistributorAccount sets the "distributor" account.
func (inst *UpdateAdminAuth) SetDistributorAccount(distributor ag_solanago.PublicKey) *UpdateAdminAuth {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(distributor).WRITE()
	return inst
}

// GetDistributorAccount gets the "distributor" account.
func (inst *UpdateAdminAuth) GetDistributorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateAdminAuth) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateAdminAuth,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateAdminAuth) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateAdminAuth) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.NewAdminAuth is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AdminAuth is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Distributor is not set")
		}
	}
	return nil
}

func (inst *UpdateAdminAuth) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateAdminAuth")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("newAdminAuth", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   adminAuth", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" distributor", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdateAdminAuth) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UpdateAdminAuth) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUpdateAdminAuthInstruction declares a new UpdateAdminAuth instruction with the provided parameters and accounts.
func NewUpdateAdminAuthInstruction(
	// Accounts:
	newAdminAuth ag_solanago.PublicKey,
	adminAuth ag_solanago.PublicKey,
	distributor ag_solanago.PublicKey) *UpdateAdminAuth {
	return NewUpdateAdminAuthInstructionBuilder().
		SetNewAdminAuthAccount(newAdminAuth).
		SetAdminAuthAccount(adminAuth).
		SetDistributorAccount(distributor)
}