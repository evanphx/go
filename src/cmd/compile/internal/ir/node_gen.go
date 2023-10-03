// Code generated by mknode.go. DO NOT EDIT.

package ir

import "fmt"

func (n *AddStringExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AddStringExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	return &c
}
func (n *AddStringExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *AddStringExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}
func (n *AddStringExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *AddrExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AddrExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *AddrExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *AddrExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}
func (n *AddrExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *AssignListStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AssignListStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Lhs = copyNodes(c.Lhs)
	c.Rhs = copyNodes(c.Rhs)
	return &c
}
func (n *AssignListStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.Lhs, do) {
		return true
	}
	if doNodes(n.Rhs, do) {
		return true
	}
	return false
}
func (n *AssignListStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Lhs, edit)
	editNodes(n.Rhs, edit)
}
func (n *AssignListStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Lhs, edit)
	editNodes(n.Rhs, edit)
}

func (n *AssignOpStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AssignOpStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *AssignOpStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *AssignOpStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}
func (n *AssignOpStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}

func (n *AssignStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *AssignStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *AssignStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *AssignStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}
func (n *AssignStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}

func (n *BasicLit) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BasicLit) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *BasicLit) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *BasicLit) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *BasicLit) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *BinaryExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BinaryExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *BinaryExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *BinaryExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}
func (n *BinaryExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
}

func (n *BlockStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BlockStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	return &c
}
func (n *BlockStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	return false
}
func (n *BlockStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
}
func (n *BlockStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
}

func (n *BranchStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *BranchStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *BranchStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *BranchStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *BranchStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *CallExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CallExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Args = copyNodes(c.Args)
	c.KeepAlive = copyNames(c.KeepAlive)
	return &c
}
func (n *CallExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Fun != nil && do(n.Fun) {
		return true
	}
	if doNodes(n.Args, do) {
		return true
	}
	if n.DeferAt != nil && do(n.DeferAt) {
		return true
	}
	if doNames(n.KeepAlive, do) {
		return true
	}
	return false
}
func (n *CallExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Fun != nil {
		n.Fun = edit(n.Fun).(Node)
	}
	editNodes(n.Args, edit)
	if n.DeferAt != nil {
		n.DeferAt = edit(n.DeferAt).(Node)
	}
	editNames(n.KeepAlive, edit)
}
func (n *CallExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Fun != nil {
		n.Fun = edit(n.Fun).(Node)
	}
	editNodes(n.Args, edit)
	if n.DeferAt != nil {
		n.DeferAt = edit(n.DeferAt).(Node)
	}
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	editNames(n.KeepAlive, edit)
}

func (n *CaseClause) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CaseClause) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	c.RTypes = copyNodes(c.RTypes)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *CaseClause) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Var != nil && do(n.Var) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	if doNodes(n.RTypes, do) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	return false
}
func (n *CaseClause) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Var != nil {
		n.Var = edit(n.Var).(*Name)
	}
	editNodes(n.List, edit)
	editNodes(n.RTypes, edit)
	editNodes(n.Body, edit)
}
func (n *CaseClause) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Var != nil {
		n.Var = edit(n.Var).(*Name)
	}
	editNodes(n.List, edit)
	editNodes(n.RTypes, edit)
	editNodes(n.Body, edit)
}

func (n *ClosureExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ClosureExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ClosureExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *ClosureExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}
func (n *ClosureExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *CommClause) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CommClause) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *CommClause) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Comm != nil && do(n.Comm) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	return false
}
func (n *CommClause) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Comm != nil {
		n.Comm = edit(n.Comm).(Node)
	}
	editNodes(n.Body, edit)
}
func (n *CommClause) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Comm != nil {
		n.Comm = edit(n.Comm).(Node)
	}
	editNodes(n.Body, edit)
}

func (n *CompLitExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *CompLitExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.List = copyNodes(c.List)
	return &c
}
func (n *CompLitExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.List, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *CompLitExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}
func (n *CompLitExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.List, edit)
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *ConvExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ConvExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ConvExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *ConvExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}
func (n *ConvExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.TypeWord != nil {
		n.TypeWord = edit(n.TypeWord).(Node)
	}
	if n.SrcRType != nil {
		n.SrcRType = edit(n.SrcRType).(Node)
	}
	if n.ElemRType != nil {
		n.ElemRType = edit(n.ElemRType).(Node)
	}
	if n.ElemElemRType != nil {
		n.ElemElemRType = edit(n.ElemElemRType).(Node)
	}
}

