# Copyright 2016 The Bazel Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# An adaptation of @bazel_tools//tools/build_defs/repo:http.bzl

"""Rule for downloading embdedded archives over HTTP.

### Setup

To use this rule, load it in your `WORKSPACE` file as follows:

```python
load(
    "@com_github_wamuir_graft//tools/build_defs/repo:http.bzl",
    "http_embedded_archive",
)
```
"""

load(
    "@bazel_tools//tools/build_defs/repo:utils.bzl",
    "patch",
    "read_netrc",
    "update_attrs",
    "use_netrc",
    "workspace_and_buildfile",
)

# Shared between http_jar, http_file and http_archive.
_AUTH_PATTERN_DOC = """An optional dict mapping host names to custom authorization patterns.

If a URL's host name is present in this dict the value will be used as a pattern when
generating the authorization header for the http request. This enables the use of custom
authorization schemes used in a lot of common cloud storage providers.

The pattern currently supports 2 tokens: <code>&lt;login&gt;</code> and
<code>&lt;password&gt;</code>, which are replaced with their equivalent value
in the netrc file for the same host name. After formatting, the result is set
as the value for the <code>Authorization</code> field of the HTTP request.

Example attribute and netrc for a http download to an oauth2 enabled API using a bearer token:

<pre>
auth_patterns = {
    "storage.cloudprovider.com": "Bearer &lt;password&gt;"
}
</pre>

netrc:

<pre>
machine storage.cloudprovider.com
        password RANDOM-TOKEN
</pre>

The final HTTP request would have the following header:

<pre>
Authorization: Bearer RANDOM-TOKEN
</pre>
"""

def _get_auth(ctx, urls):
    """Given the list of URLs obtain the correct auth dict."""
    if ctx.attr.netrc:
        netrc = read_netrc(ctx, ctx.attr.netrc)
        return use_netrc(netrc, urls, ctx.attr.auth_patterns)

    if "HOME" in ctx.os.environ and not ctx.os.name.startswith("windows"):
        netrcfile = "%s/.netrc" % (ctx.os.environ["HOME"])
        if ctx.execute(["test", "-f", netrcfile]).return_code == 0:
            netrc = read_netrc(ctx, netrcfile)
            return use_netrc(netrc, urls, ctx.attr.auth_patterns)

    if "USERPROFILE" in ctx.os.environ and ctx.os.name.startswith("windows"):
        netrcfile = "%s/.netrc" % (ctx.os.environ["USERPROFILE"])
        if ctx.path(netrcfile).exists:
            netrc = read_netrc(ctx, netrcfile)
            return use_netrc(netrc, urls, ctx.attr.auth_patterns)

    return {}

def _http_embedded_archive_impl(ctx):
    """Implementation of a http_embedded archive rule for archives-in-archives."""
    if not ctx.attr.url and not ctx.attr.urls:
        fail("At least one of url and urls must be provided")
    if ctx.attr.build_file and ctx.attr.build_file_content:
        fail("Only one of build_file and build_file_content can be provided.")

    all_urls = []
    if ctx.attr.urls:
        all_urls = ctx.attr.urls
    if ctx.attr.url:
        all_urls = [ctx.attr.url] + all_urls

    auth = _get_auth(ctx, all_urls)

    tmpdir = "_bzl.tmp"

    download_info = ctx.download_and_extract(
        all_urls,
        tmpdir,
        ctx.attr.sha256,
        ctx.attr.type,
        "",
        canonical_id = ctx.attr.canonical_id,
        auth = auth,
    )

    ctx.extract(
        "/".join([tmpdir, ctx.attr.inner_archive]),
        "",
        ctx.attr.strip_prefix,
    )
    ctx.delete(tmpdir)

    workspace_and_buildfile(ctx)
    patch(ctx)

    return update_attrs(ctx.attr, _http_embedded_archive_attrs.keys(), {"sha256": download_info.sha256})

