package filterinterface

import (
	"fmt"
	"time"
)

func FilterSlice(data []*Order, opts ...Option) []*Order {
	for _, opt := range opts {
		data = opt.Apply(data)
	}
	return data
}

type Order struct {
	Id        int64
	OrderId   string
	Nid       string
	Roleid    string
	Ctime     string
	Productid string
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

type Page struct {
	PageNo   int64 //0代表首页
	PageSize int64 //0代表不分页
}

func (p *Page) Apply(data []*Order) []*Order {
	length := len(data)
	if length == 0 || p.PageSize == 0 {
		return data
	}
	offset := p.PageNo * p.PageSize
	limit := offset + p.PageSize
	if limit > int64(length) {
		limit = int64(length)
	}
	if offset > limit {
		return make([]*Order, 0)
	}
	return data[offset:limit]
}

func Do() {
	data := make([]*Order, 0)
	data = append(data,
		&Order{Id: 1, Nid: "z1", Roleid: "a", Ctime: "2023-01-01", OrderId: "a000"}, //x
		&Order{Id: 2, Nid: "z3", Roleid: "a", Ctime: "2023-01-01", OrderId: "a000"}, //x
		&Order{Id: 3, Nid: "z3", Roleid: "b", Ctime: "2023-04-01", OrderId: "a000"}, //x
		&Order{Id: 4, Nid: "z3", Roleid: "a", Ctime: "2023-04-01", OrderId: "a001"}, //
		&Order{Id: 5, Nid: "z3", Roleid: "a", Ctime: "2024-01-01", OrderId: "a002"}, //
		&Order{Id: 6, Nid: "z3", Roleid: "a", Ctime: "2025-01-01", OrderId: "a002"}, //x
		&Order{Id: 7, Nid: "z3", Roleid: "a", Ctime: "2024-02-01", OrderId: "a003"}, //
	)
	res := FilterSlice(data,
		&Where{Nid: "z3", RoleId: "a", Ctime: "2023-03-05", Etime: "2024-3-5", OrderIds: []string{"a001", "a002", "a003"}},
		&Page{PageNo: 0, PageSize: 4},
	)

	// for _, v := range data {
	// 	fmt.Printf("%#v \n", v)
	// }
	// fmt.Println()
	for _, v := range res {
		fmt.Printf("%#v \n", v)
	}
}
