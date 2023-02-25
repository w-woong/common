package apm

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/go-wonk/si/v2/sihttp"
	"github.com/w-woong/common/utils"
	"go.elastic.co/apm/v2"
)

var _active = false

func init() {
	_active, _ = strconv.ParseBool(os.Getenv("ELASTIC_APM_ACTIVE"))
}

type Transaction interface {
	Ender
}
type Span interface {
	Ender
}
type Ender interface {
	End()
}
type NopEnder struct{}

func (NopEnder) End() {}

type TransactionType string

var (
	TransactionTypeRequest TransactionType = "request"
)

type SpanType string

var (
	SpanTypeDB           SpanType = "db"
	SpanTypeHttp         SpanType = "http"
	SpanTypeExternalHttp SpanType = "external.http"
)

type SpanName string

var (
	SpanNameSelect     SpanName = "SELECT"
	SpanNameInsert     SpanName = "INSERT"
	SpanNameUpdate     SpanName = "UPDATE"
	SpanNameDelete     SpanName = "DELETE"
	SpanNameMerge      SpanName = "MERGE"
	SpanNameStoredProc SpanName = "SP"

	SpanNameHttpGet    SpanName = "GET"
	SpanNameHttpPOST   SpanName = "POST"
	SpanNameHttpPUT    SpanName = "PUT"
	SpanNameHttpDELETE SpanName = "DELETE"
	SpanNameHttpHEAD   SpanName = "HEAD"
)

func GetOrCreateRequestTransaction(ctx context.Context) (context.Context, Transaction) {
	if !_active {
		return ctx, &NopEnder{}
	}
	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, _, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	return ctx, tran
}

// GetOrCreateTransaction tries to find apm.Transaction from ctx and return it if exist, or create a new one.
func GetOrCreateTransaction(ctx context.Context, tranName string, tranType TransactionType) (context.Context, Transaction) {
	if !_active {
		return ctx, &NopEnder{}
	}
	ctx, _, tran := getOrCreateTransaction(ctx, tranName, tranType)
	return ctx, tran
}

func getOrCreateTransaction(ctx context.Context, tranName string, tranType TransactionType) (context.Context, bool, *apm.Transaction) {
	if tx := apm.TransactionFromContext(ctx); tx == nil {
		newTx := apm.DefaultTracer().StartTransaction(tranName, string(tranType))
		ctx = apm.ContextWithTransaction(ctx, newTx)
		return ctx, true, newTx
	} else {
		return ctx, false, tx
	}
}

var NopFunc = func() {}

func StartOracleSpan(ctx context.Context, spanName SpanName, qry string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}
	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startDBSpan(ctx, spanName, "oracle", qry)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}

	return ctx, span, func() {
		span.End()
	}
}

func StartNewOracleSpan(ctx context.Context, spanName SpanName, qry string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}
	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startNewDBSpan(ctx, spanName, "oracle", qry)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}

	return ctx, span, func() {
		span.End()
	}
}

func StartMSSqlSpan(ctx context.Context, spanName SpanName, qry string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}

	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startDBSpan(ctx, spanName, "mssql", qry)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}

	return ctx, span, func() {
		span.End()
	}
}

func StartNewMSSqlSpan(ctx context.Context, spanName SpanName, qry string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}

	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startNewDBSpan(ctx, spanName, "mssql", qry)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}

	return ctx, span, func() {
		span.End()
	}
}

func StartPostgresqlSpan(ctx context.Context, spanName SpanName, qry string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}

	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startDBSpan(ctx, spanName, "postgresql", qry)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}

	return ctx, span, func() {
		span.End()
	}
}
func StartNewPostgresqlSpan(ctx context.Context, spanName SpanName, qry string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}

	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startNewDBSpan(ctx, spanName, "postgresql", qry)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}

	return ctx, span, func() {
		span.End()
	}
}

func startNewDBSpan(ctx context.Context, spanName SpanName, dbType string, qry string) (context.Context, *apm.Span) {
	span, ctx := apm.StartSpan(ctx, string(spanName), string(SpanTypeDB))
	setDbSpanLabels(span, dbType, qry)
	return ctx, span
}

func startDBSpan(ctx context.Context, spanName SpanName, dbType string, qry string) (context.Context, *apm.Span) {
	if existingSpan := apm.SpanFromContext(ctx); existingSpan == nil {
		return startNewDBSpan(ctx, spanName, dbType, qry)
	} else {
		return ctx, existingSpan
	}
}

func StartHttpSpan(ctx context.Context, spanName string, method string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}

	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startHttpSpan(ctx, spanName, method)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}
	return ctx, span, func() {
		span.End()
	}
}

func StartNewHttpSpan(ctx context.Context, spanName string, method string) (context.Context, Span, func()) {
	if !_active {
		return ctx, &NopEnder{}, NopFunc
	}

	tranName := utils.GetFunctionNameWithSkip(2)
	ctx, isNew, tran := getOrCreateTransaction(ctx, tranName, TransactionTypeRequest)
	ctx, span := startNewHttpSpan(ctx, spanName, method)
	if isNew {
		return ctx, span, func() {
			span.End()
			tran.End()
		}
	}
	return ctx, span, func() {
		span.End()
	}
}

func startNewHttpSpan(ctx context.Context, spanName string, method string) (context.Context, *apm.Span) {
	span, ctx := apm.StartSpan(ctx, spanName, string(SpanTypeExternalHttp))
	span.Subtype = string(SpanTypeHttp)
	span.Action = method
	return ctx, span
}

func startHttpSpan(ctx context.Context, spanName string, method string) (context.Context, *apm.Span) {
	if existingSpan := apm.SpanFromContext(ctx); existingSpan == nil {
		return startNewHttpSpan(ctx, spanName, method)
	} else {
		return ctx, existingSpan
	}
}

func SetHttpStatusCode(span Span, status int) {
	if !_active {
		return
	}

	if apmSpan, ok := span.(*apm.Span); ok {
		apmSpan.Context.SetHTTPStatusCode(status)
	}
}

func setDbSpanLabels(span *apm.Span, dbType string, qry string) {
	// span.Context.SetLabel("service.target.type", "oracle")
	// if span.Name == string(SpanNameSelect) {
	// 	span.Context.SetLabel("span.action", "query")
	// } else {
	// 	span.Context.SetLabel("span.action", "exec")
	// }
	// span.Context.SetLabel("span.db.statement", qry)
	// span.Context.SetLabel("span.db.type", "sql")

	// span.Context.SetLabel("span.destination.service.name", "oracle")
	// span.Context.SetLabel("span.destination.service.resource", "oracle")
	// span.Context.SetLabel("span.destination.service.type", string(SpanTypeDB))

	if span.Name == string(SpanNameSelect) {
		span.Action = "query"
	} else {
		span.Action = "exec"
	}
	span.Subtype = dbType
	span.Context.SetDestinationService(apm.DestinationServiceSpanContext{
		Name:     dbType,
		Resource: dbType,
	})
	span.Context.SetDatabase(apm.DatabaseSpanContext{
		Statement: qry,
		Type:      "sql",
	})
}

func WithApmSpan(span Span) sihttp.RequestOptionFunc {
	if !_active {
		return sihttp.RequestOptionFunc(func(req *http.Request) error {
			return nil
		})
	}

	if apmSpan, ok := span.(*apm.Span); ok {
		return sihttp.RequestOptionFunc(func(req *http.Request) error {
			apmSpan.Context.SetHTTPRequest(req)
			return nil
		})
	}

	return sihttp.RequestOptionFunc(func(req *http.Request) error {
		return nil
	})

}
