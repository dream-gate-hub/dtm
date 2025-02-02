package microservices

import (
	// load the microserver drivers
	_ "github.com/dream-gate-hub/etmarket-infra/databus/remote/dtm/driver"
	_ "github.com/dtm-labs/dtmdriver-dapr"
	_ "github.com/dtm-labs/dtmdriver-ego"
	_ "github.com/dtm-labs/dtmdriver-gozero"
	_ "github.com/dtm-labs/dtmdriver-kratos"
	_ "github.com/dtm-labs/dtmdriver-polaris"
	_ "github.com/dtm-labs/dtmdriver-springcloud"
)
