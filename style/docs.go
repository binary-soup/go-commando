/*
Package style provides semantic types and functions for styling terminal outputs using ANSI escape sequences.

# Style Type

The Style type is a string slice for storing the various modes. Use the New function to create a style:

	s := style.New(style.Bold, style.Blue)

Note: combining contradictory modes such as Bold and Dim or multiple colors results in undefined behavior.

# Pallette

Package style provides several predefined styles for convenience. For example:

	var BoldUnderline = New(Bold, Underline)
	var Success = New(Green)
	var Error = New(Red)

# Printing

Once a style is created, several fmt like functions can be used to apply the style to a string. For example:

	style.Success.Println("Task was successful!")

To format as a string:

	style.Info.FormatF("%d tasks", count)

To use multiple styles, combine with fmt formatting:

	fmt.Printf("%s: %s\n", style.BoldError.Format("ERROR"), style.Error.Format(err.Error()))

Note: there are no fmt.Fprint variants as styling only makes sense when the target is the console.
*/
package style
