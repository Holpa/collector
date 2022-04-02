package models

import (
	"fmt"
	"math/big"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	BigInt struct {
		v *big.Int
	}
)

func NewBigInt(value *big.Int) *BigInt {
	return &BigInt{
		v: value,
	}
}
func (bigInt *BigInt) Int() *big.Int {
	return bigInt.v
}
func (bigInt *BigInt) MarshalBSON() ([]byte, error) {
	txt, err := bigInt.v.MarshalText()
	if err != nil {
		return nil, err
	}

	return bson.Marshal(map[string]string{"v": string(txt)})
}
func (bigInt *BigInt) UnmarshalBSON(data []byte) error {
	var d bson.D
	err := bson.Unmarshal(data, &d)
	if err != nil {
		return err
	}

	if value, ok := d.Map()["v"]; ok {
		bigInt.v = big.NewInt(0)
		return bigInt.v.UnmarshalText([]byte(value.(string)))
	}

	return fmt.Errorf("key 'v' missing")
}
func (bigInt *BigInt) MarshalJSON() ([]byte, error) {
	return bigInt.v.MarshalText()
}
