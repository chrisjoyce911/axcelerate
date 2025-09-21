package axcelerate

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
)

// CollectDataWarnings inspects a decoded target struct and the raw JSON body
// that produced it and returns a slice of human-readable warning messages
// describing potential data issues.
//
// Purpose
//   - Help maintainers discover where struct field types may be too loose
//     (for example `interface{}`) or where the JSON shape does not match the
//     declared Go type.
//
// Behavior
//   - Flags struct fields that are declared `interface{}` and recommends a
//     concrete Go type inferred from the JSON value (when available).
//   - Flags top-level JSON keys present in the response that do not have a
//     matching `json:"..."` struct field on the target and recommends a Go
//     type for those keys.
//   - Detects simple primitive mismatches (for example JSON string vs target
//     numeric type) and includes a recommended Go type.
//
// Parameters
//   - target: a pointer to the struct instance into which the JSON was
//     unmarshaled. The function expects a struct or pointer to struct; other
//     types will produce a validation warning.
//   - body: the raw JSON bytes returned from the API (used to infer types and
//     detect unmapped keys).
//
// Returns
//   - []string: a list of warning messages. Each message is advisory and does
//     not affect program flow.
//
// Limitations
//   - Type recommendations are heuristic-based (e.g. numeric values without
//     fractional part are recommended as `int`). Dates and domain-specific
//     formats are not recognized; date-like strings will be recommended as
//     `string` unless you add custom rules.
//
// Example usage
//
//	var detail InstanceDetail
//	resp, _ := client.Courses.GetCoursesInstanceDetail(id, "w")
//	_ = json.Unmarshal([]byte(resp.Body), &detail)
//	resp.DataWarning = CollectDataWarnings(&detail, []byte(resp.Body))
//
// Thread-safety: the function is read-only and has no side effects; it is
// safe to call from multiple goroutines as long as the caller's data is not
// concurrently modified.
func CollectDataWarnings(target interface{}, body []byte) []string {
	var warnings []string
	if target == nil || len(body) == 0 {
		return warnings
	}

	// Unmarshal into a generic map to inspect JSON types
	var raw map[string]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		warnings = append(warnings, fmt.Sprintf("unable to parse body for validation: %v", err))
		return warnings
	}

	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		warnings = append(warnings, "validation target is not a struct")
		return warnings
	}

	// Recursive inspector: walk raw JSON and corresponding Go type to find
	// unmapped keys, interface{} usages and primitive mismatches at any depth.
	var inspectRawAgainstType func(path string, rawVal interface{}, typ reflect.Type)

	// helper to find struct field by json tag name
	findFieldByJSON := func(t reflect.Type, name string) (reflect.StructField, bool) {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			tag := f.Tag.Get("json")
			if idx := len(tag); idx == 0 {
				tag = f.Name
			} else if comma := index(tag, ','); comma >= 0 {
				tag = tag[:comma]
			}
			if tag == name {
				return f, true
			}
		}
		return reflect.StructField{}, false
	}

	inspectRawAgainstType = func(path string, rawVal interface{}, typ reflect.Type) {
		if typ == nil {
			return
		}
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}

		switch typ.Kind() {
		case reflect.Struct:
			// rawVal should be a map
			rawMap, ok := rawVal.(map[string]interface{})
			if !ok {
				// if JSON has different shape, recommend type
				rec := recommendTypeFromJSON(rawVal)
				warnings = append(warnings, fmt.Sprintf("%s: JSON shape does not match struct; recommended type: %s", path, rec))
				return
			}

			// iterate JSON keys and ensure matching struct fields
			for key, v := range rawMap {
				fullPath := key
				if path != "" {
					fullPath = path + "." + key
				}
				if f, ok := findFieldByJSON(typ, key); ok {
					// if field is interface{}, recommend concrete type
					if f.Type.Kind() == reflect.Interface {
						rec := recommendTypeFromJSON(v)
						warnings = append(warnings, fmt.Sprintf("field %s declared as interface{}; consider a concrete type (e.g. %s)", fullPath, rec))
						continue
					}
					// recurse into field type
					inspectRawAgainstType(fullPath, v, f.Type)
				} else {
					// unmapped key
					rec := recommendTypeFromJSON(v)
					warnings = append(warnings, fmt.Sprintf("json key %q has no matching struct field at %s; recommended Go type: %s", key, pathOrRoot(path), rec))
				}
			}

		case reflect.Slice, reflect.Array:
			// rawVal should be a slice
			rawSlice, ok := rawVal.([]interface{})
			if !ok {
				rec := recommendTypeFromJSON(rawVal)
				warnings = append(warnings, fmt.Sprintf("%s: JSON shape is not array but target is slice/array; recommended: %s", path, rec))
				return
			}
			// inspect first element to infer element structure
			if len(rawSlice) == 0 {
				return
			}
			elemType := typ.Elem()
			for idx, item := range rawSlice {
				inspectRawAgainstType(fmt.Sprintf("%s[%d]", path, idx), item, elemType)
				// limit to first few to avoid performance issues? currently inspects all
				_ = idx
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			// expect a numeric JSON value
			switch rv := rawVal.(type) {
			case float64:
				// ok; but if fractional and target is int, warn
				if math.Mod(rv, 1.0) != 0 {
					rec := recommendTypeFromJSON(rv)
					warnings = append(warnings, fmt.Sprintf("%s: JSON is non-integer number but target is integer; recommended: %s", path, rec))
				}
			case string:
				rec := recommendTypeFromJSON(rv)
				warnings = append(warnings, fmt.Sprintf("%s: JSON is string but target is numeric; recommended: %s", path, rec))
			}

		case reflect.Float32, reflect.Float64:
			switch rv := rawVal.(type) {
			case string:
				rec := recommendTypeFromJSON(rv)
				warnings = append(warnings, fmt.Sprintf("%s: JSON is string but target is float; recommended: %s", path, rec))
			}

		case reflect.String:
			switch rv := rawVal.(type) {
			case float64, bool:
				rec := recommendTypeFromJSON(rv)
				warnings = append(warnings, fmt.Sprintf("%s: JSON is %T but target is string; recommended: %s", path, rv, rec))
			}

		case reflect.Bool:
			switch rv := rawVal.(type) {
			case string:
				rec := recommendTypeFromJSON(rv)
				warnings = append(warnings, fmt.Sprintf("%s: JSON is string but target is bool; recommended: %s", path, rec))
			}

		default:
			// other kinds: nothing special
		}
	}

	// start recursive inspection from the top-level
	inspectRawAgainstType("", raw, t)

	return warnings
}

