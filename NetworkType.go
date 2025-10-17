package gxnet

// --------------------------------------------------------------------------
//
//	Gurux Ltd
//
// Filename:        $HeadURL$
//
// Version:         $Revision$,
//
//	$Date$
//	$Author$
//
// # Copyright (c) Gurux Ltd
//
// ---------------------------------------------------------------------------
//
//	DESCRIPTION
//
// This file is a part of Gurux Device Framework.
//
// Gurux Device Framework is Open Source software; you can redistribute it
// and/or modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; version 2 of the License.
// Gurux Device Framework is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// More information of Gurux products: https://www.gurux.org
//
// This code is licensed under the GNU General Public License v2.
// Full text may be retrieved at http://www.gnu.org/licenses/gpl-2.0.txt
// ---------------------------------------------------------------------------

import (
	"fmt"
	"strings"

	"github.com/Gurux/go-gxcommon"
)

// NetworkType determines which network protocol is used for data transfer.
type NetworkType int

const (
	// NetworkTypeUDP defines that UDP protocol is used.
	NetworkTypeUDP NetworkType = iota
	// NetworkTypeTCP defines that the TCP/IP protocol is used.
	NetworkTypeTCP
)

// Parse converts the given string into a NetworkType value.
//
// It returns the corresponding NetworkType constant if the string matches
// a known level name, or an error if the input is invalid.
func (g NetworkType) Parse(value string) (NetworkType, error) {
	var ret NetworkType
	var err error
	switch strings.ToUpper(value) {
	case "UDP":
		ret = NetworkTypeUDP
	case "TCP":
		ret = NetworkTypeTCP
	default:
		err = fmt.Errorf("%w: %q", gxcommon.ErrUnknownEnum, value)
	}
	return ret, err
}

// String returns the canonical name of the network type.
// It satisfies fmt.Stringer.
func (g NetworkType) String() string {
	var ret string
	switch g {
	case NetworkTypeUDP:
		ret = "UDP"
	case NetworkTypeTCP:
		ret = "TCP"
	}
	return ret
}
