load("//tools:defs.bzl", "go_library", "go_test")
load("//tools/go_generics:defs.bzl", "go_template_instance")

package(
    default_applicable_licenses = ["//:license"],
    licenses = ["notice"],
)

go_library(
    name = "bufferv2",
    srcs = [
        "buffer.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/buffer",
    ],
)
