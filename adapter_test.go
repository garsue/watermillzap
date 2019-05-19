package watermillzap

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {
	zl := zap.NewExample()
	got := NewLogger(zl)
	want := &Logger{
		backend: zl,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewLogger() = %v, want %v", got, want)
	}
}

func ExampleLogger_Error() {
	logger := NewLogger(zap.NewExample())
	logger.Error("ExampleLogger_Error", errors.New("test error"), watermill.LogFields{"test_field": 1})
	//output:
	//{"level":"error","msg":"ExampleLogger_Error","error":"test error","test_field":1}
}

func ExampleLogger_Info() {
	logger := NewLogger(zap.NewExample())
	logger.Info("ExampleLogger_Info", watermill.LogFields{"test_field": 1})
	//output:
	//{"level":"info","msg":"ExampleLogger_Info","test_field":1}
}

func ExampleLogger_Debug() {
	logger := NewLogger(zap.NewExample())
	logger.Debug("ExampleLogger_Debug", watermill.LogFields{"test_field": 1})
	//output:
	//{"level":"debug","msg":"ExampleLogger_Debug","test_field":1}
}

func ExampleLogger_Trace() {
	logger := NewLogger(zap.NewExample())
	logger.Trace("ExampleLogger_Trace", watermill.LogFields{"test_field": 1})
	//output:
	//{"level":"debug","msg":"ExampleLogger_Trace","test_field":1}
}

func TestLogger_With(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		backend := zap.NewExample()
		l := &Logger{
			backend: backend,
			fields:  watermill.LogFields{"default_field": true},
		}
		want := &Logger{
			backend: backend,
			fields:  watermill.LogFields{"default_field": true},
		}
		if got := l.With(nil); !reflect.DeepEqual(got, want) {
			t.Errorf("Logger.With() = %v, want %v", got, want)
		}
	})
	t.Run("add", func(t *testing.T) {
		backend := zap.NewExample()
		l := &Logger{
			backend: backend,
			fields:  watermill.LogFields{"default_field": true},
		}
		want := &Logger{
			backend: backend,
			fields:  watermill.LogFields{"default_field": true, "another_field": false},
		}
		if got := l.With(watermill.LogFields{"another_field": false}); !reflect.DeepEqual(got, want) {
			t.Errorf("Logger.With() = %v, want %v", got, want)
		}
	})
}
