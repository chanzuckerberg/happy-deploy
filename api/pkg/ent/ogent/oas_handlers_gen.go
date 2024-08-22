// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleDeleteAppConfigRequest handles deleteAppConfig operation.
//
// Deletes the AppConfig with the requested Key.
//
// DELETE /app-configs/{key}
func (s *Server) handleDeleteAppConfigRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteAppConfig"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/app-configs/{key}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteAppConfig",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DeleteAppConfig",
			ID:   "deleteAppConfig",
		}
	)
	params, err := decodeDeleteAppConfigParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response DeleteAppConfigRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "DeleteAppConfig",
			OperationSummary: "",
			OperationID:      "deleteAppConfig",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "app_name",
					In:   "query",
				}: params.AppName,
				{
					Name: "environment",
					In:   "query",
				}: params.Environment,
				{
					Name: "stack",
					In:   "query",
				}: params.Stack,
				{
					Name: "aws_profile",
					In:   "query",
				}: params.AWSProfile,
				{
					Name: "aws_region",
					In:   "query",
				}: params.AWSRegion,
				{
					Name: "k8s_namespace",
					In:   "query",
				}: params.K8sNamespace,
				{
					Name: "k8s_cluster_id",
					In:   "query",
				}: params.K8sClusterID,
				{
					Name: "X-Aws-Access-Key-Id",
					In:   "header",
				}: params.XAWSAccessKeyID,
				{
					Name: "X-Aws-Secret-Access-Key",
					In:   "header",
				}: params.XAWSSecretAccessKey,
				{
					Name: "X-Aws-Session-Token",
					In:   "header",
				}: params.XAWSSessionToken,
				{
					Name: "key",
					In:   "path",
				}: params.Key,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteAppConfigParams
			Response = DeleteAppConfigRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteAppConfigParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.DeleteAppConfig(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.DeleteAppConfig(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteAppConfigResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleHealthRequest handles Health operation.
//
// Simple endpoint to check if the server is up.
//
// GET /health
func (s *Server) handleHealthRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Health"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/health"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "Health",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err error
	)

	var response HealthRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "Health",
			OperationSummary: "Simple endpoint to check if the server is up",
			OperationID:      "Health",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = HealthRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.Health(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.Health(ctx)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeHealthResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListAppConfigRequest handles listAppConfig operation.
//
// GET /app-configs
func (s *Server) handleListAppConfigRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listAppConfig"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/app-configs"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListAppConfig",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ListAppConfig",
			ID:   "listAppConfig",
		}
	)
	params, err := decodeListAppConfigParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response ListAppConfigRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "ListAppConfig",
			OperationSummary: "",
			OperationID:      "listAppConfig",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "itemsPerPage",
					In:   "query",
				}: params.ItemsPerPage,
				{
					Name: "app_name",
					In:   "query",
				}: params.AppName,
				{
					Name: "environment",
					In:   "query",
				}: params.Environment,
				{
					Name: "stack",
					In:   "query",
				}: params.Stack,
				{
					Name: "aws_profile",
					In:   "query",
				}: params.AWSProfile,
				{
					Name: "aws_region",
					In:   "query",
				}: params.AWSRegion,
				{
					Name: "k8s_namespace",
					In:   "query",
				}: params.K8sNamespace,
				{
					Name: "k8s_cluster_id",
					In:   "query",
				}: params.K8sClusterID,
				{
					Name: "X-Aws-Access-Key-Id",
					In:   "header",
				}: params.XAWSAccessKeyID,
				{
					Name: "X-Aws-Secret-Access-Key",
					In:   "header",
				}: params.XAWSSecretAccessKey,
				{
					Name: "X-Aws-Session-Token",
					In:   "header",
				}: params.XAWSSessionToken,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListAppConfigParams
			Response = ListAppConfigRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListAppConfigParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListAppConfig(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListAppConfig(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListAppConfigResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleReadAppConfigRequest handles readAppConfig operation.
//
// Finds the AppConfig with the requested Key and returns it.
//
// GET /app-configs/{key}
func (s *Server) handleReadAppConfigRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readAppConfig"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/app-configs/{key}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadAppConfig",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ReadAppConfig",
			ID:   "readAppConfig",
		}
	)
	params, err := decodeReadAppConfigParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response ReadAppConfigRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "ReadAppConfig",
			OperationSummary: "",
			OperationID:      "readAppConfig",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "app_name",
					In:   "query",
				}: params.AppName,
				{
					Name: "environment",
					In:   "query",
				}: params.Environment,
				{
					Name: "stack",
					In:   "query",
				}: params.Stack,
				{
					Name: "aws_profile",
					In:   "query",
				}: params.AWSProfile,
				{
					Name: "aws_region",
					In:   "query",
				}: params.AWSRegion,
				{
					Name: "k8s_namespace",
					In:   "query",
				}: params.K8sNamespace,
				{
					Name: "k8s_cluster_id",
					In:   "query",
				}: params.K8sClusterID,
				{
					Name: "X-Aws-Access-Key-Id",
					In:   "header",
				}: params.XAWSAccessKeyID,
				{
					Name: "X-Aws-Secret-Access-Key",
					In:   "header",
				}: params.XAWSSecretAccessKey,
				{
					Name: "X-Aws-Session-Token",
					In:   "header",
				}: params.XAWSSessionToken,
				{
					Name: "key",
					In:   "path",
				}: params.Key,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ReadAppConfigParams
			Response = ReadAppConfigRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackReadAppConfigParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ReadAppConfig(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ReadAppConfig(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeReadAppConfigResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSetAppConfigRequest handles setAppConfig operation.
//
// Sets an AppConfig with the specified Key/Value.
//
// POST /app-configs
func (s *Server) handleSetAppConfigRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("setAppConfig"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/app-configs"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SetAppConfig",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SetAppConfig",
			ID:   "setAppConfig",
		}
	)
	params, err := decodeSetAppConfigParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeSetAppConfigRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response SetAppConfigRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "SetAppConfig",
			OperationSummary: "",
			OperationID:      "setAppConfig",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "itemsPerPage",
					In:   "query",
				}: params.ItemsPerPage,
				{
					Name: "app_name",
					In:   "query",
				}: params.AppName,
				{
					Name: "environment",
					In:   "query",
				}: params.Environment,
				{
					Name: "stack",
					In:   "query",
				}: params.Stack,
				{
					Name: "aws_profile",
					In:   "query",
				}: params.AWSProfile,
				{
					Name: "aws_region",
					In:   "query",
				}: params.AWSRegion,
				{
					Name: "k8s_namespace",
					In:   "query",
				}: params.K8sNamespace,
				{
					Name: "k8s_cluster_id",
					In:   "query",
				}: params.K8sClusterID,
				{
					Name: "X-Aws-Access-Key-Id",
					In:   "header",
				}: params.XAWSAccessKeyID,
				{
					Name: "X-Aws-Secret-Access-Key",
					In:   "header",
				}: params.XAWSSecretAccessKey,
				{
					Name: "X-Aws-Session-Token",
					In:   "header",
				}: params.XAWSSessionToken,
			},
			Raw: r,
		}

		type (
			Request  = *SetAppConfigReq
			Params   = SetAppConfigParams
			Response = SetAppConfigRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSetAppConfigParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SetAppConfig(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.SetAppConfig(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSetAppConfigResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
