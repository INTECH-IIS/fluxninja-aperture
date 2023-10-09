local utils = import '../../utils/policy_utils.libsonnet';
local timeSeriesPanel = import '../../utils/time_series_panel.libsonnet';

local g = import 'github.com/grafana/grafonnet/gen/grafonnet-v9.4.0/main.libsonnet';

function(policyName, infraMeterName, datasource, extraFilters) {

  local stringFilters = utils.dictToPrometheusFilter(extraFilters { policy_name: policyName }),

  local readyQuery = g.query.prometheus.new(datasource, 'sum(rabbitmq_message_current{%(filters)s, state="ready"})' % { filters: stringFilters })
                     + g.query.prometheus.withIntervalFactor(1)
                     + g.query.prometheus.withLegendFormat('Ready'),

  local ready = timeSeriesPanel('Messages Ready For Consumers', datasource.name, '', stringFilters, 'Count', targets=[readyQuery]),

  local acknowledgedQuery = g.query.prometheus.new(datasource, 'sum(rate(rabbitmq_message_acknowledged_total{%(filters)s}[$__rate_interval]))' % { filters: stringFilters })
                            + g.query.prometheus.withIntervalFactor(1)
                            + g.query.prometheus.withLegendFormat('Acknowledged'),

  local acknowledged = timeSeriesPanel('Messages Acknowledged By Consumers', datasource.name, '', stringFilters, 'Count', targets=[acknowledgedQuery]),

  local unacknowledgedQuery = g.query.prometheus.new(datasource, 'sum(rabbitmq_message_current{%(filters)s, state="unacknowledged"})' % { filters: stringFilters })
                              + g.query.prometheus.withIntervalFactor(1)
                              + g.query.prometheus.withLegendFormat('Unacknowledged'),

  local unacknowledged = timeSeriesPanel('Messages Unacknowledged By Consumers', datasource.name, '', stringFilters, 'Count', targets=[unacknowledgedQuery]),


  local readyQueryPerVhost = g.query.prometheus.new(datasource, 'sum by(rabbitmq_vhost_name) (rabbitmq_message_current{%(filters)s, state="ready"})' % { filters: stringFilters })
                             + g.query.prometheus.withIntervalFactor(1)
                             + g.query.prometheus.withLegendFormat('{{rabbitmq_vhost_name}}'),

  local readyPerVhost = timeSeriesPanel('Messages Ready For Consumers Per Vhost', datasource.name, '', stringFilters, 'Count', targets=[readyQueryPerVhost]),

  local acknowledgedQueryPerVhost = g.query.prometheus.new(datasource, 'sum by(rabbitmq_vhost_name) (rate(rabbitmq_message_acknowledged_total{%(filters)s}[$__rate_interval]))' % { filters: stringFilters })
                                    + g.query.prometheus.withIntervalFactor(1)
                                    + g.query.prometheus.withLegendFormat('{{rabbitmq_vhost_name}}'),

  local acknowledgedPerVhost = timeSeriesPanel('Messages Acknowledged By Consumers Per Vhost', datasource.name, '', stringFilters, 'Count', targets=[acknowledgedQueryPerVhost]),

  local unacknowledgedQueryPerVhost = g.query.prometheus.new(datasource, 'sum by(rabbitmq_vhost_name) (rabbitmq_message_current{%(filters)s, state="unacknowledged"})' % { filters: stringFilters })
                                      + g.query.prometheus.withIntervalFactor(1)
                                      + g.query.prometheus.withLegendFormat('{{rabbitmq_vhost_name}}'),

  local unacknowledgedPerVhost = timeSeriesPanel('Messages Unacknowledged By Consumers Per Vhost', datasource.name, '', stringFilters, 'Count', targets=[unacknowledgedQueryPerVhost]),

  local queuesGrowthQuery = g.query.prometheus.new(datasource, 'sum by (rabbitmq_queue_name) (rate(rabbitmq_message_published_total{%(filters)s}[$__rate_interval])) - sum by (rabbitmq_queue_name) (rate(rabbitmq_message_acknowledged_total{%(filters)s}[$__rate_interval]))' % { filters: stringFilters })
                            + g.query.prometheus.withIntervalFactor(1)
                            + g.query.prometheus.withLegendFormat('{{rabbitmq_queue_name}}'),

  local queuesGrowth = timeSeriesPanel('Queues Growth', datasource.name, '', stringFilters, 'Messages/Second', targets=[queuesGrowthQuery]),

  local publishedQuery = g.query.prometheus.new(datasource, 'sum(rate(rabbitmq_message_published_total{%(filters)s}[$__rate_interval]))' % { filters: stringFilters })
                         + g.query.prometheus.withIntervalFactor(1)
                         + g.query.prometheus.withLegendFormat('Published'),

  local published = timeSeriesPanel('Messages Published', datasource.name, '', stringFilters, 'Messages/Second', targets=[publishedQuery]),

  local deliveredQuery = g.query.prometheus.new(datasource, 'sum(rate(rabbitmq_message_delivered_total{%(filters)s}[$__rate_interval]))' % { filters: stringFilters })
                         + g.query.prometheus.withIntervalFactor(1)
                         + g.query.prometheus.withLegendFormat('Delivered'),

  local delivered = timeSeriesPanel('Messages Delivered', datasource.name, '', stringFilters, 'Messages/Second', targets=[deliveredQuery]),


  ready: ready.panel,
  acknowledged: acknowledged.panel,
  unacknowledged: unacknowledged.panel,
  readyPerVhost: readyPerVhost.panel,
  acknowledgedPerVhost: acknowledgedPerVhost.panel,
  unacknowledgedPerVhost: unacknowledgedPerVhost.panel,
  queuesGrowth: queuesGrowth.panel,
  published: published.panel,
  delivered: delivered.panel,

}