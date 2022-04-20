package wordplay

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
)

// SimplifyPath 化简Linux绝对路径
func SimplifyPath(path string) string {
	stack := arraystack.New()

	for i := 1; i < len(path); i++ {
		if path[i-1] == '/' || path[i] == '.' {
			if path[i] == '/' && (path[i-1] == '/' || path[i-1] == '.') {
				continue
			}
			if path[i-1] == '.' && path[i] == '.' {
				if stack.Empty() {
					continue
				}
				if v, ok := stack.Pop(); ok {
					fmt.Println("pop elem :%v", v)
				}
			}
			if path[i] <= 'z' && path[i] >= 'a' {
				tmpStr := []byte{}
				for i < len(path) && path[i] != '/' && path[i] != '.' {
					tmpStr = append(tmpStr, []byte(path)[i])
					i++
				}
				stack.Push(tmpStr)
			}
		}
	}
	res := ""
	for !stack.Empty() {
		res = "/" + res
		if str, ok := stack.Pop(); ok {
			res = string(str.([]byte)) + res
		}
	}
	if res == "" {
		return "/"
	}
	return "/" + res[:len(res)-1]
}

func SimplifyPathV2(path string) string {
	stack := []string{}
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}
