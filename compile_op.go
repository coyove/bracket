package potatolang

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/coyove/potatolang/parser"
)

const (
	errUndeclaredVariable = " %+v: undeclared variable"
)

var _nodeRegA = parser.NewNode(regA)

func (table *symtable) compileSetOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	aDest := atoms[1].Value.(parser.Atom)
	newYX, ok := table.get(aDest)
	if atoms[0].A() == parser.AMove {
		if !ok || (strings.HasPrefix(string(aDest), "$") && (newYX>>10 > 0)) {
			// variable names start with "$" will be a local variable
			newYX = table.borrowAddress()
			table.put(aDest, newYX)
		}
	} else {
		newYX = table.borrowAddress()
		table.put(aDest, newYX)
	}

	buf := newpacket()
	aSrc := atoms[2]
	switch aSrc.Type() {
	case parser.Natom:
		valueIndex, ok := table.get(aSrc.Value.(parser.Atom))
		if !ok {
			err = fmt.Errorf(errUndeclaredVariable, aSrc)
			return
		}
		buf.WriteOP(OpSet, newYX, valueIndex)
		return buf, newYX, nil
	case parser.Nnumber, parser.Nstring:
		buf.WriteOP(OpSet, newYX, table.loadK(&buf, aSrc.Value))
		return buf, newYX, nil
	}

	code, _, err = table.compileCompoundInto(aSrc, false, newYX)
	if err != nil {
		return
	}
	buf.Write(code)
	buf.WritePos(atoms[0].Meta)
	return buf, newYX, nil
}

func (table *symtable) compileMapArrayOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	if code, err = table.decompound(atoms[1].C()); err != nil {
		return
	}

	args := atoms[1].C()
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			if err = table.writeOpcode(&code, OpPush, args[i], nil); err != nil {
				return
			}
		} else {
			if err = table.writeOpcode(&code, OpPush2, args[i], args[i+1]); err != nil {
				return
			}
		}
	}

	if atoms[0].Value.(parser.Atom) == parser.AMap {
		code.WriteOP(OpMakeStruct, 0, 0)
	} else {
		code.WriteOP(OpMakeSlice, 0, 0)
	}
	code.WritePos(atoms[0].Meta)
	return code, regA, nil
}

// writeOpcode3 accepts 3 arguments at most, 2 arguments will be encoded into opcode itself, the 3rd one will be in regA
func (table *symtable) writeOpcode3(bop _Opcode, atoms []*parser.Node) (buf packet, yx uint16, err error) {
	// first atom: the op name, tail atoms: the args
	if len(atoms) > 4 {
		panic("shouldn't happen: too many arguments")
	}

	atoms = append([]*parser.Node{}, atoms...) // duplicate

	if bop == OpStore || bop == OpSlice {
		if buf, err = table.decompoundWithoutA(atoms[1:]); err != nil {
			return
		}

		if err = table.writeOpcode(&buf, OpSet, _nodeRegA, atoms[1]); err != nil {
			return
		}

		err = table.writeOpcode(&buf, bop, atoms[2], atoms[3])
		return buf, regA, err
	}

	buf, err = table.decompound(atoms[1:])
	if err != nil {
		return
	}

	switch bop {
	case OpTypeof, OpNot, OpAddressOf, OpRet, OpYield, OpLen, OpPushVararg, OpCopyStack:
		// unary op
		err = table.writeOpcode(&buf, bop, atoms[1], nil)
	default:
		// binary op
		err = table.writeOpcode(&buf, bop, atoms[1], atoms[2])
	}

	return buf, regA, err
}

func (table *symtable) compileFlatOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	head := atoms[0].Value.(parser.Atom)
	switch head {
	case "nop":
		return newpacket(), regA, nil
	}
	op, ok := flatOpMapping[head]
	if !ok {
		err = fmt.Errorf("%+v: invalid op", atoms[0])
		return
	}
	code, yx, err = table.writeOpcode3(op, atoms)
	for _, a := range atoms {
		if a.Meta.Source != "" {
			code.WritePos(a.Meta)
			break
		}
	}
	return
}

