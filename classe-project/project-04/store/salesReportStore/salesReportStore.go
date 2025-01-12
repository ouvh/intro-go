package salesReportStore

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"

	"um6p.ma/project-04/models/order"
	model "um6p.ma/project-04/models/salesReport"
	"um6p.ma/project-04/store/searchCriteria"
)

type SalesReportExecutor interface {
	CreateSalesReport(order model.SalesReport) (model.SalesReport, error)
	SearchSalesReport(criteria searchCriteria.SearchCriteria) ([]model.SalesReport, error)
	AddSale(order order.Order)
}

type SalesReportStore struct {
	sync.Mutex
	SalesReports []model.SalesReport
	Index        int64
	CurrentSales []order.Order
}

func (b *SalesReportStore) AddSale(order order.Order) {
	b.Lock()
	defer b.Unlock()
	b.CurrentSales = append(b.CurrentSales, order)

}

func (b *SalesReportStore) RetrieveSales() []order.Order {
	b.Lock()
	defer b.Unlock()
	result := b.CurrentSales
	b.CurrentSales = b.CurrentSales[:0]
	return result

}

func (b *SalesReportStore) CreateSalesReport(salesReport model.SalesReport) (model.SalesReport, error) {

	b.Lock()
	defer b.Unlock()
	b.Index++
	salesReport.Timestamp = time.Now()
	b.SalesReports = append(b.SalesReports, salesReport)
	return salesReport, nil

}

func (b *SalesReportStore) SearchSalesReport(criteria searchCriteria.SearchCriteria) ([]model.SalesReport, error) {
	b.Lock()
	defer b.Unlock()
	results := make([]model.SalesReport, 0)

	if len(criteria.Parameters) == 0 {
		return b.SalesReports, nil
	}

	for _, order := range b.SalesReports {
		matched := true
	loop:
		for key, value := range criteria.Parameters {
			comp, exist := model.ComparableFields[key]
			if exist {
				if comp == 1 {
					v, err := model.GetField(order, key)
					if err != nil {
						return results, err
					}
					if true {
						x, err := strconv.ParseFloat(value.(string), 64)
						if err != nil {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if reflect.ValueOf(v).Float() > reflect.ValueOf(value).Float() {
									matched = false
									break loop
								}
							case reflect.Float32, reflect.Float64:
								if reflect.ValueOf(v).Float() > reflect.ValueOf(value).Float() {
									matched = false
									break loop

								}
							case reflect.String:
								if reflect.ValueOf(v).String() > reflect.ValueOf(value).String() {
									matched = false
									break loop

								}

							default:
								vv := v.(time.Time)
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.After(vvvalue) {
									matched = false
									break loop

								}

							}

						} else {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if (reflect.ValueOf(v).Float)() > reflect.ValueOf(x).Float() {
									matched = false
									break loop
								}
							case reflect.Float32, reflect.Float64:
								if reflect.ValueOf(v).Float() > reflect.ValueOf(x).Float() {
									matched = false
									break loop

								}
							case reflect.String:
								if reflect.ValueOf(v).String() > reflect.ValueOf(value).String() {
									matched = false
									break loop

								}

							default:
								vv := v.(time.Time)
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.After(vvvalue) {
									matched = false
									break loop

								}

							}

						}

					} else {
						return results, fmt.Errorf("type mismatch: %T vs %T", v, value)
					}
				} else {

					v, err := model.GetField(order, key)
					if err != nil {
						return results, err
					}
					if true {
						x, err := strconv.ParseFloat(value.(string), 64)
						if err != nil {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if reflect.ValueOf(v).Float() < reflect.ValueOf(value).Float() {
									matched = false
									break loop
								}
							case reflect.Float32, reflect.Float64:
								if reflect.ValueOf(v).Float() < reflect.ValueOf(value).Float() {
									matched = false
									break loop

								}
							case reflect.String:
								if reflect.ValueOf(v).String() < reflect.ValueOf(value).String() {
									matched = false
									break loop

								}

							default:
								vv := v.(time.Time)
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.Before(vvvalue) {
									matched = false
									break loop

								}

							}

						} else {
							switch reflect.ValueOf(v).Kind() {
							case reflect.Int, reflect.Int32, reflect.Int64:
								if (reflect.ValueOf(v).Float)() < reflect.ValueOf(x).Float() {
									matched = false
									break loop
								}
							case reflect.Float32, reflect.Float64:
								if reflect.ValueOf(v).Float() < reflect.ValueOf(x).Float() {
									matched = false
									break loop

								}
							case reflect.String:
								if reflect.ValueOf(v).String() < reflect.ValueOf(value).String() {
									matched = false
									break loop

								}

							default:
								vv := v.(time.Time)
								log.Println(value)
								vvvalue, err := time.Parse("2006-01-02T15:04:05.9999999-07:00", value.(string))
								if err != nil {
									return results, fmt.Errorf("datetime format incompatible")
								}

								if vv.Before(vvvalue) {
									matched = false
									break loop

								}

							}

						}

					} else {
						return results, fmt.Errorf("type mismatch: %T vs %T", v, value)
					}

				}

			} else {
				f, err := model.GetField(order, key)
				if err != nil {
					return results, err
				}

				x, err := strconv.ParseFloat(value.(string), 64)
				if err != nil {
					if value != f {
						matched = false
						break
					}
				} else {
					if float64(x) != float64(f.(float64)) {
						matched = false
						break
					}
				}

			}

		}

		if matched {
			results = append(results, order)
		}

	}

	return results, nil

}
