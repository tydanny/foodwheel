load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "cmd",
    srcs = [
        "config.go",
        "main.go",
        "wire.go",
    ],
    importpath = "github.com/tydanny/foodwheel/cmd",
    visibility = ["//visibility:public"],
    deps = [
        "//gen/cuisines/v1:cuisines",
        "//internal/app/foodwheel/cuisine/cuisineadapter/bun",
        "//internal/cuisineserver",
        "//internal/log",
        "@com_github_google_wire//:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
        "@com_github_oklog_run//:go_default_library",
        "@com_github_spf13_pflag//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
        "@dev_emperror_emperror//:go_default_library",
        "@dev_emperror_errors//:go_default_library",
        "@org_golang_google_grpc//:grpc",
    ],
)

go_library(
    name = "cmd_lib",
    srcs = [
        "config.go",
        "main.go",
        "wire.go",
    ],
    importpath = "github.com/tydanny/foowheel/cmd",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_google_wire//:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
        "@com_github_oklog_run//:go_default_library",
        "@com_github_spf13_pflag//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
        "@dev_emperror_emperror//:go_default_library",
        "@dev_emperror_errors//:go_default_library",
        "@org_golang_google_grpc//:grpc",
    ],
)
