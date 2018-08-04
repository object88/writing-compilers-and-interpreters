package pascal

type errorCode struct {
	message string
	status  int
}

// AlreadyForwarded represents error "Already specified in FORWARD"
var AlreadyForwarded = newErrorCode("Already specified in FORWARD")

// IdentifierRedefined represents error "Redefined identifier"
var IdentifierRedefined = newErrorCode("Redefined identifier")

// IdentifierUndefined represents error "Undefined identifier"
var IdentifierUndefined = newErrorCode("Undefined identifier")

// IncompatibleAssignment represents error "Incompatible assignment"
var IncompatibleAssignment = newErrorCode("Incompatible assignment")

// IncompatibleTypes represents error "Incompatible types"
var IncompatibleTypes = newErrorCode("Incompatible types")

// InvalidAssignment represents error "Invalid assignment statement"
var InvalidAssignment = newErrorCode("Invalid assignment statement")

// InvalidCharacter represents error "Invalid character"
var InvalidCharacter = newErrorCode("Invalid character")

// InvalidConstant represents error "Invalid constant"
var InvalidConstant = newErrorCode("Invalid constant")

// InvalidExponent represents error "Invalid exponent"
var InvalidExponent = newErrorCode("Invalid exponent")

// InvalidExpression represents error "Invalid expression"
var InvalidExpression = newErrorCode("Invalid expression")

// InvalidField represents error "Invalid field"
var InvalidField = newErrorCode("Invalid field")

// InvalidFraction represents error "Invalid fraction"
var InvalidFraction = newErrorCode("Invalid fraction")

// InvalidIdentifierUsage represents error "Invalid identifier usage"
var InvalidIdentifierUsage = newErrorCode("Invalid identifier usage")

// InvalidIndexType represents error "Invalid index type"
var InvalidIndexType = newErrorCode("Invalid index type")

// InvalidNumber represents error "Invalid number"
var InvalidNumber = newErrorCode("Invalid number")

// InvalidStatement represents error "Invalid statement"
var InvalidStatement = newErrorCode("Invalid statement")

// InvalidSubrangeType represents error "Invalid subrange type"
var InvalidSubrangeType = newErrorCode("Invalid subrange type")

// InvalidTarget represents error "Invalid target"
var InvalidTarget = newErrorCode("Invalid target")

// InvalidType represents error "Invalid type"
var InvalidType = newErrorCode("Invalid type")

// InvalidVarParam represents error "Invalid VAR parameter"
var InvalidVarParam = newErrorCode("Invalid VAR parameter")

// MissingBegin represents error "Missing BEGIN"
var MissingBegin = newErrorCode("Missing BEGIN")

// MissingColon represents error "Missing :"
var MissingColon = newErrorCode("Missing :")

// MissingColonEquals represents error "Missing :="
var MissingColonEquals = newErrorCode("Missing :=")

// MissingComma represents error "Missing ,"
var MissingComma = newErrorCode("Missing ,")

// MissingConstant represents error "Missing constant"
var MissingConstant = newErrorCode("Missing constant")

// MissintDo represents error "Missing DO"
var MissintDo = newErrorCode("Missing DO")

// MissingDotDot represents error "Missing .."
var MissingDotDot = newErrorCode("Missing ..")

// MissingEnd represents error "Missing END"
var MissingEnd = newErrorCode("Missing END")

// MissingEquals represents error "Missing ="
var MissingEquals = newErrorCode("Missing =")

// MissingForControl represents error "Invalid FOR control variable"
var MissingForControl = newErrorCode("Invalid FOR control variable")

// MissingIdentifier represents error "Missing identifier"
var MissingIdentifier = newErrorCode("Missing identifier")

// MissingLeftBraket represents error "Missing ["
var MissingLeftBraket = newErrorCode("Missing [")

// MissingOf represents error "Missing OF"
var MissingOf = newErrorCode("Missing OF")

// MissingPeriod represents error "Missing ."
var MissingPeriod = newErrorCode("Missing .")

// MissingProgram represents error "Missing PROGRAM"
var MissingProgram = newErrorCode("Missing PROGRAM")

// MissingRightBracket represents error "Missing ]"
var MissingRightBracket = newErrorCode("Missing ]")

// MissingRightParen represents error "Missing )"
var MissingRightParen = newErrorCode("Missing )")

// MissingSemicolon represents error "Missing ;"
var MissingSemicolon = newErrorCode("Missing ;")

// MissingThen represents error "Missing THEN"
var MissingThen = newErrorCode("Missing THEN")

// MissingToDownto represents error "Missing TO or DOWNTO"
var MissingToDownto = newErrorCode("Missing TO or DOWNTO")

// MissingUntil represents error "Missing UNTIL"
var MissingUntil = newErrorCode("Missing UNTIL")

// MissingVariable represents error "Missing variable"
var MissingVariable = newErrorCode("Missing variable")

// CaseConstantReused represents error "CASE constant reused"
var CaseConstantReused = newErrorCode("CASE constant reused")

// NotConstantIdentifier represents error "Not a constant identifier"
var NotConstantIdentifier = newErrorCode("Not a constant identifier")

// NotRecordVariable represents error "Not a record variable"
var NotRecordVariable = newErrorCode("Not a record variable")

// NotTypeIdentifier represents error "Not a type identifier"
var NotTypeIdentifier = newErrorCode("Not a type identifier")

// RangeInteger represents error "Integer literal out of range"
var RangeInteger = newErrorCode("Integer literal out of range")

// RangeReal represents error "Real literal out of range"
var RangeReal = newErrorCode("Real literal out of range")

// StackOverflow represents error "Stack overflow"
var StackOverflow = newErrorCode("Stack overflow")

// TooManyLevels represents error "Nesting level too deep"
var TooManyLevels = newErrorCode("Nesting level too deep")

// TooManySubscripts represents error "Too many subscripts"
var TooManySubscripts = newErrorCode("Too many subscripts")

// UnexpectedEOF represents error "Unexpected EOF"
var UnexpectedEOF = newErrorCode("Unexpected EOF")

// UnexpectedToken represents error "Unexpected token"
var UnexpectedToken = newErrorCode("Unexpected token")

// Unimplemented represents error "Unimplemented feature"
var Unimplemented = newErrorCode("Unimplemented feature")

// Unrecognizable represents error "Unrecognizable input"
var Unrecognizable = newErrorCode("Unrecognizable input")

// WrongNumberOfParams represents error "Wrong number of actual parameters"
var WrongNumberOfParams = newErrorCode("Wrong number of actual parameters")

// IO_ERROR represents error "Object I/O error"
var IO_ERROR = newErrorCodeWithStatus("Object I/O error", -101)

// TOO_MANY_ERRORS represents error "Too many syntax errors"
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
