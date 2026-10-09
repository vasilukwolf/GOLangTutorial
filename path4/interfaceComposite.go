package main

import "fmt"

type Authorizer interface {
	Authorize() string
}

type Captyrer interface {
	Capture() string
}

type StripeProvider struct {
	Authorizer
	Captyrer
}

func (s *StripeProvider) String() string {
	return fmt.Sprintf("StripeProvider{Authorizer: %v, Captyrer: %v}",
		s.Authorizer, s.Captyrer)
}

func main() {
	p := &StripeProvider{}
	fmt.Println(p)
}