func (n *Decl) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *Decl) copy() Node {
	c := *n
	return &c
}
func (n *Decl) doChildren(do func(Node) bool) bool {
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *Decl) editChildren(edit func(Node) Node) {
	if n.X != nil {
		n.X = edit(n.X).(*Name)
	}
}
func (n *Decl) editChildrenWithHidden(edit func(Node) Node) {
	if n.X != nil {
		n.X = edit(n.X).(*Name)
	}
}

func (n *DynamicType) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *DynamicType) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *DynamicType) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.RType != nil && do(n.RType) {
		return true
	}
	if n.ITab != nil && do(n.ITab) {
		return true
	}
	return false
}
func (n *DynamicType) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	if n.ITab != nil {
		n.ITab = edit(n.ITab).(Node)
	}
}
func (n *DynamicType) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	if n.ITab != nil {
		n.ITab = edit(n.ITab).(Node)
	}
}

func (n *DynamicTypeAssertExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *DynamicTypeAssertExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *DynamicTypeAssertExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.SrcRType != nil && do(n.SrcRType) {
		return true
	}
	if n.RType != nil && do(n.RType) {
		return true
	}
	if n.ITab != nil && do(n.ITab) {
		return true
	}
	return false
}
func (n *DynamicTypeAssertExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.SrcRType != nil {
		n.SrcRType = edit(n.SrcRType).(Node)
	}
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	if n.ITab != nil {
		n.ITab = edit(n.ITab).(Node)
	}
}
func (n *DynamicTypeAssertExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.SrcRType != nil {
		n.SrcRType = edit(n.SrcRType).(Node)
	}
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	if n.ITab != nil {
		n.ITab = edit(n.ITab).(Node)
	}
}

func (n *ForStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ForStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *ForStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Cond != nil && do(n.Cond) {
		return true
	}
	if n.Post != nil && do(n.Post) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	return false
}
func (n *ForStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Cond != nil {
		n.Cond = edit(n.Cond).(Node)
	}
	if n.Post != nil {
		n.Post = edit(n.Post).(Node)
	}
	editNodes(n.Body, edit)
}
func (n *ForStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Cond != nil {
		n.Cond = edit(n.Cond).(Node)
	}
	if n.Post != nil {
		n.Post = edit(n.Post).(Node)
	}
	editNodes(n.Body, edit)
}

func (n *Func) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }

func (n *GoDeferStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *GoDeferStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *GoDeferStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Call != nil && do(n.Call) {
		return true
	}
	if n.DeferAt != nil && do(n.DeferAt) {
		return true
	}
	return false
}
func (n *GoDeferStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Call != nil {
		n.Call = edit(n.Call).(Node)
	}
	if n.DeferAt != nil {
		n.DeferAt = edit(n.DeferAt).(Expr)
	}
}
func (n *GoDeferStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Call != nil {
		n.Call = edit(n.Call).(Node)
	}
	if n.DeferAt != nil {
		n.DeferAt = edit(n.DeferAt).(Expr)
	}
}

func (n *Ident) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *Ident) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *Ident) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *Ident) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *Ident) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *IfStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *IfStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	c.Else = copyNodes(c.Else)
	return &c
}
func (n *IfStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Cond != nil && do(n.Cond) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	if doNodes(n.Else, do) {
		return true
	}
	return false
}
func (n *IfStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Cond != nil {
		n.Cond = edit(n.Cond).(Node)
	}
	editNodes(n.Body, edit)
	editNodes(n.Else, edit)
}
func (n *IfStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Cond != nil {
		n.Cond = edit(n.Cond).(Node)
	}
	editNodes(n.Body, edit)
	editNodes(n.Else, edit)
}

func (n *IndexExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *IndexExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *IndexExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Index != nil && do(n.Index) {
		return true
	}
	return false
}
func (n *IndexExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Index != nil {
		n.Index = edit(n.Index).(Node)
	}
}
func (n *IndexExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Index != nil {
		n.Index = edit(n.Index).(Node)
	}
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
}

