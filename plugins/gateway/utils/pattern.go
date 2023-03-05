package utils

type OpCode int

const (
	// OpNop do nothing
	OpNop = OpCode(iota)

	// OpPush push a component to stack
	OpPush

	// OpListPush pushes a component to stack if it matches to the literal
	OpListPush

	// OpPushM concatenates the remaining components and pushes it to stack
	OpPushM

	// OpConcatN pops N items from stack, concatenates them and pushes it back to stack
	OpConcatN

	// OpCapture pops an item and binds it to the variable
	OpCapture

	// OpEnd is the least positive invalid opcode.
	OpEnd
)
