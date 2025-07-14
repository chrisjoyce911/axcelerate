package axcelerate

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type StringInt int

func (i *StringInt) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch val := v.(type) {
	case float64:
		*i = StringInt(val)
		return nil
	case string:
		n, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		*i = StringInt(n)
		return nil
	case nil:
		return nil
	}
	return fmt.Errorf("unexpected type for StringInt: %T", v)
}

type StringFloat float64

func (f *StringFloat) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch val := v.(type) {
	case float64:
		*f = StringFloat(val)
		return nil
	case string:
		n, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		*f = StringFloat(n)
		return nil
	case nil:
		return nil
	}
	return fmt.Errorf("unexpected type for StringFloat: %T", v)
}
