package simulation



import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/capability/types"
)


const index = "index"


func GenIndex(r *rand.Rand) uint64 {
	return uint64(r.Int63n(1000)) + 1
}


func RandomizedGenState(simState *module.SimulationState) {
	var idx uint64

	simState.AppParams.GetOrGenerate(
		simState.Cdc, index, &idx, simState.Rand,
		func(r *rand.Rand) { idx = GenIndex(r) },
	)

	capabilityGenesis := types.GenesisState{Index: idx}

	bz, err := json.MarshalIndent(&capabilityGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&capabilityGenesis)
}
