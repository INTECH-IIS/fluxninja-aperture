// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: aperture/flowcontrol/v1/flowcontrol.proto

package com.fluxninja.generated.aperture.flowcontrol.v1;

public interface LimiterDecisionOrBuilder extends
    // @@protoc_insertion_point(interface_extends:aperture.flowcontrol.v1.LimiterDecision)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string policy_name = 1 [json_name = "policyName"];</code>
   * @return The policyName.
   */
  java.lang.String getPolicyName();
  /**
   * <code>string policy_name = 1 [json_name = "policyName"];</code>
   * @return The bytes for policyName.
   */
  com.google.protobuf.ByteString
      getPolicyNameBytes();

  /**
   * <code>string policy_hash = 2 [json_name = "policyHash"];</code>
   * @return The policyHash.
   */
  java.lang.String getPolicyHash();
  /**
   * <code>string policy_hash = 2 [json_name = "policyHash"];</code>
   * @return The bytes for policyHash.
   */
  com.google.protobuf.ByteString
      getPolicyHashBytes();

  /**
   * <code>int64 component_index = 3 [json_name = "componentIndex"];</code>
   * @return The componentIndex.
   */
  long getComponentIndex();

  /**
   * <code>bool dropped = 4 [json_name = "dropped"];</code>
   * @return The dropped.
   */
  boolean getDropped();

  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.LimiterReason reason = 5 [json_name = "reason"];</code>
   * @return The enum numeric value on the wire for reason.
   */
  int getReasonValue();
  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.LimiterReason reason = 5 [json_name = "reason"];</code>
   * @return The reason.
   */
  com.fluxninja.generated.aperture.flowcontrol.v1.LimiterDecision.LimiterReason getReason();

  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.RateLimiterInfo rate_limiter_info = 6 [json_name = "rateLimiterInfo"];</code>
   * @return Whether the rateLimiterInfo field is set.
   */
  boolean hasRateLimiterInfo();
  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.RateLimiterInfo rate_limiter_info = 6 [json_name = "rateLimiterInfo"];</code>
   * @return The rateLimiterInfo.
   */
  com.fluxninja.generated.aperture.flowcontrol.v1.LimiterDecision.RateLimiterInfo getRateLimiterInfo();
  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.RateLimiterInfo rate_limiter_info = 6 [json_name = "rateLimiterInfo"];</code>
   */
  com.fluxninja.generated.aperture.flowcontrol.v1.LimiterDecision.RateLimiterInfoOrBuilder getRateLimiterInfoOrBuilder();

  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.ConcurrencyLimiterInfo concurrency_limiter_info = 7 [json_name = "concurrencyLimiterInfo"];</code>
   * @return Whether the concurrencyLimiterInfo field is set.
   */
  boolean hasConcurrencyLimiterInfo();
  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.ConcurrencyLimiterInfo concurrency_limiter_info = 7 [json_name = "concurrencyLimiterInfo"];</code>
   * @return The concurrencyLimiterInfo.
   */
  com.fluxninja.generated.aperture.flowcontrol.v1.LimiterDecision.ConcurrencyLimiterInfo getConcurrencyLimiterInfo();
  /**
   * <code>.aperture.flowcontrol.v1.LimiterDecision.ConcurrencyLimiterInfo concurrency_limiter_info = 7 [json_name = "concurrencyLimiterInfo"];</code>
   */
  com.fluxninja.generated.aperture.flowcontrol.v1.LimiterDecision.ConcurrencyLimiterInfoOrBuilder getConcurrencyLimiterInfoOrBuilder();

  public com.fluxninja.generated.aperture.flowcontrol.v1.LimiterDecision.DetailsCase getDetailsCase();
}