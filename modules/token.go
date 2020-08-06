package modules

import (
	"fmt"
	"strings"

	sdk "github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/utils/cache"
	"github.com/irisnet/service-sdk-go/utils/log"
)

type tokenQuery struct {
	q sdk.Queries
	*log.Logger
	cache.Cache
}

func (l tokenQuery) QueryToken(denom string) (sdk.Token, error) {
	denom = strings.ToLower(denom)
	if t, err := l.Get(l.prefixKey(denom)); err == nil {
		return t.(sdk.Token), nil
	}

	param := struct {
		Denom string
	}{
		Denom: denom,
	}

	var token sdk.Token
	if err := l.q.QueryWithResponse("custom/token/token", param, &token); err != nil {
		return sdk.Token{}, err
	}

	l.SaveTokens(token)
	return token, nil
}

func (l tokenQuery) SaveTokens(tokens ...sdk.Token) {
	for _, t := range tokens {
		err1 := l.Set(l.prefixKey(t.Symbol), t)
		err2 := l.Set(l.prefixKey(t.MinUnit), t)
		if err1 != nil || err2 != nil {
			l.Warn().
				Str("symbol", t.Symbol).
				Msg("cache token failed")
		}
	}
}

func (l tokenQuery) ToMinCoin(coins ...sdk.DecCoin) (dstCoins sdk.Coins, err sdk.Error) {
	for _, coin := range coins {
		token, err := l.QueryToken(coin.Denom)
		if err != nil {
			return nil, sdk.Wrap(err)
		}

		minCoin, err := token.GetCoinType().ConvertToMinCoin(coin)
		if err != nil {
			return nil, sdk.Wrap(err)
		}
		dstCoins = append(dstCoins, minCoin)
	}
	return dstCoins.Sort(), nil
}

func (l tokenQuery) ToMainCoin(coins ...sdk.Coin) (dstCoins sdk.DecCoins, err sdk.Error) {
	for _, coin := range coins {
		token, err := l.QueryToken(coin.Denom)
		if err != nil {
			return dstCoins, sdk.Wrap(err)
		}

		mainCoin, err := token.GetCoinType().ConvertToMainCoin(coin)
		if err != nil {
			return dstCoins, sdk.Wrap(err)
		}
		dstCoins = append(dstCoins, mainCoin)
	}
	return dstCoins.Sort(), nil
}

func (l tokenQuery) prefixKey(symbol string) string {
	return fmt.Sprintf("token:%s", symbol)
}
