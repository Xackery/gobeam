// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: grpc.proto
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Gobeam {

  /// <summary>Holder for reflection information generated from grpc.proto</summary>
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute()]
  public static partial class GrpcReflection {

    #region Descriptor
    /// <summary>File descriptor for grpc.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static GrpcReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgpncnBjLnByb3RvEgZnb2JlYW0aCmJlYW0ucHJvdG8iDwoNU3RyZWFtUmVx",
            "dWVzdCJHChVQcm9ncmVzc1VwZGF0ZVJlcXVlc3QSLgoOUHJvZ3Jlc3NVcGRh",
            "dGUYASABKAsyFi5nb2JlYW0uUHJvZ3Jlc3NVcGRhdGUiGAoWUHJvZ3Jlc3NV",
            "cGRhdGVSZXNwb25zZTKcAQoMUm9ib3RTZXJ2aWNlEjkKDFN0cmVhbVJlcG9y",
            "dBIVLmdvYmVhbS5TdHJlYW1SZXF1ZXN0Gg4uZ29iZWFtLlJlcG9ydCIAMAES",
            "UQoOUHJvZ3Jlc3NVcGRhdGUSHS5nb2JlYW0uUHJvZ3Jlc3NVcGRhdGVSZXF1",
            "ZXN0Gh4uZ29iZWFtLlByb2dyZXNzVXBkYXRlUmVzcG9uc2UiAGIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Gobeam.BeamReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Gobeam.StreamRequest), global::Gobeam.StreamRequest.Parser, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Gobeam.ProgressUpdateRequest), global::Gobeam.ProgressUpdateRequest.Parser, new[]{ "ProgressUpdate" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Gobeam.ProgressUpdateResponse), global::Gobeam.ProgressUpdateResponse.Parser, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  /// Used by RobotService StreamReport
  /// </summary>
  [global::System.Diagnostics.DebuggerNonUserCodeAttribute()]
  public sealed partial class StreamRequest : pb::IMessage<StreamRequest> {
    private static readonly pb::MessageParser<StreamRequest> _parser = new pb::MessageParser<StreamRequest>(() => new StreamRequest());
    public static pb::MessageParser<StreamRequest> Parser { get { return _parser; } }

    public static pbr::MessageDescriptor Descriptor {
      get { return global::Gobeam.GrpcReflection.Descriptor.MessageTypes[0]; }
    }

    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    public StreamRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    public StreamRequest(StreamRequest other) : this() {
    }

    public StreamRequest Clone() {
      return new StreamRequest(this);
    }

    public override bool Equals(object other) {
      return Equals(other as StreamRequest);
    }

    public bool Equals(StreamRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      return true;
    }

    public override int GetHashCode() {
      int hash = 1;
      return hash;
    }

    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    public void WriteTo(pb::CodedOutputStream output) {
    }

    public int CalculateSize() {
      int size = 0;
      return size;
    }

    public void MergeFrom(StreamRequest other) {
      if (other == null) {
        return;
      }
    }

    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
        }
      }
    }

  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute()]
  public sealed partial class ProgressUpdateRequest : pb::IMessage<ProgressUpdateRequest> {
    private static readonly pb::MessageParser<ProgressUpdateRequest> _parser = new pb::MessageParser<ProgressUpdateRequest>(() => new ProgressUpdateRequest());
    public static pb::MessageParser<ProgressUpdateRequest> Parser { get { return _parser; } }

    public static pbr::MessageDescriptor Descriptor {
      get { return global::Gobeam.GrpcReflection.Descriptor.MessageTypes[1]; }
    }

    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    public ProgressUpdateRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    public ProgressUpdateRequest(ProgressUpdateRequest other) : this() {
      ProgressUpdate = other.progressUpdate_ != null ? other.ProgressUpdate.Clone() : null;
    }

    public ProgressUpdateRequest Clone() {
      return new ProgressUpdateRequest(this);
    }

    /// <summary>Field number for the "ProgressUpdate" field.</summary>
    public const int ProgressUpdateFieldNumber = 1;
    private global::Gobeam.ProgressUpdate progressUpdate_;
    public global::Gobeam.ProgressUpdate ProgressUpdate {
      get { return progressUpdate_; }
      set {
        progressUpdate_ = value;
      }
    }

    public override bool Equals(object other) {
      return Equals(other as ProgressUpdateRequest);
    }

    public bool Equals(ProgressUpdateRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(ProgressUpdate, other.ProgressUpdate)) return false;
      return true;
    }

    public override int GetHashCode() {
      int hash = 1;
      if (progressUpdate_ != null) hash ^= ProgressUpdate.GetHashCode();
      return hash;
    }

    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    public void WriteTo(pb::CodedOutputStream output) {
      if (progressUpdate_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(ProgressUpdate);
      }
    }

    public int CalculateSize() {
      int size = 0;
      if (progressUpdate_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(ProgressUpdate);
      }
      return size;
    }

    public void MergeFrom(ProgressUpdateRequest other) {
      if (other == null) {
        return;
      }
      if (other.progressUpdate_ != null) {
        if (progressUpdate_ == null) {
          progressUpdate_ = new global::Gobeam.ProgressUpdate();
        }
        ProgressUpdate.MergeFrom(other.ProgressUpdate);
      }
    }

    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
          case 10: {
            if (progressUpdate_ == null) {
              progressUpdate_ = new global::Gobeam.ProgressUpdate();
            }
            input.ReadMessage(progressUpdate_);
            break;
          }
        }
      }
    }

  }

  [global::System.Diagnostics.DebuggerNonUserCodeAttribute()]
  public sealed partial class ProgressUpdateResponse : pb::IMessage<ProgressUpdateResponse> {
    private static readonly pb::MessageParser<ProgressUpdateResponse> _parser = new pb::MessageParser<ProgressUpdateResponse>(() => new ProgressUpdateResponse());
    public static pb::MessageParser<ProgressUpdateResponse> Parser { get { return _parser; } }

    public static pbr::MessageDescriptor Descriptor {
      get { return global::Gobeam.GrpcReflection.Descriptor.MessageTypes[2]; }
    }

    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    public ProgressUpdateResponse() {
      OnConstruction();
    }

    partial void OnConstruction();

    public ProgressUpdateResponse(ProgressUpdateResponse other) : this() {
    }

    public ProgressUpdateResponse Clone() {
      return new ProgressUpdateResponse(this);
    }

    public override bool Equals(object other) {
      return Equals(other as ProgressUpdateResponse);
    }

    public bool Equals(ProgressUpdateResponse other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      return true;
    }

    public override int GetHashCode() {
      int hash = 1;
      return hash;
    }

    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    public void WriteTo(pb::CodedOutputStream output) {
    }

    public int CalculateSize() {
      int size = 0;
      return size;
    }

    public void MergeFrom(ProgressUpdateResponse other) {
      if (other == null) {
        return;
      }
    }

    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            input.SkipLastField();
            break;
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
