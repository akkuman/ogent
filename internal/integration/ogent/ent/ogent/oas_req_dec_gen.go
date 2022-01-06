// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func decodeCreateCategoryRequest(r *http.Request, span trace.Span) (req CreateCategoryReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreateCategoryReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreateCategoryPetsRequest(r *http.Request, span trace.Span) (req CreateCategoryPetsReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreateCategoryPetsReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreatePetRequest(r *http.Request, span trace.Span) (req CreatePetReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreatePetReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreatePetCategoriesRequest(r *http.Request, span trace.Span) (req CreatePetCategoriesReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreatePetCategoriesReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreatePetFriendsRequest(r *http.Request, span trace.Span) (req CreatePetFriendsReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreatePetFriendsReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreatePetOwnerRequest(r *http.Request, span trace.Span) (req CreatePetOwnerReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreatePetOwnerReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreateUserRequest(r *http.Request, span trace.Span) (req CreateUserReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreateUserReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreateUserBestFriendRequest(r *http.Request, span trace.Span) (req CreateUserBestFriendReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreateUserBestFriendReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeCreateUserPetsRequest(r *http.Request, span trace.Span) (req CreateUserPetsReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request CreateUserPetsReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeUpdateCategoryRequest(r *http.Request, span trace.Span) (req UpdateCategoryReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request UpdateCategoryReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeUpdatePetRequest(r *http.Request, span trace.Span) (req UpdatePetReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request UpdatePetReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeUpdateUserRequest(r *http.Request, span trace.Span) (req UpdateUserReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request UpdateUserReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}
		if written == 0 {
			return req, nil
		}
		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, err
		}
		return request, nil
	default:
		return req, errors.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}
