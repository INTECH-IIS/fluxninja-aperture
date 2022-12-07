// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: envoy/service/auth/v3/authz_stripped.proto

package com.fluxninja.generated.envoy.service.auth.v3;

public interface SocketAddressOrBuilder extends
    // @@protoc_insertion_point(interface_extends:envoy.service.auth.v3.SocketAddress)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.envoy.service.auth.v3.SocketAddress.Protocol protocol = 1 [json_name = "protocol", (.validate.rules) = { ... }</code>
   * @return The enum numeric value on the wire for protocol.
   */
  int getProtocolValue();
  /**
   * <code>.envoy.service.auth.v3.SocketAddress.Protocol protocol = 1 [json_name = "protocol", (.validate.rules) = { ... }</code>
   * @return The protocol.
   */
  com.fluxninja.generated.envoy.service.auth.v3.SocketAddress.Protocol getProtocol();

  /**
   * <pre>
   * The address for this socket. :ref:`Listeners &lt;config_listeners&gt;` will bind
   * to the address. An empty address is not allowed. Specify ``0.0.0.0`` or ``::``
   * to bind to any address. [#comment:TODO(zuercher) reinstate when implemented:
   * It is possible to distinguish a Listener address via the prefix/suffix matching
   * in :ref:`FilterChainMatch &lt;envoy_v3_api_msg_config.listener.v3.FilterChainMatch&gt;`.] When used
   * within an upstream :ref:`BindConfig &lt;envoy_v3_api_msg_config.core.v3.BindConfig&gt;`, the address
   * controls the source address of outbound connections. For :ref:`clusters
   * &lt;envoy_v3_api_msg_config.cluster.v3.Cluster&gt;`, the cluster type determines whether the
   * address must be an IP (``STATIC`` or ``EDS`` clusters) or a hostname resolved by DNS
   * (``STRICT_DNS`` or ``LOGICAL_DNS`` clusters). Address resolution can be customized
   * via :ref:`resolver_name &lt;envoy_v3_api_field_config.core.v3.SocketAddress.resolver_name&gt;`.
   * </pre>
   *
   * <code>string address = 2 [json_name = "address", (.validate.rules) = { ... }</code>
   * @return The address.
   */
  java.lang.String getAddress();
  /**
   * <pre>
   * The address for this socket. :ref:`Listeners &lt;config_listeners&gt;` will bind
   * to the address. An empty address is not allowed. Specify ``0.0.0.0`` or ``::``
   * to bind to any address. [#comment:TODO(zuercher) reinstate when implemented:
   * It is possible to distinguish a Listener address via the prefix/suffix matching
   * in :ref:`FilterChainMatch &lt;envoy_v3_api_msg_config.listener.v3.FilterChainMatch&gt;`.] When used
   * within an upstream :ref:`BindConfig &lt;envoy_v3_api_msg_config.core.v3.BindConfig&gt;`, the address
   * controls the source address of outbound connections. For :ref:`clusters
   * &lt;envoy_v3_api_msg_config.cluster.v3.Cluster&gt;`, the cluster type determines whether the
   * address must be an IP (``STATIC`` or ``EDS`` clusters) or a hostname resolved by DNS
   * (``STRICT_DNS`` or ``LOGICAL_DNS`` clusters). Address resolution can be customized
   * via :ref:`resolver_name &lt;envoy_v3_api_field_config.core.v3.SocketAddress.resolver_name&gt;`.
   * </pre>
   *
   * <code>string address = 2 [json_name = "address", (.validate.rules) = { ... }</code>
   * @return The bytes for address.
   */
  com.google.protobuf.ByteString
      getAddressBytes();

  /**
   * <code>uint32 port_value = 3 [json_name = "portValue", (.validate.rules) = { ... }</code>
   * @return Whether the portValue field is set.
   */
  boolean hasPortValue();
  /**
   * <code>uint32 port_value = 3 [json_name = "portValue", (.validate.rules) = { ... }</code>
   * @return The portValue.
   */
  int getPortValue();

  /**
   * <pre>
   * This is only valid if :ref:`resolver_name
   * &lt;envoy_v3_api_field_config.core.v3.SocketAddress.resolver_name&gt;` is specified below and the
   * named resolver is capable of named port resolution.
   * </pre>
   *
   * <code>string named_port = 4 [json_name = "namedPort"];</code>
   * @return Whether the namedPort field is set.
   */
  boolean hasNamedPort();
  /**
   * <pre>
   * This is only valid if :ref:`resolver_name
   * &lt;envoy_v3_api_field_config.core.v3.SocketAddress.resolver_name&gt;` is specified below and the
   * named resolver is capable of named port resolution.
   * </pre>
   *
   * <code>string named_port = 4 [json_name = "namedPort"];</code>
   * @return The namedPort.
   */
  java.lang.String getNamedPort();
  /**
   * <pre>
   * This is only valid if :ref:`resolver_name
   * &lt;envoy_v3_api_field_config.core.v3.SocketAddress.resolver_name&gt;` is specified below and the
   * named resolver is capable of named port resolution.
   * </pre>
   *
   * <code>string named_port = 4 [json_name = "namedPort"];</code>
   * @return The bytes for namedPort.
   */
  com.google.protobuf.ByteString
      getNamedPortBytes();

  /**
   * <pre>
   * The name of the custom resolver. This must have been registered with Envoy. If
   * this is empty, a context dependent default applies. If the address is a concrete
   * IP address, no resolution will occur. If address is a hostname this
   * should be set for resolution other than DNS. Specifying a custom resolver with
   * ``STRICT_DNS`` or ``LOGICAL_DNS`` will generate an error at runtime.
   * </pre>
   *
   * <code>string resolver_name = 5 [json_name = "resolverName"];</code>
   * @return The resolverName.
   */
  java.lang.String getResolverName();
  /**
   * <pre>
   * The name of the custom resolver. This must have been registered with Envoy. If
   * this is empty, a context dependent default applies. If the address is a concrete
   * IP address, no resolution will occur. If address is a hostname this
   * should be set for resolution other than DNS. Specifying a custom resolver with
   * ``STRICT_DNS`` or ``LOGICAL_DNS`` will generate an error at runtime.
   * </pre>
   *
   * <code>string resolver_name = 5 [json_name = "resolverName"];</code>
   * @return The bytes for resolverName.
   */
  com.google.protobuf.ByteString
      getResolverNameBytes();

  /**
   * <pre>
   * When binding to an IPv6 address above, this enables `IPv4 compatibility
   * &lt;https://tools.ietf.org/html/rfc3493#page-11&gt;`_. Binding to ``::`` will
   * allow both IPv4 and IPv6 connections, with peer IPv4 addresses mapped into
   * IPv6 space as ``::FFFF:&lt;IPv4-address&gt;``.
   * </pre>
   *
   * <code>bool ipv4_compat = 6 [json_name = "ipv4Compat"];</code>
   * @return The ipv4Compat.
   */
  boolean getIpv4Compat();

  public com.fluxninja.generated.envoy.service.auth.v3.SocketAddress.PortSpecifierCase getPortSpecifierCase();
}