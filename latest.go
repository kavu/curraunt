// Copyright (C) 2015 Max Riveiro <kavu13@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package curraunt

import (
	"errors"
	"github.com/openprovider/ecbrates"
	"net/url"
	"strconv"
)

func setBaseWithDefault(queryValues url.Values) ecbrates.Currency {
	if param, ok := queryValues["base"]; ok {
		return ecbrates.Currency(param[0])
	}

	return ecbrates.EUR
}

func setAmountWithDefault(queryValues url.Values) (float64, error) {
	if param, ok := queryValues["amount"]; ok {
		return strconv.ParseFloat(param[0], 64)
	}

	return 1, nil
}

func formatLatest(base ecbrates.Currency, amount float64, db *ecbrates.Rates) (data *response, err error) {
	data = &response{
		Base:   base,
		Amount: amount,
		Date:   db.Date,
		Rates:  make(map[ecbrates.Currency]float64),
	}

	if !base.IsValid() {
		return nil, errors.New("base code is invalid")
	}

	for k, v := range db.Rate {
		if base != ecbrates.EUR {
			data.Rates[k], err = db.Convert(data.Amount, base, k)
		} else {
			data.Rates[k], err = strconv.ParseFloat(v.(string), 64)
		}

		if err != nil {
			return nil, err
		}
	}

	return data, err
}
