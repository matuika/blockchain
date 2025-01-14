package simulation



import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
	"github.com/cosmos/cosmos-sdk/x/evidence/types"
)


const evidence = "evidence"


func GenEvidences(_ *rand.Rand, _ []simtypes.Account) []exported.Evidence {
	return []exported.Evidence{}
}


func RandomizedGenState(simState *module.SimulationState) {
	var ev []exported.Evidence

	simState.AppParams.GetOrGenerate(
		simState.Cdc, evidence, &ev, simState.Rand,
		func(r *rand.Rand) { ev = GenEvidences(r, simState.Accounts) },
	)

	evidenceGenesis := types.NewGenesisState(ev)

	bz, err := json.MarshalIndent(&evidenceGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(evidenceGenesis)
}
