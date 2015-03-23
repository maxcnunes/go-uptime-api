package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMonitorApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MonitorApi Suite")
}
