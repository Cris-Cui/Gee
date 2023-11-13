package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild { // 匹配n的每个子节点, 返回part相同且为非精准匹配的
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert defines the method to insert new url
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { //当前所在节点的高度等于parts, 即为子节点, 赋值pattern
		n.pattern = pattern
		return
	}
	part := parts[height]       // 取出当前的part进行匹配
	child := n.matchChild(part) // 匹配子节点
	if child == nil {           // 匹配失败, 需要新建子节点
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1) // 递归插入
}

// search defines the method to search if url already exist
func (n *node) search(parts []string, height int) *node {
	// 匹配结束判断, 匹配到了*，匹配失败，或者匹配到了第len(parts)层节点。
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" { // 为字典树上的一个中间节点
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