func (n *InlineMarkStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *InlineMarkStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *InlineMarkStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *InlineMarkStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *InlineMarkStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *InlinedCallExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *InlinedCallExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	c.ReturnVars = copyNodes(c.ReturnVars)
	return &c
}
func (n *InlinedCallExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	if doNodes(n.ReturnVars, do) {
		return true
	}
	return false
}
func (n *InlinedCallExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Body, edit)
	editNodes(n.ReturnVars, edit)
}
func (n *InlinedCallExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Body, edit)
	editNodes(n.ReturnVars, edit)
}

func (n *InterfaceSwitchStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *InterfaceSwitchStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *InterfaceSwitchStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Case != nil && do(n.Case) {
		return true
	}
	if n.Itab != nil && do(n.Itab) {
		return true
	}
	if n.RuntimeType != nil && do(n.RuntimeType) {
		return true
	}
	if n.Hash != nil && do(n.Hash) {
		return true
	}
	return false
}
func (n *InterfaceSwitchStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Case != nil {
		n.Case = edit(n.Case).(Node)
	}
	if n.Itab != nil {
		n.Itab = edit(n.Itab).(Node)
	}
	if n.RuntimeType != nil {
		n.RuntimeType = edit(n.RuntimeType).(Node)
	}
	if n.Hash != nil {
		n.Hash = edit(n.Hash).(Node)
	}
}
func (n *InterfaceSwitchStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Case != nil {
		n.Case = edit(n.Case).(Node)
	}
	if n.Itab != nil {
		n.Itab = edit(n.Itab).(Node)
	}
	if n.RuntimeType != nil {
		n.RuntimeType = edit(n.RuntimeType).(Node)
	}
	if n.Hash != nil {
		n.Hash = edit(n.Hash).(Node)
	}
}

func (n *JumpTableStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *JumpTableStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *JumpTableStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Idx != nil && do(n.Idx) {
		return true
	}
	return false
}
func (n *JumpTableStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Idx != nil {
		n.Idx = edit(n.Idx).(Node)
	}
}
func (n *JumpTableStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Idx != nil {
		n.Idx = edit(n.Idx).(Node)
	}
}

