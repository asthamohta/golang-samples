// Copyright 2021 Google LLC
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

// [START spanner_list_databases]
import (
	"context"
	"fmt"
	"io"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"google.golang.org/api/iterator"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func listDatabases(ctx context.Context, w io.Writer, instanceId string) error {
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return err
	}
	defer adminClient.Close()

	iter := adminClient.ListDatabases(ctx, &adminpb.ListDatabasesRequest{
		Parent: instanceId,
	})

	printDatabases := func(iter *database.DatabaseIterator) error {
		fmt.Printf("Databases for instance/[%s]", instanceId)
		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				return nil
			}
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "Backup %s\n", resp.Name)
		}
	}

	if err := printDatabases(iter); err != nil {
		return err
	}

	return nil

}
