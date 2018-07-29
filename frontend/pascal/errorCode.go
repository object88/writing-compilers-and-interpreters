package pascal

type errorCode struct {
	message string
	status  int
}

var AlreadyForwarded = newErrorCode("Already specified in FORWARD")
var IdentifierRedefined = newErrorCode("Redefined identifier")
var IdentifierUndefined = newErrorCode("Undefined identifier")
var IncompatibleAssignment = newErrorCode("Incompatible assignment")
var IncompatibleTypes = newErrorCode("Incompatible types")
var InvalidAssignment = newErrorCode("Invalid assignment statement")
var InvalidCharacter = newErrorCode("Invalid character")
var InvalidConstant = newErrorCode("Invalid constant")
var InvalidExponent = newErrorCode("Invalid exponent")
var InvalidExpression = newErrorCode("Invalid expression")
var InvalidField = newErrorCode("Invalid field")
var InvalidFraction = newErrorCode("Invalid fraction")
var InvalidIdentifierUsage = newErrorCode("Invalid identifier usage")
var InvalidIndexType = newErrorCode("Invalid index type")
var InvalidNumber = newErrorCode("Invalid number")
var InvalidStatement = newErrorCode("Invalid statement")
var InvalidSubrangeType = newErrorCode("Invalid subrange type")
var InvalidTarget = newErrorCode("Invalid target")
var InvalidType = newErrorCode("Invalid type")
var InvalidVarParam = newErrorCode("Invalid VAR parameter")
var MissingBegin = newErrorCode("Missing BEGIN")
var MissingColon = newErrorCode("Missing :")
var MissingColonEquals = newErrorCode("Missing :=")
var MissingComma = newErrorCode("Missing ,")
var MissingConstant = newErrorCode("Missing constant")
var MissintDo = newErrorCode("Missing DO")
var MissingDotDot = newErrorCode("Missing ..")
var MissingEnd = newErrorCode("Missing END")
var MissingEquals = newErrorCode("Missing =")
var MissingForControl = newErrorCode("Invalid FOR control variable")
var MissingIdentifier = newErrorCode("Missing identifier")
var MissingLeftBraket = newErrorCode("Missing [")
var MissingOf = newErrorCode("Missing OF")
var MissingPeriod = newErrorCode("Missing .")
var MissingProgram = newErrorCode("Missing PROGRAM")
var MissingRightBracket = newErrorCode("Missing ]")
var MissingRightParen = newErrorCode("Missing )")
var MissingSemicolon = newErrorCode("Missing ;")
var MissingThen = newErrorCode("Missing THEN")
var MissingToDownto = newErrorCode("Missing TO or DOWNTO")
var MissingUntil = newErrorCode("Missing UNTIL")
var MissingVariable = newErrorCode("Missing variable")
var CaseConstantReused = newErrorCode("CASE constant reused")
var NotConstantIdentifier = newErrorCode("Not a constant identifier")
var NotRecordVariable = newErrorCode("Not a record variable")
var NotTypeIdentifier = newErrorCode("Not a type identifier")
var RangeInteger = newErrorCode("Integer literal out of range")
var RangeReal = newErrorCode("Real literal out of range")
var StackOverflow = newErrorCode("Stack overflow")
var TooManyLevels = newErrorCode("Nesting level too deep")
var TooManySubscripts = newErrorCode("Too many subscripts")
var UnexpectedEOF = newErrorCode("Unexpected EOF")
var UnexpectedToken = newErrorCode("Unexpected token")
var Unimplemented = newErrorCode("Unimplemented feature")
var Unrecognizable = newErrorCode("Unrecognizable input")
var WrongNumberOfParams = newErrorCode("Wrong number of actual parameters")

var IO_ERROR = newErrorCodeWithStatus("Object I/O error", -101)
var TOO_MANY_ERRORS = newErrorCodeWithStatus("Too many syntax errors", -102)

// newErrorCode creates a new errorCode instance
func newErrorCode(message string) errorCode {
	return errorCode{
		message: message,
	}
}

// newErrorCodeWithStatus creates a new errorCode instance with a status
// assigned
func newErrorCodeWithStatus(message string, status int) errorCode {
	ec := newErrorCode(message)
	ec.status = status
	return ec
}

func (ec *errorCode) Error() string {
	return ec.message
}
