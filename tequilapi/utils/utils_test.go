/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteAsJSONReturnsExpectedResponse(t *testing.T) {

	respRecorder := httptest.NewRecorder()

	type TestStruct struct {
		IntField    int
		StringField string `json:"renamed"`
	}

	WriteAsJSON(TestStruct{1, "abc"}, respRecorder)

	result := respRecorder.Result()

	assert.Equal(t, "application/json; charset=utf-8", result.Header.Get("Content-type"))
	assert.JSONEq(
		t,
		`{
			"IntField" : 1,
			"renamed" : "abc"
		}`,
		respRecorder.Body.String())
}
