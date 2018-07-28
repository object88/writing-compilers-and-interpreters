package pascal

type errorCode struct {
	message string
	status  int
}

var AlreadyForwarded = NewErrorCode("Already specified in FORWARD")
var IdentifierRedefined = NewErrorCode("Redefined identifier")
var IdentifierUndefined = NewErrorCode("Undefined identifier")
var IncompatibleAssignment = NewErrorCode("Incompatible assignment")
var IncompatibleTypes = NewErrorCode("Incompatible types")
var InvalidAssignment = NewErrorCode("Invalid assignment statement")
var InvalidCharacter = NewErrorCode("Invalid character")
var InvalidConstant = NewErrorCode("Invalid constant")
var InvalidExponent = NewErrorCode("Invalid exponent")
var InvalidExpression = NewErrorCode("Invalid expression")
var InvalidField = NewErrorCode("Invalid field")
var InvalidFraction = NewErrorCode("Invalid fraction")
var InvalidIdentifierUsage = NewErrorCode("Invalid identifier usage")
var InvalidIndexType = NewErrorCode("Invalid index type")
var InvalidNumber = NewErrorCode("Invalid number")
var InvalidStatement = NewErrorCode("Invalid statement")
var InvalidSubrangeType = NewErrorCode("Invalid subrange type")
var InvalidTarget = NewErrorCode("Invalid target")
var InvalidType = NewErrorCode("Invalid type")
var InvalidVarParam = NewErrorCode("Invalid VAR parameter")
var MissingBegin = NewErrorCode("Missing BEGIN")
var MissingColon = NewErrorCode("Missing :")
var MissingColonEquals = NewErrorCode("Missing :=")
var MissingComma = NewErrorCode("Missing ,")
var MissingConstant = NewErrorCode("Missing constant")
var MissintDo = NewErrorCode("Missing DO")
var MissingDotDot = NewErrorCode("Missing ..")
var MissingEnd = NewErrorCode("Missing END")
var MissingEquals = NewErrorCode("Missing =")
var MissingForControl = NewErrorCode("Invalid FOR control variable")
var MissingIdentifier = NewErrorCode("Missing identifier")
var MissingLeftBraket = NewErrorCode("Missing [")
var MissingOf = NewErrorCode("Missing OF")
var MissingPeriod = NewErrorCode("Missing .")
var MissingProgram = NewErrorCode("Missing PROGRAM")
var MissingRightBracket = NewErrorCode("Missing ]")
var MissingRightParen = NewErrorCode("Missing )")
var MissingSemicolon = NewErrorCode("Missing ;")
var MissingThen = NewErrorCode("Missing THEN")
var MissingToDownto = NewErrorCode("Missing TO or DOWNTO")
var MissingUntil = NewErrorCode("Missing UNTIL")
var MissingVariable = NewErrorCode("Missing variable")
var CaseConstantReused = NewErrorCode("CASE constant reused")
var NotConstantIdentifier = NewErrorCode("Not a constant identifier")
var NotRecordVariable = NewErrorCode("Not a record variable")
var NotTypeIdentifier = NewErrorCode("Not a type identifier")
var RangeInteger = NewErrorCode("Integer literal out of range")
var RangeReal = NewErrorCode("Real literal out of range")
var StackOverflow = NewErrorCode("Stack overflow")
var TooManyLevels = NewErrorCode("Nesting level too deep")
var TooManySubscripts = NewErrorCode("Too many subscripts")
var UnexpectedEOF = NewErrorCode("Unexpected EOF")
var UnexpectedToken = NewErrorCode("Unexpected token")
var Unimplemented = NewErrorCode("Unimplemented feature")
var Unrecognizable = NewErrorCode("Unrecognizable input")
var WrongNumberOfParams = NewErrorCode("Wrong number of actual parameters")

var IO_ERROR = NewErrorCodeWithStatus("Object I/O error", -101)
var TOO_MANY_ERRORS = NewErrorCodeWithStatus("Too many syntax errors", -102)

// NewErrorCode creates a new errorCode instance
func NewErrorCode(message string) errorCode {
	return errorCode{
		message: message,
	}
}

// NewErrorCodeWithStatus creates a new errorCode instance with a status
// assigned
func NewErrorCodeWithStatus(message string, status int) errorCode {
	ec := NewErrorCode(message)
	ec.status = status
	return ec
}

func (ec *errorCode) Error() string {
	return ec.message
}
