---
title: Detecting Overload
keywords:
  - policies
  - signals
  - circuit
sidebar_position: 2
---

```mdx-code-block
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import Zoom from 'react-medium-image-zoom';
```

Aperture's control-loop policies are programmable "circuits" that are evaluated
periodically. One of the primary goals of these policies is to calculate the
deviation from objectives and apply counter-measures such as concurrency limits
to keep the system in a safe operational zone. The policies are used to express
where the metrics are collected from and where the actuation happens, along with
signal processing needed to translate health metrics to corrective actions.

For instance, a policy can be written to detect overload build-up at an upstream
service and trigger load-shedding at a downstream service.

## Policy

One of the most reliable metrics to detect overload state is latency of the
service requests. In Aperture, latency of service requests can be reported using
a [_Flux Meter_](/concepts/flow-control/resources/flux-meter.md).

:::tip

It's recommended to apply the _Flux Meter_ to a single type of workload to avoid
mixing the latency measurements across distinct workloads. For example, if there
are Select and Insert API calls on the same service, it is recommended to
measure the latency of only one of those workloads using a _Flux Meter_. Refer
to the [_Selector_](/concepts/flow-control/selector.md) on how to apply the
_Flux Meter_ to a subset of API calls for a service.

:::

In this example, the exponential moving average (EMA) of latency is computed,
which is gathered periodically from a
[PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) query on
_Flux Meter_ reported metrics. Further, EMA of latency will be multiplied with a
tolerance factor to calculate setpoint latency, which is a threshold to detect
overloaded state. That's, if the real-time latency of the service is more than
this setpoint (which is based on long-term EMA), then the service is considered
to be overloaded at that time.

```mdx-code-block
<Tabs>
<TabItem value="YAML">
```

```yaml
{@include: ./assets/detecting-overload/detecting-overload.yaml}
```

```mdx-code-block
</TabItem>
<TabItem value="Jsonnet">
```

```jsonnet
{@include: ./assets/detecting-overload/detecting-overload.jsonnet}
```

```mdx-code-block
</TabItem>
</Tabs>
```

### Circuit Diagram

<Zoom>

```mermaid
{@include: ./assets/detecting-overload/detecting-overload.mmd}
```

</Zoom>

### Playground

When the above policy is loaded in Aperture's
[Playground](/get-started/playground/playground.md), the various signal metrics
collected from the execution of the policy can be visualized:

<Zoom>

![LATENCY](./assets/detecting-overload/latency.png) `LATENCY`: Signal gathered
from the periodic execution of PromQL query on _Flux Meter_ metrics.

</Zoom>

<Zoom>

![LATENCY_EMA](./assets/detecting-overload/latency_ema.png) `LATENCY_EMA`:
Exponential Moving Average of `LATENCY` signal.

</Zoom>

<Zoom>

![LATENCY_SETPOINT](./assets/detecting-overload/latency_setpoint.png)
`LATENCY_SETPOINT`: Latency above which the service is considered to be
overloaded. This is calculated by multiplying the exponential moving average
with a tolerance factor (`LATENCY_EMA` \* `1.1`).

</Zoom>

<Zoom>

![IS_OVERLOAD_SWITCH](./assets/detecting-overload/is_overload_switch.png)
`IS_OVERLOAD_SWITCH` is a signal that decides whether the overload is currently
happening or not based on comparing `LATENCY` with `LATENCY_SETPOINT`. Its value
is `0` when there is no overload and `1` during overloads.

</Zoom>