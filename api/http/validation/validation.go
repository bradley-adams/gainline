package validation

import (
	"github.com/bradley-adams/gainline/http/api"
	"github.com/go-playground/validator/v10"
)

func ValidateSeasonDates(sl validator.StructLevel) {
	season := sl.Current().Interface().(api.SeasonRequest)

	if season.StartDate.After(season.EndDate) {
		sl.ReportError(
			season.StartDate,
			"StartDate",
			"start_date",
			"start_must_be_before_end",
			"",
		)
	}
}
