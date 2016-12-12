package php_serialize

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEncodeNil(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = nil
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding nil value: %v\n", err)
	} else {
		if val != "N;" {
			t.Errorf("Nil value decoded incorrectly, have got %q\n", val)
		}
	}
}

func BenchmarkEncodeNil(b *testing.B) {
	var source PhpValue = nil
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeBoolTrue(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = true
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding bool value: %v\n", err)
	} else {
		if val != "b:1;" {
			t.Errorf("Bool value decoded incorrectly, have got %q\n", val)
		}
	}
}

func BenchmarkEncodeBool(b *testing.B) {
	var source PhpValue = true
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeBoolFalse(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = false
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding bool value: %v\n", err)
	} else {
		if val != "b:0;" {
			t.Errorf("Bool value decoded incorrectly, have got %q\n", val)
		}
	}
}

func TestEncodeInt(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = 42
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding int value: %v\n", err)
	} else {
		if val != "i:42;" {
			t.Errorf("Int value decoded incorrectly, have got %q\n", val)
		}
	}
}

func BenchmarkEncodeInt(b *testing.B) {
	var source PhpValue = 42
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeIntMinus(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = -42
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding int value: %v\n", err)
	} else {
		if val != "i:-42;" {
			t.Errorf("Int value decoded incorrectly, have got %q\n", val)
		}
	}
}

func BenchmarkEncodeIntMinus(b *testing.B) {
	var source PhpValue = -42
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeFloat64(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = 42.378900000000002
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding float value: %v\n", err)
	} else {
		if val != "d:42.378900000000002;" {
			t.Errorf("Float value decoded incorrectly, have got %q\n", val)
		}
	}
}

func BenchmarkEncodeFloat(b *testing.B) {
	var source PhpValue = 42.378900000000002
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeFloat64Minus(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = -42.378900000000002
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding float value: %v\n", err)
	} else {
		if val != "d:-42.378900000000002;" {
			t.Errorf("Float value decoded incorrectly, have got %q\n", val)
		}
	}
}

func BenchmarkEncodeFloatMinus(b *testing.B) {
	var source PhpValue = -42.378900000000002
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeString(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = "foobar"
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding string value: %v\n", err)
	} else {
		if val != "s:6:\"foobar\";" {
			t.Errorf("String value decoded incorrectly, have got %q\n", val)
		}
	}
}

func BenchmarkEncodeString(b *testing.B) {
	var source PhpValue = "foobar"
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeArray(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = PhpArray{
		0: 10,
		1: 11,
		2: 12,
	}
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "i:0;i:10;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:0;i:10;", val)
		} else if !strings.Contains(val, "i:1;i:11;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:1;i:11;", val)
		} else if !strings.Contains(val, "i:2;i:12;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:2;i:12;", val)
		}
	}
}

func BenchmarkEncodeArray(b *testing.B) {
	source := PhpArray{
		0: 10,
		1: 11,
		2: 12,
	}
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeArray2(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = map[PhpValue]PhpValue{
		0: 10,
		1: 11,
		2: 12,
	}
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "i:0;i:10;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:0;i:10;", val)
		} else if !strings.Contains(val, "i:1;i:11;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:1;i:11;", val)
		} else if !strings.Contains(val, "i:2;i:12;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:2;i:12;", val)
		}
	}
}

func BenchmarkEncodeArray2(b *testing.B) {
	source := map[PhpValue]PhpValue{
		0: 10,
		1: 11,
		2: 12,
	}
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeArrayMap(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = PhpArray{
		"foo": 4,
		"bar": 2,
	}
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "s:3:\"foo\";i:4;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:3:\"foo\";i:4;", val)
		} else if !strings.Contains(val, "s:3:\"bar\";i:2;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:3:\"bar\";i:2;", val)
		}
	}
}

func BenchmarkEncodeArrayMap(b *testing.B) {
	source := PhpArray{
		"foo": 4,
		"bar": 2,
	}
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeArrayArray(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	source = PhpArray{
		"foo": PhpArray{
			0: 10,
		},
		"bar": 2,
	}
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "s:3:\"foo\";a:1:{i:0;i:10;}") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:3:\"foo\";a:1:{i:0;i:10;}", val)
		} else if !strings.Contains(val, "s:3:\"bar\";i:2;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:3:\"bar\";i:2;", val)
		}
	}
}

func BenchmarkEncodeArrayArray(b *testing.B) {
	source := PhpArray{
		"foo": PhpArray{
			0: 10,
		},
		"bar": 2,
	}
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeObject(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	obj := NewPhpObject("Test")
	obj.SetPublic("public", 1)
	obj.SetProtected("protected", 2)
	obj.SetPrivate("private", 3)

	source = obj
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "O:4:\"Test\"") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "O:4:\"Test\"", val)
		} else if !strings.Contains(val, "s:6:\"public\";i:1;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:6:\"public\";i:1;", val)
		} else if !strings.Contains(val, "s:12:\"\x00*\x00protected\";i:2;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:12:\"\x00*\x00protected\";i:2;", val)
		} else if !strings.Contains(val, "s:13:\"\x00Test\x00private\";i:3;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:13:\"\x00Test\x00private\";i:3;", val)
		}
	}
}

