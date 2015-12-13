// Copyright (C) 2015 Max Riveiro <kavu13@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package curraunt

import (
	"github.com/openprovider/ecbrates"
	"github.com/pquerna/ffjson/ffjson"
	"net/http"
	"strconv"
)

//go:generate ffjson -nodecoder -force-regenerate $GOFILE
type response struct {
	Base   ecbrates.Currency             `json:"base"`
	Amount float64                       `json:"amount"`
	Date   string                        `json:"date"`
	Rates  map[ecbrates.Currency]float64 `json:"rates"`
}

// LatestHandler is a http.HandlerFunc that renders JSON with latest parametrized Rates data.
func LatestHandler(w http.ResponseWriter, r *http.Request) {
	var (
		queryValues = r.URL.Query()
		base        = setBaseWithDefault(queryValues)
	)

	amount, err := setAmountWithDefault(queryValues)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data, err := formatLatest(base, amount, db.getECB())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("X-Last-Updated-At", strconv.Itoa(int(db.getUpdatedAt())))

	if err := ffjson.NewEncoder(w).EncodeFast(data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