_http_embedded_archive_attrs = {
    "url": attr.string(
        doc =
            """A URL to a file that will be made available to Bazel.

This must be a file, http or https URL. Redirections are followed.
Authentication is not supported.

This parameter is to simplify the transition from the native http_archive
rule. More flexibility can be achieved by the urls parameter that allows
to specify alternative URLs to fetch from.
""",
    ),
    "urls": attr.string_list(
        doc =
            """A list of URLs to a file that will be made available to Bazel.

Each entry must be a file, http or https URL. Redirections are followed.
Authentication is not supported.

URLs are tried in order until one succeeds, so you should list local mirrors first.
If all downloads fail, the rule will fail.""",
    ),
    "sha256": attr.string(
        doc = """The expected SHA-256 of the file downloaded.

This must match the SHA-256 of the file downloaded. _It is a security risk
to omit the SHA-256 as remote files can change._ At best omitting this
field will make your build non-hermetic. It is optional to make development
easier but should be set before shipping.""",
    ),
    "netrc": attr.string(
        doc = "Location of the .netrc file to use for authentication",
    ),
    "auth_patterns": attr.string_dict(
        doc = _AUTH_PATTERN_DOC,
    ),
    "canonical_id": attr.string(
        doc = """A canonical id of the archive downloaded.

If specified and non-empty, bazel will not take the archive from cache,
unless it was added to the cache by a request with the same canonical id.
""",
    ),
    "inner_archive": attr.string(
        doc = """The inner (embedded) archive to extract.""",
        mandatory = True,
    ),
    "strip_prefix": attr.string(
        doc = """A directory prefix to strip from extracted files in the inner
archive.

Many archives contain a top-level directory that contains all of the useful
files in archive. Instead of needing to specify this prefix over and over
in the `build_file`, this field can be used to strip it from all of the
extracted files.

For example, suppose you are using `foo-lib-latest.zip`, which contains the
directory `foo-lib-1.2.3/` under which there is a `WORKSPACE` file and are
`src/`, `lib/`, and `test/` directories that contain the actual code you
wish to build. Specify `strip_prefix = "foo-lib-1.2.3"` to use the
`foo-lib-1.2.3` directory as your top-level directory.

Note that if there are files outside of this directory, they will be
discarded and inaccessible (e.g., a top-level license file). This includes
files/directories that start with the prefix but are not in the directory
(e.g., `foo-lib-1.2.3.release-notes`). If the specified prefix does not
match a directory in the archive, Bazel will return an error.""",
    ),
    "type": attr.string(
        doc = """The archive type of the downloaded file.

By default, the archive type is determined from the file extension of the
URL. If the file has no extension, you can explicitly specify one of the
following: `"zip"`, `"jar"`, `"war"`, `"aar"`, `"tar"`, `"tar.gz"`, `"tgz"`,
`"tar.xz"`, or `tar.bz2`.""",
    ),
    "patches": attr.label_list(
        default = [],
        doc =
            "A list of files that are to be applied as patches after " +
            "extracting the archive. By default, it uses the Bazel-native patch implementation " +
            "which doesn't support fuzz match and binary patch, but Bazel will fall back to use " +
            "patch command line tool if `patch_tool` attribute is specified or there are " +
            "arguments other than `-p` in `patch_args` attribute.",
    ),
    "remote_patches": attr.string_dict(
        default = {},
        doc =
            "A map of patch file URL to its integrity value, they are applied after extracting " +
            "the archive and before applying patch files from the `patches` attribute. " +
            "It uses the Bazel-native patch implementation, you can specify the patch strip " +
            "number with `remote_patch_strip`",
    ),
    "remote_patch_strip": attr.int(
        default = 0,
        doc =
            "The number of leading slashes to be stripped from the file name in the remote patches.",
    ),
    "patch_tool": attr.string(
        default = "",
        doc = "The patch(1) utility to use. If this is specified, Bazel will use the specifed " +
              "patch tool instead of the Bazel-native patch implementation.",
    ),
    "patch_args": attr.string_list(
        default = ["-p0"],
        doc =
            "The arguments given to the patch tool. Defaults to -p0, " +
            "however -p1 will usually be needed for patches generated by " +
            "git. If multiple -p arguments are specified, the last one will take effect." +
            "If arguments other than -p are specified, Bazel will fall back to use patch " +
            "command line tool instead of the Bazel-native patch implementation. When falling " +
            "back to patch command line tool and patch_tool attribute is not specified, " +
            "`patch` will be used. This only affects patch files in the `patches` attribute.",
    ),
    "patch_cmds": attr.string_list(
        default = [],
        doc = "Sequence of Bash commands to be applied on Linux/Macos after patches are applied.",
    ),
    "patch_cmds_win": attr.string_list(
        default = [],
        doc = "Sequence of Powershell commands to be applied on Windows after patches are " +
              "applied. If this attribute is not set, patch_cmds will be executed on Windows, " +
              "which requires Bash binary to exist.",
    ),
    "build_file": attr.label(
        allow_single_file = True,
        doc =
            "The file to use as the BUILD file for this repository." +
            "This attribute is an absolute label (use '@//' for the main " +
            "repo). The file does not need to be named BUILD, but can " +
            "be (something like BUILD.new-repo-name may work well for " +
            "distinguishing it from the repository's actual BUILD files. " +
            "Either build_file or build_file_content can be specified, but " +
            "not both.",
    ),
    "build_file_content": attr.string(
        doc =
            "The content for the BUILD file for this repository. " +
            "Either build_file or build_file_content can be specified, but " +
            "not both.",
    ),
    "workspace_file": attr.label(
        doc =
            "The file to use as the `WORKSPACE` file for this repository. " +
            "Either `workspace_file` or `workspace_file_content` can be " +
            "specified, or neither, but not both.",
    ),
    "workspace_file_content": attr.string(
        doc =
            "The content for the WORKSPACE file for this repository. " +
            "Either `workspace_file` or `workspace_file_content` can be " +
            "specified, or neither, but not both.",
    ),
}

