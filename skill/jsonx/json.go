package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"github.com/linmingxiao/gneo/internal/bytesconv"
)

func Str2KV(str string) (map[string]interface{}, error) {
	kv := make(map[string]interface{})
	bytes := bytesconv.StringToBytes(str)
	err := json.Unmarshal(bytes, &kv)
	if err == nil{
		return kv, nil
	} else {
		return nil, err
	}
}

func KV2Str(kv map[string]interface{}) (string, error)  {
	if bytes, err := KV2Bytes(kv); err != nil {
		return "", err
	} else {
		str := bytesconv.BytesToString(bytes)
		return str, nil
	}
}

func KV2Bytes(jsonObj interface{}) ([]byte, error) {
	if bytes, err := json.Marshal(jsonObj); err == nil{
		return bytes, nil
	} else {
		return nil, err
	}
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(string(data), err)
	}

	return nil
}

func UnmarshalFromString(str string, v interface{}) error {
	decoder := json.NewDecoder(strings.NewReader(str))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(str, err)
	}

	return nil
}

func UnmarshalFromReader(reader io.Reader, v interface{}) error {
	var buf strings.Builder
	teeReader := io.TeeReader(reader, &buf)
	decoder := json.NewDecoder(teeReader)
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(buf.String(), err)
	}

	return nil
}

func unmarshalUseNumber(decoder *json.Decoder, v interface{}) error {
	decoder.UseNumber()
	return decoder.Decode(v)
}

func formatError(v string, err error) error {
	return fmt.Errorf("string: `%s`, error: `%s`", v, err.Error())
}
