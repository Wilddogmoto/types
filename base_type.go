package types

type (
	BaseTypes interface {
		IntegerTypes | FloatingTypes | SymbolsTypes | CharactersTypes | BinaryTypes
	}

	IntegerTypes interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	FloatingTypes interface {
		~float32 | ~float64
	}

	SymbolsTypes interface {
		~string
	}

	CharactersTypes interface {
		~rune
	}

	BinaryTypes interface {
		~byte
	}
)
