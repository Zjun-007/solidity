package nft

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const erc721EnumerableABI = `[
{"name":"balanceOf","type":"function","inputs":[{"name":"owner","type":"address"}],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
{"name":"tokenOfOwnerByIndex","type":"function","inputs":[{"name":"owner","type":"address"},{"name":"index","type":"uint256"}],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},
{"name":"supportsInterface","type":"function","inputs":[{"name":"interfaceId","type":"bytes4"}],"outputs":[{"name":"","type":"bool"}],"stateMutability":"view"}
]`

var interfaceIDERC721Enumerable = [4]byte{0x78, 0x0e, 0x9d, 0x63}

type ContractTokens struct {
	Contract   string   `json:"contract"`
	Balance    string   `json:"balance"`
	Enumerable bool     `json:"enumerable"`
	Tokens     []string `json:"tokens,omitempty"`
	Note       string   `json:"note,omitempty"`
}

const maxEnumerate = 500

func ListTokensByContracts(ctx context.Context, cli *ethclient.Client, owner common.Address, contractHexes []string) ([]ContractTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(erc721EnumerableABI))
	if err != nil {
		return nil, err
	}
	out := make([]ContractTokens, 0, len(contractHexes))
	for _, hex := range contractHexes {
		hex = strings.TrimSpace(hex)
		if hex == "" {
			continue
		}
		if !strings.HasPrefix(hex, "0x") {
			hex = "0x" + hex
		}
		addr := common.HexToAddress(hex)
		ct := ContractTokens{Contract: strings.ToLower(addr.Hex())}

		supData, err := parsed.Pack("supportsInterface", interfaceIDERC721Enumerable)
		if err != nil {
			return nil, err
		}
		supRes, err := cli.CallContract(ctx, ethereum.CallMsg{To: &addr, Data: supData}, nil)
		if err == nil && len(supRes) > 0 {
			vals, err := parsed.Unpack("supportsInterface", supRes)
			if err == nil && len(vals) > 0 {
				if b, ok := vals[0].(bool); ok {
					ct.Enumerable = b
				}
			}
		}

		balData, err := parsed.Pack("balanceOf", owner)
		if err != nil {
			return nil, err
		}
		balRes, err := cli.CallContract(ctx, ethereum.CallMsg{To: &addr, Data: balData}, nil)
		if err != nil {
			ct.Note = fmt.Sprintf("balanceOf failed: %v", err)
			out = append(out, ct)
			continue
		}
		bVals, err := parsed.Unpack("balanceOf", balRes)
		if err != nil || len(bVals) == 0 {
			ct.Note = "invalid balance response"
			out = append(out, ct)
			continue
		}
		bal, ok := bVals[0].(*big.Int)
		if !ok || bal == nil {
			ct.Note = "invalid balance type"
			out = append(out, ct)
			continue
		}
		ct.Balance = bal.String()

		if !ct.Enumerable || bal.Sign() == 0 {
			if !ct.Enumerable {
				ct.Note = "not ERC721Enumerable; cannot list tokenIds via RPC only"
			}
			out = append(out, ct)
			continue
		}

		limit := new(big.Int).Set(bal)
		if limit.Cmp(big.NewInt(maxEnumerate)) > 0 {
			ct.Note = fmt.Sprintf("enumerating first %d of %s tokens", maxEnumerate, bal.String())
			limit = big.NewInt(maxEnumerate)
		}
		tokens := make([]string, 0, limit.Int64())
		for i := big.NewInt(0); i.Cmp(limit) < 0; i.Add(i, big.NewInt(1)) {
			tidData, err := parsed.Pack("tokenOfOwnerByIndex", owner, i)
			if err != nil {
				break
			}
			tidRes, err := cli.CallContract(ctx, ethereum.CallMsg{To: &addr, Data: tidData}, nil)
			if err != nil {
				if ct.Note == "" {
					ct.Note = fmt.Sprintf("tokenOfOwnerByIndex at %s: %v", i.String(), err)
				}
				break
			}
			tids, err := parsed.Unpack("tokenOfOwnerByIndex", tidRes)
			if err != nil || len(tids) == 0 {
				break
			}
			tid, ok := tids[0].(*big.Int)
			if !ok || tid == nil {
				break
			}
			tokens = append(tokens, tid.String())
		}
		ct.Tokens = tokens
		out = append(out, ct)
	}
	return out, nil
}
