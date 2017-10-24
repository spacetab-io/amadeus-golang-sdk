package formats

// Format limitations: an..35
type AlphaNumericStringLength1To35 string

// Format limitations: an..14
type AlphaNumericStringLength1To14 string

// Format limitations: an1
type AlphaNumericStringLength1To1 string

// Format limitations: an..30
type AlphaNumericStringLength1To30 string

// Format limitations: n..15
type NumericIntegerLength1To15 int32

// Format limitations: an..99999
type AlphaNumericStringLength1To99999 string

// Format limitations: an..3
type AlphaNumericStringLength1To3 string

// Format limitations: an..9
type AlphaNumericStringLength1To9 string

// Format limitations: an..17
type AlphaNumericStringLength1To17 string

// Format limitations: an..25
type AlphaNumericStringLength1To25 string

// Format limitations: an..5
type AlphaNumericStringLength1To5 string

// Format limitations: an..6
type AlphaNumericStringLength1To6 string

// Format limitations: an..70
type AlphaNumericStringLength1To70 string

// Format limitations: a..6
type AlphaStringLength1To6 string

// Format limitations: an..10
type AlphaNumericStringLength1To10 string

// Format limitations: an..4
type AlphaNumericStringLength1To4 string

// Format limitations: n..3
type NumericIntegerLength1To3 int32

// Format limitations: an..2
type AlphaNumericStringLength1To2 string

// Format limitations: an3
type AlphaNumericStringLength3To3 string

// Format limitations: an2
type AlphaNumericStringLength2To2 string

// Format limitations: n..2
type NumericIntegerLength1To2 int32

// Format limitations: n..18
type NumericIntegerLength1To18 int32

// Format limitations: n..4
type NumericIntegerLength1To4 int32

// Format limitations: an..199
type AlphaNumericStringLength1To199 string

// Format limitations: n..5
type NumericIntegerLength1To5 int32

// Format limitations: an..8
type AlphaNumericStringLength1To8 string

// Value of the year. Example: 2003
type YearYYYY string

// Value of the month. Only significant digits are mandatory. Example: 7
type MonthMM string

// Value of the day in the  month. Only significant digits are mandatory. Example: 7
type DayNN string

// Value of the hours in the  time. Only significant digits are mandatory. Example: 7
type HourHH string

// Value of the minutes in the  time. Only significant digits are mandatory. Example: 7
type MinuteMM string
