package util

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"net"
)

// InterfaceNameToIp - Resolve the name of an interface to the corresponding IPv4 address
func InterfaceNameToIp(interfaceName string) (address string, err error) {
	var (
        iface *net.Interface
        addresses []net.Addr
        addr net.IP
    )
    if ief, err = net.InterfaceByName(interfaceName); err != nil { 
        return interfaceName, err
    }
    if addrs, err = ief.Addrs(); err != nil { 
        return interfaceName, err
    }
    for _, a := range addresses { 
        if addr = a.(*net.IPNet).IP.To4(); addr != nil {
            break
        }
    }
    if ipv4Addr == nil {
        return interfaceName, errors.New(fmt.Sprintf("Could not get IP for interface %s", interfaceName))
    }
    return ipv4Addr.String(), nil
}