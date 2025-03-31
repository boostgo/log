package logx

import (
	"github.com/boostgo/trace"
)

var (
	_tracer *trace.Tracer
)

func InitTracer(tracer *trace.Tracer) {
	_tracer = tracer
}

func Tracer() *trace.Tracer {
	return _tracer
}
