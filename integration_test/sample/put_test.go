package sample_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
	"github.com/sueken5/go-integration-test-example/pkg/apis/sample"
	"github.com/sueken5/go-integration-test-example/pkg/model"
)

func TestSamplePut(t *testing.T) {
	testCases := map[string]struct {
		req         *sample.PutRequest
		res         *sample.PutResponse
		expectErr   error
		expectPosts []*model.Post
	}{
		"ok": {
			req: &sample.PutRequest{
				Message: "test",
			},
			res:       &sample.PutResponse{},
			expectErr: nil,
			expectPosts: []*model.Post{
				{
					ID:        "",
					Message:   "test",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
			},
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			t.Cleanup(func() {
				if err := mysqlDB.Exec("DELETE FROM posts").Error; err != nil {
					t.Error(err)
				}
			})

			ctx := context.Background()
			_, actualErr := sampleClient.Put(ctx, tc.req)

			if d := cmp.Diff(
				actualErr,
				tc.expectErr,
			); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}

			actualPosts := make([]*model.Post, 0)
			if err := mysqlDB.Find(&actualPosts).Error; err != nil {
				t.Error(err)
			}

			if d := cmp.Diff(
				actualPosts,
				tc.expectPosts,
				cmpopts.IgnoreFields(model.Post{}, "ID", "CreatedAt", "UpdatedAt"),
			); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
