package render

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
)

// JSON writes the rows as a flat array of objects holding the given keys in that order. Each row is
// marshaled whole before its keys are picked, so the row's json tags still decide every value and
// every absence. json/v2 marshals a nil slice to [] and escapes no HTML, and MarshalWrite adds no
// trailing newline, so one is appended here.
func JSON[T any](w io.Writer, rows []T, keys []string) error {
	object := json.MarshalToFunc(func(enc *jsontext.Encoder, row T) error {
		raw, err := json.Marshal(row)
		if err != nil {
			return err
		}

		var members map[string]jsontext.Value
		if err := json.Unmarshal(raw, &members); err != nil {
			return err
		}

		if err := enc.WriteToken(jsontext.BeginObject); err != nil {
			return err
		}

		for _, k := range keys {
			v, ok := members[k]
			if !ok {
				continue
			}

			if err := enc.WriteToken(jsontext.String(k)); err != nil {
				return err
			}

			if err := enc.WriteValue(v); err != nil {
				return err
			}
		}

		return enc.WriteToken(jsontext.EndObject)
	})

	if err := json.MarshalWrite(w, rows, json.WithMarshalers(object)); err != nil {
		return fmt.Errorf("encoding rows as JSON: %w", err)
	}

	if _, err := io.WriteString(w, "\n"); err != nil {
		return fmt.Errorf("writing the JSON terminator: %w", err)
	}

	return nil
}
