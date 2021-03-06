package script

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/coyove/script/parser"
)

func inst(op byte, a, b uint16) uint32 {
	// 6 + 13 + 13
	return uint32(op)<<26 + uint32(a&0x1fff)<<13 + uint32(b&0x1fff)
}

func jmpInst(op byte, dist int) uint32 {
	if dist < -(1<<23) || dist >= 1<<23 {
		panic("long jump")
	}
	// 6 + 26
	return uint32(op)<<26 + uint32(dist+1<<23)
}

func splitInst(op uint32) (op1 byte, a, b uint16) {
	op1 = byte(op >> 26)
	a = uint16(op>>13) & 0x1fff
	b = uint16(op) & 0x1fff
	return
}

type posVByte []byte

func (p *posVByte) append(idx uint32, line uint32) {
	v := func(v uint64) {
		*p = append(*p, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		n := binary.PutUvarint((*p)[len(*p)-10:], v)
		*p = (*p)[:len(*p)-10+n]
	}
	v(uint64(idx))
	v(uint64(line))
}

func (p posVByte) read(i int) (next int, idx, line uint32) {
	rd := p[i:]
	a, n := binary.Uvarint(rd)
	b, n2 := binary.Uvarint(rd[n:])
	if n == 0 || n2 == 0 {
		next = len(p) + 1
		return
	}
	return i + n + n2, uint32(a), uint32(b)
}

type packet struct {
	Code []uint32
	Pos  posVByte
}

func (b *packet) writeInst(op byte, opa, opb uint16) {
	b.Code = append(b.Code, inst(op, opa, opb))
}

func (b *packet) writeJmpInst(op byte, d int) {
	b.Code = append(b.Code, jmpInst(op, d))
}

func (b *packet) writePos(p parser.Position) {
	if p.Line == 0 {
		// Debug Code, used to detect a null meta struct
		panicf("DEBUG: null line")
	}
	b.Pos.append(uint32(len(b.Code)), p.Line)
}

func (b *packet) truncateLast() {
	if len(b.Code) > 0 {
		b.Code = b.Code[:len(b.Code)-1]
	}
}

func (b *packet) Len() int {
	return len(b.Code)
}

var (
	biOp = map[byte]string{
		OpAdd:     parser.AAdd,
		OpSub:     parser.ASub,
		OpMul:     parser.AMul,
		OpDiv:     parser.ADiv,
		OpIDiv:    parser.AIDiv,
		OpMod:     parser.AMod,
		OpEq:      parser.AEq,
		OpNeq:     parser.ANeq,
		OpLess:    parser.ALess,
		OpLessEq:  parser.ALessEq,
		OpLoad:    parser.ALoad,
		OpStore:   parser.AStore,
		OpBitAnd:  parser.ABitAnd,
		OpBitOr:   parser.ABitOr,
		OpBitXor:  parser.ABitXor,
		OpBitLsh:  parser.ABitLsh,
		OpBitRsh:  parser.ABitRsh,
		OpBitURsh: parser.ABitURsh,
	}
	uOp = map[byte]string{
		OpBitNot:     parser.ABitNot,
		OpNot:        parser.ANot,
		OpRet:        parser.AReturn,
		OpPush:       "push",
		OpPushVararg: "pushvararg",
	}
)

func pkPrettify(c *Func, p *Program, toplevel bool) string {
	sb := &bytes.Buffer{}
	sb.WriteString("+ START " + c.String() + "\n")

	readAddr := func(a uint16, rValue bool) string {
		if a == regA {
			return "$a"
		}

		suffix := ""
		if rValue {
			if a&0xfff == p.NilIndex {
				return "nil"
			}
			if a>>12 == 1 || toplevel {
				v := (*p.Stack)[a&0xfff]
				if v != Nil {
					text := v.JSONString()
					if len(text) > 120 {
						text = text[:60] + "..." + text[len(text)-60:]
					}
					suffix = "(" + text + ")"
				}
			}
		}

		if a>>12 == 1 {
			return fmt.Sprintf("g$%d", a&0xfff) + suffix
		}
		return fmt.Sprintf("$%d", a&0xfff) + suffix
	}

	oldpos := c.Code.Pos
	lastLine := uint32(0)

	for i, inst := range c.Code.Code {
		cursor := uint32(i) + 1
		bop, a, b := splitInst(inst)

		if len(c.Code.Pos) > 0 {
			next, op, line := c.Code.Pos.read(0)
			// log.Println(cursor, splitInst, unsafe.Pointer(&Pos))
			for uint32(cursor) > op {
				c.Code.Pos = c.Code.Pos[next:]
				if len(c.Code.Pos) == 0 {
					break
				}
				if next, op, line = c.Code.Pos.read(0); uint32(cursor) <= op {
					break
				}
			}

			if op == uint32(cursor) {
				x := "."
				if line != lastLine {
					x = strconv.Itoa(int(line))
					lastLine = line
				}
				sb.WriteString(fmt.Sprintf("|%-4s % 4d| ", x, cursor-1))
				c.Code.Pos = c.Code.Pos[next:]
			} else {
				sb.WriteString(fmt.Sprintf("|     % 4d| ", cursor-1))
			}
		} else {
			sb.WriteString(fmt.Sprintf("|$    % 4d| ", cursor-1))
		}

		switch bop {
		case OpSet:
			sb.WriteString(readAddr(a, false) + " = " + readAddr(b, true))
		case OpArray:
			sb.WriteString("array")
		case OpMap:
			sb.WriteString("map")
		case OpLoadFunc:
			cls := p.Functions[a]
			sb.WriteString("loadfunc " + cls.Name + "\n")
			sb.WriteString(pkPrettify(cls, p, false))
		case OpTailCall, OpCall:
			if b != regPhantom {
				sb.WriteString("push " + readAddr(b, true) + " -> ")
			}
			if bop == OpTailCall {
				sb.WriteString("tail")
			}
			sb.WriteString("call " + readAddr(a, true))
		case OpIfNot, OpJmp:
			pos := int32(inst&0xffffff) - 1<<23
			pos2 := uint32(int32(cursor) + pos)
			if bop == OpIfNot {
				sb.WriteString("if not $a ")
			}
			sb.WriteString(fmt.Sprintf("jmp %d to %d", pos, pos2))
		case OpInc:
			sb.WriteString("inc " + readAddr(a, false) + " " + readAddr(b, true))
		default:
			if us, ok := uOp[bop]; ok {
				sb.WriteString(us + " " + readAddr(a, true))
			} else if bs, ok := biOp[bop]; ok {
				sb.WriteString(bs + " " + readAddr(a, true) + " " + readAddr(b, true))
			} else {
				sb.WriteString(fmt.Sprintf("? %02x", bop))
			}
		}

		sb.WriteString("\n")
	}

	c.Code.Pos = oldpos

	sb.WriteString("+ END " + c.String())
	return sb.String()
}
