// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package hash

import (
	"crypto/sha256"

	"github.com/medibloc/go-medibloc/crypto/signature/algorithm"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// Sha256 returns the SHA-256 digest of the data.
func Sha256(args ...[]byte) []byte {
	hasher := sha256.New()
	for _, bytes := range args {
		hasher.Write(bytes)
	}
	return hasher.Sum(nil)
}

// Sha3256 returns the SHA3-256 digest of the data.
func Sha3256(args ...[]byte) []byte {
	hasher := sha3.New256()
	for _, bytes := range args {
		hasher.Write(bytes)
	}
	return hasher.Sum(nil)
}

// Ripemd160 return the RIPEMD160 digest of the data.
func Ripemd160(args ...[]byte) []byte {
	hasher := ripemd160.New()
	for _, bytes := range args {
		hasher.Write(bytes)
	}
	return hasher.Sum(nil)
}

// CheckHashAlgorithm checks algorithm.
func CheckHashAlgorithm(alg algorithm.HashAlgorithm) error {
	switch alg {
	case algorithm.SHA3256:
		return nil
	default:
		return algorithm.ErrInvalidHashAlgorithm
	}
}
