# Go Package JSON

This package is very similar to the built-in go package called json,
implementing encoding and decoding of JSON as defined in RFC 7159. The mapping
between JSON and Go values is described in the documentation for the Marshal
and Unmarshal functions.

## Differences

The differences in this implementation are some additional functions and
interfaces.

In the Decoder type, one will find two additional functions:

- UseAutoConvert() - When enabled, this will attempt to convert any strings
  into defined element types, such as Integer, Boolean, or the CustomType
  interface
- UseSlice() - When enabled, this will automatically convert an object into a
  slice if a slice is specified in the type declaration.  When vendors provide
  JSON output, and this output is broken in that it can vary when one or more
  elements are returned, this decoder will try overcome this by creating a slice
  when specified and if one object is provided it will be a slice of one with
  that object.

Secondly, one will fine a new custom type interface.  This is useful for defining your own data type
which is encoded and decoded as a string.

```
// This interface allows custom types to be loaded via strings in JSON.  When
// this interface is implemented, the FromString must decode into the same type
// specified in the struct, likewise the FromString must return a string used
// by Marshal to create string literals.
type CustomType interface {
  FromString(string) (interface{}, error)
  ToString() string
}
```

An example of how this could be useful is encoding and decoding custom time
formats into a custom time.Duration type.  Take for example a JSON API provides
a time as "1D3H2M0S".  With the use of a custom ToString / FromString this can
be encoded / decoded when the JSON is being read in or written out.
