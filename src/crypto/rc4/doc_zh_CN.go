// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package rc4 implements RC4 encryption, as defined in Bruce Schneier's Applied
// Cryptography.
package rc4

// A Cipher is an instance of RC4 using a particular key.
type Cipher struct {
	// contains filtered or unexported fields
}

// NewCipher creates and returns a new Cipher. The key argument should be the RC4
// key, at least 1 byte and at most 256 bytes.
func NewCipher(key []byte) (*Cipher, error)

// Reset zeros the key data so that it will no longer appear in the process's
// memory.
func (c *Cipher) Reset()

// XORKeyStream sets dst to the result of XORing src with the key stream. Dst and
// src may be the same slice but otherwise should not overlap.
func (c *Cipher) XORKeyStream(dst, src []byte)

type KeySizeError int

func (k KeySizeError) Error() string
