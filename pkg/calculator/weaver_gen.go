package calculator

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:        "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Adder",
		Iface:       reflect.TypeOf((*Adder)(nil)).Elem(),
		New:         func() any { return &adder{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return adder_local_stub{impl: impl.(Adder), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return adder_client_stub{stub: stub, addMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Adder", Method: "Add"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return adder_server_stub{impl: impl.(Adder), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Divider",
		Iface: reflect.TypeOf((*Divider)(nil)).Elem(),
		New:   func() any { return &divider{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return divider_local_stub{impl: impl.(Divider), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return divider_client_stub{stub: stub, divideMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Divider", Method: "Divide"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return divider_server_stub{impl: impl.(Divider), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:     "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Multiplier",
		Iface:    reflect.TypeOf((*Multiplier)(nil)).Elem(),
		New:      func() any { return &multiplier{} },
		ConfigFn: func(i any) any { return i.(*multiplier).WithConfig.Config() },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return multiplier_local_stub{impl: impl.(Multiplier), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return multiplier_client_stub{stub: stub, multiplyMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Multiplier", Method: "Multiply"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return multiplier_server_stub{impl: impl.(Multiplier), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Subtracter",
		Iface: reflect.TypeOf((*Subtracter)(nil)).Elem(),
		New:   func() any { return &subtracter{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return subtracter_local_stub{impl: impl.(Subtracter), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return subtracter_client_stub{stub: stub, subtractMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/rynowak/serviceweaver-calculator/pkg/calculator/Subtracter", Method: "Subtract"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return subtracter_server_stub{impl: impl.(Subtracter), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type adder_local_stub struct {
	impl   Adder
	tracer trace.Tracer
}

func (s adder_local_stub) Add(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "calculator.Adder.Add", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Add(ctx, a0, a1)
}

type divider_local_stub struct {
	impl   Divider
	tracer trace.Tracer
}

func (s divider_local_stub) Divide(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "calculator.Divider.Divide", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Divide(ctx, a0, a1)
}

type multiplier_local_stub struct {
	impl   Multiplier
	tracer trace.Tracer
}

func (s multiplier_local_stub) Multiply(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "calculator.Multiplier.Multiply", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Multiply(ctx, a0, a1)
}

type subtracter_local_stub struct {
	impl   Subtracter
	tracer trace.Tracer
}

func (s subtracter_local_stub) Subtract(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "calculator.Subtracter.Subtract", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Subtract(ctx, a0, a1)
}

// Client stub implementations.

type adder_client_stub struct {
	stub       codegen.Stub
	addMetrics *codegen.MethodMetrics
}

func (s adder_client_stub) Add(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	// Update metrics.
	start := time.Now()
	s.addMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "calculator.Adder.Add", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.addMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.addMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Float64(a0)
	enc.Float64(a1)
	var shardKey uint64

	// Call the remote method.
	s.addMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.addMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Float64()
	err = dec.Error()
	return
}

type divider_client_stub struct {
	stub          codegen.Stub
	divideMetrics *codegen.MethodMetrics
}

func (s divider_client_stub) Divide(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	// Update metrics.
	start := time.Now()
	s.divideMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "calculator.Divider.Divide", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.divideMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.divideMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Float64(a0)
	enc.Float64(a1)
	var shardKey uint64

	// Call the remote method.
	s.divideMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.divideMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Float64()
	err = dec.Error()
	return
}

type multiplier_client_stub struct {
	stub            codegen.Stub
	multiplyMetrics *codegen.MethodMetrics
}

func (s multiplier_client_stub) Multiply(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	// Update metrics.
	start := time.Now()
	s.multiplyMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "calculator.Multiplier.Multiply", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.multiplyMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.multiplyMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Float64(a0)
	enc.Float64(a1)
	var shardKey uint64

	// Call the remote method.
	s.multiplyMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.multiplyMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Float64()
	err = dec.Error()
	return
}

type subtracter_client_stub struct {
	stub            codegen.Stub
	subtractMetrics *codegen.MethodMetrics
}

func (s subtracter_client_stub) Subtract(ctx context.Context, a0 float64, a1 float64) (r0 float64, err error) {
	// Update metrics.
	start := time.Now()
	s.subtractMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "calculator.Subtracter.Subtract", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.subtractMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.subtractMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Float64(a0)
	enc.Float64(a1)
	var shardKey uint64

	// Call the remote method.
	s.subtractMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.subtractMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Float64()
	err = dec.Error()
	return
}

// Server stub implementations.

type adder_server_stub struct {
	impl    Adder
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s adder_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Add":
		return s.add
	default:
		return nil
	}
}

func (s adder_server_stub) add(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 float64
	a0 = dec.Float64()
	var a1 float64
	a1 = dec.Float64()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Add(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Float64(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type divider_server_stub struct {
	impl    Divider
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s divider_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Divide":
		return s.divide
	default:
		return nil
	}
}

func (s divider_server_stub) divide(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 float64
	a0 = dec.Float64()
	var a1 float64
	a1 = dec.Float64()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Divide(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Float64(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type multiplier_server_stub struct {
	impl    Multiplier
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s multiplier_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Multiply":
		return s.multiply
	default:
		return nil
	}
}

func (s multiplier_server_stub) multiply(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 float64
	a0 = dec.Float64()
	var a1 float64
	a1 = dec.Float64()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Multiply(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Float64(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type subtracter_server_stub struct {
	impl    Subtracter
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s subtracter_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Subtract":
		return s.subtract
	default:
		return nil
	}
}

func (s subtracter_server_stub) subtract(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 float64
	a0 = dec.Float64()
	var a1 float64
	a1 = dec.Float64()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Subtract(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Float64(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}
