package modbus

type Filter struct {
	Name       string `json:"name"`       //字段
	Expression string `json:"expression"` //表达式
	//Entire     bool   `json:"entire"`
}

type Calculator struct {
	Name       string `json:"name"`       //赋值
	Expression string `json:"expression"` //表达式
}

type Poller struct {
	Code    uint8  `json:"code"`
	Address uint16 `json:"address"`
	Length  uint16 `json:"length"` //长度
}

func (p *Poller) Parse(mappers []*Mapper, buf []byte, values map[string]any) error {
	for _, m := range mappers {
		if p.Code == m.Code &&
			p.Address <= m.Address &&
			p.Length > m.Address-p.Address {
			ret, err := m.Parse(p.Address, buf)
			if err != nil {
				//log.Error(err)
				return err
			}
			//03指令 的 位类型
			if rets, ok := ret.(map[string]bool); ok {
				for k, v := range rets {
					values[k] = v
				}
			} else {
				values[m.Name] = ret
			}
		}
	}
	return nil
}
