package release

// Operator This type correspond to a determinated operator (byte-value) that can be used to compare two numbers.
type Operator byte

const (
	Equal          Operator = iota // Equal is the operator that checks if two numbers are equal.
	Less                           // Less is the operator that checks if a number is less than another.
	LessOrEqual                    // LessOrEqual is the operator that checks if a number is less or equal than another.
	Greater                        // Greater is the operator that checks if a number is greater than another.
	GreaterOrEqual                 // GreaterOrEqual is the operator that checks if a number is greater or equal than another.
)
