func calculations

func NobilityErrors(nob string, tc []string, ix int) []error {
var allErrs []error
if nob == "Nobl?" {
	allErrs = append(allErrs, fmt.Printf("Nobility undefined"))
}
if nob == "" {
	allErrs = append(allErrs, fmt.Printf("Nobility undefined"))
}

return allErrs
}