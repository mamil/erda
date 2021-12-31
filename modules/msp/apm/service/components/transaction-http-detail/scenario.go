// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transaction_http_detail

import (
	_ "github.com/erda-project/erda/modules/msp/apm/service/components/transaction-http-detail/avg_duration"
	_ "github.com/erda-project/erda/modules/msp/apm/service/components/transaction-http-detail/error_count"
	_ "github.com/erda-project/erda/modules/msp/apm/service/components/transaction-http-detail/grid"
	_ "github.com/erda-project/erda/modules/msp/apm/service/components/transaction-http-detail/kv_grid"
	_ "github.com/erda-project/erda/modules/msp/apm/service/components/transaction-http-detail/page"
	_ "github.com/erda-project/erda/modules/msp/apm/service/components/transaction-http-detail/rps"
	_ "github.com/erda-project/erda/modules/msp/apm/service/components/transaction-http-detail/slow_count"
)