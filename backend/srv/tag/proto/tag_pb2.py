# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tag.proto
import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='tag.proto',
  package='tag',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\ttag.proto\x12\x03tag\"!\n\nTagRequest\x12\x13\n\x0b\x64\x65scription\x18\x01 \x01(\t\"\x1b\n\x0bTagResponse\x12\x0c\n\x04tags\x18\x01 \x03(\t2/\n\x03Tag\x12(\n\x03Get\x12\x0f.tag.TagRequest\x1a\x10.tag.TagResponseb\x06proto3')
)




_TAGREQUEST = _descriptor.Descriptor(
  name='TagRequest',
  full_name='tag.TagRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='description', full_name='tag.TagRequest.description', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=18,
  serialized_end=51,
)


_TAGRESPONSE = _descriptor.Descriptor(
  name='TagResponse',
  full_name='tag.TagResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='tags', full_name='tag.TagResponse.tags', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=53,
  serialized_end=80,
)

DESCRIPTOR.message_types_by_name['TagRequest'] = _TAGREQUEST
DESCRIPTOR.message_types_by_name['TagResponse'] = _TAGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TagRequest = _reflection.GeneratedProtocolMessageType('TagRequest', (_message.Message,), {
  'DESCRIPTOR' : _TAGREQUEST,
  '__module__' : 'tag_pb2'
  # @@protoc_insertion_point(class_scope:tag.TagRequest)
  })
_sym_db.RegisterMessage(TagRequest)

TagResponse = _reflection.GeneratedProtocolMessageType('TagResponse', (_message.Message,), {
  'DESCRIPTOR' : _TAGRESPONSE,
  '__module__' : 'tag_pb2'
  # @@protoc_insertion_point(class_scope:tag.TagResponse)
  })
_sym_db.RegisterMessage(TagResponse)



_TAG = _descriptor.ServiceDescriptor(
  name='Tag',
  full_name='tag.Tag',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=82,
  serialized_end=129,
  methods=[
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='tag.Tag.Get',
    index=0,
    containing_service=None,
    input_type=_TAGREQUEST,
    output_type=_TAGRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_TAG)

DESCRIPTOR.services_by_name['Tag'] = _TAG

# @@protoc_insertion_point(module_scope)
