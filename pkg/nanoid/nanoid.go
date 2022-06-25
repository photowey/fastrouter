package nanoid

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	single      = 1
	DefaultSize = 21
)

func New(sizes ...int) (string, error) {
	size := determineSize(sizes...)
	id, err := gonanoid.New(size)

	return id, err
}

func MustNew(sizes ...int) string {
	id, err := New(sizes...)
	if err != nil {
		size := determineSize(sizes...)
		return Random(size)
	}

	return id
}

func Generate(alphabet string, sizes ...int) (string, error) {
	size := determineSize(sizes...)
	id, err := gonanoid.Generate(alphabet, size)

	return id, err
}

func MustGenerate(alphabet string, sizes ...int) string {
	id, err := Generate(alphabet, sizes...)
	if err != nil {
		size := determineSize(sizes...)
		return RandomAlphabet(alphabet, size)
	}

	return id
}

func determineSize(sizes ...int) int {
	size := DefaultSize
	switch len(sizes) {
	case single:
		size = sizes[0]
	}

	return size
}
