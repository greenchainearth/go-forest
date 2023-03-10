package eventcheck

import (
	"github.com/greenchainearth/go-forest/eventcheck/basiccheck"
	"github.com/greenchainearth/go-forest/eventcheck/epochcheck"
	"github.com/greenchainearth/go-forest/eventcheck/gaspowercheck"
	"github.com/greenchainearth/go-forest/eventcheck/heavycheck"
	"github.com/greenchainearth/go-forest/eventcheck/parentscheck"
	"github.com/greenchainearth/go-forest/inter"
)

// Checkers is collection of all the checkers
type Checkers struct {
	Basiccheck    *basiccheck.Checker
	Epochcheck    *epochcheck.Checker
	Parentscheck  *parentscheck.Checker
	Gaspowercheck *gaspowercheck.Checker
	Heavycheck    *heavycheck.Checker
}

// Validate runs all the checks except Poset-related
func (v *Checkers) Validate(e inter.EventPayloadI, parents inter.EventIs) error {
	if err := v.Basiccheck.Validate(e); err != nil {
		return err
	}
	if err := v.Epochcheck.Validate(e); err != nil {
		return err
	}
	if err := v.Parentscheck.Validate(e, parents); err != nil {
		return err
	}
	var selfParent inter.EventI
	if e.SelfParent() != nil {
		selfParent = parents[0]
	}
	if err := v.Gaspowercheck.Validate(e, selfParent); err != nil {
		return err
	}
	if err := v.Heavycheck.Validate(e); err != nil {
		return err
	}
	return nil
}
