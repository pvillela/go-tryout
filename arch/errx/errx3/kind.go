package errx3

import (
	"fmt"
	"runtime/debug"
)

type Kind struct {
	msg              string
	directSuperKinds []*Kind
}

func KindOf(err error) *Kind {
	errx, ok := err.(*errxImpl)
	if !ok {
		return nil
	}
	return errx.kind
}

func NewKind(msg string, superKinds ...*Kind) *Kind {
	return &Kind{msg, superKinds}
}

func (s *Kind) makeInternal(
	cause error,
	stackLinesToSuppress int,
	args ...any,
) *errxImpl {
	err := errxImpl{kind: s}
	err.args = args
	err.cause = cause
	err.stack = debug.Stack()
	err.stackLinesToSuppress = stackLinesToSuppress
	return &err
}

func (s *Kind) Make(cause error, args ...any) Errx {
	err := s.makeInternal(cause, 3, args...)
	return err
}

func (s *Kind) Decorate(cause Errx, args ...interface{}) Errx {
	err := errxImpl{kind: s}
	err.args = args
	err.cause = cause
	return &err
}

func newErrxInternal(cause error, msg string, stackLinesToSuppress int) Errx {
	kind := NewKind(msg)
	err := kind.makeInternal(cause, stackLinesToSuppress)
	return err
}

func NewErrx(cause error, msg string) Errx {
	return newErrxInternal(cause, msg, 4)
}

func ErrxOf(r any) Errx {
	if r == nil {
		return nil
	}
	var err error
	switch r.(type) {
	case error:
		err = r.(error)
	default:
		err = nil
	}
	errX, ok := err.(Errx)
	if !ok {
		if err != nil {
			errX = newErrxInternal(err, ".", 4)
		} else {
			errX = newErrxInternal(nil, fmt.Sprintf("%v", r), 4)
		}
	}
	return errX
}

func put(m map[*Kind]struct{}, kind *Kind) {
	m[kind] = struct{}{}
}

// SuperKinds returns the set of all super-kinds of the receiver, including the receiver itself.
func (s *Kind) SuperKinds() map[*Kind]struct{} {
	result := make(map[*Kind]struct{}, len(s.directSuperKinds))
	seen := make(map[*Kind]struct{}, len(s.directSuperKinds))
	workQueue := make([]*Kind, 1, len(s.directSuperKinds)+1)
	workQueue[0] = s
	for len(workQueue) > 0 {
		kind := workQueue[0]
		if _, ok := seen[kind]; !ok {
			put(seen, kind)
			put(result, kind)
			workQueue = append(workQueue[1:], kind.directSuperKinds...)
		} else {
			workQueue = workQueue[1:]
		}
	}
	return result
}

func (s *Kind) IsSubKindOf(kind *Kind) bool {
	superKinds := s.SuperKinds()
	_, ok := superKinds[kind]
	return ok
}
