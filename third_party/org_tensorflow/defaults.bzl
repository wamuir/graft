def _libtensorflow_defaults_impl(repository_ctx):
    repository_ctx.file("BUILD")
    v = repository_ctx.os.environ.get("LIBTENSORFLOW_PKG_URL", "file:///opt/libtensorflow.tar.gz")
    repository_ctx.file("config.bzl", content = """LIBTENSORFLOW_PKG_URL = {}""".format(repr(v)))

libtensorflow_defaults = repository_rule(
    implementation = _libtensorflow_defaults_impl,
    environ = ["LIBTENSORFLOW_PKG_URL"],
)
