package types

type (
	BaseTypes interface {
		IntegerTypes | FloatingTypes | SymbolsTypes | PointerTypes
	}

	IntegerTypes interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
	}

	PointerTypes interface {
		~uintptr
	}

	FloatingTypes interface {
		~float32 | ~float64
	}

	SymbolsTypes interface {
		~string
	}
)
