package calculator

import (
	"context"
	"math/rand"

	"github.com/ServiceWeaver/weaver"
)

type MathError struct {
	X       float64
	Y       float64
	Message string
}

func (e *MathError) Error() string {
	return e.Message
}

type Adder interface {
	Add(ctx context.Context, x float64, y float64) (float64, error)
}

type Subtracter interface {
	Subtract(ctx context.Context, x float64, y float64) (float64, error)
}

type Multiplier interface {
	Multiply(ctx context.Context, x float64, y float64) (float64, error)
}

type Divider interface {
	Divide(ctx context.Context, x float64, y float64) (float64, error)
}

type adder struct {
	weaver.Implements[Adder]
}

func (c *adder) Add(ctx context.Context, x float64, y float64) (float64, error) {
	c.Logger().Info("adding %v + %v", x, y)
	return x + y, nil
}

type subtracter struct {
	weaver.Implements[Subtracter]
}

func (c *subtracter) Subtract(ctx context.Context, x float64, y float64) (float64, error) {
	c.Logger().Info("subtracting %v - %v", x, y)
	return x - y, nil
}

type multiplierConfig struct {
	FailureRate float64
}

type multiplier struct {
	weaver.Implements[Multiplier]
	weaver.WithConfig[multiplierConfig]
}

func (c *multiplier) Init(ctx context.Context) {
	c.Logger().Info("initializing multiplier with config %+v", c.Config())
}

func (c *multiplier) Multiply(ctx context.Context, x float64, y float64) (float64, error) {
	c.Logger().Info("multiplying %v * %v", x, y)

	if rand.Float64()*c.Config().FailureRate > 0.5 {
		return 0, &MathError{X: x, Y: y, Message: "failed to multiply. Sorry :-/"}
	}

	return x * y, nil
}

type divider struct {
	weaver.Implements[Divider]
}

func (c *divider) Divide(ctx context.Context, x float64, y float64) (float64, error) {
	c.Logger().Info("dividing %v / %v", x, y)
	if y == 0 {
		return 0, &MathError{X: x, Y: y, Message: "cannot divide by zero"}
	}

	return x / y, nil
}