func (n *KeyExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *KeyExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *KeyExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Key != nil && do(n.Key) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	return false
}
func (n *KeyExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Key != nil {
		n.Key = edit(n.Key).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}
func (n *KeyExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Key != nil {
		n.Key = edit(n.Key).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}

func (n *LabelStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *LabelStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *LabelStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *LabelStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *LabelStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *LinksymOffsetExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *LinksymOffsetExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *LinksymOffsetExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *LinksymOffsetExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *LinksymOffsetExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *LogicalExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *LogicalExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *LogicalExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Y != nil && do(n.Y) {
		return true
	}
	return false
}
func (n *LogicalExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}
func (n *LogicalExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Y != nil {
		n.Y = edit(n.Y).(Node)
	}
}

func (n *MakeExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *MakeExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *MakeExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Len != nil && do(n.Len) {
		return true
	}
	if n.Cap != nil && do(n.Cap) {
		return true
	}
	return false
}
func (n *MakeExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
	if n.Cap != nil {
		n.Cap = edit(n.Cap).(Node)
	}
}
func (n *MakeExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
	if n.Cap != nil {
		n.Cap = edit(n.Cap).(Node)
	}
}

func (n *Name) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }

func (n *NilExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *NilExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *NilExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *NilExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *NilExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *ParenExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ParenExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ParenExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *ParenExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}
func (n *ParenExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *RangeStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *RangeStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Body = copyNodes(c.Body)
	return &c
}
func (n *RangeStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Key != nil && do(n.Key) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	if doNodes(n.Body, do) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *RangeStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Key != nil {
		n.Key = edit(n.Key).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
	editNodes(n.Body, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}
func (n *RangeStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.RType != nil {
		n.RType = edit(n.RType).(Node)
	}
	if n.Key != nil {
		n.Key = edit(n.Key).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
	editNodes(n.Body, edit)
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
	if n.KeyTypeWord != nil {
		n.KeyTypeWord = edit(n.KeyTypeWord).(Node)
	}
	if n.KeySrcRType != nil {
		n.KeySrcRType = edit(n.KeySrcRType).(Node)
	}
	if n.ValueTypeWord != nil {
		n.ValueTypeWord = edit(n.ValueTypeWord).(Node)
	}
	if n.ValueSrcRType != nil {
		n.ValueSrcRType = edit(n.ValueSrcRType).(Node)
	}
}

func (n *ResultExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ResultExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *ResultExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	return false
}
func (n *ResultExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
}
func (n *ResultExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
}

func (n *ReturnStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *ReturnStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Results = copyNodes(c.Results)
	return &c
}
func (n *ReturnStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doNodes(n.Results, do) {
		return true
	}
	return false
}
func (n *ReturnStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Results, edit)
}
func (n *ReturnStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	editNodes(n.Results, edit)
}

func (n *SelectStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SelectStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Cases = copyCommClauses(c.Cases)
	c.Compiled = copyNodes(c.Compiled)
	return &c
}
func (n *SelectStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if doCommClauses(n.Cases, do) {
		return true
	}
	if doNodes(n.Compiled, do) {
		return true
	}
	return false
}
func (n *SelectStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	editCommClauses(n.Cases, edit)
	editNodes(n.Compiled, edit)
}
func (n *SelectStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	editCommClauses(n.Cases, edit)
	editNodes(n.Compiled, edit)
}

func (n *SelectorExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SelectorExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SelectorExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Prealloc != nil && do(n.Prealloc) {
		return true
	}
	return false
}
func (n *SelectorExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}
func (n *SelectorExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Prealloc != nil {
		n.Prealloc = edit(n.Prealloc).(*Name)
	}
}

func (n *SendStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SendStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SendStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Chan != nil && do(n.Chan) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	return false
}
func (n *SendStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Chan != nil {
		n.Chan = edit(n.Chan).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}
func (n *SendStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Chan != nil {
		n.Chan = edit(n.Chan).(Node)
	}
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}

func (n *SliceExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SliceExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SliceExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	if n.Low != nil && do(n.Low) {
		return true
	}
	if n.High != nil && do(n.High) {
		return true
	}
	if n.Max != nil && do(n.Max) {
		return true
	}
	return false
}
func (n *SliceExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Low != nil {
		n.Low = edit(n.Low).(Node)
	}
	if n.High != nil {
		n.High = edit(n.High).(Node)
	}
	if n.Max != nil {
		n.Max = edit(n.Max).(Node)
	}
}
func (n *SliceExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.Low != nil {
		n.Low = edit(n.Low).(Node)
	}
	if n.High != nil {
		n.High = edit(n.High).(Node)
	}
	if n.Max != nil {
		n.Max = edit(n.Max).(Node)
	}
}

func (n *SliceHeaderExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SliceHeaderExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *SliceHeaderExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Ptr != nil && do(n.Ptr) {
		return true
	}
	if n.Len != nil && do(n.Len) {
		return true
	}
	if n.Cap != nil && do(n.Cap) {
		return true
	}
	return false
}
func (n *SliceHeaderExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Ptr != nil {
		n.Ptr = edit(n.Ptr).(Node)
	}
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
	if n.Cap != nil {
		n.Cap = edit(n.Cap).(Node)
	}
}
func (n *SliceHeaderExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Ptr != nil {
		n.Ptr = edit(n.Ptr).(Node)
	}
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
	if n.Cap != nil {
		n.Cap = edit(n.Cap).(Node)
	}
}

func (n *StarExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *StarExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *StarExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *StarExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}
func (n *StarExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *StringHeaderExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *StringHeaderExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *StringHeaderExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Ptr != nil && do(n.Ptr) {
		return true
	}
	if n.Len != nil && do(n.Len) {
		return true
	}
	return false
}
func (n *StringHeaderExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Ptr != nil {
		n.Ptr = edit(n.Ptr).(Node)
	}
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
}
func (n *StringHeaderExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Ptr != nil {
		n.Ptr = edit(n.Ptr).(Node)
	}
	if n.Len != nil {
		n.Len = edit(n.Len).(Node)
	}
}

func (n *StructKeyExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *StructKeyExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *StructKeyExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Value != nil && do(n.Value) {
		return true
	}
	return false
}
func (n *StructKeyExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}
func (n *StructKeyExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Value != nil {
		n.Value = edit(n.Value).(Node)
	}
}

func (n *SwitchStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *SwitchStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	c.Cases = copyCaseClauses(c.Cases)
	c.Compiled = copyNodes(c.Compiled)
	return &c
}
func (n *SwitchStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Tag != nil && do(n.Tag) {
		return true
	}
	if doCaseClauses(n.Cases, do) {
		return true
	}
	if doNodes(n.Compiled, do) {
		return true
	}
	return false
}
func (n *SwitchStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Tag != nil {
		n.Tag = edit(n.Tag).(Node)
	}
	editCaseClauses(n.Cases, edit)
	editNodes(n.Compiled, edit)
}
func (n *SwitchStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Tag != nil {
		n.Tag = edit(n.Tag).(Node)
	}
	editCaseClauses(n.Cases, edit)
	editNodes(n.Compiled, edit)
}

func (n *TailCallStmt) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *TailCallStmt) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *TailCallStmt) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.Call != nil && do(n.Call) {
		return true
	}
	return false
}
func (n *TailCallStmt) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Call != nil {
		n.Call = edit(n.Call).(*CallExpr)
	}
}
func (n *TailCallStmt) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.Call != nil {
		n.Call = edit(n.Call).(*CallExpr)
	}
}