func BenchmarkEncodeObject(b *testing.B) {
	b.StopTimer()
	source := NewPhpObject("Test")
	source.SetPublic("public", 1)
	source.SetProtected("protected", 2)
	source.SetPrivate("private", 3)
	b.ReportAllocs()
	b.StartTimer()
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeArrayOfObjects(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	obj1 := NewPhpObject("Test1")
	obj1.SetPublic("public", 11)
	obj1.SetProtected("protected", 12)
	obj1.SetPrivate("private", 13)

	obj2 := NewPhpObject("Test2")
	obj2.SetPublic("public", 21)
	obj2.SetProtected("protected", 22)
	obj2.SetPrivate("private", 23)

	source = PhpArray{
		0: obj1,
		1: obj2,
	}
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "a:2:") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "a:2:", val)
		} else if !strings.Contains(val, "i:0;O:5:\"Test1\"") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:0;O:5:\"Test1\"", val)
		} else if !strings.Contains(val, "s:6:\"public\";i:11;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:6:\"public\";i:11;", val)
		} else if !strings.Contains(val, "s:12:\"\x00*\x00protected\";i:12;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:12:\"\x00*\x00protected\";i:12;", val)
		} else if !strings.Contains(val, "s:14:\"\x00Test1\x00private\";i:13;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:14:\"\x00Test1\x00private\";i:13;", val)
		} else if !strings.Contains(val, "i:1;O:5:\"Test2\"") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "i:1;O:5:\"Test2\"", val)
		} else if !strings.Contains(val, "s:6:\"public\";i:21;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:6:\"public\";i:21;", val)
		} else if !strings.Contains(val, "s:12:\"\x00*\x00protected\";i:22;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:12:\"\x00*\x00protected\";i:22;", val)
		} else if !strings.Contains(val, "s:14:\"\x00Test2\x00private\";i:23;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:14:\"\x00Test2\x00private\";i:23;", val)
		}
	}
}

func BenchmarkEncodeArrayOfObjects(b *testing.B) {
	b.StopTimer()
	obj1 := NewPhpObject("Test1")
	obj1.SetPublic("public", 11)
	obj1.SetProtected("protected", 12)
	obj1.SetPrivate("private", 13)

	obj2 := NewPhpObject("Test2")
	obj2.SetPublic("public", 21)
	obj2.SetProtected("protected", 22)
	obj2.SetPrivate("private", 23)

	source := PhpArray{
		0: obj1,
		1: obj2,
	}
	b.ReportAllocs()
	b.StartTimer()
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}

func TestEncodeObjectSerializable(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	obj := NewPhpObjectSerialized("TestSerializable")
	obj.SetData("foobar")

	source = obj
	encoder := NewSerializer()
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "C:16:\"TestSerializable\":6:{foobar}") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "C:16:\"TestSerializable\":6:{foobar}", val)
		}
	}
}

func TestEncodeObjectSerializableArray(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
	)

	obj := NewPhpObjectSerialized("TestSerializable1")
	obj.SetValue(PhpArray{
		"foo": 4,
		"bar": 2,
	})

	source = obj
	encoder := NewSerializer()
	encoder.SetSerializedEncodeFunc(SerializedEncodeFunc(Serialize))
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "C:17:\"TestSerializable1\"") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "C:17:\"TestSerializable1\"", val)
		} else if !strings.Contains(val, "s:3:\"foo\";i:4;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:3:\"foo\";i:4;", val)
		} else if !strings.Contains(val, "s:3:\"bar\";i:2;") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "s:3:\"bar\";i:2;", val)
		}
	}
}

func TestEncodeObjectSerializableJSON(t *testing.T) {
	var (
		source PhpValue
		val    string
		err    error
		f      SerializedEncodeFunc
	)

	f = func(v PhpValue) (string, error) {
		var (
			res []byte
			err error
		)
		res, err = json.Marshal(v)
		return string(res), err
	}

	obj := NewPhpObjectSerialized("TestSerializable2")
	obj.SetValue(map[string]int{
		"foo": 4,
		"bar": 2,
	})

	source = obj
	encoder := NewSerializer()
	encoder.SetSerializedEncodeFunc(f)
	if val, err = encoder.Encode(source); err != nil {
		t.Errorf("Error while encoding array value: %v\n", err)
	} else {
		if !strings.Contains(val, "C:17:\"TestSerializable2\"") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "C:17:\"TestSerializable2\"", val)
		} else if !strings.Contains(val, "\"foo\":4") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "\"foo\":4", val)
		} else if !strings.Contains(val, "\"bar\":2") {
			t.Errorf("Array value decoded incorrectly, expected substring %q but have got %q\n", "\"bar\":2", val)
		}
	}
}

func TestEncodeSplArray(t *testing.T) {
	obj := NewPhpSplArray(PhpArray{"foo": 42}, nil)

	data, err := Serialize(obj)
	if err != nil {
		t.Errorf("Error while encoding array object: %v\n", err)
	}

	expected := "x:i:0;a:1:{s:3:\"foo\";i:42;};m:a:0:{}"

	if data != expected {
		t.Errorf("SplArray decoded incorrectly, expected: %q, got: %q\n", expected, data)
	}
}

func BenchmarkEncodeSplArray(b *testing.B) {
	source := NewPhpSplArray(PhpArray{"foo": 42}, nil)
	encoder := NewSerializer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(source)
	}
}