package filterinterface

import (
	"fmt"
	"time"
)

type Order struct {
	Id        int64
	OrderId   string
	Nid       string
	Roleid    string
	Ctime     string
	Productid string
}

type Option interface {
	Apply([]*Order) []*Order
}

type Where struct {
	Nid      string
	RoleId   string
	Ctime    string
	Etime    string
	OrderIds []string
}

func ParseTimeString(input string) (time.Time, error) {
	formats := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05-07:00",
		"2006-1-2",
		"2006-1-2 15:04:05",
		"2006-1-2T15:04:05-07:00",
	}

	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, input)
		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Printf("ParseTimeString(input=[%s]) err:%v", input, err)
		return parsedTime, err
	}

	return parsedTime, nil
}

func (w *Where) Apply(data []*Order) []*Order {
	ret := make([]*Order, 0)
	if len(data) == 0 {
		return ret
	}
	for _, v := range data {
		if v == nil {
			continue
		}
		if w.Nid != "" && v.Nid != w.Nid {
			continue
		}
		if w.RoleId != "" && v.Roleid != w.RoleId {
			continue
		}
		if w.Ctime != "" {
			paramCtm, err1 := ParseTimeString(w.Ctime)
			if err1 != nil {
				continue
			}
			dataCtm, err2 := ParseTimeString(v.Ctime)
			if err2 != nil {
				continue
			}
			if dataCtm.Before(paramCtm) {
				continue
			}
		}
		if w.Etime != "" {
			paramCtm, err1 := ParseTimeString(w.Etime)
			if err1 != nil {
				continue
			}
			dataCtm, err2 := ParseTimeString(v.Ctime)
			if err2 != nil {
				continue
			}
			if dataCtm.After(paramCtm) {
				continue
			}
		}
		if len(w.OrderIds) > 0 {
			find := false
			for _, order := range w.OrderIds {
				if v.OrderId == order {
					find = true
					break
				}
			}
			if !find {
				continue
			}
		}
		ret = append(ret, v)
	}
	return ret
}

func FilterSlice(data []*Order, opts ...Option) []*Order {
	ret := make([]*Order, 0)
	for _, opt := range opts {
		ret = opt.Apply(data)
	}
	return ret
}
