load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix testpubsub
# gazelle:proto disable_global
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "testpubsub",
    visibility = ["//visibility:private"],
    deps = ["@com_google_cloud_go_firestore//:go_default_library"],
)

go_binary(
    name = "test",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
