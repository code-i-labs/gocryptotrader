package currency

import (
	"fmt"
	"strings"
)

// NewPairDelimiter splits the desired currency string at delimeter, the returns
// a Pair struct
func NewPairDelimiter(currencyPair, delimiter string) (Pair, error) {
	if !strings.Contains(currencyPair, delimiter) {
		return Pair{},
			fmt.Errorf("delimiter: [%s] not found in currencypair string", delimiter)
	}
	result := strings.Split(currencyPair, delimiter)
	if len(result) < 2 {
		return Pair{},
			fmt.Errorf("supplied pair: [%s] cannot be split with %s",
				currencyPair,
				delimiter)
	}
	if len(result) > 2 {
		result[1] = strings.Join(result[1:], delimiter)
	}
	return Pair{
		Delimiter: delimiter,
		Base:      NewCode(result[0]),
		Quote:     NewCode(result[1]),
	}, nil
}

// NewPairFromStrings returns a CurrencyPair without a delimiter
func NewPairFromStrings(base, quote string) (Pair, error) {
	if strings.Contains(base, " ") {
		return Pair{},
			fmt.Errorf("cannot create pair, invalid base currency string [%s]",
				base)
	}

	if strings.Contains(quote, " ") {
		return Pair{},
			fmt.Errorf("cannot create pair, invalid quote currency string [%s]",
				quote)
	}

	return Pair{Base: NewCode(base), Quote: NewCode(quote)}, nil
}

// NewPair returns a currency pair from currency codes
func NewPair(baseCurrency, quoteCurrency Code) Pair {
	return Pair{
		Base:  baseCurrency,
		Quote: quoteCurrency,
	}
}

// NewPairWithDelimiter returns a CurrencyPair with a delimiter
func NewPairWithDelimiter(base, quote, delimiter string) Pair {
	return Pair{
		Base:      NewCode(base),
		Quote:     NewCode(quote),
		Delimiter: delimiter,
	}
}

// NewPairFromIndex returns a CurrencyPair via a currency string and specific
// index
func NewPairFromIndex(currencyPair, index string) (Pair, error) {
	i := strings.Index(currencyPair, index)
	if i == -1 {
		return Pair{},
			fmt.Errorf("index %s not found in currency pair string", index)
	}
	if i == 0 {
		return NewPairFromStrings(currencyPair[0:len(index)],
			currencyPair[len(index):])
	}
	return NewPairFromStrings(currencyPair[0:i], currencyPair[i:])
}

// NewPairFromString converts currency string into a new CurrencyPair
// with or without delimeter
func NewPairFromString(currencyPair string) (Pair, error) {
	for x := range delimiters {
		if strings.Contains(currencyPair, delimiters[x]) {
			return NewPairDelimiter(currencyPair, delimiters[x])
		}
	}
	if len(currencyPair) < 3 {
		return Pair{},
			fmt.Errorf("cannot create pair from %s string",
				currencyPair)
	}
	return NewPairFromStrings(currencyPair[0:3], currencyPair[3:])
}

// NewPairFromFormattedPairs matches a supplied currency pair to a list of pairs
// with a specific format. This is helpful for exchanges which
// provide currency pairs with no delimiter so we can match it with a list and
// apply the same format
func NewPairFromFormattedPairs(currencyPair string, pairs Pairs, pairFmt PairFormat) (Pair, error) {
	for x := range pairs {
		fPair := pairs[x].Format(pairFmt.Delimiter, pairFmt.Uppercase)
		if strings.EqualFold(fPair.String(), currencyPair) {
			return pairs[x], nil
		}
	}
	return NewPairFromString(currencyPair)
}