http_embedded_archive = repository_rule(
    implementation = _http_embedded_archive_impl,
    attrs = _http_embedded_archive_attrs,
    doc =
        """Downloads a Bazel repository as a compressed archive file, extracts
and decompresses an inner (embedded) archive, and makes targets available for
binding.

It supports the following file extensions: `"zip"`, `"jar"`, `"war"`, `"aar"`,
`"tar"`, `"tar.gz"`, `"tgz"`, `"tar.xz"`, and `tar.bz2`.

Examples:
  Suppose the current repository contains the source code for a chat program,
  rooted at the directory `~/chat-app`. It needs to depend on an SSL library
  which is available from http://example.com/bundle.tar.gz. This `.tar.gz` file
  contains several archive files, including `.zip` file for the SSL library
  at `foo/bar/openssl.zip` which itself has the following directory structure:

  ```
  WORKSPACE
  src/
    openssl.cc
    openssl.h
  ```

  In the local repository, the user creates a `openssl.BUILD` file which
  contains the following target definition:

  ```python
  cc_library(
      name = "openssl-lib",
      srcs = ["src/openssl.cc"],
      hdrs = ["src/openssl.h"],
  )
  ```

  Targets in the `~/chat-app` repository can depend on this target if the
  following lines are added to `~/chat-app/WORKSPACE`:

  ```python
  load(
      "@com_github_wamuir_graft//tools/build_defs/repo:http.bzl",
      "http_embedded_archive",
   )

  http_embedded_archive(
      name = "my_ssl",
      urls = ["http://example.com/bundle.tar.gz"],
      sha256 = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
      inner_archive = "foo/bar/openssl.zip",
      build_file = "@//:openssl.BUILD",
  )
  ```

  Then targets would specify `@my_ssl//:openssl-lib` as a dependency.
""",
)
