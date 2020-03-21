# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: storm.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='storm.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\x0bstorm.proto\" \n\x10SummarizeRequest\x12\x0c\n\x04text\x18\x01 \x01(\t\"$\n\x11SummarizeResponse\x12\x0f\n\x07summary\x18\x01 \x01(\t2B\n\x0cStormService\x12\x32\n\tSummarize\x12\x11.SummarizeRequest\x1a\x12.SummarizeResponseb\x06proto3'
)




_SUMMARIZEREQUEST = _descriptor.Descriptor(
  name='SummarizeRequest',
  full_name='SummarizeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='SummarizeRequest.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=15,
  serialized_end=47,
)


_SUMMARIZERESPONSE = _descriptor.Descriptor(
  name='SummarizeResponse',
  full_name='SummarizeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='summary', full_name='SummarizeResponse.summary', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=49,
  serialized_end=85,
)

DESCRIPTOR.message_types_by_name['SummarizeRequest'] = _SUMMARIZEREQUEST
DESCRIPTOR.message_types_by_name['SummarizeResponse'] = _SUMMARIZERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SummarizeRequest = _reflection.GeneratedProtocolMessageType('SummarizeRequest', (_message.Message,), {
  'DESCRIPTOR' : _SUMMARIZEREQUEST,
  '__module__' : 'storm_pb2'
  # @@protoc_insertion_point(class_scope:SummarizeRequest)
  })
_sym_db.RegisterMessage(SummarizeRequest)

SummarizeResponse = _reflection.GeneratedProtocolMessageType('SummarizeResponse', (_message.Message,), {
  'DESCRIPTOR' : _SUMMARIZERESPONSE,
  '__module__' : 'storm_pb2'
  # @@protoc_insertion_point(class_scope:SummarizeResponse)
  })
_sym_db.RegisterMessage(SummarizeResponse)



_STORMSERVICE = _descriptor.ServiceDescriptor(
  name='StormService',
  full_name='StormService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=87,
  serialized_end=153,
  methods=[
  _descriptor.MethodDescriptor(
    name='Summarize',
    full_name='StormService.Summarize',
    index=0,
    containing_service=None,
    input_type=_SUMMARIZEREQUEST,
    output_type=_SUMMARIZERESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_STORMSERVICE)

DESCRIPTOR.services_by_name['StormService'] = _STORMSERVICE

# @@protoc_insertion_point(module_scope)
