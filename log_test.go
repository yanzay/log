package log

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

type MockWriter struct {
	lastLog []byte
}

func (mw *MockWriter) Write(v []byte) (int, error) {
	mw.lastLog = v
	return 0, nil
}

func (mw *MockWriter) GetLastLog() string {
	return string(mw.lastLog)
}

func TestPrintln(t *testing.T) {
	c.Convey("Given a MockWriter", t, func() {
		writer := &MockWriter{}
		Writer = writer
		Println("test")
		c.So(writer.GetLastLog(), c.ShouldEqual, "test")
	})
}

func TestTrace(t *testing.T) {
	c.Convey("Given a MockWriter", t, func() {
		writer := &MockWriter{}
		Writer = writer
		c.Convey("When Level is LevelTrace", func() {
			Level = LevelTrace
			Trace("test")
			c.So(writer.GetLastLog(), c.ShouldEqual, "test")
		})
		c.Convey("When Level is LevelDebug", func() {
			Level = LevelDebug
			Trace("badtest")
			c.So(writer.GetLastLog(), c.ShouldNotEqual, "badtest")
		})
	})
}

func TestDebug(t *testing.T) {
	c.Convey("Given a MockWriter", t, func() {
		writer := &MockWriter{}
		Writer = writer
		c.Convey("When Level is LevelDebug", func() {
			Level = LevelDebug
			Debug("test")
			c.So(writer.GetLastLog(), c.ShouldEqual, "test")
		})
		c.Convey("When Level is LevelInfo", func() {
			Level = LevelInfo
			Debug("badtest")
			c.So(writer.GetLastLog(), c.ShouldNotEqual, "badtest")
		})
	})
}

func TestInfo(t *testing.T) {
	c.Convey("Given a MockWriter", t, func() {
		writer := &MockWriter{}
		Writer = writer
		c.Convey("When Level is LevelInfo", func() {
			Level = LevelInfo
			Info("test")
			c.So(writer.GetLastLog(), c.ShouldEqual, "test")
		})
		c.Convey("Without Level setting", func() {
			Level = 0
			Info("test")
			c.So(writer.GetLastLog(), c.ShouldEqual, "test")
		})
		c.Convey("When Level is LevelWarning", func() {
			Level = LevelWarning
			Info("badtest")
			c.So(writer.GetLastLog(), c.ShouldNotEqual, "badtest")
		})
	})
}

func TestWarning(t *testing.T) {
	c.Convey("Given a MockWriter", t, func() {
		writer := &MockWriter{}
		Writer = writer
		c.Convey("When Level is LevelWarning", func() {
			Level = LevelWarning
			Warning("test")
			c.So(writer.GetLastLog(), c.ShouldEqual, "test")
		})
		c.Convey("When Level is LevelError", func() {
			Level = LevelError
			Warning("badtest")
			c.So(writer.GetLastLog(), c.ShouldNotEqual, "badtest")
		})
	})
}

func TestError(t *testing.T) {
	c.Convey("Given a MockWriter", t, func() {
		writer := &MockWriter{}
		Writer = writer
		c.Convey("When Level is LevelError", func() {
			Level = LevelError
			Error("test")
			c.So(writer.GetLastLog(), c.ShouldEqual, "test")
		})
		c.Convey("When Level is LevelCritical", func() {
			Level = LevelCritical
			Error("badtest")
			c.So(writer.GetLastLog(), c.ShouldNotEqual, "badtest")
		})
	})
}

func TestCritical(t *testing.T) {
	c.Convey("Given a MockWriter", t, func() {
		writer := &MockWriter{}
		Writer = writer
		c.Convey("When Level is LevelCritical", func() {
			Level = LevelCritical
			c.So(func() { Fatal("fatal") }, c.ShouldPanic)
			c.So(writer.GetLastLog(), c.ShouldEqual, "fatal")
		})
	})
}

func TestFlags(t *testing.T) {
	c.Convey("Flags should be parsed and translated to log levels", t, func() {
		Level = 0
		c.Convey("When flag is 'trace'", func() {
			f := "trace"
			logFlag = &f
			Info("test")
			c.So(Level, c.ShouldEqual, LevelTrace)
		})
		c.Convey("When flag is 'debug'", func() {
			f := "debug"
			logFlag = &f
			Debug("test")
			c.So(Level, c.ShouldEqual, LevelDebug)
		})
		c.Convey("When flag is 'info'", func() {
			f := "info"
			logFlag = &f
			Debug("test")
			c.So(Level, c.ShouldEqual, LevelInfo)
		})
		c.Convey("When flag is 'warning'", func() {
			f := "warning"
			logFlag = &f
			Debug("test")
			c.So(Level, c.ShouldEqual, LevelWarning)
		})
		c.Convey("When flag is 'error'", func() {
			f := "error"
			logFlag = &f
			Debug("test")
			c.So(Level, c.ShouldEqual, LevelError)
		})
		c.Convey("When flag is 'critical'", func() {
			f := "critical"
			logFlag = &f
			Debug("test")
			c.So(Level, c.ShouldEqual, LevelCritical)
		})
		c.Convey("When flag is 'incorrect'", func() {
			f := "incorrect"
			logFlag = &f
			Debug("test")
			c.So(Level, c.ShouldEqual, LevelInfo)
		})
	})
}

func TestFormat(t *testing.T) {
	c.Convey("Given format funcs", t, func() {
		funcs := map[string]func(string, ...interface{}){
			"Printf":   Printf,
			"Tracef":   Tracef,
			"Debugf":   Debugf,
			"Infof":    Infof,
			"Warningf": Warningf,
			"Errorf":   Errorf,
		}
		writer := &MockWriter{}
		Writer = writer
		Level = LevelTrace
		for name, fun := range funcs {
			c.Convey(name+" format should work", func() {
				fun("%s answer: %d", name, 42)
				c.So(writer.GetLastLog(), c.ShouldEqual, name+" answer: 42")
			})
		}
		c.Convey("Fatalf format should work", func() {
			c.So(func() { Fatalf("answer: %d", 42) }, c.ShouldPanic)
			c.So(writer.GetLastLog(), c.ShouldEqual, "answer: 42")
		})

	})
}
