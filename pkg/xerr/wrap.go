package xerr

import "github.com/teamgram/teamgram-server/v2/pkg/xerr/stack"

const stackSkip = 4

func Unwrap(err error) error {
	for err != nil {
		unwrap, ok := err.(interface {
			error
			Unwrap() error
		})
		if !ok {
			break
		}
		err = unwrap.Unwrap()
		if err == nil {
			return unwrap
		}
	}
	return err
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return stack.New(err, stackSkip)
}

func WrapMsg(err error, msg string, kv ...any) error {
	if err == nil {
		return nil
	}
	return stack.New(NewErrorWrapper(err, toString(msg, kv)), stackSkip)
}
