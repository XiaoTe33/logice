package main

import "gorm_demo1/src/lg3"

func main() {
	lg3.SetFiles(lg3.Paths{
		All:   "./src/lg3/test/all.log",
		Debug: "./src/lg3/test/debug.log",
		Info:  "./src/lg3/test/info.log",
		Warn:  "./src/lg3/test/warn.log",
		Error: "./src/lg3/test/error.log",
	})
	lg3.Info("info")
	lg3.Debug("debug")
	lg3.Warn("warn")
	lg3.Error("err")
}
