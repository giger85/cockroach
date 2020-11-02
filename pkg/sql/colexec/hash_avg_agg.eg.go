// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"unsafe"

	"github.com/cockroachdb/apd/v2"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase/colexecerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/cockroachdb/errors"
)

func newAvgHashAggAlloc(
	allocator *colmem.Allocator, t *types.T, allocSize int64,
) (aggregateFuncAlloc, error) {
	allocBase := aggAllocBase{allocator: allocator, allocSize: allocSize}
	switch t.Family() {
	case types.IntFamily:
		switch t.Width() {
		case 16:
			return &avgInt16HashAggAlloc{aggAllocBase: allocBase}, nil
		case 32:
			return &avgInt32HashAggAlloc{aggAllocBase: allocBase}, nil
		default:
			return &avgInt64HashAggAlloc{aggAllocBase: allocBase}, nil
		}
	case types.DecimalFamily:
		return &avgDecimalHashAggAlloc{aggAllocBase: allocBase}, nil
	case types.FloatFamily:
		return &avgFloat64HashAggAlloc{aggAllocBase: allocBase}, nil
	case types.IntervalFamily:
		return &avgIntervalHashAggAlloc{aggAllocBase: allocBase}, nil
	default:
		return nil, errors.Errorf("unsupported avg agg type %s", t.Name())
	}
}

type avgInt16HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &avgInt16HashAgg{}

const sizeOfAvgInt16HashAgg = int64(unsafe.Sizeof(avgInt16HashAgg{}))

