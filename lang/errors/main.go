package main

import (
	"errors"
	"fmt"
	"os"
)

func f3() error {
	err := DeepError{
		errorCode:    69,
		wrappedError: &os.SyscallError{Syscall: "open", Err: errors.New("syscall open does not exist")},
	}

	return fmt.Errorf("f3 failed: %w", err)
}

type DeepError struct {
	errorCode    int
	wrappedError error
}

func (e DeepError) Error() string {
	return fmt.Sprintf("very deeeeeep error with errorCode %d", e.errorCode)
}

func (e DeepError) Unwrap() error {
	return e.wrappedError
}

func f2() error {
	if err := f3(); err != nil {
		return fmt.Errorf("f2 failed: %w", err)
	}
	return nil
}

func f1() error {
	if err := f2(); err != nil {
		return fmt.Errorf("f1 failed: %w", err)
	}
	return nil
}

func main() {
	syscallError := &os.SyscallError{}
	deepError := DeepError{}

	fmt.Println("deep error before:", deepError)
	err := f1()
	if errors.As(err, &syscallError) {
		fmt.Printf("failed to call f1(): error of type %T: %v\n", syscallError, err)
	}
	fmt.Println("deep error after:", deepError)
}
