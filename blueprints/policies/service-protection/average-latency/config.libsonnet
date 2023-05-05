local serviceProtectionDefaults = import '../base/config-defaults.libsonnet';

/**
* @param (common.policy_name: string required) Name of the policy.
* @param (policy.components: []aperture.spec.v1.Component) List of additional circuit components.
* @param (policy.resources: aperture.spec.v1.Resources) Additional resources.
* @param (policy.evaluation_interval: string) The interval between successive evaluations of the Circuit.
* @param (policy.service_protection_core.overload_confirmations: []overload_confirmation) List of overload confirmation criteria. Load scheduler can shed flows when all of the specified overload confirmation criteria are met.
* @schema (overload_confirmation.query_string: string required) The Prometheus query to be run. Must return a scalar or a vector with a single element.
* @schema (overload_confirmation.threshold: float64) The threshold for the overload confirmation criteria.
* @schema (overload_confirmation.operator: string) The operator for the overload confirmation criteria. oneof: `gt | lt | gte | lte | eq | neq`
* @param (policy.service_protection_core.adaptive_load_scheduler: aperture.spec.v1.AdaptiveLoadSchedulerParameters required) Parameters for Adaptive Load Scheduler.
* @param (policy.service_protection_core.dry_run: bool) Default configuration for setting dry run mode on Load Scheduler. In dry run mode, the Load Scheduler acts as a passthrough and does not throttle flows. This config can be updated at runtime without restarting the policy.
*/

/**
* @param (dashboard.refresh_interval: string) Refresh interval for dashboard panels.
* @param (dashboard.time_from: string) From time of dashboard.
* @param (dashboard.time_to: string) To time of dashboard.
* @param (dashboard.datasource.name: string) Datasource name.
* @param (dashboard.datasource.filter_regex: string) Datasource filter regex.
*/

serviceProtectionDefaults {
  policy+: {
    latency_baseliner: {
      /**
      * @param (policy.latency_baseliner.flux_meter: aperture.spec.v1.FluxMeter required) Flux Meter defines the scope of latency measurements.
      */
      flux_meter: {
        selectors: serviceProtectionDefaults.selectors,
      },
      /**
      * @param (policy.latency_baseliner.ema: aperture.spec.v1.EMAParameters) EMA parameters.
      * @param (policy.latency_baseliner.latency_tolerance_multiplier: float64) Tolerance factor beyond which the service is considered to be in overloaded state. E.g. if EMA of latency is 50ms and if Tolerance is 1.1, then service is considered to be in overloaded state if current latency is more than 55ms.
      * @param (policy.latency_baseliner.latency_ema_limit_multiplier: float64) Current latency value is multiplied with this factor to calculate maximum envelope of Latency EMA.
      */
      ema: {
        ema_window: '1500s',
        warmup_window: '60s',
        correction_factor_on_max_envelope_violation: 0.95,
      },
      latency_tolerance_multiplier: 1.1,
      latency_ema_limit_multiplier: 2.0,
    },
  },
}