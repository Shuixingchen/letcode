package helper

import (
	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
)

const (
	PolygonRPC = "https://polygon-mumbai.g.alchemy.com/v2/_4xDtlTKWmynPDVaX1JfRvysRif0wZ85"

	MumbaiChainID = 80001
	WMaticAddr    = "0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889"
	STAddr        = "0x4047DeF04Ae123Bf223808119B9685104A33D5f7"
	GTAddr        = "0x8dc67E514f29D099c06BBF23a227031c0b3808ec"
)

const (
	ContractV3Factory            = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	ContractV3SwapRouterV1       = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	ContractV3SwapRouterV2       = "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"
	ContractV3NFTPositionManager = "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"
	ContractV3Quoter             = "0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6"
)

var (
	WMATIC = coreEntities.NewToken(MumbaiChainID, common.HexToAddress(WMaticAddr), 18, "Matic", "Matic Network(PolyGon)")
	ST     = coreEntities.NewToken(MumbaiChainID, common.HexToAddress(STAddr), 18, "ST", "SportToken (ST)")
	GT     = coreEntities.NewToken(MumbaiChainID, common.HexToAddress(GTAddr), 18, "GT", "GoveranceToken (GT)")
)
