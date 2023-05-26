%% -*- coding: utf-8 -*-
%% Automatically generated, do not edit
%% Generated by gpb_compile version 4.19.7

-ifndef(gpb_descriptor).
-define(gpb_descriptor, true).

-define(gpb_descriptor_gpb_version, "4.19.7").

-ifndef('FILEDESCRIPTORSET_PB_H').
-define('FILEDESCRIPTORSET_PB_H', true).
-record('FileDescriptorSet',
        {file = []              :: [gpb_descriptor:'FileDescriptorProto'()] | undefined % = 1, repeated
        }).
-endif.

-ifndef('FILEDESCRIPTORPROTO_PB_H').
-define('FILEDESCRIPTORPROTO_PB_H', true).
-record('FileDescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         package                :: unicode:chardata() | undefined, % = 2, optional
         dependency = []        :: [unicode:chardata()] | undefined, % = 3, repeated
         public_dependency = [] :: [integer()] | undefined, % = 10, repeated, 32 bits
         weak_dependency = []   :: [integer()] | undefined, % = 11, repeated, 32 bits
         message_type = []      :: [gpb_descriptor:'DescriptorProto'()] | undefined, % = 4, repeated
         enum_type = []         :: [gpb_descriptor:'EnumDescriptorProto'()] | undefined, % = 5, repeated
         service = []           :: [gpb_descriptor:'ServiceDescriptorProto'()] | undefined, % = 6, repeated
         extension = []         :: [gpb_descriptor:'FieldDescriptorProto'()] | undefined, % = 7, repeated
         options                :: gpb_descriptor:'FileOptions'() | undefined, % = 8, optional
         source_code_info       :: gpb_descriptor:'SourceCodeInfo'() | undefined, % = 9, optional
         syntax                 :: unicode:chardata() | undefined % = 12, optional
        }).
-endif.

-ifndef('DESCRIPTORPROTO.EXTENSIONRANGE_PB_H').
-define('DESCRIPTORPROTO.EXTENSIONRANGE_PB_H', true).
-record('DescriptorProto.ExtensionRange',
        {start                  :: integer() | undefined, % = 1, optional, 32 bits
         'end'                  :: integer() | undefined, % = 2, optional, 32 bits
         options                :: gpb_descriptor:'ExtensionRangeOptions'() | undefined % = 3, optional
        }).
-endif.

-ifndef('DESCRIPTORPROTO.RESERVEDRANGE_PB_H').
-define('DESCRIPTORPROTO.RESERVEDRANGE_PB_H', true).
-record('DescriptorProto.ReservedRange',
        {start                  :: integer() | undefined, % = 1, optional, 32 bits
         'end'                  :: integer() | undefined % = 2, optional, 32 bits
        }).
-endif.

-ifndef('DESCRIPTORPROTO_PB_H').
-define('DESCRIPTORPROTO_PB_H', true).
-record('DescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         field = []             :: [gpb_descriptor:'FieldDescriptorProto'()] | undefined, % = 2, repeated
         extension = []         :: [gpb_descriptor:'FieldDescriptorProto'()] | undefined, % = 6, repeated
         nested_type = []       :: [gpb_descriptor:'DescriptorProto'()] | undefined, % = 3, repeated
         enum_type = []         :: [gpb_descriptor:'EnumDescriptorProto'()] | undefined, % = 4, repeated
         extension_range = []   :: [gpb_descriptor:'DescriptorProto.ExtensionRange'()] | undefined, % = 5, repeated
         oneof_decl = []        :: [gpb_descriptor:'OneofDescriptorProto'()] | undefined, % = 8, repeated
         options                :: gpb_descriptor:'MessageOptions'() | undefined, % = 7, optional
         reserved_range = []    :: [gpb_descriptor:'DescriptorProto.ReservedRange'()] | undefined, % = 9, repeated
         reserved_name = []     :: [unicode:chardata()] | undefined % = 10, repeated
        }).
-endif.

