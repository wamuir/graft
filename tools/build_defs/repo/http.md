<!-- Generated with Stardoc: http://skydoc.bazel.build -->

Rule for downloading embdedded archives over HTTP.

### Setup

To use this rule, load it in your `WORKSPACE` file as follows:

```python
load(
    "@com_github_wamuir_graft//tools/build_defs/repo:http.bzl",
    "http_embedded_archive",
)
```


<a id="#http_embedded_archive"></a>

## http_embedded_archive

<pre>
http_embedded_archive(<a href="#http_embedded_archive-name">name</a>, <a href="#http_embedded_archive-auth_patterns">auth_patterns</a>, <a href="#http_embedded_archive-build_file">build_file</a>, <a href="#http_embedded_archive-build_file_content">build_file_content</a>, <a href="#http_embedded_archive-canonical_id">canonical_id</a>,
                      <a href="#http_embedded_archive-inner_archive">inner_archive</a>, <a href="#http_embedded_archive-netrc">netrc</a>, <a href="#http_embedded_archive-patch_args">patch_args</a>, <a href="#http_embedded_archive-patch_cmds">patch_cmds</a>, <a href="#http_embedded_archive-patch_cmds_win">patch_cmds_win</a>, <a href="#http_embedded_archive-patch_tool">patch_tool</a>,
                      <a href="#http_embedded_archive-patches">patches</a>, <a href="#http_embedded_archive-remote_patch_strip">remote_patch_strip</a>, <a href="#http_embedded_archive-remote_patches">remote_patches</a>, <a href="#http_embedded_archive-repo_mapping">repo_mapping</a>, <a href="#http_embedded_archive-sha256">sha256</a>, <a href="#http_embedded_archive-strip_prefix">strip_prefix</a>,
                      <a href="#http_embedded_archive-type">type</a>, <a href="#http_embedded_archive-url">url</a>, <a href="#http_embedded_archive-urls">urls</a>, <a href="#http_embedded_archive-workspace_file">workspace_file</a>, <a href="#http_embedded_archive-workspace_file_content">workspace_file_content</a>)
</pre>

Downloads a Bazel repository as a compressed archive file, extracts
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


**ATTRIBUTES**


