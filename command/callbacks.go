package command

import "strings"

type callbackDispatch struct {
	callbacks map[string]*callback
}

func (cd callbackDispatch) Registry(action, name string, h func(*Context)) {
	c := &callback{
		name:    name,
		handler: h,
	}
	if strings.HasPrefix(action, "before:") {
		c.before = action[7:]
		cd.callbacks[name] = c
	} else if strings.HasPrefix(action, "after:") {
		c.after = action[6:]
		cd.callbacks[name] = c
	}
}

func (cd callbackDispatch) before(command string) []callback {
	var res []callback
	for _, item := range cd.callbacks {
		if item.before != "" && item.before == command {
			res = append(res, *item)
		}
	}
	return res
}
func (cd callbackDispatch) after(command string) []callback {
	var res []callback
	for _, item := range cd.callbacks {
		if item.after != "" && item.after == command {
			res = append(res, *item)
		}
	}
	return res
}

type callback struct {
	name    string
	before  string
	after   string
	handler func(*Context)
}
