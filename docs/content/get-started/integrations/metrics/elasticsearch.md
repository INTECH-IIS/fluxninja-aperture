---
title: Elasticsearch
description: Integrating Elasticsearch Metrics
keywords:
  - elasticsearch
  - otel
  - opentelemetry
  - collector
  - metrics
---

Before proceeding, ensure that you have [built][build] Aperture Agent with the
`elasticsearchreceiver` extension enabled, so that
[elasticsearchreceiver][receiver] is available.

You can configure [Custom metrics][custom-metrics] for Elasticsearch using the
following configuration in the [Aperture Agent's config][agent-config]:

```yaml
otel:
  custom_metrics:
    elasticsearch:
      per_agent_group: true
      pipeline:
        processors:
          - batch
        receivers:
          - elasticsearch
      processors:
        batch:
          send_batch_size: 10
          timeout: 10s
      receivers:
        elasticsearch: [elasticsearchreceiver configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/elasticsearchreceiver
[custom-metrics]: /reference/configuration/agent.md#custom-metrics-config
[agent-config]: /reference/configuration/agent.md#agent-o-t-e-l-config