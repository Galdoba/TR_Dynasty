package Trade

import (
	"fmt"
	"strconv"
)

type Cargo struct {
	cargo map[string]int
}

type Cargoer interface {
	Add(TradeGoodR, int)
	Remove(TradeGoodR)
	Info(string) string
}

func NewCargo() *Cargo {
	c := &Cargo{}
	c.cargo = make(map[string]int)
	return c
}

func (c *Cargo) Add(tgr *TradeGoodR, addVolume int) {
	key := tgr.code
	c.cargo[key] = c.cargo[key] + addVolume
	if c.cargo[key] < 0 {
		delete(c.cargo, key)
	}
}

func (c *Cargo) Remove(tgr *TradeGoodR) {
	key := tgr.code
	if _, ok := c.cargo[key]; ok {
		delete(c.cargo, key)
	}
}

func (c *Cargo) Info(code string) (data string) {
	for k, v := range c.cargo {
		if k == code {
			tgr := NewTradeGoodR(code)
			data = tgr.description + "	(" + code + ")	" + strconv.Itoa(v) + " tons\n"
		}
	}
	if code == "All" {
		allCodes := allTradeGoodsRCodes()
		for _, aCode := range allCodes {
			if v, ok := c.cargo[aCode]; ok {
				tgr := NewTradeGoodR(aCode)
				data = data + fmt.Sprintf("%v	 %v 		%v tons\n", aCode, tgr.description, strconv.Itoa(v))
			}
		}
	}
	return data
}
