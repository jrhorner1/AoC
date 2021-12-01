package utils

type Computer struct {
    memory          *[]int      // this is the programs memory space
    pointer         int         // pointer to current memory address
    input           chan int    // input channel
    output          chan int    // output channel
    instruction     int         // the raw opcode instruction
    relativeBase    int         // relative base for parameter 2 operations
}

const (
    memoryMultiplier = 50
)

// initialize a new computer instance
func NewComputer(program *[]int) Computer {
    // create the extended memory space and copy the program (intcode) into it
    memory := make([]int, len(*program) * memoryMultiplier)
    copy(memory, *program)

    return Computer{
        memory: &memory, 
        pointer: 0,
        input: make(chan int, 1), 
        output: make(chan int), 
        instruction: 0,
        relativeBase: 0}
}

func (c *Computer) Run() {
loop: // loop until the program ends
    for {
        c.instruction = c.Read(1) // read the first instruction in parameter 1 mode
        opcode := c.instruction % 100 // get the opcode without the parameters
        switch opcode {
        case 1: c.Add()
        case 2: c.Multiply()
        case 3: c.PutInput()
        case 4: c.PutOutput()
        case 5: c.JumpIfTrue()
        case 6: c.JumpIfFalse()
        case 7: c.LessThan()
        case 8: c.Equals()
        case 9: c.RelativeBaseOffset()
        case 99: break loop // self explanatory, break the loop created above
        }
    }
    close(c.output) // close the output channel once the program ends
    return 
}
// return the correct address based on the parameter mode
func (c *Computer) Parameter(mode int) int { 
    var address int
    switch mode {
    case 0: // position mode
        address = (*c.memory)[c.pointer]
    case 1: // immediate mode
        address = c.pointer
    case 2: // relative mode
        address = (*c.memory)[c.pointer] + c.relativeBase
    }
    c.pointer++ // move the pointer to the next address
    return address
}
// returns the value at the correct memory address
func (c *Computer) Read(mode int) int {
    address := c.Parameter(mode)
    return (*c.memory)[address]
}
// writes the value to the correct memory address
func (c *Computer) Write(value int, mode int) {
    address := c.Parameter(mode)
    (*c.memory)[address] = value
    return
}
// add 2 values, write them to memory
func (c *Computer) Add() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)

    c.Write(a + b, (c.instruction / 10000) % 10)
}

func (c *Computer) Multiply() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)

    c.Write(a * b, (c.instruction / 10000) % 10)
}

func (c *Computer) PutInput() {
    in := <- c.input
    c.Write(in, (c.instruction / 100) % 10)
}

func (c *Computer) GetInput() chan int {
    return c.input
}

func (c *Computer) PutOutput() {
    out := c.Read((c.instruction / 100) % 10)
    c.output <- out
}

func (c *Computer) GetOutput() chan int {
    return c.output
}

func (c *Computer) GetMemory() *[]int {
    return c.memory
}

func (c *Computer) JumpIfTrue() {
    test := c.Read((c.instruction / 100) % 10)
    newPointer := c.Read((c.instruction / 1000) % 10)
    if test != 0 {
        c.pointer = newPointer
    }
}

func (c *Computer) JumpIfFalse() {
    test := c.Read((c.instruction / 100) % 10)
    newPointer := c.Read((c.instruction / 1000) % 10)
    if test == 0 {
        c.pointer = newPointer
    }
}

func (c *Computer) LessThan() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)
    if a < b {
        c.Write(1, (c.instruction / 10000) % 10)
    } else {
        c.Write(0, (c.instruction / 10000) % 10)
    }
}

func (c *Computer) Equals() {
    a := c.Read((c.instruction / 100) % 10)
    b := c.Read((c.instruction / 1000) % 10)
    if a == b {
        c.Write(1, (c.instruction / 10000) % 10)
    } else {
        c.Write(0, (c.instruction / 10000) % 10)
    }
}

func (c *Computer) RelativeBaseOffset() {
    offset := c.Read((c.instruction / 100) % 10)
    c.relativeBase += offset
}



