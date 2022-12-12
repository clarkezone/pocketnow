/*
Copyright © 2022 clarkezone
*/
package main

import (
	"github.com/clarkezone/geocache/cmd"
	clarkezoneLog "github.com/clarkezone/geocache/pkg/log"
	"github.com/sirupsen/logrus"
)

func main() {
	clarkezoneLog.Init(logrus.WarnLevel)
	cmd.Execute()
}
