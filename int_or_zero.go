package axcelerate

import (
    "bytes"
    "encoding/json"
    "strconv"
)

// IntOrZero is a tolerant integer type that unmarshals from JSON numbers,
// quoted numeric strings, empty strings, or null. Empty string and null
// become zero instead of causing an error.
type IntOrZero int

// UnmarshalJSON implements a tolerant unmarshaler for IntOrZero.
func (i *IntOrZero) UnmarshalJSON(b []byte) error {
    s := bytes.TrimSpace(b)
    if len(s) == 0 || bytes.Equal(s, []byte("null")) {
        *i = 0
        return nil
    }

    // Quoted string (could be empty or numeric)
    if s[0] == '"' && s[len(s)-1] == '"' {
        var str string
        if err := json.Unmarshal(s, &str); err != nil {
            *i = 0
            return nil
        }
        if str == "" {
            *i = 0
            return nil
        }
        if v, err := strconv.Atoi(str); err == nil {
            *i = IntOrZero(v)
            return nil
        }
        // fallback: set zero on parse failure
        *i = 0
        return nil
    }

    // Numeric (unquoted)
    if v, err := strconv.Atoi(string(s)); err == nil {
        *i = IntOrZero(v)
        return nil
    }

    // Try float decode then cast
    var f float64
    if err := json.Unmarshal(s, &f); err == nil {
        *i = IntOrZero(int(f))
        return nil
    }

    // If we get here, be tolerant and set zero.
    *i = 0
    return nil
}

// Int returns the underlying int value.
func (i IntOrZero) Int() int { return int(i) }
