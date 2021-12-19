load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")

cc_library(
    name = "hash",
    srcs = ["Hash.cpp"],
    hdrs = ["Hash.h"],
    visibility = ["//2017/14:__pkg__"],
)

cc_library(
    name = "hash_internal",
    srcs = ["Hash.cpp"],
    hdrs = [
        "Hash.h",
        "HashInternal.h",
    ],
    visibility = ["//visibility:private"],
)

cc_binary(
    name = "10",
    srcs = ["Day10.cpp"],
    deps = [
        "hash_internal",
        "//2017/lib",
    ],
)
