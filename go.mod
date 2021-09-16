module github.com/hashicorp/terraform-plugin-sdk/v2

go 1.15

require (
	cloud.google.com/go/storage v1.10.0
	github.com/Azure/azure-sdk-for-go v57.2.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.19
	github.com/agext/levenshtein v1.2.2
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1258
	github.com/aliyun/aliyun-oss-go-sdk v2.1.10+incompatible
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0
	github.com/apparentlymart/go-shquot v0.0.1
	github.com/apparentlymart/go-versions v1.0.1
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2
	github.com/aws/aws-sdk-go v1.40.25
	github.com/bgentry/speakeasy v0.1.0
	github.com/bmatcuk/doublestar v1.3.4
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/davecgh/go-spew v1.1.1
	github.com/dylanmei/winrmtest v0.0.0-20210303004826-fbc9ae56efb6
	github.com/fatih/color v1.9.0 // indirect
	github.com/frankban/quicktest v1.11.0 // indirect
	github.com/go-test/deep v1.0.3
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.3
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.2.0
	github.com/hashicorp/aws-sdk-go-base v0.7.1
	github.com/hashicorp/errwrap v1.1.0
	github.com/hashicorp/go-azure-helpers v0.16.5
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/go-getter v1.5.2
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-plugin v1.4.1
	github.com/hashicorp/go-retryablehttp v0.6.7
	github.com/hashicorp/go-tfe v0.18.0
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.2.1
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/hcl/v2 v2.10.1
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform v1.0.7
	github.com/hashicorp/terraform-config-inspect v0.0.0-20210625153042-09f34846faab
	github.com/hashicorp/terraform-exec v0.13.0
	github.com/hashicorp/terraform-json v0.8.0
	github.com/hashicorp/terraform-plugin-go v0.3.0
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce // indirect
	github.com/jmespath/go-jmespath v0.4.0
	github.com/joyent/triton-go v1.7.1-0.20200416154420-6801d15b779f
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/keybase/go-crypto v0.0.0-20161004153544-93f5b35093ba
	github.com/lib/pq v1.8.0
	github.com/lusis/go-artifactory v0.0.0-20180304164534-a47f63f234b2
	github.com/masterzen/winrm v0.0.0-20210623064412-3b76017826b0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12
	github.com/mitchellh/cli v1.1.2
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
	github.com/mitchellh/copystructure v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-linereader v0.0.0-20190213213312-1b945b3263eb
	github.com/mitchellh/go-testing-interface v1.0.4
	github.com/mitchellh/go-wordwrap v1.0.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/mitchellh/panicwrap v1.0.0
	github.com/mitchellh/reflectwalk v1.0.1
	github.com/packer-community/winrmcp v0.0.0-20180921211025-c76d91c1e7db
	github.com/pierrec/lz4 v2.5.2+incompatible
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/pkg/errors v0.9.1
	github.com/posener/complete v1.2.3
	github.com/spf13/afero v1.2.2
	github.com/tombuildsstuff/giovanni v0.16.0
	github.com/xanzy/ssh-agent v0.2.1
	github.com/xlab/treeprint v1.1.0
	github.com/zclconf/go-cty v1.9.1
	github.com/zclconf/go-cty-debug v0.0.0-20191215020915-b22d67c1ba0b
	github.com/zclconf/go-cty-yaml v1.0.2
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/mod v0.4.2
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	golang.org/x/text v0.3.6
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	golang.org/x/tools v0.1.5
	google.golang.org/api v0.56.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
