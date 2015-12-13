// Copyright (C) 2015 Max Riveiro <kavu13@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package curraunt

import (
	"github.com/openprovider/ecbrates"
	"log"
	"sync/atomic"
	"time"
)

var db = &ratesDB{}

type ratesDB struct {
	atomic.Value
}

type fields struct {
	ecb       *ecbrates.Rates
	updatedAt int64
}

func (db *ratesDB) updateECB() {
	ecb, err := ecbrates.New()
	if err != nil {
		log.Panicln(err)
	}

	db.Store(&fields{
		ecb:       ecb,
		updatedAt: time.Now().Unix(),
	})
}

func (db *ratesDB) getECB() *ecbrates.Rates {
	return db.Load().(*fields).ecb
}

func (db *ratesDB) getUpdatedAt() int64 {
	return db.Load().(*fields).updatedAt
}

func (db *ratesDB) startExpirationMonitor() {
	for {
		select {
		case <-time.After(1 * time.Hour):
			db.updateECB()
		}
	}
}

// InitDB will pre-populate Rates database and start Rates expiration monitor.
func InitDB() {
	db.updateECB()

	go db.startExpirationMonitor()
}
