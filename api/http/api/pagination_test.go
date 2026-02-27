package api

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PaginatedResponse and PaginationMeta", func() {
	Describe("PaginationMeta", func() {
		It("should initialize correctly", func() {
			meta := PaginationMeta{
				Page:       2,
				PageSize:   10,
				Total:      25,
				TotalPages: 3,
			}

			Expect(meta.Page).To(Equal(2))
			Expect(meta.PageSize).To(Equal(10))
			Expect(meta.Total).To(Equal(int64(25)))
			Expect(meta.TotalPages).To(Equal(3))
		})

		It("should marshal to JSON with correct field names", func() {
			meta := PaginationMeta{
				Page:       1,
				PageSize:   20,
				Total:      100,
				TotalPages: 5,
			}

			bytes, err := json.Marshal(meta)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(bytes)).To(ContainSubstring(`"page":1`))
			Expect(string(bytes)).To(ContainSubstring(`"page_size":20`))
			Expect(string(bytes)).To(ContainSubstring(`"total":100`))
			Expect(string(bytes)).To(ContainSubstring(`"total_pages":5`))
		})
	})

	Describe("PaginatedResponse", func() {
		It("should wrap data and pagination correctly", func() {
			data := []string{"a", "b", "c"}
			meta := PaginationMeta{
				Page:       1,
				PageSize:   3,
				Total:      10,
				TotalPages: 4,
			}

			resp := PaginatedResponse[string]{
				Data:       data,
				Pagination: meta,
			}

			Expect(resp.Data).To(Equal(data))
			Expect(resp.Pagination).To(Equal(meta))
		})

		It("should marshal generic response to JSON correctly", func() {
			data := []string{"x", "y"}
			meta := PaginationMeta{
				Page:       1,
				PageSize:   2,
				Total:      5,
				TotalPages: 3,
			}

			resp := PaginatedResponse[string]{
				Data:       data,
				Pagination: meta,
			}

			bytes, err := json.Marshal(resp)
			Expect(err).NotTo(HaveOccurred())
			jsonStr := string(bytes)
			Expect(jsonStr).To(ContainSubstring(`"data":["x","y"]`))
			Expect(jsonStr).To(ContainSubstring(`"page":1`))
			Expect(jsonStr).To(ContainSubstring(`"page_size":2`))
			Expect(jsonStr).To(ContainSubstring(`"total":5`))
			Expect(jsonStr).To(ContainSubstring(`"total_pages":3`))
		})
	})
})
