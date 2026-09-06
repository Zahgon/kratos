package encoding

type Codec interface {
	Marshal(v any) ([]byte, error)

	Unmarshal(data []byte, v any) error

	Name() string
}

var registeredCodecs = make(map[string]Codec)

func RegisterCodec(codec Codec) { _ = "STUB: not implemented"; return }

func GetCodec(contentSubtype string) Codec { _ = "STUB: not implemented"; return *new(Codec) }
