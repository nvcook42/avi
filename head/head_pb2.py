# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: head/head.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)




DESCRIPTOR = _descriptor.FileDescriptor(
  name='head/head.proto',
  package='head',
  serialized_pb='\n\x0fhead/head.proto\x12\x04head\")\n\x06Vector\x12\t\n\x01x\x18\x01 \x02(\x02\x12\t\n\x01y\x18\x02 \x02(\x02\x12\t\n\x01z\x18\x03 \x02(\x02\"o\n\x06Object\x12\n\n\x02ID\x18\x01 \x02(\x03\x12\x19\n\x03pos\x18\x02 \x02(\x0b\x32\x0c.head.Vector\x12\x0e\n\x06radius\x18\x03 \x02(\x02\x12\x1a\n\x03tex\x18\x04 \x02(\x0e\x32\r.head.Texture\x12\x12\n\ntex_custom\x18\x05 \x01(\t\"%\n\x05\x46rame\x12\x1c\n\x06object\x18\x01 \x03(\x0b\x32\x0c.head.Object\"$\n\x06Stream\x12\x1a\n\x05\x66rame\x18\x02 \x03(\x0b\x32\x0b.head.Frame*D\n\x07Texture\x12\x08\n\x04SHIP\x10\x00\x12\x0c\n\x08\x41STEROID\x10\x01\x12\x11\n\rCONTROL_POINT\x10\x02\x12\x0e\n\nPROJECTILE\x10\x03')

_TEXTURE = _descriptor.EnumDescriptor(
  name='Texture',
  full_name='head.Texture',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SHIP', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ASTEROID', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONTROL_POINT', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROJECTILE', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=258,
  serialized_end=326,
)

Texture = enum_type_wrapper.EnumTypeWrapper(_TEXTURE)
SHIP = 0
ASTEROID = 1
CONTROL_POINT = 2
PROJECTILE = 3



_VECTOR = _descriptor.Descriptor(
  name='Vector',
  full_name='head.Vector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='x', full_name='head.Vector.x', index=0,
      number=1, type=2, cpp_type=6, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='y', full_name='head.Vector.y', index=1,
      number=2, type=2, cpp_type=6, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='z', full_name='head.Vector.z', index=2,
      number=3, type=2, cpp_type=6, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=25,
  serialized_end=66,
)


_OBJECT = _descriptor.Descriptor(
  name='Object',
  full_name='head.Object',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ID', full_name='head.Object.ID', index=0,
      number=1, type=3, cpp_type=2, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='pos', full_name='head.Object.pos', index=1,
      number=2, type=11, cpp_type=10, label=2,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='radius', full_name='head.Object.radius', index=2,
      number=3, type=2, cpp_type=6, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tex', full_name='head.Object.tex', index=3,
      number=4, type=14, cpp_type=8, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tex_custom', full_name='head.Object.tex_custom', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=68,
  serialized_end=179,
)


_FRAME = _descriptor.Descriptor(
  name='Frame',
  full_name='head.Frame',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='object', full_name='head.Frame.object', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=181,
  serialized_end=218,
)


_STREAM = _descriptor.Descriptor(
  name='Stream',
  full_name='head.Stream',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='frame', full_name='head.Stream.frame', index=0,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=220,
  serialized_end=256,
)

_OBJECT.fields_by_name['pos'].message_type = _VECTOR
_OBJECT.fields_by_name['tex'].enum_type = _TEXTURE
_FRAME.fields_by_name['object'].message_type = _OBJECT
_STREAM.fields_by_name['frame'].message_type = _FRAME
DESCRIPTOR.message_types_by_name['Vector'] = _VECTOR
DESCRIPTOR.message_types_by_name['Object'] = _OBJECT
DESCRIPTOR.message_types_by_name['Frame'] = _FRAME
DESCRIPTOR.message_types_by_name['Stream'] = _STREAM

class Vector(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _VECTOR

  # @@protoc_insertion_point(class_scope:head.Vector)

class Object(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _OBJECT

  # @@protoc_insertion_point(class_scope:head.Object)

class Frame(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _FRAME

  # @@protoc_insertion_point(class_scope:head.Frame)

class Stream(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _STREAM

  # @@protoc_insertion_point(class_scope:head.Stream)


# @@protoc_insertion_point(module_scope)
