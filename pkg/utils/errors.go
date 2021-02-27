package utils

//FlagError struct
type FlagError struct {
	Err error
}

//PrintCmdError print
func PrintCmdError(fe FlagError) string {
	return "Hello World"
}
