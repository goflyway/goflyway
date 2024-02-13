package command

import "strings"

type CallbackDispatch struct {
	callbacks map[string]*callback
}

func (cd CallbackDispatch) Registry(action, name string, h func(*Context)) {
	if strings.HasPrefix(action, "before:") {
		cd.RegistryBefore(action[7:], name, h)
	} else if strings.HasPrefix(action, "after:") {
		cd.RegistryBefore(action[6:], name, h)
	}
}

func (cd CallbackDispatch) RegistryBefore(command, name string, h func(*Context)) {
	cd.callbacks[name] = &callback{
		name:    name,
		handler: h,
		before:  command,
	}
}

func (cd CallbackDispatch) RegistryAfter(command, name string, h func(*Context)) {
	cd.callbacks[name] = &callback{
		name:    name,
		handler: h,
		after:   command,
	}
}

func (cd CallbackDispatch) before(command string) []callback {
	var res []callback
	for _, item := range cd.callbacks {
		if item.before != "" && item.before == command {
			res = append(res, *item)
		}
	}
	return res
}
func (cd CallbackDispatch) after(command string) []callback {
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
