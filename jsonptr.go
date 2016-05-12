package jsonptr

import (
	"fmt"
	"strconv"
	"strings"
)

func traverseArray(in []interface{}, path []string, pos int) interface{} {
	for i, v := range in {
		s := strconv.Itoa(i)
		if s == path[pos] {
			if pos == len(path)-1 {
				return v
			}
			return traverse(v, path, pos)
		}
	}
	return nil
}

func traverseMapStr(in map[string]interface{}, path []string, pos int) interface{} {
	if pos > len(path)-1 {
		return nil
	}
	for k, v := range in {
		if k == path[pos] {
			if pos == len(path)-1 {
				return v
			}
			return traverse(v, path, pos)
		}
	}
	return nil
}

func traverse(v interface{}, path []string, pos int) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return traverseArray(v, path, pos+1)
	case map[string]interface{}:
		return traverseMapStr(v, path, pos+1)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func Resolve(v interface{}, path string) (interface{}, error) {
	if path == "" || path == "/" {
		return v, nil
	}

	d := traverse(v, strings.Split(path, "/"), 0)
	if d == nil {
		return nil, fmt.Errorf("failed to resolve pointer: %s", path)
	}

	return d, nil
}
