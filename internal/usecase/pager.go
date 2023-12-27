package usecase

import (
	"context"
	api "prel/api/prel_api"
	"prel/pkg/custom_error"
	"prel/pkg/utils"

	"github.com/cockroachdb/errors"
)

type pagerInterface interface {
	Count(ctx context.Context) (int, error)
}

func calcPager(ctx context.Context, repo pagerInterface, page, pageSize int) (totalPage, currentPage, start, end int, err error) {
	totalPage = 1
	currentPage = 1
	if pageSize > api.LargestPageSize() {
		err = errors.WithDetail(
			errors.Newf("pageSize must be less than %d", api.LargestPageSize()),
			string(custom_error.InvalidArgument),
		)
		return
	}

	if page < 1 {
		err = errors.WithDetail(
			errors.Newf("page must be greater than 0"),
			string(custom_error.InvalidArgument),
		)
		return
	}

	count, err := repo.Count(ctx)
	if err != nil {
		return
	}

	if count == 0 {
		return
	}

	start, end = utils.CalcPageSize(page, pageSize, count)
	if start > end {
		err = errors.Newf("page must be less than %d", count)
		return
	}

	if start == end {
		return
	}

	if end > count {
		err = errors.Newf("page must be less than %d", count)
		return
	}

	totalPage = count / pageSize
	if count%pageSize != 0 {
		totalPage++
	}
	currentPage = page

	return
}
