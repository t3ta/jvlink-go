package main

import (
	"fmt"
	"strings"
)

type DataRecordType struct {
	dataType   DataType
	recordType RecordType
}

type DataType string

const (
	TOKU DataType = "TOKU"
	RACE DataType = "RACE"
	DIFF DataType = "DIFF"
	BLOD DataType = "BLOD"
	SNAP DataType = "SNAP"
	SLOP DataType = "SLOP"
	WOOD DataType = "WOOD"
	YSCH DataType = "YSCH"
	HOSE DataType = "HOSE"
	HOYU DataType = "HOYU"
	COMM DataType = "COMM"
	MING DataType = "MING"
)

type RecordType string

const (
	TK RecordType = "TK" // No.1
	RA RecordType = "RA" // No.2
	SE RecordType = "SE" // No.3
	HR RecordType = "HR" // No.4
	H1 RecordType = "H1" // No.5
	H6 RecordType = "H6" // No.6
	O1 RecordType = "O1" // No.7
	O2 RecordType = "O2" // No.8
	O3 RecordType = "O3" // No.9
	O4 RecordType = "O4" // No.10
	O5 RecordType = "O5" // No.11
	O6 RecordType = "O6" // No.12
	UM RecordType = "UM" // No.13
	KS RecordType = "KS" // No.14
	CH RecordType = "CH" // No.15
	BR RecordType = "BR" // No.16
	BN RecordType = "BN" // No.17
	HN RecordType = "HN" // No.18
	SK RecordType = "SK" // No.19
	CK RecordType = "CK" // No.20
	RC RecordType = "RC" // No.21
	HC RecordType = "HC" // No.22
	HS RecordType = "HS" // No.23
	HY RecordType = "HY" // No.24
	YS RecordType = "YS" // No.25
	BT RecordType = "BT" // No.26
	CS RecordType = "CS" // No.27
	DM RecordType = "DM" // No.28
	TM RecordType = "TM" // No.29
	WF RecordType = "WF" // No.30
	JG RecordType = "JG" // No.31
	WC RecordType = "WC" // No.32
)

// type attr struct {
// 	dataTypes []DataType
// 	objType Type
// }

// var num = 32 // number of RecordType
// var attr = [num]attr{
// }

var ErrorInvalidDataType = fmt.Errorf("Invalid DataType")
var ErrorInvalidRecordType = fmt.Errorf("Invalid RecordType")

func NewDataRecordType(typeString string) (*DataRecordType, error) {
	str := strings.Split(typeString, ":")
	dataType := DataType(str[0])
	recordType := RecordType(str[1])

	err := dataType.Valid()
	err = fmt.Errorf(recordType.Valid().Error(), err)

	dataRecordType := &DataRecordType{
		dataType:   dataType,
		recordType: recordType,
	}

	return dataRecordType, err
}

func (dr DataRecordType) Valid() error

func (d DataType) Valid() error {
	switch d {
	case TOKU, RACE, DIFF, BLOD, SNAP, SLOP, WOOD, YSCH, HOSE, HOYU, COMM, MING:
		return nil
	default:
		return ErrorInvalidDataType
	}
}

func (r RecordType) Valid() error {
	switch r {
	case TK, RA, SE, HR, H1, H6, O1, O2, O3, O4, O5, O6, UM, KS, CH, BR, BN, HN, SK, CK, RC, HC, HS, HY, YS, BT, CS, DM, TM, WF, JG, WC:
		return nil
	default:
		return ErrorInvalidRecordType
	}
}
