package constants

import "strings"

type (
	TransferDirection int
	TransferMethod    int
	TransferMethodId  string
)

const (
	TransferDirectionFromUser TransferDirection = iota
	TransferDirectionToUser
)

const (
	TransferMethodAny TransferMethod = iota
	TransferMethodClaim
	TransferMethodLevelUp
	TransferMethodMultiLevelUp
	TransferMethodVeFlyVote
	TransferMethodFlyStakeDeposit
	TransferMethodFlyStakeWithdraw
	TransferMethodBreedingEnter
	TransferMethodChangeName
	TransferMethodSellFly
	TransferMethodProvideLiquidity
	TransferMethodRemoveLiquidity
)

const (
	METHOD_ID_ANY                TransferMethodId = ""
	METHOD_ID_CLAIM              TransferMethodId = "0x4e71d92d"
	METHOD_ID_LEVEl_UP           TransferMethodId = "0x0c679fa0"
	METHOD_ID_MULTI_LEVEl_UP     TransferMethodId = "0x6a369d6c"
	METHOD_ID_VE_FLY_VOTE        TransferMethodId = "0xc9d27afe"
	METHOD_ID_FLY_STAKE_DEPOSIT  TransferMethodId = "0xb6b55f25"
	METHOD_ID_FLY_STAKE_WITHDRAW TransferMethodId = "0x2e1a7d4d"
	METHOD_ID_BREEDING_ENTER     TransferMethodId = "0xa59f3e0c"
	METHOD_ID_CHANGE_NAME        TransferMethodId = "0x8eb3be28"
	METHOD_ID_SELL_FLY           TransferMethodId = "0x676528d1"
	METHOD_ID_PROVIDE_LIQUIDITY  TransferMethodId = "0xf91b3f72"
	METHOD_ID_REMOVE_LIQUIDITY   TransferMethodId = "0x2c407024"
)

func TransferDirectionFromString(value string) TransferDirection {
	lowerCased := strings.ToLower(value)

	switch lowerCased {
	case "in", "to":
		return TransferDirectionToUser
	case "out", "from":
		return TransferDirectionFromUser
	default:
		return TransferDirectionToUser
	}
}
func (transferDirection TransferDirection) String() string {
	switch transferDirection {
	case TransferDirectionToUser:
		return "in"
	case TransferDirectionFromUser:
		return "out"
	default:
		return "any"
	}
}

func TransferMethodFromString(transferMethod string) TransferMethod {
	lowerCased := strings.ToLower(transferMethod)

	switch lowerCased {
	case "claim":
		return TransferMethodClaim
	case "level-up":
		return TransferMethodLevelUp
	case "multi-level-up":
		return TransferMethodMultiLevelUp
	case "vefly-vote":
		return TransferMethodVeFlyVote
	case "stake-deposit":
		return TransferMethodFlyStakeDeposit
	case "stake-withdraw":
		return TransferMethodFlyStakeWithdraw
	case "breeding":
		return TransferMethodBreedingEnter
	case "change-name":
		return TransferMethodChangeName
	case "sell-fly":
		return TransferMethodSellFly
	case "provide-liquidity":
		return TransferMethodProvideLiquidity
	case "remove-liquidity":
		return TransferMethodRemoveLiquidity
	default:
		return TransferMethodAny
	}
}
func TransferMethodFromMethodId(methodId string) TransferMethod {
	lowerCased := strings.ToLower(methodId)

	switch lowerCased {
	case string(METHOD_ID_CLAIM):
		return TransferMethodClaim
	case string(METHOD_ID_LEVEl_UP):
		return TransferMethodLevelUp
	case string(METHOD_ID_MULTI_LEVEl_UP):
		return TransferMethodMultiLevelUp
	case string(METHOD_ID_VE_FLY_VOTE):
		return TransferMethodVeFlyVote
	case string(METHOD_ID_FLY_STAKE_DEPOSIT):
		return TransferMethodFlyStakeDeposit
	case string(METHOD_ID_FLY_STAKE_WITHDRAW):
		return TransferMethodFlyStakeWithdraw
	case string(METHOD_ID_BREEDING_ENTER):
		return TransferMethodBreedingEnter
	case string(METHOD_ID_CHANGE_NAME):
		return TransferMethodChangeName
	case string(METHOD_ID_SELL_FLY):
		return TransferMethodSellFly
	case string(METHOD_ID_PROVIDE_LIQUIDITY):
		return TransferMethodProvideLiquidity
	case string(METHOD_ID_REMOVE_LIQUIDITY):
		return TransferMethodRemoveLiquidity
	default:
		return TransferMethodAny
	}
}
func (transferMethod TransferMethod) String() string {
	switch transferMethod {
	case TransferMethodClaim:
		return "claim"
	case TransferMethodLevelUp:
		return "level-up"
	case TransferMethodMultiLevelUp:
		return "multi-level-up"
	case TransferMethodVeFlyVote:
		return "vefly-vote"
	case TransferMethodFlyStakeDeposit:
		return "stake-deposit"
	case TransferMethodFlyStakeWithdraw:
		return "stake-withdraw"
	case TransferMethodBreedingEnter:
		return "breeding"
	case TransferMethodChangeName:
		return "change-name"
	case TransferMethodSellFly:
		return "sell-fly"
	case TransferMethodProvideLiquidity:
		return "provide-liquidity"
	case TransferMethodRemoveLiquidity:
		return "remove-liquidity"
	default:
		return "any"
	}
}
func (transferMethod TransferMethod) ToMethodId() TransferMethodId {
	switch transferMethod {
	case TransferMethodClaim:
		return METHOD_ID_CLAIM
	case TransferMethodLevelUp:
		return METHOD_ID_LEVEl_UP
	case TransferMethodMultiLevelUp:
		return METHOD_ID_MULTI_LEVEl_UP
	case TransferMethodVeFlyVote:
		return METHOD_ID_VE_FLY_VOTE
	case TransferMethodFlyStakeDeposit:
		return METHOD_ID_FLY_STAKE_DEPOSIT
	case TransferMethodFlyStakeWithdraw:
		return METHOD_ID_FLY_STAKE_WITHDRAW
	case TransferMethodBreedingEnter:
		return METHOD_ID_BREEDING_ENTER
	case TransferMethodChangeName:
		return METHOD_ID_CHANGE_NAME
	case TransferMethodSellFly:
		return METHOD_ID_SELL_FLY
	case TransferMethodProvideLiquidity:
		return METHOD_ID_PROVIDE_LIQUIDITY
	case TransferMethodRemoveLiquidity:
		return METHOD_ID_REMOVE_LIQUIDITY
	default:
		return METHOD_ID_ANY
	}
}
