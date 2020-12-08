package cpu

type Computer struct {
	Ip, Acc int
	Inst    []Instruction
}

func NewComputer(program []Instruction) Computer {
	return Computer{
		Ip:   0,
		Acc:  0,
		Inst: program,
	}
}

func (c *Computer) Step() bool {
	i := &c.Inst[c.Ip]
	switch i.Op {
	case "nop":
		c.Ip++
	case "acc":
		c.Acc += i.Num
		c.Ip++
	case "jmp":
		c.Ip += i.Num
	}
	return c.Ip != len(c.Inst)
}

func (c *Computer) FindInfiniteLoop() bool {
	seen := make(map[int]bool)
	for !seen[c.Ip] {
		seen[c.Ip] = true
		stillRunning := c.Step()
		if !stillRunning {
			return false
		}
	}
	return true
}