// [and a b] => $a = a if not a then return else $a = b end
// [or a b]  => $a = a if a then do nothing else $a = b end
func (table *symtable) compileAndOrOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	bop := OpIfNot
	if atoms[0].Value.(parser.Atom) == parser.AOr {
		bop = OpIf
	}

	buf := newpacket()

	if err = table.writeOpcode(&buf, OpSet, _nodeRegA, atoms[1]); err != nil {
		return
	}
	buf.WriteOP(bop, regA, 0)
	c2 := buf.Len()

	if err = table.writeOpcode(&buf, OpSet, _nodeRegA, atoms[2]); err != nil {
		return
	}

	_, yx, _ = op(buf.data[c2-1])
	buf.data[c2-1] = makejmpop(bop, yx, buf.Len()-c2)
	buf.WritePos(atoms[0].Meta)
	return buf, regA, nil
}

// [if condition [true-chain ...] [false-chain ...]]
func (table *symtable) compileIfOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	condition := atoms[1]
	trueBranch, falseBranch := atoms[2], atoms[3]
	buf := newpacket()

	code, yx, err = table.compileNode(condition)
	if err != nil {
		return newpacket(), 0, err
	}
	buf.Write(code)
	condyx := yx

	var trueCode, falseCode packet
	trueCode, yx, err = table.compileChainOp(trueBranch)
	if err != nil {
		return
	}
	falseCode, yx, err = table.compileChainOp(falseBranch)
	if err != nil {
		return
	}
	if len(falseCode.data) > 0 {
		buf.WriteJmpOP(OpIfNot, condyx, len(trueCode.data)+1)
		buf.WritePos(atoms[0].Meta)
		buf.Write(trueCode)
		buf.WriteJmpOP(OpJmp, 0, len(falseCode.data))
		buf.Write(falseCode)
	} else {
		buf.WriteJmpOP(OpIfNot, condyx, len(trueCode.data))
		buf.WritePos(atoms[0].Meta)
		buf.Write(trueCode)
	}
	return buf, regA, nil
}

// [call callee [args ...]]
func (table *symtable) compileCallOp(nodes []*parser.Node) (code packet, yx uint16, err error) {
	tmp := append([]*parser.Node{nodes[1]}, nodes[2].C()...)
	code, err = table.decompound(tmp)
	if err != nil {
		return
	}

	var varIndex uint16
	var ok bool
	switch callee := tmp[0]; callee.Type() {
	case parser.Natom:
		varIndex, ok = table.get(callee.Value.(parser.Atom))
		if !ok {
			err = fmt.Errorf(errUndeclaredVariable, callee)
			return
		}
	case parser.Naddr:
		varIndex = callee.Value.(uint16)
	default:
		err = fmt.Errorf("%+v: invalid callee", callee)
		return
	}

	for i := 1; i < len(tmp); i += 2 {
		if i+1 >= len(tmp) {
			if err = table.writeOpcode(&code, OpPush, tmp[i], nil); err != nil {
				return
			}
		} else {
			if err = table.writeOpcode(&code, OpPush2, tmp[i], tmp[i+1]); err != nil {
				return
			}
		}
	}

	for i := 1; i < len(tmp); i++ {
		if tmp[i].Type() == parser.Naddr {
			table.returnAddress(tmp[i].Value.(uint16))
		}
	}

	code.WriteOP(OpCall, varIndex, 0)
	code.WritePos(nodes[0].Meta)
	return code, regA, nil
}

