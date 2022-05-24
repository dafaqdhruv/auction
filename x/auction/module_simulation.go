package auction

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/deepstack/auction/testutil/sample"
	auctionsimulation "github.com/deepstack/auction/x/auction/simulation"
	"github.com/deepstack/auction/x/auction/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = auctionsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateObject = "op_weight_msg_object"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateObject int = 100

	opWeightMsgUpdateObject = "op_weight_msg_object"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateObject int = 100

	opWeightMsgDeleteObject = "op_weight_msg_object"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteObject int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	auctionGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ObjectList: []types.Object{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ObjectCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&auctionGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateObject, &weightMsgCreateObject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateObject = defaultWeightMsgCreateObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateObject,
		auctionsimulation.SimulateMsgCreateObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateObject, &weightMsgUpdateObject, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateObject = defaultWeightMsgUpdateObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateObject,
		auctionsimulation.SimulateMsgUpdateObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteObject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteObject, &weightMsgDeleteObject, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteObject = defaultWeightMsgDeleteObject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteObject,
		auctionsimulation.SimulateMsgDeleteObject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
