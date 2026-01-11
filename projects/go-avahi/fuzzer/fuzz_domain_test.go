package fuzzer

import (
	"testing"
	"github.com/OpenPrinting/go-avahi"
)

// FuzzDomainNormalize tests the Go -> C -> Go round trip.
func FuzzDomainNormalize(f *testing.F) {
	// 1. Seed Corpus
	f.Add("example.com")
	f.Add("My Printer._ipp._tcp.local")
	f.Add("foo.bar.")
	
	// 2. The Fuzz Target
	f.Fuzz(func(t *testing.T, data string) {
		_ = avahi.DomainNormalize(data)
	})
}

// FuzzDomainSplit tests parsing of complex service strings.
func FuzzDomainSplit(f *testing.F) {
	f.Add("Kyocera ECOSYS M2040dn._ipp._tcp.local")
	f.Add("InvalidString")
	
	f.Fuzz(func(t *testing.T, data string) {
		instance, svctype, domain := avahi.DomainServiceNameSplit(data)
		if instance != "" && svctype != "" && domain != "" {
			return
		}
	})
}
