package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		assertEqual(t, 1, 1)
		assertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		assertEqual(t, "hello", "hello")
		assertNotEqual(t, "hello", "Grace")
	})
	//assertEqual(t, 1, "1")
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		//stackOfInts := new(StackOfInts)
		stackOfInts := NewStack[int]()
		// assert stack is empty
		assertTrue(t, stackOfInts.IsEmpty())
		// add something, assert stack is not empty
		stackOfInts.Push(123)
		assertFalse(t, stackOfInts.IsEmpty())
		// add another, pop it
		stackOfInts.Push(456)
		value, _ := stackOfInts.Pop()
		assertEqual(t, value, 456)
		value, _ = stackOfInts.Pop()
		assertEqual(t, value, 123)
		assertTrue(t, stackOfInts.IsEmpty())

		// can get the numbers we put in numbers, not untyped interface
		stackOfInts.Push(1)
		stackOfInts.Push(2)
		fNum, _ := stackOfInts.Pop()
		sNum, _ := stackOfInts.Pop()
		assertEqual(t, fNum+sNum, 3)
	})
	t.Run("string stack", func(t *testing.T) {
		//stackOfStrings := new(StackOfStrings)
		stackOfStrings := NewStack[string]()
		// assert stack is empty
		assertTrue(t, stackOfStrings.IsEmpty())
		// add something, assert stack is not empty
		stackOfStrings.Push("123")
		assertFalse(t, stackOfStrings.IsEmpty())
		// add another, pop it
		stackOfStrings.Push("456")
		value, _ := stackOfStrings.Pop()
		assertEqual(t, value, "456")
		value, _ = stackOfStrings.Pop()
		assertEqual(t, value, "123")
		assertTrue(t, stackOfStrings.IsEmpty())
	})
}

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func assertNotEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected == actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func assertTrue(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Errorf("expected %v, got false", actual)
	}
}

func assertFalse(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		t.Errorf("expected %v, got true", actual)
	}
}