// recommendTypeFromJSON returns a best-effort Go type recommendation as a string
// based on the JSON value. It is conservative and uses simple heuristics.
func recommendTypeFromJSON(v interface{}) string {
	if v == nil {
		return "nullable (e.g. *string, *int, json.RawMessage)"
	}
	switch t := v.(type) {
	case string:
		return "string"
	case float64:
		// prefer int when value has no fractional part
		if math.Trunc(t) == t {
			return "int"
		}
		return "float64"
	case bool:
		return "bool"
	case map[string]interface{}:
		return "struct or map[string]interface{}"
	case []interface{}:
		// infer element type from first element
		if len(t) == 0 {
			return "[]interface{}"
		}
		switch t0 := t[0].(type) {
		case string:
			return "[]string"
		case float64:
			// check if integer-like
			if math.Trunc(t0) == t0 {
				return "[]int"
			}
			return "[]float64"
		case map[string]interface{}:
			return "[]struct{} or []map[string]interface{}"
		default:
			return "[]interface{}"
		}
	default:
		_ = t
		return "interface{}"
	}
}

// index returns the index of sep in s or -1 if not found. A small helper to avoid
// importing strings for this trivial use.
func index(s string, sep byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			return i
		}
	}
	return -1
}

// pathOrRoot returns the path or the string "root" when path is empty.
func pathOrRoot(p string) string {
	if p == "" {
		return "root"
	}
	return p
}