-ifndef('EXTENSIONRANGEOPTIONS_PB_H').
-define('EXTENSIONRANGEOPTIONS_PB_H', true).
-record('ExtensionRangeOptions',
        {uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('FIELDDESCRIPTORPROTO_PB_H').
-define('FIELDDESCRIPTORPROTO_PB_H', true).
-record('FieldDescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         number                 :: integer() | undefined, % = 3, optional, 32 bits
         label                  :: 'LABEL_OPTIONAL' | 'LABEL_REQUIRED' | 'LABEL_REPEATED' | integer() | undefined, % = 4, optional, enum FieldDescriptorProto.Label
         type                   :: 'TYPE_DOUBLE' | 'TYPE_FLOAT' | 'TYPE_INT64' | 'TYPE_UINT64' | 'TYPE_INT32' | 'TYPE_FIXED64' | 'TYPE_FIXED32' | 'TYPE_BOOL' | 'TYPE_STRING' | 'TYPE_GROUP' | 'TYPE_MESSAGE' | 'TYPE_BYTES' | 'TYPE_UINT32' | 'TYPE_ENUM' | 'TYPE_SFIXED32' | 'TYPE_SFIXED64' | 'TYPE_SINT32' | 'TYPE_SINT64' | integer() | undefined, % = 5, optional, enum FieldDescriptorProto.Type
         type_name              :: unicode:chardata() | undefined, % = 6, optional
         extendee               :: unicode:chardata() | undefined, % = 2, optional
         default_value          :: unicode:chardata() | undefined, % = 7, optional
         oneof_index            :: integer() | undefined, % = 9, optional, 32 bits
         json_name              :: unicode:chardata() | undefined, % = 10, optional
         options                :: gpb_descriptor:'FieldOptions'() | undefined, % = 8, optional
         proto3_optional        :: boolean() | 0 | 1 | undefined % = 17, optional
        }).
-endif.

-ifndef('ONEOFDESCRIPTORPROTO_PB_H').
-define('ONEOFDESCRIPTORPROTO_PB_H', true).
-record('OneofDescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         options                :: gpb_descriptor:'OneofOptions'() | undefined % = 2, optional
        }).
-endif.

-ifndef('ENUMDESCRIPTORPROTO.ENUMRESERVEDRANGE_PB_H').
-define('ENUMDESCRIPTORPROTO.ENUMRESERVEDRANGE_PB_H', true).
-record('EnumDescriptorProto.EnumReservedRange',
        {start                  :: integer() | undefined, % = 1, optional, 32 bits
         'end'                  :: integer() | undefined % = 2, optional, 32 bits
        }).
-endif.

-ifndef('ENUMDESCRIPTORPROTO_PB_H').
-define('ENUMDESCRIPTORPROTO_PB_H', true).
-record('EnumDescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         value = []             :: [gpb_descriptor:'EnumValueDescriptorProto'()] | undefined, % = 2, repeated
         options                :: gpb_descriptor:'EnumOptions'() | undefined, % = 3, optional
         reserved_range = []    :: [gpb_descriptor:'EnumDescriptorProto.EnumReservedRange'()] | undefined, % = 4, repeated
         reserved_name = []     :: [unicode:chardata()] | undefined % = 5, repeated
        }).
-endif.

-ifndef('ENUMVALUEDESCRIPTORPROTO_PB_H').
-define('ENUMVALUEDESCRIPTORPROTO_PB_H', true).
-record('EnumValueDescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         number                 :: integer() | undefined, % = 2, optional, 32 bits
         options                :: gpb_descriptor:'EnumValueOptions'() | undefined % = 3, optional
        }).
-endif.

-ifndef('SERVICEDESCRIPTORPROTO_PB_H').
-define('SERVICEDESCRIPTORPROTO_PB_H', true).
-record('ServiceDescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         method = []            :: [gpb_descriptor:'MethodDescriptorProto'()] | undefined, % = 2, repeated
         options                :: gpb_descriptor:'ServiceOptions'() | undefined % = 3, optional
        }).
-endif.

