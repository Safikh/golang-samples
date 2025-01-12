// Copyright 2023 Google LLC
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

package snippets

import (
	"bytes"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestTranscribeFileV2(t *testing.T) {
	tc := testutil.SystemTest(t)

	var buf bytes.Buffer
	err := transcribeFileV2(&buf, tc.ProjectID, "../testdata/Google_Gnome.wav")
	if err != nil {
		t.Fatal(err)
	}

	if got := buf.String(); !strings.Contains(got, "stream stranger things from Netflix to my TV") {
		t.Fatalf(`transcribeFileV2(../testdata/Google_Gnome.wav) = %q; want "stream stranger things from Netflix to my TV"`, got)
	}
}
