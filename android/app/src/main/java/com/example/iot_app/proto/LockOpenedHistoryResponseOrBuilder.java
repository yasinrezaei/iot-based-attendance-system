// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: service.proto

package com.example.iot_app.proto;

public interface LockOpenedHistoryResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   *  string first_name = 2;
   *  string last_name = 3 ;
   * </pre>
   *
   * <code>repeated .ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.history histories = 1;</code>
   */
  java.util.List<LockOpenedHistoryResponse.history>
      getHistoriesList();
  /**
   * <pre>
   *  string first_name = 2;
   *  string last_name = 3 ;
   * </pre>
   *
   * <code>repeated .ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.history histories = 1;</code>
   */
  LockOpenedHistoryResponse.history getHistories(int index);
  /**
   * <pre>
   *  string first_name = 2;
   *  string last_name = 3 ;
   * </pre>
   *
   * <code>repeated .ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.history histories = 1;</code>
   */
  int getHistoriesCount();
  /**
   * <pre>
   *  string first_name = 2;
   *  string last_name = 3 ;
   * </pre>
   *
   * <code>repeated .ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.history histories = 1;</code>
   */
  java.util.List<? extends LockOpenedHistoryResponse.historyOrBuilder>
      getHistoriesOrBuilderList();
  /**
   * <pre>
   *  string first_name = 2;
   *  string last_name = 3 ;
   * </pre>
   *
   * <code>repeated .ir.mahmoud.iot_attendance_system.LockOpenedHistoryResponse.history histories = 1;</code>
   */
  LockOpenedHistoryResponse.historyOrBuilder getHistoriesOrBuilder(
      int index);
}