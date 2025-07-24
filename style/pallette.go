package style

// Several common styles for ease of use.

var None = New()

var Bolded = New(Bold)
var Underlined = New(Underline)
var BoldUnderline = New(Bold, Underline)

var BoldSuccess = New(Bold, Green)
var Success = New(Green)

var BoldWarning = New(Bold, Yellow)
var Warning = New(Yellow)

var BoldError = New(Bold, Red)
var Error = New(Red)

var BoldCreate = New(Bold, Green)
var Create = New(Green)

var BoldDelete = New(Bold, Red)
var Delete = New(Red)

var BoldInfo = New(Bold, Cyan)
var Info = New(Cyan)

var BoldHighlight = New(Bold, Yellow)
var Highlight = New(Yellow)
