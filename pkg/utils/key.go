// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"encoding/json"
	"fmt"
)

type KeyProvider interface {
	Key() (string, error)
}

func Key(keyProvider interface{}) (string, error) {
	if k, ok := keyProvider.(KeyProvider); ok {
		return k.Key()
	}
	data, err := json.Marshal(keyProvider)
	if err != nil {
		return "", fmt.Errorf("cannot marshal spec %w, consider implementing a Key() function", err)
	}
	return string(data), err
}