func (n *TypeAssertExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *TypeAssertExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *TypeAssertExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *TypeAssertExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}
func (n *TypeAssertExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
	if n.ITab != nil {
		n.ITab = edit(n.ITab).(Node)
	}
}

func (n *TypeSwitchGuard) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *TypeSwitchGuard) copy() Node {
	c := *n
	return &c
}
func (n *TypeSwitchGuard) doChildren(do func(Node) bool) bool {
	if n.Tag != nil && do(n.Tag) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *TypeSwitchGuard) editChildren(edit func(Node) Node) {
	if n.Tag != nil {
		n.Tag = edit(n.Tag).(*Ident)
	}
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}
func (n *TypeSwitchGuard) editChildrenWithHidden(edit func(Node) Node) {
	if n.Tag != nil {
		n.Tag = edit(n.Tag).(*Ident)
	}
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *UnaryExpr) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *UnaryExpr) copy() Node {
	c := *n
	c.init = copyNodes(c.init)
	return &c
}
func (n *UnaryExpr) doChildren(do func(Node) bool) bool {
	if doNodes(n.init, do) {
		return true
	}
	if n.X != nil && do(n.X) {
		return true
	}
	return false
}
func (n *UnaryExpr) editChildren(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}
func (n *UnaryExpr) editChildrenWithHidden(edit func(Node) Node) {
	editNodes(n.init, edit)
	if n.X != nil {
		n.X = edit(n.X).(Node)
	}
}

func (n *typeNode) Format(s fmt.State, verb rune) { fmtNode(n, s, verb) }
func (n *typeNode) copy() Node {
	c := *n
	return &c
}
func (n *typeNode) doChildren(do func(Node) bool) bool {
	return false
}
func (n *typeNode) editChildren(edit func(Node) Node) {
}
func (n *typeNode) editChildrenWithHidden(edit func(Node) Node) {
}

func copyCaseClauses(list []*CaseClause) []*CaseClause {
	if list == nil {
		return nil
	}
	c := make([]*CaseClause, len(list))
	copy(c, list)
	return c
}
func doCaseClauses(list []*CaseClause, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editCaseClauses(list []*CaseClause, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(*CaseClause)
		}
	}
}

func copyCommClauses(list []*CommClause) []*CommClause {
	if list == nil {
		return nil
	}
	c := make([]*CommClause, len(list))
	copy(c, list)
	return c
}
func doCommClauses(list []*CommClause, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editCommClauses(list []*CommClause, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(*CommClause)
		}
	}
}

func copyNames(list []*Name) []*Name {
	if list == nil {
		return nil
	}
	c := make([]*Name, len(list))
	copy(c, list)
	return c
}
func doNames(list []*Name, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editNames(list []*Name, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(*Name)
		}
	}
}

func copyNodes(list []Node) []Node {
	if list == nil {
		return nil
	}
	c := make([]Node, len(list))
	copy(c, list)
	return c
}
func doNodes(list []Node, do func(Node) bool) bool {
	for _, x := range list {
		if x != nil && do(x) {
			return true
		}
	}
	return false
}
func editNodes(list []Node, edit func(Node) Node) {
	for i, x := range list {
		if x != nil {
			list[i] = edit(x).(Node)
		}
	}
}