// [lambda name? [namelist] [chain ...]]
func (table *symtable) compileLambdaOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	table.envescape = true
	name := atoms[1].Value.(parser.Atom)
	newtable := newsymtable()
	newtable.parent = table

	params := atoms[2]
	if params.Type() != parser.Ncompound {
		err = fmt.Errorf("%+v: invalid arguments list", atoms[0])
		return
	}

	for i, p := range params.C() {
		argname := p.Value.(parser.Atom)
		if strings.HasSuffix(string(argname), "...") {
			if i != params.Cn()-1 {
				return code, 0, fmt.Errorf("%+v: vararg must be the last parameter", atoms[0])
			}
			argname = argname[:len(argname)-3]
			atoms[3] = parser.CompNode(
				parser.AChain,
				parser.CompNode(
					parser.AMove,
					argname,
					parser.CompNode(parser.AForeach, parser.NewNode(uint16(i))).SetPos0(atoms[0]),
				).SetPos0(atoms[0]),
				atoms[3],
			)
		}
		if _, ok := newtable.sym[argname]; ok {
			return newpacket(), 0, fmt.Errorf("%+v: duplicated parameter: %s", atoms[0], argname)
		}
		newtable.put(argname, uint16(i))
	}

	ln := len(newtable.sym)
	if ln > 255 {
		return newpacket(), 0, fmt.Errorf("%+v: do you really need more than 255 arguments?", atoms[0])
	}

	newtable.vp = uint16(ln)

	code, yx, err = newtable.compileChainOp(atoms[3])
	if err != nil {
		return
	}

	code.WriteOP(OpEOB, 0, 0)
	buf := newpacket()
	cls := Closure{}
	cls.ArgsCount = byte(ln)
	if newtable.y {
		cls.Set(ClsYieldable)
	}
	if !newtable.envescape {
		cls.Set(ClsNoEnvescape)
	}
	buf.WriteOP(OpLambda, uint16(ln), uint16(cls.options))
	buf.Write32(uint32(len(newtable.consts)))
	buf.WriteConsts(newtable.consts)

	cls.Code = code.data
	src := string(name) + cls.String() + "@" + code.source
	if len(src) > 4095 {
		src = src[:4095]
	}

	// 26bit Code Pos length, 26bit Code data length, 12bit Code source length
	buf.Write64(uint64(uint32(len(code.pos))&0x03ffffff)<<38 +
		uint64(uint32(len(code.data))&0x03ffffff)<<12 +
		uint64(uint16(len(src))&0x0fff))
	buf.WriteRaw(u32FromBytes([]byte(src)))
	buf.WriteRaw(u32FromBytes(code.pos))

	// Note buf.source will be set to Code.source in buf.Write
	// but buf.source shouldn't be changed
	code.source = buf.source
	buf.Write(code)
	buf.WritePos(atoms[0].Meta)
	return buf, regA, nil
}

var staticWhileHack [8]uint32

func init() {
	rand.Seed(time.Now().Unix())
	for i := range staticWhileHack {
		staticWhileHack[i] = makeop(OpNOP, uint16(rand.Uint32()), uint16(rand.Uint32()))
	}
}

// [continue | break]
func (table *symtable) compileContinueBreakOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	buf := newpacket()
	if len(table.continueNode) == 0 {
		err = fmt.Errorf("%+v: invalid statement outside a for loop", atoms[0])
		return
	}

	if atoms[0].Value.(parser.Atom) == parser.AContinue {
		cn := table.continueNode[len(table.continueNode)-1]
		code, yx, err = table.compileChainOp(cn)
		if err != nil {
			return
		}
		buf.Write(code)
		buf.WriteRaw(staticWhileHack[:4]) // write a 'continue' placeholder
		return buf, regA, nil
	}

	buf.WriteRaw(staticWhileHack[4:]) // write a 'break' placeholder
	return buf, regA, nil
}

// [for condition incr [chain ...]]
func (table *symtable) compileWhileOp(atoms []*parser.Node) (code packet, yx uint16, err error) {
	condition := atoms[1]
	buf := newpacket()

	code, yx, err = table.compileNode(condition)
	if err != nil {
		return
	}
	buf.Write(code)
	cond := yx

	table.continueNode = append(table.continueNode, atoms[2])
	code, yx, err = table.compileChainOp(atoms[3])
	if err != nil {
		return
	}
	var icode packet
	icode, yx, err = table.compileChainOp(atoms[2])
	if err != nil {
		return
	}
	table.continueNode = table.continueNode[:len(table.continueNode)-1]

	code.Write(icode)
	buf.WriteJmpOP(OpIfNot, cond, len(code.data)+1) // +---> test cond  ---+
	buf.Write(code)                                 // |     Code ...      | (break)
	buf.WriteJmpOP(OpJmp, 0, -buf.Len()-1)          // +---- continue      |
	//                                                   following Code <--+

	code = buf
	code2 := u32Bytes(code.data)

	// search for special 'continue' placeholder and replace it with a OP_JMP to the
	// beginning of the Code
	continueflag := u32Bytes(staticWhileHack[:4])
	for i := 0; i < len(code2); {
		x := bytes.Index(code2[i:], continueflag)
		if x == -1 {
			break
		}
		idx := (i + x) / 4
		code.data[idx] = makejmpop(OpJmp, 0, -idx-1)
		i += x + 4
	}

	// search for special 'break' placeholder and replace it with a OP_JMP to the
	// end of the Code
	breakflag := u32Bytes(staticWhileHack[4:])
	for i := 0; i < len(code2); {
		x := bytes.Index(code2[i:], breakflag)
		if x == -1 {
			break
		}
		idx := (i + x) / 4
		code.data[idx] = makejmpop(OpJmp, 0, len(code.data)-idx-1)
		i += x + 4
	}
	return buf, regA, nil
}