-ifndef('METHODDESCRIPTORPROTO_PB_H').
-define('METHODDESCRIPTORPROTO_PB_H', true).
-record('MethodDescriptorProto',
        {name                   :: unicode:chardata() | undefined, % = 1, optional
         input_type             :: unicode:chardata() | undefined, % = 2, optional
         output_type            :: unicode:chardata() | undefined, % = 3, optional
         options                :: gpb_descriptor:'MethodOptions'() | undefined, % = 4, optional
         client_streaming = false :: boolean() | 0 | 1 | undefined, % = 5, optional
         server_streaming = false :: boolean() | 0 | 1 | undefined % = 6, optional
        }).
-endif.

-ifndef('FILEOPTIONS_PB_H').
-define('FILEOPTIONS_PB_H', true).
-record('FileOptions',
        {java_package           :: unicode:chardata() | undefined, % = 1, optional
         java_outer_classname   :: unicode:chardata() | undefined, % = 8, optional
         java_multiple_files = false :: boolean() | 0 | 1 | undefined, % = 10, optional
         java_generate_equals_and_hash :: boolean() | 0 | 1 | undefined, % = 20, optional
         java_string_check_utf8 = false :: boolean() | 0 | 1 | undefined, % = 27, optional
         optimize_for = 'SPEED' :: 'SPEED' | 'CODE_SIZE' | 'LITE_RUNTIME' | integer() | undefined, % = 9, optional, enum FileOptions.OptimizeMode
         go_package             :: unicode:chardata() | undefined, % = 11, optional
         cc_generic_services = false :: boolean() | 0 | 1 | undefined, % = 16, optional
         java_generic_services = false :: boolean() | 0 | 1 | undefined, % = 17, optional
         py_generic_services = false :: boolean() | 0 | 1 | undefined, % = 18, optional
         php_generic_services = false :: boolean() | 0 | 1 | undefined, % = 42, optional
         deprecated = false     :: boolean() | 0 | 1 | undefined, % = 23, optional
         cc_enable_arenas = true :: boolean() | 0 | 1 | undefined, % = 31, optional
         objc_class_prefix      :: unicode:chardata() | undefined, % = 36, optional
         csharp_namespace       :: unicode:chardata() | undefined, % = 37, optional
         swift_prefix           :: unicode:chardata() | undefined, % = 39, optional
         php_class_prefix       :: unicode:chardata() | undefined, % = 40, optional
         php_namespace          :: unicode:chardata() | undefined, % = 41, optional
         php_metadata_namespace :: unicode:chardata() | undefined, % = 44, optional
         ruby_package           :: unicode:chardata() | undefined, % = 45, optional
         uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('MESSAGEOPTIONS_PB_H').
-define('MESSAGEOPTIONS_PB_H', true).
-record('MessageOptions',
        {message_set_wire_format = false :: boolean() | 0 | 1 | undefined, % = 1, optional
         no_standard_descriptor_accessor = false :: boolean() | 0 | 1 | undefined, % = 2, optional
         deprecated = false     :: boolean() | 0 | 1 | undefined, % = 3, optional
         map_entry              :: boolean() | 0 | 1 | undefined, % = 7, optional
         uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('FIELDOPTIONS_PB_H').
-define('FIELDOPTIONS_PB_H', true).
-record('FieldOptions',
        {ctype = 'STRING'       :: 'STRING' | 'CORD' | 'STRING_PIECE' | integer() | undefined, % = 1, optional, enum FieldOptions.CType
         packed                 :: boolean() | 0 | 1 | undefined, % = 2, optional
         jstype = 'JS_NORMAL'   :: 'JS_NORMAL' | 'JS_STRING' | 'JS_NUMBER' | integer() | undefined, % = 6, optional, enum FieldOptions.JSType
         lazy = false           :: boolean() | 0 | 1 | undefined, % = 5, optional
         deprecated = false     :: boolean() | 0 | 1 | undefined, % = 3, optional
         weak = false           :: boolean() | 0 | 1 | undefined, % = 10, optional
         uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('ONEOFOPTIONS_PB_H').
-define('ONEOFOPTIONS_PB_H', true).
-record('OneofOptions',
        {uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('ENUMOPTIONS_PB_H').
-define('ENUMOPTIONS_PB_H', true).
-record('EnumOptions',
        {allow_alias            :: boolean() | 0 | 1 | undefined, % = 2, optional
         deprecated = false     :: boolean() | 0 | 1 | undefined, % = 3, optional
         uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('ENUMVALUEOPTIONS_PB_H').
-define('ENUMVALUEOPTIONS_PB_H', true).
-record('EnumValueOptions',
        {deprecated = false     :: boolean() | 0 | 1 | undefined, % = 1, optional
         uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('SERVICEOPTIONS_PB_H').
-define('SERVICEOPTIONS_PB_H', true).
-record('ServiceOptions',
        {deprecated = false     :: boolean() | 0 | 1 | undefined, % = 33, optional
         uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('METHODOPTIONS_PB_H').
-define('METHODOPTIONS_PB_H', true).
-record('MethodOptions',
        {deprecated = false     :: boolean() | 0 | 1 | undefined, % = 33, optional
         idempotency_level = 'IDEMPOTENCY_UNKNOWN' :: 'IDEMPOTENCY_UNKNOWN' | 'NO_SIDE_EFFECTS' | 'IDEMPOTENT' | integer() | undefined, % = 34, optional, enum MethodOptions.IdempotencyLevel
         uninterpreted_option = [] :: [gpb_descriptor:'UninterpretedOption'()] | undefined % = 999, repeated
        }).
-endif.

-ifndef('UNINTERPRETEDOPTION.NAMEPART_PB_H').
-define('UNINTERPRETEDOPTION.NAMEPART_PB_H', true).
-record('UninterpretedOption.NamePart',
        {name_part              :: unicode:chardata() | undefined, % = 1, required
         is_extension           :: boolean() | 0 | 1 | undefined % = 2, required
        }).
-endif.

-ifndef('UNINTERPRETEDOPTION_PB_H').
-define('UNINTERPRETEDOPTION_PB_H', true).
-record('UninterpretedOption',
        {name = []              :: [gpb_descriptor:'UninterpretedOption.NamePart'()] | undefined, % = 2, repeated
         identifier_value       :: unicode:chardata() | undefined, % = 3, optional
         positive_int_value     :: non_neg_integer() | undefined, % = 4, optional, 64 bits
         negative_int_value     :: integer() | undefined, % = 5, optional, 64 bits
         double_value           :: float() | integer() | infinity | '-infinity' | nan | undefined, % = 6, optional
         string_value           :: iodata() | undefined, % = 7, optional
         aggregate_value        :: unicode:chardata() | undefined % = 8, optional
        }).
-endif.

-ifndef('SOURCECODEINFO.LOCATION_PB_H').
-define('SOURCECODEINFO.LOCATION_PB_H', true).
-record('SourceCodeInfo.Location',
        {path = []              :: [integer()] | undefined, % = 1, repeated, 32 bits
         span = []              :: [integer()] | undefined, % = 2, repeated, 32 bits
         leading_comments       :: unicode:chardata() | undefined, % = 3, optional
         trailing_comments      :: unicode:chardata() | undefined, % = 4, optional
         leading_detached_comments = [] :: [unicode:chardata()] | undefined % = 6, repeated
        }).
-endif.

-ifndef('SOURCECODEINFO_PB_H').
-define('SOURCECODEINFO_PB_H', true).
-record('SourceCodeInfo',
        {location = []          :: [gpb_descriptor:'SourceCodeInfo.Location'()] | undefined % = 1, repeated
        }).
-endif.

-ifndef('GENERATEDCODEINFO.ANNOTATION_PB_H').
-define('GENERATEDCODEINFO.ANNOTATION_PB_H', true).
-record('GeneratedCodeInfo.Annotation',
        {path = []              :: [integer()] | undefined, % = 1, repeated, 32 bits
         source_file            :: unicode:chardata() | undefined, % = 2, optional
         'begin'                :: integer() | undefined, % = 3, optional, 32 bits
         'end'                  :: integer() | undefined % = 4, optional, 32 bits
        }).
-endif.

-ifndef('GENERATEDCODEINFO_PB_H').
-define('GENERATEDCODEINFO_PB_H', true).
-record('GeneratedCodeInfo',
        {annotation = []        :: [gpb_descriptor:'GeneratedCodeInfo.Annotation'()] | undefined % = 1, repeated
        }).
-endif.

-endif.
