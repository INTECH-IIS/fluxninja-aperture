---
title: Pure Storage FlashArray
description: Integrating Pure Storage FlashArray Metrics
keywords:
  - purefa
  - otel
  - opentelemetry
  - collector
  - metrics
---

Before proceeding, ensure that you have [built][build] Aperture Agent with the
`purefareceiver` extension enabled, so that [purefareceiver][receiver] is
available.

You can configure [Custom metrics][custom-metrics] for Pure Storage FlashArray
using the following configuration in the [Aperture Agent's
config][agent-config]:

```yaml
otel:
  custom_metrics:
    purefa:
      per_agent_group: true
      pipeline:
        processors:
          - batch
        receivers:
          - purefa
      processors:
        batch:
          send_batch_size: 10
          timeout: 10s
      receivers:
        purefa: [purefareceiver configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/purefareceiver
[custom-metrics]: /reference/configuration/agent.md#custom-metrics-config
[agent-config]: /reference/configuration/agent.md#agent-o-t-e-l-config