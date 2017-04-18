package log

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestSyslogWrite(t *testing.T) {
	c.Convey("Given SyslogWriter", t, func() {
		var err error
		Writer, err = NewSyslogWriter(1, "test")
		c.So(err, c.ShouldBeNil)
		c.Convey("Write should go to stdout", func() {
			count, err := Writer.Write([]byte("test"))
			c.So(err, c.ShouldBeNil)
			c.So(count, c.ShouldEqual, 4)
		})
	})
}

func TestSyslogBad(t *testing.T) {
	c.Convey("Given SyslogWriter", t, func() {
		var err error
		Writer, err = NewSyslogWriter(100000, "incorrect")
		c.So(err, c.ShouldNotBeNil)
	})
}