func (a *avgInt16HashAgg) Init(groups []bool, vec coldata.Vec) {
	a.hashAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *avgInt16HashAgg) Reset() {
	a.hashAggregateFuncBase.Reset()
	a.scratch.curSum = zeroDecimalValue
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *avgInt16HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int16(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgInt16HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgInt16HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgInt16HashAgg
}

var _ aggregateFuncAlloc = &avgInt16HashAggAlloc{}

func (a *avgInt16HashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfAvgInt16HashAgg * a.allocSize)
		a.aggFuncs = make([]avgInt16HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgInt32HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &avgInt32HashAgg{}

const sizeOfAvgInt32HashAgg = int64(unsafe.Sizeof(avgInt32HashAgg{}))

func (a *avgInt32HashAgg) Init(groups []bool, vec coldata.Vec) {
	a.hashAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *avgInt32HashAgg) Reset() {
	a.hashAggregateFuncBase.Reset()
	a.scratch.curSum = zeroDecimalValue
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *avgInt32HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int32(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgInt32HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgInt32HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgInt32HashAgg
}

var _ aggregateFuncAlloc = &avgInt32HashAggAlloc{}

func (a *avgInt32HashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfAvgInt32HashAgg * a.allocSize)
		a.aggFuncs = make([]avgInt32HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgInt64HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
	overloadHelper overloadHelper
}

var _ aggregateFunc = &avgInt64HashAgg{}

const sizeOfAvgInt64HashAgg = int64(unsafe.Sizeof(avgInt64HashAgg{}))

func (a *avgInt64HashAgg) Init(groups []bool, vec coldata.Vec) {
	a.hashAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *avgInt64HashAgg) Reset() {
	a.hashAggregateFuncBase.Reset()
	a.scratch.curSum = zeroDecimalValue
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *avgInt64HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	// In order to inline the templated code of overloads, we need to have a
	// "_overloadHelper" local variable of type "overloadHelper".
	_overloadHelper := a.overloadHelper
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Int64(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						tmpDec := &_overloadHelper.tmpDec1
						tmpDec.SetInt64(int64(col[i]))
						if _, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, tmpDec); err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgInt64HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgInt64HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgInt64HashAgg
}

var _ aggregateFuncAlloc = &avgInt64HashAggAlloc{}

func (a *avgInt64HashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfAvgInt64HashAgg * a.allocSize)
		a.aggFuncs = make([]avgInt64HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgDecimalHashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum apd.Decimal
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []apd.Decimal
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &avgDecimalHashAgg{}

const sizeOfAvgDecimalHashAgg = int64(unsafe.Sizeof(avgDecimalHashAgg{}))

func (a *avgDecimalHashAgg) Init(groups []bool, vec coldata.Vec) {
	a.hashAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Decimal()
	a.Reset()
}

func (a *avgDecimalHashAgg) Reset() {
	a.hashAggregateFuncBase.Reset()
	a.scratch.curSum = zeroDecimalValue
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *avgDecimalHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Decimal(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						_, err := tree.ExactCtx.Add(&a.scratch.curSum, &a.scratch.curSum, &col[i])
						if err != nil {
							colexecerror.ExpectedError(err)
						}
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgDecimalHashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {

		a.scratch.vec[outputIdx].SetInt64(a.scratch.curCount)
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[outputIdx], &a.scratch.curSum, &a.scratch.vec[outputIdx]); err != nil {
			colexecerror.InternalError(err)
		}
	}
}

type avgDecimalHashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgDecimalHashAgg
}

var _ aggregateFuncAlloc = &avgDecimalHashAggAlloc{}

func (a *avgDecimalHashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfAvgDecimalHashAgg * a.allocSize)
		a.aggFuncs = make([]avgDecimalHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgFloat64HashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum float64
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []float64
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &avgFloat64HashAgg{}

const sizeOfAvgFloat64HashAgg = int64(unsafe.Sizeof(avgFloat64HashAgg{}))

func (a *avgFloat64HashAgg) Init(groups []bool, vec coldata.Vec) {
	a.hashAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Float64()
	a.Reset()
}

func (a *avgFloat64HashAgg) Reset() {
	a.hashAggregateFuncBase.Reset()
	a.scratch.curSum = zeroFloat64Value
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *avgFloat64HashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Float64(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {

					{

						a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {

					{

						a.scratch.curSum = float64(a.scratch.curSum) + float64(col[i])
					}

					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgFloat64HashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curSum / float64(a.scratch.curCount)
	}
}

type avgFloat64HashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgFloat64HashAgg
}

var _ aggregateFuncAlloc = &avgFloat64HashAggAlloc{}

func (a *avgFloat64HashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfAvgFloat64HashAgg * a.allocSize)
		a.aggFuncs = make([]avgFloat64HashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}

type avgIntervalHashAgg struct {
	hashAggregateFuncBase
	scratch struct {
		// curSum keeps track of the sum of elements belonging to the current group,
		// so we can index into the slice once per group, instead of on each
		// iteration.
		curSum duration.Duration
		// curCount keeps track of the number of elements that we've seen
		// belonging to the current group.
		curCount int64
		// vec points to the output vector.
		vec []duration.Duration
		// foundNonNullForCurrentGroup tracks if we have seen any non-null values
		// for the group that is currently being aggregated.
		foundNonNullForCurrentGroup bool
	}
}

var _ aggregateFunc = &avgIntervalHashAgg{}

const sizeOfAvgIntervalHashAgg = int64(unsafe.Sizeof(avgIntervalHashAgg{}))

func (a *avgIntervalHashAgg) Init(groups []bool, vec coldata.Vec) {
	a.hashAggregateFuncBase.Init(groups, vec)
	a.scratch.vec = vec.Interval()
	a.Reset()
}

func (a *avgIntervalHashAgg) Reset() {
	a.hashAggregateFuncBase.Reset()
	a.scratch.curSum = zeroIntervalValue
	a.scratch.curCount = 0
	a.scratch.foundNonNullForCurrentGroup = false
}

func (a *avgIntervalHashAgg) Compute(
	vecs []coldata.Vec, inputIdxs []uint32, inputLen int, sel []int,
) {
	vec := vecs[inputIdxs[0]]
	col, nulls := vec.Interval(), vec.Nulls()
	{
		sel = sel[:inputLen]
		if nulls.MaybeHasNulls() {
			for _, i := range sel {

				var isNull bool
				isNull = nulls.NullAt(i)
				if !isNull {
					a.scratch.curSum = a.scratch.curSum.Add(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		} else {
			for _, i := range sel {

				var isNull bool
				isNull = false
				if !isNull {
					a.scratch.curSum = a.scratch.curSum.Add(col[i])
					a.scratch.curCount++
					a.scratch.foundNonNullForCurrentGroup = true
				}
			}
		}
	}
}

func (a *avgIntervalHashAgg) Flush(outputIdx int) {
	// The aggregation is finished. Flush the last value. If we haven't found
	// any non-nulls for this group so far, the output for this group should be
	// NULL.
	if !a.scratch.foundNonNullForCurrentGroup {
		a.nulls.SetNull(outputIdx)
	} else {
		a.scratch.vec[outputIdx] = a.scratch.curSum.Div(int64(a.scratch.curCount))
	}
}

type avgIntervalHashAggAlloc struct {
	aggAllocBase
	aggFuncs []avgIntervalHashAgg
}

var _ aggregateFuncAlloc = &avgIntervalHashAggAlloc{}

func (a *avgIntervalHashAggAlloc) newAggFunc() aggregateFunc {
	if len(a.aggFuncs) == 0 {
		a.allocator.AdjustMemoryUsage(sizeOfAvgIntervalHashAgg * a.allocSize)
		a.aggFuncs = make([]avgIntervalHashAgg, a.allocSize)
	}
	f := &a.aggFuncs[0]
	a.aggFuncs = a.aggFuncs[1:]
	return f
}