package main

import (
	"encoding/binary"
)

func main() {}

// LendoMensagem ...
func LendoMensagem(
	unitAddr byte, regAddr uint16, length uint16, TransactionID uint16,
) (b []byte) {

	/*
		o endereço inicial dos dados requeridos,
		bem como o número de valores pode
		ser representado em dois bytes cada
	*/
	addrSlice := make([]byte, 2)
	lenSlice := make([]byte, 2)
	transSlice := make([]byte, 2)

	binary.BigEndian.PutUint16(addrSlice, regAddr)
	binary.BigEndian.PutUint16(lenSlice, length)
	binary.BigEndian.PutUint16(transSlice, TransactionID)

	/*
		criando uma matriz de bytes com o formato necessario que
		inclui todas as informações
	*/

	ReadMsg := []byte{
		transSlice[0], transSlice[1],
		0x00, 0x00, 0x00, 0x06, unitAddr, 0x03,
		addrSlice[0], addrSlice[1],
		lenSlice[0], lenSlice[1],
	}

	return ReadMsg

}
