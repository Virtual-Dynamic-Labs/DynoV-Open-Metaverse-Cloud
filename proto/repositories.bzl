"""
Dependencies that are needed for proto tests and tools.
"""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")

def proto_deps():
    """
    Fetches all required dependencies for proto tests and tools.
    """
    COM_GOOGLE_PROTOBUF_SHA = "d7d204a59fd0d2d2387bd362c2155289d5060f32122c4d1d922041b61191d522"
    COM_GOOGLE_PROTOBUF_VERSION = "3.21.5"

    maybe(
        http_archive,
        name = "com_google_protobuf",
        sha256 = COM_GOOGLE_PROTOBUF_SHA,
        strip_prefix = "protobuf-%s" % COM_GOOGLE_PROTOBUF_VERSION,
        urls = [
            "https://github.com/protocolbuffers/protobuf/archive/v%s.tar.gz" % COM_GOOGLE_PROTOBUF_VERSION
        ],
    )

    RULES_PROTO_SHA = "e017528fd1c91c5a33f15493e3a398181a9e821a804eb7ff5acdd1d2d6c2b18d"
    RULES_PROTO_VERSION = "4.0.0-3.20.0"

    maybe(
        http_archive,
        name = "rules_proto",
        sha256 = RULES_PROTO_SHA,
        strip_prefix = "rules_proto-%s" % RULES_PROTO_VERSION,
        urls = [
            "https://github.com/bazelbuild/rules_proto/archive/refs/tags/%s.tar.gz" % RULES_PROTO_VERSION,
        ],
    )