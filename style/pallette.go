package style

// Several common styles for ease of use.

// No style
var None = New()

// Bold only
var Bolded = New(Bold)

// Underline only
var Underlined = New(Underline)

// Bold and Underline
var BoldUnderline = New(Bold, Underline)

// Bold and Green
var BoldSuccess = New(Bold, Green)

// Green
var Success = New(Green)

// Bold and Yellow
var BoldWarning = New(Bold, Yellow)

// Yellow
var Warning = New(Yellow)

// Bold and Red
var BoldError = New(Bold, Red)

// Red
var Error = New(Red)

// Bold and Green
var BoldCreate = New(Bold, Green)

// Green
var Create = New(Green)

// Bold and Red
var BoldDelete = New(Bold, Red)

// Red
var Delete = New(Red)

// Bold and Cyan
var BoldInfo = New(Bold, Cyan)

// Cyan
var Info = New(Cyan)
