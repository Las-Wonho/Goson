package goson

// JSON is goson type
type JSON struct{}

// New create JSON instance
func New() JSON {
	return JSON{}
}

// SetElement set eliment to JSON instance
func (json JSON) SetElement(name string, value string) JSON {
	return json
}

// GetEliment transformation JSON instance to golang struct
func (json JSON) GetEliment(name string) string {
	return "value1"
}
