// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanner

// [START spanner_update_data_with_json_column]
import (
	"context"
	"fmt"
	"io"
	"regexp"

	"cloud.google.com/go/spanner"
)

// updateDataWithJsonColumn updates database with Json type values
func updateDataWithJsonColumn(w io.Writer, db string) error {
	// db = `projects/<project>/instances/<instance-id>/database/<database-id>`
	matches := regexp.MustCompile("^(.*)/databases/(.*)$").FindStringSubmatch(db)
	if matches == nil || len(matches) != 3 {
		return fmt.Errorf("addJsonColumn: invalid database id %s", db)
	}

	ctx := context.Background()

	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		return err
	}
	defer client.Close()

	type VenueDetails struct {
		Name   interface{} `json:"name"`
		Rating interface{} `json:"rating"`
		Open   interface{} `json:"open"`
		Tags   interface{} `json:"tags"`
	}

	details_1, _ := NullJSON([]VenueDetails{
		{Name: "room1", Open: true},
		{Name: "room2", Open: false},
	}, true)
	details_2, _ := NullJSON(VenueDetails{
		Rating: 9,
		Open:   true,
	}, true)
	details_3, _ := NullJSON(VenueDetails{
		Name: nil,
		Open: map[string]bool{"monday": true, "tuesday": false},
	}, true)

	cols := []string{"VenueId", "VenueDetails"}
	_, err = client.Apply(ctx, []*spanner.Mutation{
		spanner.Update("VenueDetails", cols, []interface{}{4, string(details_1)}),
		spanner.Update("VenueDetails", cols, []interface{}{19, string(details_2)}),
		spanner.Update("VenueDetails", cols, []interface{}{42, string(details_3)}),
	})

	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Updated data to VenueDetails column\n")

	return nil
}

// [END spanner_update_data_with_json_column]