| Name  | Description | Type | Mandatory | Default |
| :------------- | :------------- | :------------- | :------------- | :------------- |
| <a id="http_embedded_archive-name"></a>name |  A unique name for this repository.   | <a href="https://bazel.build/docs/build-ref.html#name">Name</a> | required |  |
| <a id="http_embedded_archive-auth_patterns"></a>auth_patterns |  An optional dict mapping host names to custom authorization patterns.<br><br>If a URL's host name is present in this dict the value will be used as a pattern when generating the authorization header for the http request. This enables the use of custom authorization schemes used in a lot of common cloud storage providers.<br><br>The pattern currently supports 2 tokens: &lt;code&gt;&lt;login&gt;&lt;/code&gt; and &lt;code&gt;&lt;password&gt;&lt;/code&gt;, which are replaced with their equivalent value in the netrc file for the same host name. After formatting, the result is set as the value for the &lt;code&gt;Authorization&lt;/code&gt; field of the HTTP request.<br><br>Example attribute and netrc for a http download to an oauth2 enabled API using a bearer token:<br><br>&lt;pre&gt; auth_patterns = {     "storage.cloudprovider.com": "Bearer &lt;password&gt;" } &lt;/pre&gt;<br><br>netrc:<br><br>&lt;pre&gt; machine storage.cloudprovider.com         password RANDOM-TOKEN &lt;/pre&gt;<br><br>The final HTTP request would have the following header:<br><br>&lt;pre&gt; Authorization: Bearer RANDOM-TOKEN &lt;/pre&gt;   | <a href="https://bazel.build/docs/skylark/lib/dict.html">Dictionary: String -> String</a> | optional | {} |
| <a id="http_embedded_archive-build_file"></a>build_file |  The file to use as the BUILD file for this repository.This attribute is an absolute label (use '@//' for the main repo). The file does not need to be named BUILD, but can be (something like BUILD.new-repo-name may work well for distinguishing it from the repository's actual BUILD files. Either build_file or build_file_content can be specified, but not both.   | <a href="https://bazel.build/docs/build-ref.html#labels">Label</a> | optional | None |
| <a id="http_embedded_archive-build_file_content"></a>build_file_content |  The content for the BUILD file for this repository. Either build_file or build_file_content can be specified, but not both.   | String | optional | "" |
| <a id="http_embedded_archive-canonical_id"></a>canonical_id |  A canonical id of the archive downloaded.<br><br>If specified and non-empty, bazel will not take the archive from cache, unless it was added to the cache by a request with the same canonical id.   | String | optional | "" |
| <a id="http_embedded_archive-inner_archive"></a>inner_archive |  The inner (embedded) archive to extract.   | String | required |  |
| <a id="http_embedded_archive-netrc"></a>netrc |  Location of the .netrc file to use for authentication   | String | optional | "" |
| <a id="http_embedded_archive-patch_args"></a>patch_args |  The arguments given to the patch tool. Defaults to -p0, however -p1 will usually be needed for patches generated by git. If multiple -p arguments are specified, the last one will take effect.If arguments other than -p are specified, Bazel will fall back to use patch command line tool instead of the Bazel-native patch implementation. When falling back to patch command line tool and patch_tool attribute is not specified, <code>patch</code> will be used. This only affects patch files in the <code>patches</code> attribute.   | List of strings | optional | ["-p0"] |
| <a id="http_embedded_archive-patch_cmds"></a>patch_cmds |  Sequence of Bash commands to be applied on Linux/Macos after patches are applied.   | List of strings | optional | [] |
| <a id="http_embedded_archive-patch_cmds_win"></a>patch_cmds_win |  Sequence of Powershell commands to be applied on Windows after patches are applied. If this attribute is not set, patch_cmds will be executed on Windows, which requires Bash binary to exist.   | List of strings | optional | [] |
| <a id="http_embedded_archive-patch_tool"></a>patch_tool |  The patch(1) utility to use. If this is specified, Bazel will use the specifed patch tool instead of the Bazel-native patch implementation.   | String | optional | "" |
| <a id="http_embedded_archive-patches"></a>patches |  A list of files that are to be applied as patches after extracting the archive. By default, it uses the Bazel-native patch implementation which doesn't support fuzz match and binary patch, but Bazel will fall back to use patch command line tool if <code>patch_tool</code> attribute is specified or there are arguments other than <code>-p</code> in <code>patch_args</code> attribute.   | <a href="https://bazel.build/docs/build-ref.html#labels">List of labels</a> | optional | [] |
| <a id="http_embedded_archive-remote_patch_strip"></a>remote_patch_strip |  The number of leading slashes to be stripped from the file name in the remote patches.   | Integer | optional | 0 |
| <a id="http_embedded_archive-remote_patches"></a>remote_patches |  A map of patch file URL to its integrity value, they are applied after extracting the archive and before applying patch files from the <code>patches</code> attribute. It uses the Bazel-native patch implementation, you can specify the patch strip number with <code>remote_patch_strip</code>   | <a href="https://bazel.build/docs/skylark/lib/dict.html">Dictionary: String -> String</a> | optional | {} |
| <a id="http_embedded_archive-repo_mapping"></a>repo_mapping |  A dictionary from local repository name to global repository name. This allows controls over workspace dependency resolution for dependencies of this repository.&lt;p&gt;For example, an entry <code>"@foo": "@bar"</code> declares that, for any time this repository depends on <code>@foo</code> (such as a dependency on <code>@foo//some:target</code>, it should actually resolve that dependency within globally-declared <code>@bar</code> (<code>@bar//some:target</code>).   | <a href="https://bazel.build/docs/skylark/lib/dict.html">Dictionary: String -> String</a> | required |  |
| <a id="http_embedded_archive-sha256"></a>sha256 |  The expected SHA-256 of the file downloaded.<br><br>This must match the SHA-256 of the file downloaded. _It is a security risk to omit the SHA-256 as remote files can change._ At best omitting this field will make your build non-hermetic. It is optional to make development easier but should be set before shipping.   | String | optional | "" |
| <a id="http_embedded_archive-strip_prefix"></a>strip_prefix |  A directory prefix to strip from extracted files in the inner archive.<br><br>Many archives contain a top-level directory that contains all of the useful files in archive. Instead of needing to specify this prefix over and over in the <code>build_file</code>, this field can be used to strip it from all of the extracted files.<br><br>For example, suppose you are using <code>foo-lib-latest.zip</code>, which contains the directory <code>foo-lib-1.2.3/</code> under which there is a <code>WORKSPACE</code> file and are <code>src/</code>, <code>lib/</code>, and <code>test/</code> directories that contain the actual code you wish to build. Specify <code>strip_prefix = "foo-lib-1.2.3"</code> to use the <code>foo-lib-1.2.3</code> directory as your top-level directory.<br><br>Note that if there are files outside of this directory, they will be discarded and inaccessible (e.g., a top-level license file). This includes files/directories that start with the prefix but are not in the directory (e.g., <code>foo-lib-1.2.3.release-notes</code>). If the specified prefix does not match a directory in the archive, Bazel will return an error.   | String | optional | "" |
| <a id="http_embedded_archive-type"></a>type |  The archive type of the downloaded file.<br><br>By default, the archive type is determined from the file extension of the URL. If the file has no extension, you can explicitly specify one of the following: <code>"zip"</code>, <code>"jar"</code>, <code>"war"</code>, <code>"aar"</code>, <code>"tar"</code>, <code>"tar.gz"</code>, <code>"tgz"</code>, <code>"tar.xz"</code>, or <code>tar.bz2</code>.   | String | optional | "" |
| <a id="http_embedded_archive-url"></a>url |  A URL to a file that will be made available to Bazel.<br><br>This must be a file, http or https URL. Redirections are followed. Authentication is not supported.<br><br>This parameter is to simplify the transition from the native http_archive rule. More flexibility can be achieved by the urls parameter that allows to specify alternative URLs to fetch from.   | String | optional | "" |
| <a id="http_embedded_archive-urls"></a>urls |  A list of URLs to a file that will be made available to Bazel.<br><br>Each entry must be a file, http or https URL. Redirections are followed. Authentication is not supported.<br><br>URLs are tried in order until one succeeds, so you should list local mirrors first. If all downloads fail, the rule will fail.   | List of strings | optional | [] |
| <a id="http_embedded_archive-workspace_file"></a>workspace_file |  The file to use as the <code>WORKSPACE</code> file for this repository. Either <code>workspace_file</code> or <code>workspace_file_content</code> can be specified, or neither, but not both.   | <a href="https://bazel.build/docs/build-ref.html#labels">Label</a> | optional | None |
| <a id="http_embedded_archive-workspace_file_content"></a>workspace_file_content |  The content for the WORKSPACE file for this repository. Either <code>workspace_file</code> or <code>workspace_file_content</code> can be specified, or neither, but not both.   | String | optional | "" |


